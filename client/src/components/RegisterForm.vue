<script lang="ts" setup>
import { reactive, ref } from 'vue';
import Swal from 'sweetalert2';

// Define the form data interface and create a reactive object for form data
interface FormData {
  email: string;
  password: string;
}

const formData = reactive<FormData>({
  email: '',
  password: ''
});

const passwordStrength = ref(0);
const errorMessage = ref('');
const successMessage = ref('');

function checkPasswordStrength() {
  const hasUppercase = /[A-Z]/.test(formData.password);
  const hasNumber = /\d/.test(formData.password);
  const hasSymbol = /[!@#$%^&*(),.?":{}|<>]/.test(formData.password);
  const isLongEnough = formData.password.length >= 8;

  let strength = 0;

  // Increment strength based on criteria met
  if (isLongEnough) strength += 1;
  if (hasUppercase) strength += 1;
  if (hasNumber) strength += 1;
  if (hasSymbol) strength += 1;

  passwordStrength.value = strength;
}

const handleSubmit = async () => {
  errorMessage.value = '';
  successMessage.value = '';

  // Show the loading spinner with SweetAlert2
  Swal.fire({
    title: 'Registering...',
    text: 'Please wait while we process your registration',
    allowOutsideClick: false,
    showConfirmButton: false,
    willOpen: () => {
      Swal.showLoading();
    }
  });

  try {
    const payload = {
      email: formData.email,
      password: formData.password
    };

    console.log("payload: ", payload);

    // Simulate fetch request to API
    const response = await fetch('http://localhost:8080/api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });

    const data = await response.json();

    // Close the loading spinner
    Swal.close();

    if (response.ok) {
      successMessage.value = 'Registration successful! You will be redirected to the login page.';
      formData.email = '';
      formData.password = '';
      passwordStrength.value = 0;

      // Show success message with SweetAlert2
      Swal.fire({
        icon: 'success',
        title: 'Registration Complete',
        text: 'You will be redirected to the login page shortly.',
        timer: 2000,
        showConfirmButton: false
      });

      setTimeout(() => {
        window.location.href = '/login';
      }, 2000);

    } else {
      errorMessage.value = data.message || 'Registration failed.';
      passwordStrength.value = 0;

      // Show error message with SweetAlert2
      Swal.fire({
        icon: 'error',
        title: 'Registration Failed',
        text: errorMessage.value
      });
    }
  } catch (error) {
    // Close the loading spinner and show an error message
    Swal.close();
    errorMessage.value = 'An error occurred. Please try again.';
    passwordStrength.value = 0;
    console.error("Error in Register: ", error);

    // Show error message with SweetAlert2
    Swal.fire({
      icon: 'error',
      title: 'Error',
      text: 'An error occurred during registration. Please try again.'
    });
  }
};
</script>


<template>
  <section class="bg-gray-50 dark:bg-gray-900 h-[80%] w-[30%] flex items-center">
    <div class="w-full flex flex-col items-center justify-center px-6 py-8">
      <div class="w-full rounded-lg shadow dark:border dark:bg-gray-800 dark:border-gray-700">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
            Sign up to get started
          </h1>

          <!-- Display success or error messages -->
          <p v-if="errorMessage" class="text-red-500 text-sm mt-4">{{ errorMessage }}</p>
          <p v-if="successMessage" class="text-green-500 text-sm mt-4">{{ successMessage }}</p>

          <!-- Form starts here -->
          <form class="space-y-4 md:space-y-6" @submit.prevent="handleSubmit">
            <div>
              <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your email</label>
              <input
                type="email"
                name="email"
                id="email"
                v-model="formData.email"
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
                v-model="formData.password"
                @input="checkPasswordStrength"
                placeholder="••••••••"
                class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                required
              />

              <!-- Password Strength Progress Bar -->
              <div v-if="passwordStrength > 0" class="mt-2 w-full h-2 bg-gray-200 rounded-lg">
                <div
                  class="h-full rounded-lg"
                  :class="{
                    'bg-red-500': passwordStrength === 1,
                    'bg-yellow-500': passwordStrength === 2,
                    'bg-blue-500': passwordStrength === 3,
                    'bg-green-500': passwordStrength === 4
                  }"
                  :style="{ width: passwordStrength * 25 + '%' }"
                ></div>
              </div>

              <!-- Password Strength Text -->
              <div class="mt-2 text-sm">
                <span v-if="passwordStrength === 1" class="text-red-500">Weak</span>
                <span v-if="passwordStrength === 2" class="text-yellow-500">Medium</span>
                <span v-if="passwordStrength === 3" class="text-blue-500">Good</span>
                <span v-if="passwordStrength === 4" class="text-green-500">Strong</span>
              </div>
            </div>

            <button
              type="submit"
              class="w-full text-white bg-violet-600 hover:opacity-75 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Sign up
            </button>
            
            

            <p class="text-sm font-light text-gray-500 dark:text-gray-400">
              Already have an account?
              <router-link to="/login" class="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign in</router-link>
            </p>
          </form>
          <!-- Form ends here -->
        </div>
      </div>
    </div>
  </section>
</template>
