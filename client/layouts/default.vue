<template>
  <div class="min-h-[100dvh] font-poppins text-sm bg-gray-50 select-none">
    <nav
      class="bg-white h-24 flex border-b justify-center items-center gap-x-10"
    >
      <input
        class="bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:99%_60%] transition duration-300 ease-in-out ml-64 w-[550px] h-10 border-none shadow focus:ring-0 focus:scale-[1.01] focus:shadow-lg focus:shadow-black/10 bg-gray-100 rounded-md text-sm"
        type="text"
        placeholder="Search products"
      />
      <div class="flex items-center group h-16">
        <NuxtLink
          :to="user.token ? '/account' : '/login'"
          :class="button + (user.token ? '' : ' hover:-translate-y-0.5')"
        >
          <div class="w-6 h-6 bg-[url(/icons/account.svg)]"></div>
          <span class="mt-0.5">{{ user.token ? "Account" : "Login" }}</span>
        </NuxtLink>
        <div
          v-if="user.token"
          @click="userstate.$reset(); navigateTo('/');"
          class="cursor-pointer flex justify-center transition-visibility duration-300 ease-in-out delay-0 group-hover:delay-200 absolute w-[100px] h-12 invisible group-hover:visible opacity-0 group-hover:opacity-100 bg-white border rounded-lg shadow-xl mt-20 group-hover:mt-24 p-3"
        >
          Logout
        </div>
      </div>

      <button :class="button + ' hover:-translate-y-0.5'">
        <div class="w-6 h-6 bg-[url(/icons/favorite.svg)]"></div>
        <span class="mt-0.5">Favorites</span>
      </button>
      <button :class="button + ' hover:-translate-y-0.5'">
        <div class="w-6 h-6 bg-[url(/icons/cart.svg)]"></div>
        <span class="mt-0.5">Cart</span>
      </button>
    </nav>
    <slot />
    <footer></footer>
  </div>
</template>
<script setup lang="ts">
import { useUser } from "~/store/user";
const userstate = useUser();
const { user } = storeToRefs(userstate);
const button = ref(
  "transition duration-200 ease-in-out flex h-10 items-center gap-x-2"
);
</script>
