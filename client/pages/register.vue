<template>
  <main class="bg-gray-50 flex flex-col items-center justify-center gap-y-4">
    <h1 class="mt-8 text-xl">Create new account</h1>
    <div
      class="flex justify-center bg-white w-[400px] h-[550px] border rounded-xl"
    >
      <div class="mt-16 flex flex-col gap-y-8 w-[300px]">
        <div class="flex flex-col">
          <label for="email">E-mail</label>
          <input :class="input" v-model="email" name="email" type="text" />
        </div>
        <div class="flex flex-col">
          <label for="password">Password</label>
          <input
            :class="input"
            v-model="password"
            name="password"
            type="password"
          />
        </div>
        <h1 class="grow text-center text-red-500">{{ register_error }}</h1>
        <div>
          <button
            @click="register"
            :disabled="email == ''"
            class="mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
          >
            Create Account
          </button>
          <h1 class="grow text-center">or</h1>
          <NuxtLink
            to="/login"
            class="flex justify-center items-center my-4 w-full transition duration-300 ease-in-out h-12 rounded-full border border-black bg-white hover:bg-black/10"
          >
          Sign In
          </NuxtLink>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
const input = ref(
  "transition duration-200 ease-in-out h-11 border-gray-300 bg-gray-50 rounded-md text-sm"
);
const email = ref("");
const password = ref("");
const register_error = ref("");
const register = async () => {
  const { data: response } = await useFetch("/api/users", {
    method: "post",
    body: {
      email: email.value,
      password: password.value,
    },
  });
  if ((response.value as any).email == email.value) {
    navigateTo("/login");
  } else register_error.value = (response.value as any).error as string;
};
</script>
