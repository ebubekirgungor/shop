<template>
  <main
    class="sm:bg-gray-50 flex flex-col items-center justify-center gap-y-4 m-4"
  >
    <h1 class="mt-4 text-xl">{{ $t("create_new_account") }}</h1>
    <div
      class="flex justify-center bg-white sm:pt-8 sm:pb-4 sm:w-[25rem] sm:h-[31rem] sm:border sm:rounded-xl"
    >
      <form
        v-if="step == 1"
        @submit.prevent="check_email"
        class="flex flex-col gap-y-6 w-[300px]"
      >
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
          <label for="first_name">{{ $t("first_name") }}</label>
          <input
            class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent"
            v-model="form.first_name"
            name="first_name"
            type="text"
            required
          />
        </div>
        <div class="flex flex-col gap-y-1 grow">
          <label for="last_name">{{ $t("last_name") }}</label>
          <input
            class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent"
            v-model="form.last_name"
            name="last_name"
            type="text"
            required
          />
        </div>
        <div>
          <button
            type="submit"
            :disabled="
              form.email == '' || form.first_name == '' || form.last_name == ''
            "
            class="sm:mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
          >
            {{ $t("next") }}
          </button>
          <h1 class="grow text-center hidden sm:block">{{ $t("or") }}</h1>
          <NuxtLinkLocale
            to="login"
            class="flex justify-center items-center my-4 w-full transition duration-300 ease-in-out h-12 rounded-full border border-gray-200 bg-white hover:bg-gray-200"
          >
            {{ $t("sign_in") }}
          </NuxtLinkLocale>
        </div>
      </form>
      <form
        v-else
        @submit.prevent="register"
        class="flex flex-col gap-y-6 w-[300px]"
      >
        <div class="flex flex-col gap-y-1">
          <label for="phone">{{ $t("phone") }}</label>
          <input
            class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent"
            v-model="form.phone"
            name="phone"
            minlength="14"
            maxlength="14"
            type="tel"
            @input="phone_format"
            required
          />
        </div>
        <div class="flex flex-col gap-y-1">
          <label for="phone">{{ $t("birthdate") }}</label>
          <div class="flex gap-x-2">
            <select
              class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent w-full"
              v-model="form.birthdate.day"
              required
            >
              <option value="">{{ $t("day") }}</option>
              <option v-for="i in 31" :value="i.toString()">{{ i }}</option>
            </select>
            <select
              class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent w-full"
              v-model="form.birthdate.month"
              required
            >
              <option value="">{{ $t("month") }}</option>
              <option v-for="i in 12" :value="i.toString()">{{ i }}</option>
            </select>
            <select
              class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent w-full"
              v-model="form.birthdate.year"
              required
            >
              <option value="">{{ $t("year") }}</option>
              <option
                v-for="i in Array.from(
                  { length: new Date().getFullYear() - 1919 },
                  (_, index) => index + 1920
                ).reverse()"
                :value="i.toString()"
              >
                {{ i }}
              </option>
            </select>
          </div>
        </div>
        <div class="flex flex-col gap-y-1">
          <label for="phone">{{ $t("gender") }}</label>
          <div class="flex gap-x-8">
            <label class="flex items-center gap-x-3 cursor-pointer">
              <input
                class="transition duration-200 ease-in-out text-black cursor-pointer focus:ring-0 focus:ring-offset-0"
                type="radio"
                v-model="form.gender"
                name="gender"
                id="male"
                value="m"
                required
              />
              {{ $t("male") }}
            </label>
            <label class="flex items-center gap-x-3 cursor-pointer">
              <input
                class="transition duration-200 ease-in-out text-black cursor-pointer focus:ring-0 focus:ring-offset-0"
                type="radio"
                v-model="form.gender"
                name="gender"
                id="female"
                value="f"
              />
              {{ $t("female") }}
            </label>
          </div>
        </div>
        <div class="flex flex-col gap-y-1 grow">
          <label for="password">{{ $t("password") }}</label>
          <div class="flex items-center">
            <input
              class="transition duration-200 ease-in-out w-full h-11 border-slate-300 bg-gray-50 rounded-md text-sm focus:ring-slate-300 focus:ring-2 focus:border-transparent"
              v-model="form.password"
              name="password"
              :type="eye ? 'text' : 'password'"
              required
            />
            <button
              v-if="form.password != ''"
              @click="eye = !eye"
              type="button"
              :class="
                eye
                  ? 'size-6 absolute ml-[265px] bg-[url(/icons/eye.svg)]'
                  : 'size-6 absolute ml-[265px] bg-[url(/icons/eye_off.svg)]'
              "
            ></button>
          </div>
        </div>
        <button
          type="submit"
          :disabled="
            form.phone == '' ||
            form.birthdate.day == '' ||
            form.birthdate.month == '' ||
            form.birthdate.year == '' ||
            form.gender == '' ||
            form.password == ''
          "
          class="sm:mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
        >
          {{ $t("create_account") }}
        </button>
      </form>
    </div>
  </main>
</template>
<script setup lang="ts">
import { useToast, POSITION } from "vue-toastification";
const { t } = useI18n();
const localePath = useLocalePath();
const toast = useToast();
const config = useRuntimeConfig().public;

definePageMeta({
  middleware: "auth",
});

const step = ref(1);
const eye = ref(false);

const form = ref({
  email: "",
  first_name: "",
  last_name: "",
  phone: "",
  birthdate: {
    day: "",
    month: "",
    year: "",
  },
  gender: "",
  password: "",
});

const check_email = async () => {
  await useFetch(config.apiBase + "/auth/check_user", {
    method: "post",
    body: {
      email: form.value.email,
    },
    onResponse({ response }) {
      if (response._data) {
        toast.warning(t("user_with_email_exists"), {
          bodyClassName: "toast-font",
          position: POSITION.TOP_CENTER,
        });
      } else {
        step.value++;
      }
    },
  });
};

const register = async () => {
  await useFetch(config.apiBase + "/users", {
    method: "post",
    body: {
      email: form.value.email,
      first_name: form.value.first_name,
      last_name: form.value.last_name,
      phone: form.value.phone.replace(/\D/g, ""),
      birthdate: `${form.value.birthdate.year}-${form.value.birthdate.month}-${form.value.birthdate.day}`,
      gender: form.value.gender,
      password: form.value.password,
      cart: [],
    },
    onResponse({ response }) {
      if (response._data.status == "success") {
        navigateTo(localePath("login"));
      } else {
        toast.error(t("server_error"), {
          bodyClassName: "toast-font",
          position: POSITION.TOP_CENTER,
        });
      }
    },
  });
};

const phone_format = () => {
  let x: RegExpMatchArray | null = form.value.phone
    .replace(/\D/g, "")
    .match(/(\d{0,3})(\d{0,3})(\d{0,4})/) as RegExpMatchArray;
  form.value.phone = !x[2]
    ? x[1]
    : "(" + x[1] + ") " + x[2] + (x[3] ? "-" + x[3] : "");
};
</script>
