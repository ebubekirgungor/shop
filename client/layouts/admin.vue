<template>
  <NuxtLayout name="default">
    <main class="flex justify-center m-4 gap-x-4">
      <nav class="flex flex-col bg-white rounded-xl shadow-md min-w-64 h-min gap-y-4 p-4">
        <NuxtLink :class="$route.path.includes('products') ? link + active : link" to="/admin/products">
          <div class="size-6 bg-[url(/icons/product.svg)]"></div>
          <span>Products</span>
        </NuxtLink>
        <NuxtLink :class="$route.path.includes('categories') ? link + active : link" to="/admin/categories">
          <div class="size-6 bg-[url(/icons/category.svg)]"></div>
          <span>Categories</span>
        </NuxtLink>
      </nav>
      <slot />
    </main>
  </NuxtLayout>
</template>
<script setup lang="ts">
const link = "transition duration-300 ease-in-out flex items-center gap-x-4 p-4 h-12 hover:bg-black/5 rounded-full ";
const active = "bg-black/10 pointer-events-none font-medium";
</script>