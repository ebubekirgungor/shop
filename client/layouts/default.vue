<template>
  <div class="min-h-[100dvh] font-poppins text-sm sm:bg-gray-50 select-none">
    <nav
      class="bg-white h-24 flex flex-col sm:flex-row sm:border-b justify-center items-center gap-x-10"
    >
      <input
        class="order-2 sm:order-1 bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:99%_60%] transition duration-300 ease-in-out sm:ml-[14vw] w-[90%] sm:w-[550px] h-10 border-none shadow focus:ring-0 sm:focus:scale-[1.01] sm:focus:shadow-lg sm:focus:shadow-black/10 bg-gray-100 rounded-md text-sm"
        type="text"
        placeholder="Search products"
      />
      <div
        class="flex order-1 sm:order-2 gap-x-2 sm:gap-x-10 items-center self-end sm:self-auto mr-4"
      >
        <div :class="'flex flex-col group items-center w-10 sm:w-20 h-16 ' + (user.token ? 'mt-10' : 'mt-6')">
          <NuxtLink
            :to="user.token ? '/account/orders' : '/login'"
            :class="button + (user.token ? '' : ' sm:hover:-translate-y-0.5')"
          >
            <div class="w-6 h-6 bg-[url(/icons/account.svg)]"></div>
            <span class="mt-0.5 hidden sm:block">{{
              user.token ? "Account" : "Login"
            }}</span>
          </NuxtLink>
          <div
            v-if="user.token"
            class="flex flex-col gap-y-2 justify-center items-center transition-visibility duration-300 ease-in-out delay-0 group-hover:delay-500 w-[100px] h-auto invisible group-hover:visible opacity-0 group-hover:opacity-100 bg-white border rounded-lg shadow-xl  group-hover:mt-4 -ml-8 sm:ml-0 p-3 z-10"
          >
            <NuxtLink v-if="user.role != 0" to="/admin" :class="menu_item"
              >Admin</NuxtLink
            >
            <button
              @click="
                userstate.$reset();
                navigateTo('/');
              "
              :class="menu_item"
            >
              Logout
            </button>
          </div>
        </div>
        <button :class="button + ' sm:hover:-translate-y-0.5'">
          <div class="w-6 h-6 bg-[url(/icons/favorite.svg)]"></div>
          <span class="mt-0.5 hidden sm:block">Favorites</span>
        </button>
        <button :class="button + ' sm:hover:-translate-y-0.5'">
          <div class="w-6 h-6 bg-[url(/icons/cart.svg)]"></div>
          <span class="mt-0.5 hidden sm:block">Cart</span>
        </button>
      </div>
    </nav>
    <slot />
    <footer></footer>
  </div>
</template>
<script setup lang="ts">
import { useUser } from "~/store/user";
const userstate = useUser();
const { user } = storeToRefs(userstate);
const button =
  "transition duration-200 ease-in-out flex h-10 w-10 sm:w-auto justify-center items-center gap-x-2 rounded-full hover:bg-black/10 sm:hover:bg-transparent";
const menu_item = "flex justify-center items-center w-full h-8";
</script>
<style>
* {
  -webkit-tap-highlight-color: transparent;
  user-select: none;
  -webkit-user-drag: none;
}
</style>
