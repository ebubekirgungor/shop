<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div
      class="mt-10 sm:mt-0 flex justify-center sm:justify-start items-center py-4 sm:p-6 text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      {{ $t("personal_details") }}
    </div>
    <form
      v-if="fetch_complete"
      @submit.prevent="update"
      class="flex flex-col sm:grid sm:grid-cols-2 sm:gap-x-[7%] gap-y-4 sm:gap-y-8 sm:items-center sm:p-6 sm:min-w-[40rem] h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <label :class="label"
        >{{ $t("first_name")
        }}<input :class="input" type="text" v-model="form.first_name"
      /></label>
      <label :class="label"
        >{{ $t("last_name")
        }}<input :class="input" type="text" v-model="form.last_name"
      /></label>
      <label :class="label"
        >{{ $t("phone")
        }}<input
          :class="input"
          type="tel"
          minlength="14"
          maxlength="14"
          v-model="form.phone"
          @input="phone_format"
      /></label>
      <label :class="label"
        >{{ $t("email")
        }}<input
          disabled
          :class="input + ' bg-gray-300 cursor-not-allowed'"
          :placeholder="form.email"
          type="text"
      /></label>
      <label :class="label">
        {{ $t("birthdate") }}
        <div class="flex gap-x-[7%]">
          <select :class="input + input_select" v-model="form.birthdate.day">
            <option value="">{{ $t("day") }}</option>
            <option v-for="i in 31" :value="i.toString()">{{ i }}</option>
          </select>
          <select :class="input + input_select" v-model="form.birthdate.month">
            <option value="">{{ $t("month") }}</option>
            <option v-for="i in 12" :value="i.toString()">{{ i }}</option>
          </select>
          <select :class="input + input_select" v-model="form.birthdate.year">
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
      </label>
      <div :class="label">
        {{ $t("gender") }}
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
            {{ $t("male") }}
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
            {{ $t("female") }}
          </label>
        </div>
      </div>
      <button
        :disabled="JSON.stringify(form) == JSON.stringify(form_old)"
        type="submit"
        class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
      >
        {{ $t("update") }}
      </button>
    </form>
    <div
      v-else
      class="flex flex-col sm:grid sm:grid-cols-2 sm:gap-x-[7%] gap-y-4 sm:gap-y-8 sm:items-center sm:p-6 sm:min-w-[40rem] h-auto bg-white rounded-xl shadow-md animate-pulse"
    >
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
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const { t } = useI18n();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
interface Form {
  email: string;
  first_name: string;
  last_name: string;
  phone: string;
  birthdate: {
    day: string;
    month: string;
    year: string;
  };
  gender: string;
}
const fetch_complete = ref(false);
const form = ref<Form>({
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
});
const form_old = ref<Form>({
  email: "",
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
const data_to_form = async (data: Form) => {
  form.value = { ...data };
  const birthdate = (data.birthdate as unknown as string).slice(0, 10);
  form.value.birthdate = {
    day: birthdate == "" ? "0" : new Date(birthdate).getDate().toString(),
    month:
      birthdate == "" ? "0" : (new Date(birthdate).getMonth() + 1).toString(),
    year: birthdate == "" ? "0" : new Date(birthdate).getFullYear().toString(),
  };
  phone_format();
  form_old.value = { ...form.value };
  form_old.value.birthdate = { ...form.value.birthdate };
};
onMounted(() => {
  nextTick(async () => {
    await useFetch(config.apiBase + "/users", {
      headers: {
        Authorization: config.apiKey,
      },
      onResponse({ response }) {
        data_to_form(response._data);
        fetch_complete.value = true;
      },
    });
  });
});
const update = async () => {
  await useFetch(config.apiBase + "/users", {
    headers: {
      Authorization: config.apiKey,
    },
    method: "patch",
    body: {
      first_name: form.value.first_name,
      last_name: form.value.last_name,
      phone: form.value.phone.replace(/\D/g, ""),
      birthdate: `${form.value.birthdate.year}-${form.value.birthdate.month}-${form.value.birthdate.day}`,
      gender: form.value.gender,
    },
    onResponse({ response }) {
      if (response._data.ID) {
        toast.success(t("user_updated"), {
          bodyClassName: "toast-font",
        });
        data_to_form(response._data);
      } else {
        toast.warning(response._data.error, {
          bodyClassName: "toast-font",
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
const input =
  "transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const input_select = " w-full cursor-pointer";
const label = "flex flex-col gap-y-2";
const radio =
  "transition duration-200 ease-in-out text-black cursor-pointer focus:ring-0 focus:ring-offset-0";
const skeleton_title = "w-32 h-5 bg-gray-200 rounded-full";
const skeleton_input = "w-full h-9 bg-gray-200 rounded-full";
</script>
