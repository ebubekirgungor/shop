<template>
  <main class="flex justify-center gap-x-4 sm:m-4">
    <div
      class="fixed overflow-auto top-0 z-10 sm:relative sm:flex sm:flex-col sm:p-6 min-w-full sm:min-w-64 h-screen sm:h-fit bg-white rounded-xl shadow-md"
      :class="{ hidden: !filters_menu }"
    >
      <div
        class="fixed flex items-center px-3 sm:hidden bg-white border w-full h-16"
      >
        <button
          @click="filters_menu = false"
          class="absolute size-7 bg-[url(/icons/close.svg)] bg-cover bg-no-repeat"
        ></button>
        <div class="w-full text-lg text-center">Filters</div>
      </div>
      <div class="mt-16 sm:mt-0 flex flex-col p-3 sm:p-0 divide-y">
        <div v-for="[key, values] in filters" class="flex flex-col">
          <input type="checkbox" :id="key" class="peer hidden" />
          <label
            :for="key"
            class="transition duration-200 ease-in-out flex justify-between items-center px-2 py-4 cursor-pointer hover:bg-gray-100 bg-no-repeat bg-right bg-[url(/icons/expand_more.svg)] peer-checked:bg-[url(/icons/expand_less.svg)]"
            >{{ key }}
          </label>
          <div
            class="transition-all duration-300 ease-in-out flex flex-col gap-y-2 bg-white max-h-0 overflow-hidden peer-checked:max-h-72 peer-checked:overflow-auto"
          >
            <div class="flex flex-col gap-y-2 py-2">
              <div v-for="value in values" class="flex flex-col">
                <div
                  class="transition duration-200 ease-in-out flex items-center gap-x-2 pl-1.5 rounded-lg cursor-pointer hover:bg-gray-100"
                >
                  <input
                    type="checkbox"
                    :id="value.filter"
                    v-model="value.selected"
                    @change="filters_menu ? null : filter_product()"
                    class="transition duration-200 ease-in-out size-[1.1rem] cursor-pointer rounded-md border-gray-300 text-gray-800 hover:border-gray-400 focus:ring-0 focus:ring-offset-0"
                  /><label
                    :for="value.filter"
                    class="flex items-center w-full h-7 cursor-pointer"
                    >{{ value.filter }}</label
                  >
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div
        class="fixed bottom-0 flex items-center p-3 sm:hidden bg-white border-t w-full"
      >
        <button
          @click="
            filter_product();
            filters_menu = false;
          "
          class="w-full h-12 text-white border bg-black rounded-full"
        >
          Show results
        </button>
      </div>
    </div>
    <div class="flex flex-col gap-y-4 w-[clamp(32rem,65rem,65rem)]">
      <div
        class="flex justify-center sm:justify-start sm:p-6 text-xl h-auto bg-white rounded-xl sm:shadow-md"
      >
        {{ category_title }}
      </div>
      <div class="sm:hidden flex border-y divide-x text-md">
        <button
          @click="filters_menu = true"
          class="flex justify-center items-center gap-x-1 w-full p-2"
        >
          <div class="size-6 bg-[url(/icons/filter.svg)]"></div>
          Filters
        </button>
        <button
          @click="filters_menu = true"
          class="flex justify-center items-center gap-x-1 w-full p-2"
        >
          <div class="size-6 bg-[url(/icons/sort.svg)]"></div>
          Sort
        </button>
      </div>
      <div class="flex m-2 sm:m-0 h-fit bg-white sm:rounded-xl sm:shadow-md">
        <div
          class="w-full sm:min-h-[27rem] flex flex-wrap gap-2 sm:gap-6 sm:p-6"
        >
          <div
            v-for="product in products"
            class="flex flex-col pb-2 sm:pb-0 w-[calc(50vw-12px)] sm:w-[14.25rem] bg-white rounded-xl border"
          >
            <button
              v-if="role != undefined"
              @click="toggle_favorite(product)"
              class="self-end absolute m-2 transition duration-200 ease-in-out flex justify-center items-center size-10 border bg-white rounded-full sm:hover:bg-gray-200"
            >
              <div
                class="bg-no-repeat bg-center size-6 bg-contain"
                :class="
                  product.is_favorite
                    ? 'bg-[url(/icons/favorite_filled.svg)]'
                    : 'bg-[url(/icons/favorite.svg)] contrast-0'
                "
              ></div>
            </button>
            <NuxtLink :to="'/product/' + product.url">
              <img
                class="min-w-[calc(50vw-14px)] sm:min-w-full min-h-[calc(50vw-14px)] h-[calc(50vw-14px)] sm:size-[14.25rem] sm:min-h-[14.25rem] object-contain rounded-t-xl"
                :src="'/images/products/' + product.image"
              />
              <div class="flex flex-col gap-y-6 pt-4 px-4">
                <div class="text-md">{{ product.title }}</div>
                <div class="text-2xl">
                  {{ product.list_price.toLocaleString("tr-TR") }}
                  TL
                </div>
              </div>
            </NuxtLink>
            <div class="hidden sm:block p-4">
              <button
                @click="add_to_cart(product.id)"
                :disabled="product.stock_quantity == 0 || cart.find((item: Cart) => item.id == product.id)?.quantity! >= product.stock_quantity"
                class="transition duration-300 ease-in-out w-full h-10 font-medium border-2 rounded-full hover:bg-gray-200 disabled:text-gray-400 disabled:bg-gray-100 disabled:border-0 disabled:pointer-events-none"
              >
                {{
                  product.stock_quantity == 0 ? "Out of stock" : "Add to cart"
                }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
const config = useRuntimeConfig().public;
const route = useRoute();
const role = useCookie<number>("role");
const cart_unregistered = useCookie<Cart[]>("cart");
if (!cart_unregistered.value) cart_unregistered.value = [];
interface Cart {
  id: number;
  quantity: number;
  selected: boolean;
}
interface Product {
  id: number;
  title: string;
  url: string;
  image: string;
  list_price: number;
  stock_quantity: number;
  is_favorite: boolean;
  filters: ProductFilter[];
}
interface ProductFilter {
  name: string;
  value: string;
}
const filters_menu = ref(false);
const filters = ref(new Map<string, { filter: string; selected: boolean }[]>());
const category_title = ref("");
const products = ref<Product[]>([]);
const cart = ref<Cart[]>([]);
const favorites_ids = ref<number[]>([]);
onMounted(() => {
  nextTick(async () => {
    if (role.value != undefined) {
      await useFetch(config.apiBase + "/favorites/ids", {
        headers: {
          Authorization: config.apiKey,
        },
        onResponse({ response }) {
          favorites_ids.value = response._data;
        },
      });
    }
    await useFetch(config.apiBase + "/categories/" + route.params.category, {
      onResponse({ response }) {
        category_title.value = response._data.category_title;
        products.value = response._data.products;
        products.value.map((product: Product) => {
          product.filters.map((filter) => {
            const existing = filters.value.get(filter.name) || [];
            if (!existing.find((f) => f.filter == filter.value)) {
              existing[existing.length] = {
                filter: filter.value,
                selected: false,
              };
              filters.value.set(filter.name, existing);
            }
          });
          if (role.value != undefined)
            product.is_favorite = favorites_ids.value.includes(product.id);
        });
      },
    });
    if (role.value != undefined) {
      await useFetch(config.apiBase + "/users", {
        headers: {
          Authorization: config.apiKey,
        },
        onResponse({ response }) {
          cart.value = response._data.cart;
        },
      });
    } else {
      cart.value = cart_unregistered.value;
    }
  });
});
const filter_product = async () => {
  const selected_filters: { [index: string]: string } = {};
  let selected_count = 0;
  filters.value.forEach((filter, key) => {
    const selected = filter.filter((f) => f.selected == true);
    selected_count += selected.length;
    selected_filters[key] = selected.map((f) => f.filter).join(",");
  });
  useFetch(config.apiBase + "/categories/" + route.params.category, {
    query: selected_count != 0 ? selected_filters : undefined,
    onResponse({ response }) {
      products.value = response._data.products;
      if (role.value != undefined) {
        products.value.map((product: Product) => {
          product.is_favorite = favorites_ids.value.includes(product.id);
        });
      }
    },
  });
};
const add_to_cart = async (id: number) => {
  if (role.value != undefined) {
    let item: Cart = cart.value.find((item: Cart) => item.id == id)!;
    if (item) item.quantity += 1;
    else cart.value.push({ id: id, quantity: 1, selected: true });
    await useFetch(config.apiBase + "/users/cart/" + id, {
      headers: {
        Authorization: config.apiKey,
      },
      method: "post",
      onResponse() {},
    });
  } else {
    for (let i = 0; i < cart_unregistered.value.length; i++) {
      if (cart_unregistered.value[i].id == id) {
        cart_unregistered.value[i].quantity++;
        return true;
      }
    }
    cart_unregistered.value.push({
      id: id,
      quantity: 1,
      selected: true,
    });
  }
};
const toggle_favorite = async (product: Product) => {
  await useFetch(config.apiBase + "/favorites/" + product.id, {
    headers: {
      Authorization: config.apiKey,
    },
    method: product.is_favorite ? "delete" : "post",
    onResponse({ response }) {
      if (response._data == "ok") {
        product.is_favorite = !product.is_favorite;
      }
    },
  });
};
</script>
