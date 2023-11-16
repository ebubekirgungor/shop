<template>
  <main class="sm:bg-gray-50 flex flex-col items-center justify-center gap-y-4">
    <h1 class="mt-8 text-xl">Create new account</h1>
    <div
      class="flex justify-center bg-white sm:w-[400px] sm:h-[520px] sm:border sm:rounded-xl"
    >
      <form v-if="step == 1" @submit.prevent="check_email" :class="form_step">
        <div :class="form_field">
          <label for="email">E-mail</label>
          <input
            :class="input"
            v-model="form.email"
            name="email"
            type="email"
            required
          />
        </div>
        <div :class="form_field">
          <label for="first_name">First Name</label>
          <input
            :class="input"
            v-model="form.first_name"
            name="first_name"
            type="text"
            required
          />
        </div>
        <div :class="form_field">
          <label for="last_name">Last Name</label>
          <input
            :class="input"
            v-model="form.last_name"
            name="last_name"
            type="text"
            required
          />
        </div>
        <h1 class="grow text-center text-red-500">{{ register_error }}</h1>
        <div>
          <button
            type="submit"
            :disabled="
              form.email == '' || form.first_name == '' || form.last_name == ''
            "
            :class="button"
          >
            Next
          </button>
          <h1 class="grow text-center hidden sm:block">or</h1>
          <NuxtLink
            to="/login"
            class="flex justify-center items-center my-4 w-full transition duration-300 ease-in-out h-12 rounded-full border border-gray-200 bg-white hover:bg-gray-200"
          >
            Sign In
          </NuxtLink>
        </div>
      </form>
      <form v-else @submit.prevent="register" :class="form_step">
        <div :class="form_field">
          <label for="phone">Phone</label>
          <input
            :class="input"
            v-model="form.phone"
            name="phone"
            minlength="10"
            maxlength="10"
            type="tel"
            required
          />
        </div>
        <div :class="form_field">
          <label for="phone">Birthdate</label>
          <div class="flex gap-x-2">
            <select
              :class="input + ' w-full'"
              v-model="form.birthdate.day"
              required
            >
              <option value="">Day</option>
              <option v-for="i in 31" :value="i.toString()">{{ i }}</option>
            </select>
            <select
              :class="input + ' w-full'"
              v-model="form.birthdate.month"
              required
            >
              <option value="">Month</option>
              <option v-for="i in 12" :value="i.toString()">{{ i }}</option>
            </select>
            <select
              :class="input + ' w-full'"
              v-model="form.birthdate.year"
              required
            >
              <option value="">Year</option>
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
        <div :class="form_field">
          <label for="phone">Gender</label>
          <div class="flex gap-x-8">
            <label class="flex items-center gap-x-3 cursor-pointer">
              <input
                :class="radio"
                type="radio"
                v-model="form.gender"
                name="gender"
                id="male"
                value="m"
                required
              />
              Male
            </label>
            <label class="flex items-center gap-x-3 cursor-pointer">
              <input
                :class="radio"
                type="radio"
                v-model="form.gender"
                name="gender"
                id="female"
                value="f"
              />
              Female
            </label>
          </div>
        </div>
        <div :class="form_field + ' grow'">
          <label for="password">Password</label>
          <div class="flex items-center">
            <input
              :class="input"
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
                  ? showpassword + 'bg-[url(/icons/eye.svg)]'
                  : showpassword + 'bg-[url(/icons/eye_off.svg)]'
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
          :class="button"
        >
          Create Account
        </button>
      </form>
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
const step = ref(1);
const form_step = "sm:mt-8 flex flex-col gap-y-6 w-[300px]";
const form_field = "flex flex-col gap-y-1";
const button =
  "sm:mb-4 w-full transition duration-300 ease-in-out h-12 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none";
const showpassword = "w-6 h-6 absolute ml-[265px] ";
const eye = ref(false);
const input =
  "transition duration-200 ease-in-out w-full h-11 border-gray-300 bg-gray-50 rounded-md text-sm";
const radio = "transition duration-200 ease-in-out cursor-pointer focus:ring-0";
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
const register_error = ref("");
const check_email = async () => {
  const { data: check } = await useFetch("/api/auth/check_user", {
    method: "post",
    body: {
      email: form.value.email,
    },
  });
  if (check.value) {
    register_error.value = "User with given e-mail is already exists";
  } else {
    step.value++;
  }
};
const register = async () => {
  const { data: response } = await useFetch("/api/users", {
    method: "post",
    body: {
      email: form.value.email,
      first_name: form.value.first_name,
      last_name: form.value.last_name,
      phone: form.value.phone,
      birthdate: `${form.value.birthdate.year}-${form.value.birthdate.month}-${form.value.birthdate.day}`,
      gender: form.value.gender,
      password: form.value.password,
    },
  });
  if ((response.value as any).status == "success") {
    navigateTo("/login");
  } else register_error.value = (response.value as any).message as string;
};
</script>
