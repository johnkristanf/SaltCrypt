<script lang="ts" setup>
import { ref } from 'vue';
import Swal from 'sweetalert2';

// Reactive variables for the form inputs
const email = ref('');
const password = ref('');
const errorMessage = ref('');
const successMessage = ref('');

// Handle form submission
const handleSubmit = async () => {
    errorMessage.value = '';
    successMessage.value = '';

    const payload = {
        email: email.value,
        password: password.value
    };

    console.log("login payload: ", payload);

    // Show SweetAlert2 loader when the login process starts
    Swal.fire({
        title: 'Logging in...',
        text: 'Please wait while we process your login',
        allowOutsideClick: false,
        showConfirmButton: false,
        willOpen: () => {
            Swal.showLoading();
        }
    });

    try {
        const response = await fetch('https://saltcrypt.onrender.com/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload),
        });

        if (response.ok) {
            const data = await response.json();
            successMessage.value = 'Login successful!';
            console.log('Login successful:', data);

            // Store the email or user data in localStorage
            localStorage.setItem('email', JSON.stringify(data));

            // Close the Swal loader
            Swal.close();

            // Show success message with Swal
            Swal.fire({
                icon: 'success',
                title: 'Login Successful',
                text: 'Redirecting to the dashboard...',
                timer: 2000,
                showConfirmButton: false,
            });

            // Redirect to dashboard after delay
            setTimeout(() => {
                window.location.href = '/dashboard';
            }, 2000);
        } else {
            const errorData = await response.json();
            errorMessage.value = errorData.message || 'The credentials you entered are incorrect. Please try again.';

            // Close the Swal loader
            Swal.close();

            // Show error message with Swal
            Swal.fire({
                icon: 'error',
                title: 'Login Failed',
                text: errorMessage.value,
            });
        }
    } catch (error) {
        console.error("Error in Login: ", error);
        errorMessage.value = 'An error occurred during login. Please try again.';

        // Close the Swal loader
        Swal.close();

        // Show error message with Swal
        Swal.fire({
            icon: 'error',
            title: 'Error',
            text: 'An error occurred during login. Please try again.',
        });
    }
};
</script>



<template>
  <section class="bg-gray-50 dark:bg-gray-900 h-[80%] w-[30%] flex items-center">
    <div class="w-full flex flex-col items-center justify-center px-6 py-8">
      <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
            Sign in to your account
          </h1>

          <!-- Display success or error message here, below the heading -->
          <p v-if="errorMessage" class="text-red-500 text-sm">{{ errorMessage }}</p>
          <p v-if="successMessage" class="text-green-500 text-sm">{{ successMessage }}</p>

          <!-- Form starts here -->
          <form class="space-y-4 md:space-y-6" @submit.prevent="handleSubmit">
            <div>
              <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your email</label>
              <input
                type="email"
                name="email"
                id="email"
                v-model="email"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="name@company.com"
                required
              />
            </div>

            <div>
              <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
              <input
                type="password"
                name="password"
                id="password"
                v-model="password"
                placeholder="••••••••"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />
            </div>

            <button
              type="submit"
              class="w-full text-white bg-violet-600 hover:opacity-75 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Sign in
            </button>

            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
              Don’t have an account yet?
              <router-link to="/register" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</router-link>
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
