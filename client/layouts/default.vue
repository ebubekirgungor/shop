<template>
  <div
    @click="!result_focus ? (search_focus = false) : null"
    class="min-h-[100dvh] font-poppins text-sm sm:bg-gray-50 select-none"
  >
    <nav
      class="bg-white h-24 flex flex-col sm:flex-row sm:border-b justify-center items-center gap-x-10"
    >
      <div
        class="sm:ml-[14vw]"
        @mouseover="result_focus = true"
        @mouseout="result_focus = false"
      >
        <input
          class="order-2 sm:order-1 bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:99%_60%] transition duration-300 ease-in-out w-[90%] sm:w-[550px] h-10 border-none shadow focus:ring-0 sm:focus:scale-[1.01] sm:focus:shadow-lg sm:focus:shadow-black/10 bg-gray-100 rounded-md text-sm"
          type="text"
          placeholder="Search products"
          v-model="search"
          @focus="search_focus = true"
          @click="$event.stopPropagation()"
        />
        <transition name="background" mode="out-in"
          ><div
            v-if="search_focus && products.length > 0"
            class="fixed flex flex-col w-[550px] h-auto scale-[1.01] p-4 bg-white shadow-2xl rounded-b-lg z-10"
          >
            <NuxtLink
              :to="'/' + product.url"
              class="transition duration-200 flex justify-between px-4 py-3 rounded-full hover:bg-black/5"
              v-for="product in products"
            >
              <div v-html="product.title"></div>
              <div class="text-gray-400">{{ product.category }}</div>
            </NuxtLink>
          </div>
        </transition>
      </div>
      <div
        class="flex order-1 sm:order-2 gap-x-2 sm:gap-x-10 items-center self-end sm:self-auto mr-4"
      >
        <div
          :class="
            'flex flex-col group items-center w-10 sm:w-20 h-16 ' +
            (role != undefined ? 'mt-10' : 'mt-6')
          "
        >
          <NuxtLink
            :to="role != undefined ? '/account/orders' : '/login'"
            :class="
              button + (role != undefined ? '' : ' sm:hover:-translate-y-0.5')
            "
          >
            <div class="size-6 bg-[url(/icons/account.svg)]"></div>
            <span class="mt-0.5 hidden sm:block">{{
              role != undefined ? "Account" : "Login"
            }}</span>
          </NuxtLink>
          <div
            v-if="role != undefined"
            class="flex flex-col gap-y-2 justify-center items-center transition-visibility duration-300 ease-in-out delay-0 group-hover:delay-500 w-[100px] h-auto invisible group-hover:visible opacity-0 group-hover:opacity-100 bg-white border rounded-lg shadow-xl group-hover:mt-4 -ml-8 sm:ml-0 p-3 z-10"
          >
            <NuxtLink v-if="role != 0" to="/admin" :class="menu_item"
              >Admin</NuxtLink
            >
            <button @click="logout()" :class="menu_item">Logout</button>
          </div>
        </div>
        <button :class="button + ' sm:hover:-translate-y-0.5'">
          <div class="size-6 bg-[url(/icons/favorite.svg)]"></div>
          <span class="mt-0.5 hidden sm:block">Favorites</span>
        </button>
        <NuxtLink to="/cart" :class="button + ' sm:hover:-translate-y-0.5'">
          <div class="size-6 bg-[url(/icons/cart.svg)]"></div>
          <span class="mt-0.5 hidden sm:block">Cart</span>
        </NuxtLink>
      </div>
    </nav>
    <slot />
    <footer></footer>
  </div>
</template>
<script setup lang="ts">
const config = useRuntimeConfig().public;
const role = useCookie<number | null>("role");
interface Product {
  category: string;
  title: string;
  url: string;
}
const logout = async () => {
  role.value = null;
  await useFetch(config.apiBase + "/auth/logout", {
    onResponse({ response }) {
      if (response._data.status == "success") {
        navigateTo("/");
      }
    },
  });
};
const search = ref("");
const search_focus = ref(false);
const result_focus = ref(false);
const products = ref<Product[]>([]);
watch(search, async () => {
  if (search.value.length > 1) {
    await useFetch(config.apiBase + "/search", {
      query: { q: search.value },
      onResponse({ response }) {
        products.value = response._data;
      },
    });
  } else {
    products.value = [];
  }
});
watch(products, async () => {
  products.value.map((product: Product) => {
    let out = "";
    product.title.split(" ").map((word) => {
      if (word.toLowerCase().startsWith(search.value.toLowerCase())) {
        word =
          "<strong>" +
          word.substring(0, search.value.length) +
          "</strong>" +
          word.substring(search.value.length);
      }
      out += " " + word;
    });
    product.title = out;
  });
});
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
