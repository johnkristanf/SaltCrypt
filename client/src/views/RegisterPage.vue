
<script setup lang="ts">
    import RegisterForm from '../components/RegisterForm.vue';
</script>

<template>
    <div class="h-screen w-full flex justify-center items-center font-semibold bg-violet-600">

        <div class="flex flex-col justify-center items-center gap-5 bg-black h-[80%] p-8 ">

            <div class="flex items-center">
                <img src="/img/saltcrypt_logo.png" alt="Encyptify Logo" width="80" height="80">
                <h1 class="text-white text-5xl">SaltCrypt</h1>
            </div>

            <h1 class="text-md opacity-75 text-gray-200 text-center">
               Welcome to SaltCrypt, offering advanced <br> 
               data protection with <br> secure hashing, salting, and peppering 
               <br> to safeguard your information.
           </h1>

        </div>

        <RegisterForm />
        
    </div>
</template>


<style>
    .background-image {
        background-image: url('/img/encryptify_bg.jpg');
        background-size: cover;
        background-position: center;
        background-repeat: no-repeat;        
    }
</style>
