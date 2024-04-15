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
          class="flex flex-col p-3 bg-white -mt-48 w-[28rem] h-auto rounded-xl"
        >
          <button
            @click="close_create_update_dialog"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl grow">
            {{ is_add ? "Create new category" : "Edit category" }}
          </h1>
          <form
            class="flex flex-col items-center gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="create_update"
          >
            <input
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              placeholder="Title"
              v-model="category_title"
              required
            />
            <label
              v-if="category_image_url == ''"
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
              <h1>Click to upload</h1>
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
                class="size-52 border rounded-xl"
                :src="category_image_url"
              />
            </div>
            <button
              :disabled="
              is_add ? 
              !category_title || !!categories.find((category: Category) => category.title == category_title) || !category_image_url : 
              !category_title || !category_image_url || (!!categories.find((category: Category) => category.title == category_title) && category_image_url.startsWith('/images'))
              "
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              {{ is_add ? "Create" : "Update" }}
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
        <div class="flex flex-col p-2 bg-white -mt-48 w-96 h-auto rounded-xl">
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">Delete category</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_category_id" required />
            <div class="text-center">Are you sure?</div>
            <button
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              Delete
            </button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center gap-x-4 p-6 text-xl h-auto bg-white rounded-xl shadow-md"
    >
      Categories
      <button
        @click="
          is_add = true;
          create_update_dialog = true;
        "
        class="transition duration-300 ease-in-out w-16 h-7 border border-black rounded-full hover:bg-black/10 text-sm"
      >
        Add
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
        You have not created a category yet
      </div>
    </div>
    <div
      v-else
      class="flex p-6 min-w-[35rem] h-auto bg-white rounded-xl shadow-md"
    >
      <div
        class="w-full grid grid-cols-4 grid-cols-auto_box gap-6 items-center"
      >
        <div
          v-for="category in categories"
          class="flex flex-col w-56 bg-white rounded-xl border"
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
            class="size-56 object-contain rounded-t-xl"
            :src="'/images/categories/' + category.image"
          />
          <div class="p-5 text-lg">{{ category.title }}</div>
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
definePageMeta({
  middleware: "auth",
  layout: "admin",
});
interface Category {
  ID: number | null;
  title: string;
  image: string;
}
const is_add = ref(true);
const create_update_dialog = ref(false);
const delete_dialog = ref(false);
const delete_category_id = ref<number | null>(null);
let category_image: File | null;
const category_title = ref("");
const category_image_url = ref("");
const update_category_id = ref<number>();
const categories = ref<Category[]>([]);
onMounted(() => {
  nextTick(async () => {
    await useFetch<Category[]>(config.apiBase + "/categories", {
      headers: {
        Authorization: config.apiKey,
      },
      onResponse({ response }) {
        if (response._data) {
          categories.value = response._data;
        }
      },
    });
  });
});
const open_edit_dialog = (category: Category) => {
  is_add.value = false;
  update_category_id.value = category.ID!;
  category_title.value = category.title;
  category_image_url.value = "/images/categories/" + category.image;
  create_update_dialog.value = true;
};
const close_create_update_dialog = () => {
  create_update_dialog.value = false;
  category_title.value = "";
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
            is_add.value ? "Category created" : "Category updated",
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
        toast.success("Category removed", {
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
