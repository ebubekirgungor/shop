<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <transition name="background" mode="out-in">
      <div
        v-if="delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 size-full fixed z-[2]"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="delete_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed z-[3]"
      >
        <div
          class="flex flex-col p-2 bg-white sm:-mt-48 w-80 h-auto rounded-xl"
        >
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">{{ $t("delete_image") }}</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_image_order" required />
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
      class="mt-12 sm:mt-0 flex justify-center sm:justify-start items-center gap-x-4 sm:p-6 sm:h-20 text-xl bg-white sm:rounded-xl sm:shadow-md"
    >
      <button
        @click="navigateTo(localePath('admin-products'))"
        class="transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/previous.svg)] ml-4 sm:ml-0 size-10 bg-black/5 rounded-full hover:bg-black/10"
      ></button>
      {{ is_add ? $t("add_product") : $t("edit_product") }}
    </div>
    <form
      @submit.prevent="create_update"
      class="flex flex-col sm:grid grid-cols-2 sm:gap-x-[7%] gap-y-4 sm:gap-y-8 items-center p-4 sm:p-6 w-screen sm:w-auto sm:min-w-[30rem] h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <label class="w-full flex flex-col gap-y-2"
        >{{ $t("title")
        }}<input
          class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
          type="text"
          v-model="form.title"
      /></label>
      <label class="w-full flex flex-col gap-y-2"
        >{{ $t("list_price")
        }}<input
          class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
          step=".01"
          type="number"
          min="1"
          v-model="form.list_price"
      /></label>
      <label class="w-full flex flex-col gap-y-2"
        >{{ $t("stock_quantity")
        }}<input
          class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
          type="number"
          min="0"
          v-model="form.stock_quantity"
      /></label>
      <label class="w-full flex flex-col gap-y-2"
        >{{ $t("category") }}
        <div class="flex gap-x-4">
          <select
            class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300 w-full cursor-pointer"
            v-model="form.category_id"
            @change="category_change"
          >
            <option value="" selected>{{ $t("select_category") }}</option>
            <option v-for="c in categories" :value="c.ID">{{ c.title }}</option>
          </select>
          <button
            type="button"
            @click="navigateTo(localePath('admin-categories'))"
            class="transition duration-300 ease-in-out w-32 rounded-md bg-black/5 text-sm hover:bg-black/10"
          >
            {{ $t("add_new") }}
          </button>
        </div>
      </label>
      <label
        v-for="filter in form.filters"
        class="w-full flex flex-col gap-y-2"
      >
        {{ filter.name }}
        <input
          class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
          type="text"
          v-model="filter.value"
      /></label>
      <div
        class="w-full flex flex-col gap-y-2 w-screen sm:w-auto p-2 sm:p-0 col-span-2"
      >
        <label class="px-2 sm:px-0">{{ $t("images") }}</label>
        <div class="flex flex-wrap gap-2 sm:gap-4 mt-2">
          <div
            v-for="image in images"
            class="flex flex-col w-[calc(50vw-12px)] h-[calc(50vw-12px)] sm:size-44 border rounded-xl cursor-pointer"
          >
            <button
              type="button"
              @click="open_delete_dialog($event, image.order)"
              class="absolute m-2 self-end transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/delete.svg)] size-10 bg-gray-200 rounded-full sm:hover:bg-gray-300"
            ></button>
            <img
              class="w-[calc(50vw-14px)] h-[calc(50vw-14px)] sm:size-full object-center object-contain rounded-xl"
              :src="image.url"
            />
          </div>
          <label class="sm:w-44">
            <div
              class="transition duration-200 ease-in-out flex flex-col justify-center items-center gap-y-3 w-[calc(50vw-12px)] h-[calc(50vw-12px)] sm:size-44 bg-gray-50 border-2 border-gray-300 border-dashed rounded-xl cursor-pointer hover:bg-gray-100"
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
              <h1>{{ $t("click_to_upload") }}</h1>
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
              form.stock_quantity.toString() == '' ||
              form.filters.some((filter: Filter) => filter.value === '')
            : form.title == '' ||
              form.category_id == '' ||
              !form.list_price ||
              form.stock_quantity.toString() == '' ||
              form.filters.some((filter: Filter) => filter.value === '') ||
              JSON.stringify(form) == JSON.stringify(form_old) &&
              JSON.stringify(images) == JSON.stringify(images_old)
        "
        type="submit"
        class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
      >
        {{ is_add ? $t("create") : $t("update") }}
      </button>
    </form>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const localePath = useLocalePath();
const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const role = useCookie<number>("role");

definePageMeta({
  middleware: "auth",
  layout: "admin",
});

interface Category {
  ID: number;
  title: string;
  filters: string[];
}

interface Image {
  order: number | null;
  name: string;
  url: string;
}

interface Filter {
  name: string;
  value: string;
}

interface Product {
  ID?: number | null;
  title: string;
  category_id: string;
  list_price: number | null;
  stock_quantity: string;
  filters: Filter[];
  images?: Image[];
}

const is_add = route.params.product == t("add_url");

const open_delete_dialog = (event: Event, order: Image["order"]) => {
  event.stopPropagation();
  delete_image_order.value = order;
  delete_dialog.value = true;
};

const delete_image_order = ref<Product["ID"]>(null);
const delete_dialog = ref(false);
const categories = ref<Category[]>([]);

const form = ref<Product>({
  ID: null,
  title: "",
  category_id: "",
  list_price: null,
  stock_quantity: "",
  filters: [],
});
let form_old = <Product>{
  ID: null,
  title: "",
  category_id: "",
  list_price: null,
  stock_quantity: "",
  filters: [],
};

let images_to_upload: File[] = [];
const images = ref<Image[]>([]);
let images_old: Image[] = [];
const image_current = ref("");

interface EventTarget {
  files: unknown;
}

const upload = (event: Event) => {
  Array.prototype.slice
    .call((event.target as unknown as EventTarget).files)
    .forEach((file: File) => {
      images.value.push({
        order: images.value.length,
        name: file.name,
        url: URL.createObjectURL(file),
      });
      images_to_upload.push(file);
    });
};

const data_to_form = (response: Product) => {
  form.value = response;
  const images_array: Image[] = [];
  (response.images as Image[]).forEach((image: Image) => {
    images_array.push({
      order: image.order,
      name: image.name,
      url: "/images/products/" + image.name,
    });
  });
  images.value = images_array;
  form_old = { ...form.value };
  form_old.filters = JSON.parse(JSON.stringify(response.filters));
  images_old = [...images.value];
};

onMounted(() => {
  nextTick(async () => {
    if (role.value === 1) {
      await useFetch(config.apiBase + "/categories", {
        onResponse({ response }) {
          categories.value = response._data;
        },
      });
      if (!is_add) {
        await useFetch(
          config.apiBase +
            "/products/" +
            route.params.product.toString().match(/-(\d+)$/)![1],
          {
            onResponse({ response }) {
              data_to_form(response._data);
              if (form.value.filters.length == 0) category_change();
            },
          }
        );
      }
    }
  });
});

const category_change = async () => {
  const filters_temp: Filter[] = [];
  categories.value
    .find((category) => category.ID == parseInt(form.value.category_id))
    ?.filters.map((filter) =>
      filters_temp.push({
        name: filter,
        value: "",
      })
    );
  form.value.filters = filters_temp;
};

const create_update = async () => {
  const form_values = {
    title: form.value.title,
    url: form.value.title
      .toLowerCase()
      .replaceAll(" ", "-")
      .replaceAll("ç", "c")
      .replaceAll("ğ", "g")
      .replaceAll("ı", "i")
      .replaceAll("ö", "o")
      .replaceAll("ş", "s")
      .replaceAll("ü", "u"),
    category_id: form.value.category_id,
    list_price: form.value.list_price,
    stock_quantity: form.value.stock_quantity,
    filters: JSON.stringify(form.value.filters),
    images: JSON.stringify(images.value.map(({ url, ...rest }: Image) => rest)),
  };
  const form_data = new FormData();
  for (const item in form_values) {
    form_data.append(
      item,
      (form_values as unknown as { [key: string]: string })[item]
    );
  }
  for (const image in images_to_upload) {
    form_data.append("file", images_to_upload[image]);
  }
  await useFetch(
    is_add
      ? config.apiBase + "/products"
      : config.apiBase + "/products/" + form.value.ID,
    {
      headers: {
        Authorization: config.apiKey,
      },
      method: is_add ? "post" : "patch",
      body: form_data,
      onResponse({ response }) {
        if (response._data.ID) {
          toast.success(is_add ? t("product_created") : t("product_updated"), {
            bodyClassName: "toast-font",
          });
          is_add
            ? router.push(
                localePath("admin-products") + "/" + response._data.url
              )
            : history.replaceState(
                {},
                "",
                localePath("admin-products") + "/" + response._data.url
              );
          images_to_upload = [];
          data_to_form(response._data);
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
  images_to_upload.splice(delete_image_order.value as number);
  images.value.splice(
    images.value
      .map((image: Image) => image.order)
      .indexOf(delete_image_order.value as number),
    1
  );
  for (let i = 0; i < images.value.length; i++) {
    images.value[i].order = i;
  }
  delete_dialog.value = false;
  delete_image_order.value = null;
  toast.success(t("image_removed"), {
    bodyClassName: "toast-font",
  });
};
</script>
