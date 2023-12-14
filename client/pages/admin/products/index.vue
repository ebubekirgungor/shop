<template>
  <main class="flex flex-col gap-y-4">
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
      class="flex items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md"
    ></div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: "admin",
});
</script>
