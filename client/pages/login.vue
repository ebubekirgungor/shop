<template>
  <main class="sm:bg-gray-50 flex flex-col items-center justify-center gap-y-4">
    <h1 class="mt-8 text-xl">Sign in to your account</h1>
    <div
      class="flex justify-center bg-white sm:w-[400px] sm:h-[520px] sm:border sm:rounded-xl"
    >
      <form
        @submit.prevent="login"
        class="sm:mt-8 flex flex-col gap-y-6 w-[300px]"
      >
        <div class="flex flex-col gap-y-1">
          <label for="email">E-mail</label>
          <input
            :class="input"
            v-model="form.email"
            name="email"
            type="email"
            required
          />
        </div>
        <div class="flex flex-col gap-y-1">
          <label for="password">Password</label>
          <div class="flex items-center">
            <input
              :class="input"
              v-model="form.password"
              name="password"
              :type="eye ? 'text' : 'password'"
            />
            <button
              @click="eye = !eye"
              type="button"
              :class="
                eye
                  ? showpassword + 'bg-[url(/icons/eye.svg)]'
                  : showpassword + 'bg-[url(/icons/eye_off.svg)]'
              "
            ></button>
          </div>
        </div>
        <NuxtLink to="/reset-password" class="text-center"
          >Forgot password?</NuxtLink
        >
        <h1 class="grow mt-5 text-center text-red-500">{{ login_error }}</h1>
        <div>
          <button
            type="submit"
            :disabled="form.email == '' || form.password == ''"
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
      </form>
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
import { useUser } from "@/store/user";
const showpassword = "w-6 h-6 absolute ml-[265px] ";
const eye = ref(false);
const input =
  "transition duration-200 ease-in-out w-full h-11 border-gray-300 bg-gray-50 rounded-md text-sm";
const form = ref({
  email: "",
  password: "",
});
const login_error = ref("");
const login = async () => {
  const { data: response } = await useFetch("/api/auth/login", {
    method: "post",
    body: {
      email: form.value.email,
      password: form.value.password,
    },
  });
  if ((response.value as any).status == "success") {
    const user = useUser();
    user.user = (response.value as any).data;
    navigateTo("/account");
  } else login_error.value = (response.value as any).message as string;
};
</script>
