<template>
  <main class="flex flex-col gap-y-4">
    <div
      class="flex items-center p-6 text-xl w-[50vw] h-auto bg-white rounded-xl shadow-md"
    >
      Personal Details
    </div>
    <div
      class="grid grid-cols-2 gap-x-16 gap-y-8 items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md"
    >
      <label :class="label">Name<input :class="input" type="text" v-model="userdata!.first_name" /></label>
      <label :class="label">Surname<input :class="input" type="text" /></label>
      <label :class="label"
        >Phone<input
          :class="input"
          type="tel"
          pattern="[0-9]{3}-[0-9]{3}-[0-9]{4}"
          v-model="userdata!.phone"
      /></label>
      <label :class="label"
        >E-mail<input
          disabled
          :class="input + ' bg-black/20 cursor-not-allowed'"
          type="text"
      /></label>
      <label :class="label">
        Birthdate
        <div class="flex gap-x-8">
          <select :class="input + ' w-full'">
            <option>Day</option>
            <option v-for="i in 31">{{ i }}</option>
          </select>
          <select :class="input + ' w-full'">
            <option>Month</option>
            <option v-for="i in 12">{{ i }}</option>
          </select>
          <select :class="input + ' w-full'">
            <option>Year</option>
            <option
              v-for="i in Array.from(
                { length: new Date().getFullYear() - 1919 },
                (_, index) => index + 1920
              ).reverse()"
            >
              {{ i }}
            </option>
          </select>
        </div>
      </label>
      <div :class="label">
        Gender
        <div class="flex gap-x-8">
          <label class="flex items-center gap-x-3 cursor-pointer">
            <input :class="radio" type="radio" name="gender" id="male" />
            Male
          </label>
          <label class="flex items-center gap-x-3 cursor-pointer">
            <input :class="radio" type="radio" name="gender" id="female" />
            Female
          </label>
        </div>
      </div>
      <button
        class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
      >
        Update
      </button>
    </div>
  </main>
</template>
<script setup lang="ts">
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
interface User {
  first_name: string;
  last_name: string;
  phone: number;
  gender: string;
}
const input =
  "transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const label = "flex flex-col gap-y-2";
const radio = "transition duration-200 ease-in-out cursor-pointer focus:ring-0";
const { data: userdata } = await useFetch<User>("/api/users", {
  headers: {
    Authorization: `Bearer ${user.token}`,
  },
  query: { email: user.email },
});
</script>
