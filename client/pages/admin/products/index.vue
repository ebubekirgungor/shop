<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <transition name="background" mode="out-in">
      <div
        v-if="delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 size-full fixed"
      ></div>
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
          <h1 class="text-center text-xl">Delete product</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_product_id" required />
            <div class="text-center">Are you sure?</div>
            <button type="submit" :class="button">Delete</button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center gap-x-4 p-6 text-xl bg-white rounded-xl shadow-md"
    >
      Products
      <NuxtLink
        to="/admin/products/add"
        class="flex justify-center items-center transition duration-300 ease-in-out w-16 h-7 border border-black rounded-full hover:bg-black/10 text-sm"
      >
        Add
      </NuxtLink>
    </div>
    <div
      class="flex flex-col justify-center p-5 min-w-[40rem] h-full min-h-[14rem] bg-white rounded-xl shadow-md"
    >
      <div v-if="fetch_complete" class="flex flex-col gap-y-4">
        <input
          class="transition duration-300 ease-in-out bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:98%_60%] w-64 rounded-md border-0 bg-gray-100 text-sm focus:ring-2 focus:ring-slate-300"
          v-model="filter"
          type="text"
          placeholder="Search"
        />
        <table
          class="w-full table-fixed text-left overflow-hidden bg-gray-100 shadow-lg rounded-md"
        >
          <thead>
            <tr class="bg-gray-200 whitespace-nowrap">
              <th class="w-1/3 pl-4">
                <div @click="sort('title')" :class="th_button">
                  Title
                  <div
                    :class="
                      arrow_active != 'title'
                        ? th_icon + 'contrast-0'
                        : sort_direction == -1
                        ? th_icon + 'rotate-180'
                        : th_icon
                    "
                  ></div>
                </div>
              </th>
              <th class="w-1/4">
                <div @click="sort('category')" :class="th_button">
                  Category
                  <div
                    :class="
                      arrow_active != 'category'
                        ? th_icon + 'contrast-0'
                        : sort_direction == -1
                        ? th_icon + 'rotate-180'
                        : th_icon
                    "
                  ></div>
                </div>
              </th>
              <th class="w-1/4">
                <div @click="sort('list_price')" :class="th_button">
                  List Price
                  <div
                    :class="
                      arrow_active != 'list_price'
                        ? th_icon + 'contrast-0'
                        : sort_direction == -1
                        ? th_icon + 'rotate-180'
                        : th_icon
                    "
                  ></div>
                </div>
              </th>
              <th class="w-1/4">
                <div @click="sort('stock_quantity')" :class="th_button">
                  Stock Quantity
                  <div
                    :class="
                      arrow_active != 'stock_quantity'
                        ? th_icon + 'contrast-0'
                        : sort_direction == -1
                        ? th_icon + 'rotate-180'
                        : th_icon
                    "
                  ></div>
                </div>
              </th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="product in displayed_products"
              class="border-y odd:bg-gray-50"
            >
              <td class="p-4">
                <NuxtLink :to="'/' + product.url">{{ product.title }}</NuxtLink>
              </td>
              <td>{{ product.category }}</td>
              <td>{{ product.list_price }}</td>
              <td>{{ product.stock_quantity }}</td>
              <td class="flex gap-x-1 pr-2 py-2 float-right">
                <NuxtLink
                  :to="'/admin/products/' + product.url"
                  :class="icon + 'bg-[url(/icons/edit.svg)]'"
                ></NuxtLink>
                <button
                  @click="open_delete_dialog(product.id)"
                  :class="icon + 'bg-[url(/icons/delete.svg)]'"
                ></button>
              </td>
            </tr>
            <tr
              v-if="displayed_products.length == 0"
              class="bg-gray-50 text-center"
            >
              <td colspan="5" class="p-4 border-y">Product not found</td>
            </tr>
          </tbody>
          <tfoot>
            <td colspan="5">
              <div class="flex justify-end items-center gap-x-2 p-2">
                <select
                  class="transition duration-300 ease-in-out rounded-md border-0 bg-gray-200 text-sm focus:ring-2 focus:ring-slate-300 cursor-pointer"
                  v-model="items_per_page"
                >
                  <option
                    v-for="option in items_per_page_options"
                    :key="option"
                    :value="option"
                  >
                    {{ option }}
                  </option>
                </select>
                <div class="mr-4">items per page</div>
                <button
                  @click="current_page = 1"
                  :disabled="current_page == 1"
                  :class="icon + 'bg-[url(/icons/first.svg)]'"
                ></button>
                <button
                  @click="current_page > 1 ? current_page-- : ''"
                  :disabled="current_page == 1"
                  :class="icon + 'bg-[url(/icons/previous.svg)]'"
                ></button>
                <button
                  @click="current_page < total_pages ? current_page++ : ''"
                  :disabled="current_page == total_pages || total_pages == 0"
                  :class="icon + 'bg-[url(/icons/next.svg)]'"
                ></button>
                <button
                  @click="current_page = total_pages"
                  :disabled="current_page == total_pages || total_pages == 0"
                  :class="icon + 'bg-[url(/icons/last.svg)]'"
                ></button>
              </div>
            </td>
          </tfoot>
        </table>
      </div>
      <div v-else class="flex justify-center">
        <div
          class="animate-spin size-24 bg-[url(/icons/loading.svg)] bg-no-repeat bg-cover"
        ></div>
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
interface Product {
  id: number | null;
  title: string;
  url: string;
  category: string;
  list_price: number | null;
  stock_quantity: string;
}
const fetch_complete = ref(false);
const open_delete_dialog = (id: Product["id"]) => {
  delete_product_id.value = id;
  delete_dialog.value = true;
};
const delete_product_id = ref<Product["id"]>(null);
const delete_dialog = ref(false);
const products = ref<Product[]>([]);
const filter = ref("");
const current_page = ref(1);
const items_per_page = ref(5);
const items_per_page_options = [5, 10, 25, 50, 100];
const sort_by_field = ref("");
const sort_direction = ref(0);
const arrow_active = ref("");
onMounted(() => {
  nextTick(async () => {
    await useFetch<Product[]>(config.apiBase + "/products", {
      onResponse({ response }) {
        if (response._data) {
          products.value = response._data;
          fetch_complete.value = true;
        }
      },
    });
  });
});
const filtered_products = computed(() => {
  return products.value.filter((product: Product) =>
    product.title.toLowerCase().includes(filter.value.toLowerCase())
  );
});
const total_pages = computed(() =>
  Math.ceil(filtered_products.value.length / items_per_page.value)
);
const displayed_products = computed(() => {
  const start_index = (current_page.value - 1) * items_per_page.value;
  let sorted_products = [...filtered_products.value];
  if (sort_by_field.value) {
    sorted_products.sort((a, b) => {
      const fieldA = a[sort_by_field.value as keyof Product];
      const fieldB = b[sort_by_field.value as keyof Product];
      if (typeof fieldA === "number" && typeof fieldB === "number") {
        return (fieldA - fieldB) * sort_direction.value;
      } else {
        const strA = String(fieldA).toLowerCase();
        const strB = String(fieldB).toLowerCase();
        return strA.localeCompare(strB) * sort_direction.value;
      }
    });
  }
  return sorted_products.slice(start_index, start_index + items_per_page.value);
});
const sort = async (field: string) => {
  arrow_active.value = field;
  if (current_page.value != 1) current_page.value = 1;
  if (sort_by_field.value === field) {
    sort_direction.value *= -1;
  } else {
    sort_by_field.value = field;
    sort_direction.value = 1;
  }
};
const remove = async () => {
  await useFetch(config.apiBase + "/products/" + delete_product_id.value, {
    headers: {
      Authorization: config.apiKey,
    },
    method: "delete",
    onResponse({ response }) {
      if (response._data) {
        products.value.splice(
          products.value
            .map((product: Product) => product.id)
            .indexOf(delete_product_id.value),
          1
        );
        delete_dialog.value = false;
        delete_product_id.value = null;
        toast.success("Product removed", {
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
const button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none";
const icon =
  "transition duration-300 ease-in-out size-9 bg-no-repeat bg-center rounded-full hover:bg-black/10 disabled:pointer-events-none disabled:contrast-0 ";
const th_button = "flex gap-x-1 items-center h-12 cursor-pointer group";
const th_icon =
  "transition duration-200 ease-in-out size-6 bg-no-repeat bg-center bg-[url(/icons/arrow.svg)] opacity-0 group-hover:opacity-100 ";
</script>
<style>
td {
  user-select: text;
}
</style>
