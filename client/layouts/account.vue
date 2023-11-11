<template>
  <NuxtLayout name="default">
    <main class="flex justify-center mt-4">
      <nav class="flex flex-col justify-center bg-white rounded-xl shadow-md w-64 h-auto gap-y-4 p-4">
        <NuxtLink :class="link" to="/account/orders">
          <div class="w-6 h-6 bg-[url(/icons/order.svg)]"></div>
          <span>Orders</span>
        </NuxtLink>
        <NuxtLink :class="link" to="/account/personal-details">
          <div class="w-6 h-6 bg-[url(/icons/account.svg)]"></div>
          <span>Personal Details</span>
        </NuxtLink>
        <NuxtLink :class="link" to="/account/payment-methods">
          <div class="w-6 h-6 bg-[url(/icons/payment.svg)]"></div>
          <span>Payment Methods</span>
        </NuxtLink>
        <NuxtLink :class="link" to="/account/addresses">
          <div class="w-6 h-6 bg-[url(/icons/address.svg)]"></div>
          <span>Addresses</span>
        </NuxtLink>
      </nav>
      <slot />
    </main>
  </NuxtLayout>
</template>
<script setup lang="ts">
const link = "transition duration-200 ease-in-out flex items-center gap-x-4 p-3 h-12 hover:bg-black/5 rounded-full";
</script>