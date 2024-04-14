<template>
  <main></main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
const config = useRuntimeConfig().public;
const role = useCookie("role");
interface Product {
  url: string;
  title: string;
  list_price: number;
  image: string;
}
const products = ref<Product[]>([]);
onMounted(() => {
  nextTick(async () => {
    await useFetch<Product[]>(config.apiBase + "/products", {
      onResponse({ response }) {
        if (response._data) {
          products.value = response._data;
        }
      },
    });
  });
});
</script>
