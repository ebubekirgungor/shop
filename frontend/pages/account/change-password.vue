<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div
      class="mt-12 sm:mt-0 flex justify-center sm:justify-start items-center sm:p-6 sm:h-20 text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      {{ $t("change_password") }}
    </div>
    <form
      @submit.prevent="update"
      class="flex flex-col items-center justify-center p-6 gap-y-5 sm:min-w-[25rem] h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <input
        class="transition duration-300 ease-in-out w-80 rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
        type="password"
        :placeholder="$t('old_password')"
        v-model="form.old_password"
      />
      <input
        class="transition duration-300 ease-in-out w-80 rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
        type="password"
        :placeholder="$t('new_password')"
        v-model="form.new_password"
      />
      <input
        class="transition duration-300 ease-in-out w-80 rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
        type="password"
        :placeholder="$t('retype_new_password')"
        v-model="form.new_password2"
      />
      <button
        :disabled="
          !form.old_password ||
          !form.new_password ||
          form.new_password != form.new_password2
        "
        class="transition duration-300 ease-in-out w-full sm:w-80 h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
        type="submit"
      >
        {{ $t("change_password") }}
      </button>
    </form>
  </main>
</template>
<script setup lang="ts">
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const { t } = useI18n();

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
  await useFetch(config.apiBase + "/users/password", {
    headers: {
      Authorization: config.apiKey,
    },
    method: "patch",
    body: {
      old_password: form.value.old_password,
      new_password: form.value.new_password,
    },
    onResponse({ response }) {
      if (response._data.ID) {
        toast.success(t("password_changed"), {
          bodyClassName: "toast-font",
        });
      } else {
        toast.warning(response._data.error, {
          bodyClassName: "toast-font",
        });
      }
    },
  });
};
</script>
