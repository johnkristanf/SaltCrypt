package database

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"saltcrypt/types"

	"golang.org/x/crypto/bcrypt"
)


var pepper = "s3cr3tP3pp3r"


func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}


func HashPassword(password, salt, pepper string) (string, error) {
	fullPassword := password + salt + pepper
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fullPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterDB(user types.UserDTO) error {

	conn := ConnectDB();
	defer conn.Close(context.Background())

	salt, saltErr := GenerateSalt()
	if saltErr != nil {
		return fmt.Errorf("unable to generate salt: %w", saltErr)
	}

	// Hash the password with salt and pepper
	hashedPassword, hashErr := HashPassword(user.Password, salt, pepper)
	if hashErr != nil {
		return fmt.Errorf("unable to hash password: %w", hashErr)
	}

	query := `INSERT INTO users (email, password, salt, pepper) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := conn.QueryRow(context.Background(), query, user.Email, hashedPassword, salt, pepper).Scan(&id)
	if err != nil {
		return fmt.Errorf("unable to execute the query: %w", err)
	}

	fmt.Printf("Inserted user with ID: %d\n", id)
	return nil

}


func VerifyPassword(inputPassword, storedHash, salt, pepper string) error {
	// Hash the input password using the salt and pepper
	fullPassword := inputPassword + salt + pepper

	// Compare the hashed input password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(fullPassword))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}


func LoginDB(user types.UserDTO) error {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	fmt.Printf("Email sa login: %s", user.Email)
	fmt.Printf("Password sa login: %s", user.Password)

	// Query to get the user's hashed password and salt from the database
	query := `SELECT password, salt FROM users WHERE email = $1`

	var storedHash, storedSalt string
	err := conn.QueryRow(context.Background(), query, user.Email).Scan(&storedHash, &storedSalt)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	err = VerifyPassword(user.Password, storedHash, storedSalt, pepper)
	if err != nil {
		return fmt.Errorf("invalid_credentials")
	}

	fmt.Println("Login successful!")
	return nil
}

func FetchUsers()([]types.FetchUsers, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	query := `SELECT email, password, salt, pepper, created_at FROM public.users`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var users []types.FetchUsers

	for rows.Next() {
		var user types.FetchUsers
		err := rows.Scan(&user.Email, &user.Password, &user.Salt, &user.Pepper, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	// Check for any errors encountered during iteration
	if rows.Err() != nil {
		return nil, fmt.Errorf("error during row iteration: %w", rows.Err())
	}

	return users, nil
}