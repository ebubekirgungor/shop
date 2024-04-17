<template>
  <main class="flex justify-center m-4">
    <div
      class="flex w-[clamp(32rem,65rem,65rem)] min-h-[24rem] h-auto bg-white rounded-xl shadow-md"
    >
      <div class="w-full grid grid-cols-auto_box gap-6 p-6 items-center">
        <NuxtLink
          v-for="product in products"
          :to="'/product/' + product.url"
          class="flex flex-col w-[14.25rem] h-[21rem] bg-white rounded-xl border"
        >
          <img
            class="size-[14.25rem] object-contain rounded-t-xl"
            :src="'/images/products/' + product.image"
          />
          <div class="flex flex-col gap-y-6 p-4">
            <div class="text-md">{{ product.title }}</div>
            <div class="text-2xl">
              {{ product.list_price.toLocaleString("tr-TR") }}
              TL
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const route = useRoute();
const router = useRouter();
const role = useCookie<number>("role");
interface Image {
  order: number;
  name: string;
  url: string;
}
interface Product {
  title: string;
  url: string;
  image: string;
  list_price: number;
  stock_quantity: string;
}
const products = ref<Product[]>([]);
onMounted(() => {
  nextTick(async () => {
    await useFetch(config.apiBase + "/categories/" + route.params.category, {
      onResponse({ response }) {
        products.value = response._data;
      },
    });
  });
});
/*const toggle_favorite = async () => {
  await useFetch(config.apiBase + "/favorites/" + product.value.ID, {
    headers: {
      Authorization: config.apiKey,
    },
    method: product.value.is_favorite ? "delete" : "post",
    onResponse({ response }) {
      if (response._data == "ok") {
        product.value.is_favorite = !product.value.is_favorite;
        if (product.value.is_favorite) {
          animate.value = true;
          setTimeout(() => {
            animate.value = false;
          }, 500);
        }
      }
    },
  });
};*/
const icon =
  "transition duration-300 ease-in-out bg-no-repeat bg-center size-10 bg-white rounded-full hover:bg-gray-200 disabled:pointer-events-none disabled:opacity-0 ";
const button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 font-medium disabled:bg-black/60 disabled:pointer-events-none";
</script>
