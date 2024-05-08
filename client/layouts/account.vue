<template>
  <NuxtLayout name="default">
    <main class="flex justify-center m-4 gap-x-4">
      <button
        @click="nav_menu = true"
        class="sm:hidden absolute left-0 ml-5 size-8 bg-[url(/icons/menu.svg)] bg-cover bg-no-repeat"
      ></button>
      <nav
        class="transition-all duration-500 ease-in-out right-[100vw] sm:right-0 z-10 sm:z-0 top-0 fixed sm:relative flex flex-col bg-white sm:rounded-xl sm:shadow-md w-screen sm:w-auto min-w-64 h-screen sm:h-min gap-y-4 p-4"
        :class="{ '!right-[0vw]': nav_menu }"
      >
        <button
          @click="nav_menu = false"
          class="sm:hidden ml-3 mt-3 size-7 bg-[url(/icons/previous.svg)] bg-cover bg-no-repeat"
        ></button>
        <NuxtLink
          :class="$route.path.endsWith('orders') ? link + active : link"
          to="/account/orders"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/order.svg)]"></div>
          <span>Orders</span>
        </NuxtLink>
        <NuxtLink
          :class="
            $route.path.endsWith('personal-details') ? link + active : link
          "
          to="/account/personal-details"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/account.svg)]"></div>
          <span>Personal Details</span>
        </NuxtLink>
        <NuxtLink
          :class="
            $route.path.endsWith('payment-methods') ? link + active : link
          "
          to="/account/payment-methods"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/payment.svg)]"></div>
          <span>Payment Methods</span>
        </NuxtLink>
        <NuxtLink
          :class="$route.path.endsWith('addresses') ? link + active : link"
          to="/account/addresses"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/address.svg)]"></div>
          <span>Addresses</span>
        </NuxtLink>
        <NuxtLink
          :class="
            $route.path.endsWith('change-password') ? link + active : link
          "
          to="/account/change-password"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/password.svg)]"></div>
          <span>Change Password</span>
        </NuxtLink>
        <NuxtLink
          :class="$route.path.endsWith('favorites') ? link + active : link"
          to="/account/favorites"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/favorite.svg)]"></div>
          <span>Favorites</span>
        </NuxtLink>
      </nav>
      <slot />
    </main>
  </NuxtLayout>
</template>
<script setup lang="ts">
const link =
  "transition duration-300 ease-in-out flex items-center gap-x-4 p-4 h-12 hover:bg-black/5 rounded-full ";
const active = "bg-black/10 pointer-events-none font-medium";
const nav_menu = ref(false);
</script>
