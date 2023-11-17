<template>
  <NuxtLayout name="default">
    <main class="flex justify-center mt-4 gap-x-4">
      <nav class="flex flex-col bg-white rounded-xl shadow-md w-64 h-auto gap-y-4 p-4">
        <NuxtLink :class="$route.path.endsWith('products') ? link + active : link" to="/admin/products">
          <div class="w-6 h-6 bg-[url(/icons/product.svg)]"></div>
          <span>Products</span>
        </NuxtLink>
      </nav>
      <slot />
    </main>
  </NuxtLayout>
</template>
<script setup lang="ts">
import { useUser } from "@/store/user";
const { user } = useUser();
const link = "transition duration-300 ease-in-out flex items-center gap-x-4 p-4 h-12 hover:bg-black/5 rounded-full ";
const active = "bg-black/10 pointer-events-none font-medium";
</script>