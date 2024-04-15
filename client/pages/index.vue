<template>
  <main class="flex justify-center">
    <div class="grid grid-cols-4 justify-items-center gap-6 p-4">
      <NuxtLink
        v-for="category in categories"
        :to="category.url"
        class="size-72 border rounded-xl bg-white"
      >
        <img
          class="size-full rounded-xl object-contain"
          :src="'/images/categories/' + category.image"
        />
      </NuxtLink>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
const config = useRuntimeConfig().public;
const role = useCookie("role");
/*interface Product {
  url: string;
  title: string;
  list_price: number;
  image: string;
}*/
interface Category {
  title: string;
  url: string;
  image: string;
}
const categories = ref<Category[]>([]);
//const products = ref<Product[]>([]);
onMounted(() => {
  nextTick(async () => {
    /*await useFetch(config.apiBase + "/products", {
      onResponse({ response }) {
        if (response._data) {
          products.value = response._data;
        }
      },
    });*/
    await useFetch(config.apiBase + "/categories", {
      onResponse({ response }) {
        if (response._data) {
          categories.value = response._data;
        }
      },
    });
  });
});
</script>
