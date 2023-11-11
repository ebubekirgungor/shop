<template>
  <main class="sm:bg-gray-50 flex flex-col items-center justify-center gap-y-4">
    <h1 class="mt-8 text-xl">Create new account</h1>
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
        <div class="flex flex-col gap-y-1">
          <label for="confirm_password">Confirm Password</label>
          <input
            :class="input"
            v-model="confirm_password"
            name="confirm_password"
            type="password"
          />
        </div>
        <h1 class="grow text-center text-red-500">{{ register_error }}</h1>
        <div>
          <button
            @click="register"
            :disabled="
              email == '' || password == '' || password != confirm_password
            "
            class="sm:mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
          >
            Create Account
          </button>
          <h1 class="grow text-center hidden sm:block">or</h1>
          <NuxtLink
            to="/login"
            class="flex justify-center items-center my-4 w-full transition duration-300 ease-in-out h-12 rounded-full border border-gray-200 bg-white hover:bg-gray-200"
          >
            Sign In
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
const input =
  "transition duration-200 ease-in-out h-11 border-gray-300 bg-gray-50 rounded-md text-sm";
const email = ref("");
const password = ref("");
const confirm_password = ref("");
const register_error = ref("");
const register = async () => {
  const { data: check } = await useFetch("/api/auth/check_user", {
    method: "post",
    body: {
      email: email.value,
    },
  });
  if (check.value) {
    register_error.value = "User with given e-mail is already exists";
  } else {
    const { data: response } = await useFetch("/api/users", {
      method: "post",
      body: {
        email: email.value,
        password: password.value,
      },
    });
    if ((response.value as any).status == "success") {
      navigateTo("/login");
    } else register_error.value = (response.value as any).message as string;
  }
};
</script>
