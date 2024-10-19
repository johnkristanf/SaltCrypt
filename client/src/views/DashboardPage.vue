<script lang="ts" setup>
import { ref, onMounted } from 'vue';

interface User {
  email: string;
  password: string;
  salt: string;
  pepper: string;
  created_at: string;
}

const users = ref<User[]>([]); 

const emailString = localStorage.getItem('email');
const userEmail = emailString ? JSON.parse(emailString) : ''; 



const fetchUsers = async () => {
  try {
    const response = await fetch('https://saltcrypt.onrender.com/api/users');
    if (!response.ok) {
      throw new Error('Failed to fetch users');
    }

    const data: User[] = await response.json(); 
    users.value = data; 
  } catch (error) {
    console.error(error);
  }
};

const signOut = () => window.location.href = '/login'

onMounted(() => {
  fetchUsers();
});
</script>

<template>

  <div class="fixed top-0 w-full flex justify-between items-center mb-3 p-3 bg-black px-[5rem]">

    <div class="flex items-center">
      <img src="/img/saltcrypt_logo.png" alt="SaltCrypt Logo" width="50">
      <h1 class="font-semibold my-5 text-4xl text-white">SaltCrypt</h1>
    </div>

    <div class="flex items-center gap-5 text-white ">
      <h1 class="font-semibold my-5 text-xl ">{{ userEmail || 'No email found' }}</h1>

      <div @click="signOut" class="flex items-center gap-1 hover:opacity-75 hover:cursor-pointer">
        <img src="/img/signout.png" alt="User Icon" width="30" height="30">
        <h1>Sign Out</h1>
      </div>

    </div>


  </div>

  <div class="w-full h-screen flex justify-center p-12 bg-violet-600">
    <div class="w-full mt-4">

      <h1 class="font-semibold my-5 text-4xl text-white mt-[5rem]">Registered Accounts</h1>

      <div class="flex flex-col">
        <div class="-m-1.5 overflow-x-auto h-[400px] ">
          <div class="p-1.5 min-w-full inline-block align-middle">
            <div class="border rounded-lg shadow overflow-hidden">
              <table class="min-w-full divide-y divide-gray-200">
                <thead class="bg-white">
                  <tr>
                    <th scope="col" class="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Email</th>
                    <th scope="col" class="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Password</th>
                    <th scope="col" class="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Salt</th>
                    <th scope="col" class="px-6 py-3 text-end text-xs font-medium text-gray-500 uppercase">Pepper</th>
                  </tr>
                </thead>
                <tbody class="divide-y bg-white divide-gray-200">
                  <tr
                    v-for="(user, index) in users"
                    :key="index"
                  >
                    <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-800">{{ user.email }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{{ user.password }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{{ user.salt }}</td>
                    <td class="px-6 py-4 whitespace-nowrap text-end text-sm font-medium">
                      {{ user.pepper }}
                    </td>

                    

                  </tr>

                 
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
