<template>
  <main class="flex flex-col gap-y-4">
    <div
      class="flex items-center p-6 text-xl h-auto bg-white rounded-xl shadow-md"
    >
      Personal Details
    </div>
    <form v-if="fetch_complete" @submit.prevent="update" :class="form_div">
      <label :class="label"
        >First Name<input :class="input" type="text" v-model="form.first_name"
      /></label>
      <label :class="label"
        >Last Name<input :class="input" type="text" v-model="form.last_name"
      /></label>
      <label :class="label"
        >Phone<input
          :class="input"
          type="tel"
          minlength="10"
          maxlength="10"
          v-model="form.phone"
      /></label>
      <label :class="label"
        >E-mail<input
          disabled
          :class="input + ' bg-gray-500/40 cursor-not-allowed'"
          :placeholder="(user.email as any)"
          type="text"
      /></label>
      <label :class="label">
        Birthdate
        <div class="flex gap-x-[7%]">
          <select :class="input + ' w-full'" v-model="form.birthdate.day">
            <option value="0">Day</option>
            <option v-for="i in 31" :value="i.toString()">{{ i }}</option>
          </select>
          <select :class="input + ' w-full'" v-model="form.birthdate.month">
            <option value="0">Month</option>
            <option v-for="i in 12" :value="i.toString()">{{ i }}</option>
          </select>
          <select :class="input + ' w-full'" v-model="form.birthdate.year">
            <option value="0">Year</option>
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
      </label>
      <div :class="label">
        Gender
        <div class="flex gap-x-[7%]">
          <label class="flex items-center gap-x-3 cursor-pointer">
            <input
              :class="radio"
              type="radio"
              v-model="form.gender"
              name="gender"
              id="male"
              value="m"
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
      <button
        :disabled="JSON.stringify(form) == JSON.stringify(form_old)"
        type="submit"
        class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
      >
        Update
      </button>
    </form>
    <div v-else :class="form_div + ' animate-pulse'">
      <div :class="label">
        <div :class="skeleton_title"></div>
        <div :class="skeleton_input"></div>
      </div>
      <div :class="label">
        <div :class="skeleton_title"></div>
        <div :class="skeleton_input"></div>
      </div>
      <div :class="label">
        <div :class="skeleton_title"></div>
        <div :class="skeleton_input"></div>
      </div>
      <div :class="label">
        <div :class="skeleton_title"></div>
        <div :class="skeleton_input"></div>
      </div>
      <div :class="label">
        <div :class="skeleton_title"></div>
        <div class="flex gap-x-8">
          <div :class="skeleton_input"></div>
          <div :class="skeleton_input"></div>
          <div :class="skeleton_input"></div>
        </div>
      </div>
      <div :class="label">
        <div :class="skeleton_title + ' !w-16'"></div>
        <div class="flex gap-x-8">
          <div class="flex gap-x-4">
            <div :class="skeleton_title + ' !w-5'"></div>
            <div :class="skeleton_title + ' !w-12'"></div>
          </div>
          <div class="flex gap-x-4">
            <div :class="skeleton_title + ' !w-5'"></div>
            <div :class="skeleton_title + ' !w-12'"></div>
          </div>
        </div>
      </div>
      <div class="w-full h-12 bg-gray-200 col-span-2 rounded-full"></div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
const fetch_complete = ref(false);
const form = ref({
  first_name: "",
  last_name: "",
  phone: "",
  birthdate: {
    day: "0",
    month: "0",
    year: "0",
  },
  gender: "",
});
const form_old = ref({
  first_name: "",
  last_name: "",
  phone: "",
  birthdate: {
    day: "0",
    month: "0",
    year: "0",
  },
  gender: "",
});

const form_div =
  "grid grid-cols-2 gap-x-16 gap-y-8 items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md";
const input =
  "transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const label = "flex flex-col gap-y-2";
const radio = "transition duration-200 ease-in-out cursor-pointer focus:ring-0";
const skeleton_title = "w-32 h-5 bg-gray-200 rounded-full";
const skeleton_input = "w-full h-9 bg-gray-200 rounded-full";
onMounted(() => {
  nextTick(async () => {
    const { data } = await useFetch<any>("/api/users", {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
      query: { email: user.email },
    });
    form.value = { ...data.value };
    const birthdate = data!.value!.birthdate.slice(0, 10);
    form.value.birthdate = {
      day: birthdate == "" ? "0" : new Date(birthdate).getDate().toString(),
      month:
        birthdate == "" ? "0" : (new Date(birthdate).getMonth() + 1).toString(),
      year:
        birthdate == "" ? "0" : new Date(birthdate).getFullYear().toString(),
    };
    form_old.value = { ...form.value };
    form_old.value.birthdate = { ...form.value.birthdate };
    fetch_complete.value = true;
  });
});
const update = async () => {
  /*console.log(form.value);
  console.log(form_old.value);*/
  await useFetch("/api/users", {
    method: "patch",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { email: user.email },
    body: {
      first_name: form.value.first_name,
      last_name: form.value.last_name,
      phone: form.value.phone,
      birthdate: `${form.value.birthdate.year}-${form.value.birthdate.month}-${form.value.birthdate.day}`,
      gender: form.value.gender,
    },
  });
};
</script>
