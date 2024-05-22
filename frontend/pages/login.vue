<template>
  <main
    class="sm:bg-gray-50 flex flex-col items-center justify-center gap-y-4 m-4"
  >
    <h1 class="mt-4 text-xl">{{ $t("sign_in_to_your_account") }}</h1>
    <div
      class="flex justify-center bg-white sm:pt-8 sm:pb-4 sm:w-[25rem] sm:h-[31rem] sm:border sm:rounded-xl"
    >
      <form @submit.prevent="login" class="flex flex-col gap-y-6 w-[300px]">
        <div class="flex flex-col gap-y-1">
          <label for="email">{{ $t("email") }}</label>
          <input
            class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent"
            v-model="form.email"
            name="email"
            type="email"
            required
          />
        </div>
        <div class="flex flex-col gap-y-1">
          <label for="password">{{ $t("password") }}</label>
          <div class="flex items-center">
            <input
              class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent"
              v-model="form.password"
              name="password"
              :type="eye ? 'text' : 'password'"
            />
            <button
              v-if="form.password != ''"
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
        <div class="grow">
          <div class="flex justify-between">
            <div class="flex items-center gap-x-2">
              <input
                type="checkbox"
                v-model="form.remember_me"
                id="remember_me"
                class="transition duration-200 ease-in-out size-5 cursor-pointer rounded-md border-gray-300 text-gray-800 hover:border-gray-500 focus:ring-0 focus:ring-offset-0"
              /><label for="remember_me" class="cursor-pointer">{{
                $t("remember_me")
              }}</label>
            </div>
            <NuxtLink to="/reset-password">{{
              $t("forgot_password")
            }}</NuxtLink>
          </div>
        </div>
        <div>
          <button
            type="submit"
            :disabled="form.email == '' || form.password == ''"
            class="sm:mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
          >
            {{ $t("sign_in") }}
          </button>
          <h1 class="grow text-center hidden sm:block">{{ $t("or") }}</h1>
          <NuxtLinkLocale
            to="register"
            class="flex justify-center items-center my-4 w-full transition duration-300 ease-in-out h-12 rounded-full border border-gray-200 bg-white hover:bg-gray-200"
          >
            {{ $t("create_account") }}
          </NuxtLinkLocale>
        </div>
      </form>
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
import { useToast, POSITION } from "vue-toastification";
const { t } = useI18n();
const localePath = useLocalePath();
const toast = useToast();
const config = useRuntimeConfig().public;
const showpassword = "size-6 absolute ml-[265px] ";
const eye = ref(false);
const form = ref({
  email: "",
  password: "",
  remember_me: false,
});
const login = async () => {
  await useFetch(config.apiBase + "/auth/login", {
    method: "post",
    body: {
      email: form.value.email,
      password: form.value.password,
      remember_me: form.value.remember_me,
    },
    onResponse({ response }) {
      if (response._data.status == "success") {
        navigateTo(localePath("account"));
      } else {
        toast.error(
          response._data.message == "login"
            ? t("email_password_incorrect")
            : t("server_error"),
          {
            bodyClassName: "toast-font",
            position: POSITION.TOP_CENTER,
          }
        );
      }
    },
  });
};
</script>
