<template>
  <main class="flex flex-col gap-y-4 pb-8">
    <transition name="background" mode="out-in">
      <div
        v-if="category_dialog || image_gallery || delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 size-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="category_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed"
      >
        <div
          class="flex flex-col p-3 bg-white -mt-48 w-[28rem] h-auto rounded-xl z-3"
        >
          <button
            @click="category_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
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
            <button :disabled="!category_new" type="submit" :class="button">
              Create
            </button>
          </form>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="image_gallery"
        class="flex flex-col justify-center items-center gap-y-4 inset-x-0 inset-y-0 size-full fixed"
      >
        <div
          class="flex flex-col p-3 -mt-48 w-[48rem] h-[28rem] z-3 bg-no-repeat bg-cover"
          :style="'background-image: url(' + image_current + ');'"
        >
          <button
            @click="image_gallery = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-white bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-white/60"
          ></button>
        </div>
        <div class="flex gap-x-1">
          <button
            v-for="image in images"
            @click="image_current = image.url"
            :class="
              image_current == image.url
                ? gallery_button + 'bg-white'
                : gallery_button + 'bg-gray-300'
            "
          ></button>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="delete_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed"
      >
        <div
          class="flex flex-col p-2 bg-white -mt-48 w-80 h-auto rounded-xl z-3"
        >
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">Delete image</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_image_order" required />
            <div class="text-center">Are you sure?</div>
            <button type="submit" :class="button">Delete</button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center gap-x-4 p-6 text-xl bg-white rounded-xl shadow-md"
    >
      {{ is_add ? "Add product" : "Edit product" }}
    </div>
    <form @submit.prevent="create_update" :class="form_div">
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
            @click="open_gallery(image.url)"
            class="w-52 h-32 rounded-xl cursor-pointer bg-no-repeat bg-cover"
            :style="'background-image: url(' + image.url + ');'"
          >
            <div
              class="transition duration-200 ease-in-out flex flex-col p-3 w-52 h-32 rounded-xl hover:bg-black/20"
            >
              <button
                type="button"
                @click="open_delete_dialog($event, image.order)"
                class="self-end size-7 bg-white bg-[url(/icons/delete.svg)] bg-no-repeat bg-center rounded-md"
              ></button>
            </div>
          </div>
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
                class="size-12 bg-[url(/icons/upload.svg)] bg-no-repeat bg-cover"
              ></div>
              <h1>Click to upload</h1>
            </div>
          </label>
        </div>
      </div>
      <button
        :disabled="
          is_add
            ? form.title == '' ||
              form.category_id == '' ||
              !form.list_price ||
              form.stock_quantity.toString() == ''
            : JSON.stringify(form) == JSON.stringify(form_old) &&
              JSON.stringify(images) == JSON.stringify(images_old)
        "
        type="submit"
        :class="button"
      >
        {{ is_add ? "Create" : "Update" }}
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
const route = useRoute();
const router = useRouter();
definePageMeta({
  middleware: "auth",
  layout: "admin",
});
const is_add = route.params.product == "add";
const open_delete_dialog = (event: any, order: string) => {
  event.stopPropagation();
  delete_image_order.value = order;
  delete_dialog.value = true;
};
const delete_image_order = ref("");
const delete_dialog = ref(false);
const categories = ref<any>([]);
const category_dialog = ref(false);
const category_new = ref("");
const image_gallery = ref(false);
const form = ref({
  ID: "",
  title: "",
  category_id: "",
  list_price: null,
  stock_quantity: "",
});
let form_old = {
  ID: "",
  title: "",
  category_id: "",
  list_price: null,
  stock_quantity: "",
};
let images_to_upload: any = [];
const images = ref<any>([]);
let images_old: any = [];
const image_current = ref("");
const upload = (event: any) => {
  Array.prototype.slice.call(event.target.files).forEach((file: any) => {
    images.value.push({
      order: images.value.length,
      name: file.name,
      url: URL.createObjectURL(file),
    });
    images_to_upload.push(file);
  });
};
const open_gallery = (image: any) => {
  image_gallery.value = true;
  image_current.value = image;
};
const data_to_form = (response: any) => {
  form.value = response;
  const images_array: any = [];
  JSON.parse(response.images).forEach((image: any) => {
    images_array.push({
      order: image.order,
      name: image.name,
      url: "/images/products/" + image.name,
    });
  });
  images.value = images_array;
  form_old = { ...form.value };
  images_old = [...images.value];
};
onMounted(() => {
  nextTick(async () => {
    await useFetch<any>("/api/categories", {
      onResponse({ response }) {
        categories.value = response._data;
      },
    });
    if (!is_add) {
      await useFetch<any>("/api/products/" + route.params.product, {
        onResponse({ response }) {
          data_to_form(response._data);
        },
      });
    }
  });
});
const create_update = async () => {
  const form_values = <any>{
    title: form.value.title,
    url: form.value.title.toLowerCase().replaceAll(" ", "-"),
    category_id: form.value.category_id,
    list_price: form.value.list_price,
    stock_quantity: form.value.stock_quantity,
    images: JSON.stringify(images.value.map(({ url, ...rest }: any) => rest)),
  };
  const form_data = new FormData();
  for (const item in form_values) {
    form_data.append(item, form_values[item]);
  }
  for (const image in images_to_upload) {
    form_data.append("file", images_to_upload[image]);
  }
  await useFetch(is_add ? "/api/products" : "/api/products/" + form.value.ID, {
    method: is_add ? "post" : "patch",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { id: user.id },
    body: form_data,
    onResponse({ response }) {
      if (response._data.ID) {
        toast.success("Product " + (is_add ? "created" : "updated"), {
          bodyClassName: "toast-font",
        });
        router.push("/admin/products/" + response._data.url);
        images_to_upload = [];
        data_to_form(response._data);
      } else {
        toast.warning(response._data.error, {
          bodyClassName: "toast-font",
        });
      }
    },
  });
};
const remove = async () => {
  images_to_upload.splice(delete_image_order.value);
  images.value.splice(
    images.value
      .map((image: any) => image.order)
      .indexOf(delete_image_order.value),
    1
  );
  for (let i = 0; i < images.value.length; i++) {
    images.value[i].order = i;
  }
  delete_dialog.value = false;
  delete_image_order.value = "";
  toast.success("Image removed", {
    bodyClassName: "toast-font",
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
const button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none";
const label = "flex flex-col gap-y-2";
const gallery_button =
  "transition duration-300 ease-in-out size-3 rounded-full ";
</script>
