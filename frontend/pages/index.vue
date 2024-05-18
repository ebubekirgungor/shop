<template>
  <main class="flex justify-center">
    <div class="w-[80rem] flex flex-wrap gap-6 p-6">
      <NuxtLink
        v-for="category in categories"
        :to="locale == 'tr' ? category.url : locale + '/' + category.url"
        class="min-h-[19rem] max-h-[19rem] sm:size-72 border rounded-xl bg-white"
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
const { locale } = useI18n();
const config = useRuntimeConfig().public;
interface Category {
  title: string;
  url: string;
  image: string;
}
const categories = ref<Category[]>([]);
onMounted(() => {
  nextTick(async () => {
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
