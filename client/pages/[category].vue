<template>
  <main class="flex justify-center m-4">
    <div
      class="flex w-[clamp(32rem,65rem,65rem)] h-auto bg-white rounded-xl shadow-md"
    >
      <div
        class="w-full min-h-[27rem] grid grid-cols-auto_box gap-6 p-6 items-center"
      >
        <div
          v-for="product in products"
          class="flex flex-col w-[14.25rem] bg-white rounded-xl border"
        >
          <button
            v-if="role != undefined"
            @click="toggle_favorite(product)"
            class="self-end absolute m-2 transition duration-200 ease-in-out flex justify-center items-center size-10 border bg-white rounded-full hover:bg-gray-200"
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
              class="size-[14.25rem] min-h-[14.25rem] object-contain rounded-t-xl"
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
          <div class="p-4">
            <button
              @click="add_to_cart(product.id)"
              :disabled="product.stock_quantity == 0 || cart.find((item: Cart) => item.id == product.id)?.quantity! >= product.stock_quantity"
              class="transition duration-300 ease-in-out w-full h-10 font-medium border-2 rounded-full hover:bg-gray-200 disabled:text-gray-400 disabled:bg-gray-100 disabled:border-0 disabled:pointer-events-none"
            >
              {{ product.stock_quantity == 0 ? "Out of stock" : "Add to cart" }}
            </button>
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
}
const products = ref<Product[]>([]);
const cart = ref<Cart[]>([]);
onMounted(() => {
  nextTick(async () => {
    await useFetch(config.apiBase + "/categories/" + route.params.category, {
      onResponse({ response }) {
        products.value = response._data;
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
      onResponse({ response }) {},
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
