<template>
  <main class="sm:bg-gray-50 flex flex-col items-center justify-center gap-y-4">
    <h1 class="mt-8 text-xl">Sign in to your account</h1>
    <div
      class="flex justify-center bg-white sm:w-[400px] sm:h-[520px] sm:border sm:rounded-xl"
    >
      <div class="sm:mt-8 flex flex-col gap-y-6 w-[300px]">
        <div class="flex flex-col gap-y-1">
          <label for="email">E-mail</label>
          <input :class="input" v-model="email" name="email" type="text" />
        </div>
        <div class="flex flex-col gap-y-1">
          <label for="password">Password</label>
          <input
            :class="input"
            v-model="password"
            name="password"
            type="password"
          />
        </div>
        <NuxtLink to="/reset-password" class="text-center"
          >Forgot password?</NuxtLink
        >
        <h1 class="grow mt-5 text-center text-red-500">{{ login_error }}</h1>
        <div>
          <button
            @click="login"
            :disabled="email == '' || password == ''"
            class="sm:mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
          >
            Sign In
          </button>
          <h1 class="grow text-center hidden sm:block">or</h1>
          <NuxtLink
            to="/register"
            class="flex justify-center items-center my-4 w-full transition duration-300 ease-in-out h-12 rounded-full border border-gray-200 bg-white hover:bg-gray-200"
          >
            Create Account
          </NuxtLink>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
import { useUser } from "@/store/user";
const input =
  "transition duration-200 ease-in-out h-11 border-gray-300 bg-gray-50 rounded-md text-sm";
const email = ref("");
const password = ref("");
const login_error = ref("");
const login = async () => {
  const { data: response } = await useFetch("/api/auth/login", {
    method: "post",
    body: {
      email: email.value,
      password: password.value,
    },
  });
  if ((response.value as any).status == "success") {
    const user = useUser();
    user.user = (response.value as any).data;
    navigateTo("/account");
  } else login_error.value = (response.value as any).message as string;
};
</script>
