<template>
  <div class="min-h-[100dvh] font-poppins text-sm bg-gray-50 select-none">
    <nav class="bg-white h-24 flex border-b justify-center items-center gap-x-10">
      <input
        class="transition duration-200 ease-in-out ml-64 mr-6 w-[550px] h-10 border-none bg-gray-100 rounded-md text-sm"
        type="text"
        placeholder="Search products"
      />
      <div class="-ml-24 w-6 h-6 bg-[url(/icons/search.svg)]"></div>
      <NuxtLink to="/login" :class="button">
        <div class="w-6 h-6 bg-[url(/icons/account.svg)]"></div>
        <span class="mt-0.5">Account</span>
      </NuxtLink>
      <button :class="button">
        <div class="w-6 h-6 bg-[url(/icons/favorite.svg)]"></div>
        <span class="mt-0.5">Favorites</span>
      </button>
      <button :class="button">
        <div class="w-6 h-6 bg-[url(/icons/cart.svg)]"></div>
        <span class="mt-0.5">Cart</span>
      </button>
    </nav>
    <slot />
    <footer></footer>
  </div>
</template>
<script setup lang="ts">
const button = ref("transition duration-200 ease-in-out flex gap-x-2 hover:-translate-y-0.5");
</script>