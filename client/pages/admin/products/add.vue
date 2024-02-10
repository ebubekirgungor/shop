<template>
  <main class="flex flex-col gap-y-4 pb-8">
    <transition name="background" mode="out-in">
      <div
        v-if="category_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 w-full h-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="category_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 w-full h-full fixed"
      >
        <div
          class="flex flex-col p-3 bg-white -mt-48 w-[28rem] h-auto rounded-xl z-3"
        >
          <button
            @click="category_dialog = false"
            class="self-end transition duration-300 ease-in-out w-8 h-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl grow">Create new category</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="create_category"
          >
            <input
              :class="input"
              type="text"
              placeholder="Title"
              v-model="category_new"
              required
            />
            <button
              :disabled="!category_new"
              type="submit"
              :class="create_button"
            >
              Create
            </button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center gap-x-4 p-6 text-xl bg-white rounded-xl shadow-md"
    >
      Add Product
    </div>
    <form @submit.prevent="create" :class="form_div">
      <label :class="label"
        >Title<input :class="input" type="text" v-model="form.title"
      /></label>
      <label :class="label"
        >Category
        <div class="flex gap-x-4">
          <select
            :class="input + ' w-full cursor-pointer'"
            v-model="form.category_id"
          >
            <option value="" selected>Select category</option>
            <option v-for="c in categories" :value="c.ID">{{ c.title }}</option>
          </select>
          <button
            type="button"
            @click="category_dialog = true"
            :class="add_category_button"
          >
            Add New
          </button>
        </div>
      </label>
      <label :class="label"
        >List Price<input
          :class="input"
          step="any"
          type="number"
          min="1"
          v-model="form.list_price"
      /></label>
      <label :class="label"
        >Stock Quantity<input
          :class="input"
          type="number"
          min="0"
          v-model="form.stock_quantity"
      /></label>
      <div :class="label + ' col-span-2'">
        <label>Images</label>
        <div class="grid grid-cols-auto gap-4">
          <div
            v-for="image in images"
            class="transition duration-200 ease-in-out w-52 h-32 rounded-xl cursor-pointer bg-no-repeat bg-cover hover:brightness-90"
            :style="'background-image: url(' + image + ');'"
          ></div>
          <label class="w-52">
            <div
              class="transition duration-200 ease-in-out flex flex-col justify-center items-center gap-y-3 w-52 h-32 bg-gray-50 border-2 border-gray-300 border-dashed rounded-xl cursor-pointer hover:bg-gray-100"
            >
              <input
                @change="upload"
                class="hidden"
                type="file"
                accept=".jpg,.jpeg,.png"
                multiple
              />
              <div
                class="w-12 h-12 bg-[url(/icons/upload.svg)] bg-no-repeat bg-cover"
              ></div>
              <h1>Click to upload</h1>
            </div>
          </label>
        </div>
      </div>
      <button
        :disabled="
          form.title == '' ||
          form.category_id == '' ||
          !form.list_price ||
          !form.stock_quantity
        "
        type="submit"
        :class="create_button"
      >
        Create
      </button>
    </form>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useUser } from "@/store/user";
import { useToast } from "vue-toastification";
const { user } = useUser();
const toast = useToast();
definePageMeta({
  middleware: "auth",
  layout: "admin",
});
const categories = ref<any>([]);
const category_dialog = ref(false);
const category_new = ref("");
const form = ref({
  title: "",
  category_id: "",
  list_price: null,
  stock_quantity: null,
});
let form_data = new FormData();
const images = ref<any>([]);
const image_filenames = ref<any>([]);
const upload = (event: any) => {
  Array.prototype.slice.call(event.target.files).forEach((file: any) => {
    images.value.push(URL.createObjectURL(file));
    form_data.append("file", file);
    image_filenames.value.push({
      order: image_filenames.value.length + 1,
      name: file.name,
    });
  });
};
onMounted(() => {
  nextTick(async () => {
    await useFetch<any>("/api/categories", {
      onResponse({ response }) {
        categories.value = response._data;
      },
    });
  });
});
const create = async () => {
  const form_values = <any>{
    title: form.value.title,
    url: form.value.title.toLowerCase().replaceAll(" ", "-"),
    category_id: form.value.category_id,
    list_price: form.value.list_price,
    stock_quantity: form.value.stock_quantity,
    images: JSON.stringify(image_filenames.value),
  };
  for (const item in form_values) {
    form_data.append(item, form_values[item]);
  }
  await useFetch("/api/products", {
    method: "post",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { id: user.id },
    body: form_data,
    onResponse({ response }) {
      if (response._data.ID) {
        toast.success("Product created", {
          bodyClassName: "toast-font",
        });
        form_data = new FormData();
      } else {
        toast.warning(response._data.error, {
          bodyClassName: "toast-font",
        });
      }
    },
  });
};
const create_category = async () => {
  await useFetch("/api/categories", {
    method: "post",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { id: user.id },
    body: {
      title: category_new.value,
    },
    onResponse({ response }) {
      if (response._data.title) {
        category_dialog.value = false;
        categories.value.push(response._data);
        toast.success("Category created", {
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
const form_div =
  "grid grid-cols-2 gap-x-[7%] gap-y-8 items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md";
const input =
  "transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const add_category_button =
  "transition duration-300 ease-in-out w-32 rounded-md bg-black/5 text-sm hover:bg-black/10";
const create_button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none";
const label = "flex flex-col gap-y-2";
</script>
