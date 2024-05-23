<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <transition name="background" mode="out-in">
      <div
        v-if="create_update_dialog || delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 size-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="create_update_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed z-[3]"
      >
        <div
          class="flex flex-col justify-center sm:justify-start sm:p-3 bg-white sm:-mt-48 w-full sm:w-[28rem] h-screen sm:h-auto sm:rounded-xl"
        >
          <button
            @click="close_create_update_dialog"
            class="absolute sm:relative top-0 mt-4 mr-4 sm:m-0 self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-cover sm:bg-auto bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl sm:grow">
            {{ is_add ? $t("create_new_category") : $t("edit_category") }}
          </h1>
          <form
            class="flex flex-col items-center gap-y-6 mt-8 mb-4 mx-4 sm:mx-8"
            @submit.prevent="create_update"
          >
            <div class="flex flex-col gap-y-4 w-full">
              <input
                class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
                type="text"
                :placeholder="$t('title')"
                v-model="category_title"
                required
              />
              <div class="flex gap-x-4">
                <input
                  class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
                  type="text"
                  :placeholder="$t('filter_name')"
                  v-model="new_filter_input"
                />
                <button
                  class="transition duration-200 ease-in-out w-20 bg-gray-200 rounded-md hover:bg-gray-300 disabled:pointer-events-none"
                  type="button"
                  @click="
                    category_filters.push(new_filter_input);
                    new_filter_input = '';
                  "
                  :disabled="!new_filter_input"
                >
                  {{ $t("add") }}
                </button>
              </div>
              <div class="flex flex-wrap gap-2">
                <div
                  v-for="(filter, i) in category_filters"
                  class="w-min flex items-center pl-2 h-8 bg-gray-200 rounded-md"
                >
                  {{ filter }}
                  <button
                    type="button"
                    class="size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center"
                    @click="category_filters.splice(i, 1)"
                  ></button>
                </div>
              </div>
            </div>
            <label
              v-if="!category_image_url"
              class="transition duration-200 ease-in-out flex flex-col justify-center items-center gap-y-3 size-52 bg-gray-50 border-2 border-gray-300 border-dashed rounded-xl cursor-pointer hover:bg-gray-100"
            >
              <input
                @change="category_image_upload"
                class="hidden"
                type="file"
                accept=".jpg,.jpeg,.png"
              />
              <div
                class="size-12 bg-[url(/icons/upload.svg)] bg-no-repeat bg-cover"
              ></div>
              <h1>{{ $t("click_to_upload") }}</h1>
            </label>
            <div v-else class="flex justify-end">
              <button
                @click="
                  category_image_url = '';
                  category_image = null;
                "
                class="absolute m-2 transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/delete.svg)] size-10 bg-black/10 rounded-full hover:bg-black/25"
              ></button>
              <img
                class="object-contain size-52 border rounded-xl"
                :src="category_image_url"
              />
            </div>
            <button
              :disabled="
              is_add ? 
              !category_title || !!categories.find((category: Category) => category.title == category_title) || !category_image_url : 
              !category_title || !category_image_url || (!!categories.find((category: Category) => category.title == category_title) && JSON.stringify(category_filters) == JSON.stringify(category_filters_copy) && category_image_url.startsWith('/images'))
              "
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              {{ is_add ? $t("create") : $t("update") }}
            </button>
          </form>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="delete_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed z-10"
      >
        <div
          class="flex flex-col p-2 bg-white sm:-mt-48 w-80 sm:w-96 h-auto rounded-xl"
        >
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">{{ $t("delete_category") }}</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_category_id" required />
            <div class="text-center">{{ $t("are_you_sure") }}</div>
            <button
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              {{ $t("delete") }}
            </button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="mt-12 sm:mt-0 flex justify-center sm:justify-start items-center gap-x-4 sm:p-6 sm:h-20 text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      {{ $t("categories") }}
      <button
        @click="
          is_add = true;
          create_update_dialog = true;
        "
        class="transition duration-300 ease-in-out w-16 h-7 border border-black rounded-full hover:bg-black/10 text-sm"
      >
        {{ $t("add") }}
      </button>
    </div>
    <div
      v-if="categories.length == 0"
      class="flex flex-col justify-center items-center gap-y-8 min-w-[35rem] min-h-80 h-auto bg-white rounded-xl shadow-md"
    >
      <div
        class="size-36 bg-no-repeat bg-center bg-contain bg-[url(/icons/category.svg)] contrast-0"
      ></div>
      <div class="text-xl text-gray-500">
        {{ $t("have_not_created_category_yet") }}
      </div>
    </div>
    <div
      v-else
      class="flex p-2 sm:p-6 sm:min-w-[35rem] w-screen sm:w-auto h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <div class="w-full flex flex-wrap gap-2 sm:gap-6 items-center">
        <div
          v-for="category in categories"
          class="flex flex-col w-[calc(50vw-12px)] sm:w-56 bg-white rounded-xl border"
        >
          <div class="flex self-end gap-x-2 p-2 absolute">
            <button
              @click="open_edit_dialog(category)"
              class="transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/edit.svg)] size-10 bg-black/10 rounded-full hover:bg-black/25"
            ></button>
            <button
              @click="
                delete_dialog = true;
                delete_category_id = category.ID;
              "
              class="transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/delete.svg)] size-10 bg-black/10 rounded-full hover:bg-black/25"
            ></button>
          </div>
          <img
            class="w-[calc(50vw-14px)] h-[calc(50vw-14px)] sm:size-56 object-contain rounded-t-xl"
            :src="'/images/categories/' + category.image"
          />
          <NuxtLink :to="'/' + category.url" class="p-5 text-lg">{{
            category.title
          }}</NuxtLink>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const { t } = useI18n();
const role = useCookie<number>("role");

definePageMeta({
  middleware: "auth",
  layout: "admin",
});

interface Category {
  ID: number | null;
  title: string;
  url: string;
  filters: string[];
  image: string;
}

const is_add = ref(true);
const create_update_dialog = ref(false);
const delete_dialog = ref(false);
const delete_category_id = ref<number | null>(null);
let category_image: File | null;
const category_title = ref("");
const new_filter_input = ref("");
const category_filters = ref<string[]>([]);
const category_filters_copy = ref<string[]>([]);
const category_image_url = ref("");
const update_category_id = ref<number>();
const categories = ref<Category[]>([]);

onMounted(() => {
  nextTick(async () => {
    if (role.value === 1) {
      await useFetch(config.apiBase + "/categories", {
        headers: {
          Authorization: config.apiKey,
        },
        onResponse({ response }) {
          if (response._data) {
            categories.value = response._data;
          }
        },
      });
    }
  });
});

const open_edit_dialog = (category: Category) => {
  is_add.value = false;
  update_category_id.value = category.ID!;
  category_title.value = category.title;
  category_filters.value = [...category.filters];
  category_filters_copy.value = [...category.filters];
  category_image_url.value = "/images/categories/" + category.image;
  create_update_dialog.value = true;
};

const close_create_update_dialog = () => {
  create_update_dialog.value = false;
  category_title.value = "";
  category_filters.value = [];
  category_filters_copy.value = [];
  category_image = null;
  category_image_url.value = "";
};

const category_image_upload = (event: Event) => {
  category_image = (event.target as any).files[0];
  category_image_url.value = URL.createObjectURL(
    (event.target as any).files[0]
  );
};

const create_update = async () => {
  const form_data = new FormData();
  form_data.append("title", category_title.value);
  form_data.append(
    "url",
    category_title.value
      .toLowerCase()
      .replaceAll(" ", "-")
      .replaceAll("ç", "c")
      .replaceAll("ğ", "g")
      .replaceAll("ı", "i")
      .replaceAll("ö", "o")
      .replaceAll("ş", "s")
      .replaceAll("ü", "u")
  );
  form_data.append("filters", JSON.stringify(category_filters.value));
  if (category_image) {
    form_data.append("image", category_image!.name);
    form_data.append("file", category_image!);
  } else {
    form_data.append("image", "");
  }
  await useFetch(
    is_add.value
      ? config.apiBase + "/categories"
      : config.apiBase + "/categories/" + update_category_id.value,
    {
      headers: {
        Authorization: config.apiKey,
      },
      method: is_add.value ? "post" : "patch",
      body: form_data,
      onResponse({ response }) {
        if (response._data.title) {
          create_update_dialog.value = false;
          if (is_add.value) categories.value.push(response._data);
          else {
            categories.value[
              categories.value
                .map((category: Category) => category.ID)
                .indexOf(update_category_id.value!)
            ] = response._data;
          }
          category_title.value = "";
          toast.success(
            is_add.value ? t("category_created") : t("category_updated"),
            {
              bodyClassName: "toast-font",
            }
          );
        } else {
          toast.warning(response._data.error, {
            bodyClassName: "toast-font",
          });
        }
      },
    }
  );
};

const remove = async () => {
  await useFetch(config.apiBase + "/categories/" + delete_category_id.value, {
    headers: {
      Authorization: config.apiKey,
    },
    method: "delete",
    onResponse({ response }) {
      if (response._data == "ok") {
        categories.value.splice(
          categories.value
            .map((category: Category) => category.ID)
            .indexOf(delete_category_id.value),
          1
        );
        delete_dialog.value = false;
        delete_category_id.value = null;
        toast.success(t("category_removed"), {
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
