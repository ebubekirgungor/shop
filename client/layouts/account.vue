<template>
  <NuxtLayout name="default">
    <main class="flex justify-center m-4 gap-x-4">
      <nav class="flex flex-col bg-white rounded-xl shadow-md min-w-64 h-min gap-y-4 p-4">
        <NuxtLink :class="$route.path.endsWith('orders') ? link + active : link" to="/account/orders">
          <div class="size-6 bg-[url(/icons/order.svg)]"></div>
          <span>Orders</span>
        </NuxtLink>
        <NuxtLink :class="$route.path.endsWith('personal-details') ? link + active : link" to="/account/personal-details">
          <div class="size-6 bg-[url(/icons/account.svg)]"></div>
          <span>Personal Details</span>
        </NuxtLink>
        <NuxtLink :class="$route.path.endsWith('payment-methods') ? link + active : link" to="/account/payment-methods">
          <div class="size-6 bg-[url(/icons/payment.svg)]"></div>
          <span>Payment Methods</span>
        </NuxtLink>
        <NuxtLink :class="$route.path.endsWith('addresses') ? link + active : link" to="/account/addresses">
          <div class="size-6 bg-[url(/icons/address.svg)]"></div>
          <span>Addresses</span>
        </NuxtLink>
        <NuxtLink :class="$route.path.endsWith('change-password') ? link + active : link" to="/account/change-password">
          <div class="size-6 bg-[url(/icons/password.svg)]"></div>
          <span>Change Password</span>
        </NuxtLink>
        <NuxtLink :class="$route.path.endsWith('favorites') ? link + active : link" to="/account/favorites">
          <div class="size-6 bg-[url(/icons/favorite.svg)]"></div>
          <span>Favorites</span>
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