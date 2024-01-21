<template>
  <main class="flex flex-col gap-y-4">
    <div
      class="flex items-center p-6 text-xl h-auto bg-white rounded-xl shadow-md"
    >
      Change Password
    </div>
    <form
      @submit.prevent="update"
      class="flex flex-col items-center justify-center p-6 gap-y-5 w-[50vw] h-auto bg-white rounded-xl shadow-md"
    >
      <input
        :class="input"
        type="password"
        placeholder="Old Password"
        v-model="form.old_password"
      />
      <input
        :class="input"
        type="password"
        placeholder="New Password"
        v-model="form.new_password"
      />
      <input
        :class="input"
        type="password"
        placeholder="Retype New Password"
        v-model="form.new_password2"
      />
      <button
        :disabled="
          !form.old_password ||
          !form.new_password ||
          form.new_password != form.new_password2
        "
        class="transition duration-300 ease-in-out w-80 h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
        type="submit"
      >
        Change Password
      </button>
    </form>
  </main>
</template>
<script setup lang="ts">
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
const form = ref({
  old_password: "",
  new_password: "",
  new_password2: "",
});
const update = async () => {
  await useFetch(`/api/users/password/${user.id}`, {
    method: "patch",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    body: {
      old_password: form.value.old_password,
      new_password: form.value.new_password,
    },
  });
};
const input =
  "transition duration-300 ease-in-out w-80 rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
</script>
