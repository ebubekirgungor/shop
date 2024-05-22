<template>
  <div
    class="min-h-dvh font-poppins text-sm sm:bg-gray-50 select-none overflow-hidden"
  >
    <nav
      class="bg-white h-16 sm:h-24 flex sm:border-b justify-end sm:justify-center sm:items-center sm:gap-x-4"
    >
      <div
        class="ml-20 flex sm:justify-center sm:items-center sm:gap-x-12 sm:grow"
      >
        <div class="flex items-center sm:w-80 sm:-ml-16 mr-8 sm:mr-16">
          <NuxtLinkLocale
            to="index"
            class="w-32 sm:w-52 h-16 sm:h-24 bg-center sm:bg-left bg-cover bg-[url(/images/logo.png)]"
          ></NuxtLinkLocale>
        </div>
        <div class="flex sm:justify-center">
          <input
            class="sm:block peer order-2 sm:order-1 bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:99%_60%] transition duration-300 ease-in-out sm:w-[550px] h-10 border-none shadow focus:ring-0 sm:focus:scale-[1.01] focus:shadow-lg focus:shadow-black/10 bg-gray-100 sm:rounded-md text-sm sm:will-change-transform"
            :class="{ hidden: !mobile_search }"
            type="text"
            :placeholder="$t('search_products')"
            v-model="search"
          />
          <button
            class="overflow-y-auto transition-visibility duration-300 ease-in-out sm:opacity-0 sm:invisible absolute hidden sm:flex flex-col divide-y sm:divide-y-0 sm:w-[550px] sm:h-auto sm:mt-10 sm:scale-[1.01] p-4 bg-white sm:shadow-2xl sm:rounded-b-lg cursor-default z-10"
            :class="{
              'active:visible active:opacity-100 focus:visible peer-focus:visible focus:opacity-100 peer-focus:opacity-100':
                search.length > 1 && products.length,
            }"
            :style="
              mobile_search ? 'inset: 0px; display:flex; position:fixed;' : ''
            "
          >
            <div
              class="sm:hidden mb-4 flex items-center shadow w-full bg-gray-100"
            >
              <div
                @click="mobile_search = false"
                class="min-w-10 h-10 bg-[url(/icons/previous.svg)] bg-no-repeat bg-center contrast-0 rounded-l-md"
              ></div>
              <input
                class="bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:99%_60%] transition duration-300 ease-in-out w-full h-10 border-none focus:ring-0 bg-gray-100 rounded-r-md text-sm sm:will-change-transform"
                type="text"
                :placeholder="$t('search_products')"
                v-model="search"
              />
            </div>
            <NuxtLink
              :to="$t('product_url') + product.url"
              class="transition duration-200 flex justify-between w-full px-4 py-3 sm:rounded-full hover:bg-black/5"
              v-for="product in products"
            >
              <div v-html="product.title"></div>
              <div class="text-gray-400">{{ product.category }}</div>
            </NuxtLink>
          </button>
        </div>
        <div class="sm:w-80 flex flex-col justify-center">
          <div
            class="flex order-1 sm:order-2 gap-x-2 sm:gap-x-10 items-center self-end sm:self-auto mr-4"
          >
            <button
              @click="mobile_search = true"
              class="transition duration-200 ease-in-out flex h-10 w-10 sm:w-auto justify-center items-center gap-x-2 rounded-full sm:hover:bg-black/10 sm:hover:bg-transparent will-change-transform sm:hidden sm:hover:-translate-y-0.5"
            >
              <div class="size-6 bg-[url(/icons/search_mobile.svg)]"></div>
            </button>
            <div
              :class="
                'flex flex-col group items-center w-10 sm:w-20 h-6 sm:h-16 ' +
                (role != undefined ? 'sm:mt-10' : 'sm:mt-6')
              "
            >
              <NuxtLinkLocale
                :to="role != undefined ? 'account-orders' : 'login'"
                :class="
                  'transition duration-200 ease-in-out flex h-10 w-10 sm:w-auto justify-center items-center gap-x-2 rounded-full sm:hover:bg-black/10 sm:hover:bg-transparent will-change-transform' +
                  (role != undefined ? '' : ' sm:hover:-translate-y-0.5')
                "
              >
                <div class="size-6 bg-[url(/icons/account.svg)]"></div>
                <span class="mt-0.5 hidden sm:block whitespace-nowrap">{{
                  role != undefined ? $t("account") : $t("login")
                }}</span>
              </NuxtLinkLocale>
              <div
                v-if="role != undefined"
                class="hidden sm:flex flex-col gap-y-2 justify-center items-center transition-visibility duration-300 ease-in-out delay-0 group-hover:delay-500 w-[100px] h-auto invisible group-hover:visible opacity-0 group-hover:opacity-100 bg-white border rounded-lg shadow-xl group-hover:mt-4 -ml-8 sm:ml-0 p-3 z-10"
              >
                <NuxtLinkLocale
                  v-if="role != 0"
                  to="admin"
                  class="flex justify-center items-center w-full h-8"
                  >{{ $t("admin") }}</NuxtLinkLocale
                >
                <button
                  @click="logout()"
                  class="flex justify-center items-center w-full h-8"
                >
                  {{ $t("logout") }}
                </button>
              </div>
            </div>
            <NuxtLinkLocale
              to="account-favorites"
              class="transition duration-200 ease-in-out flex h-10 w-10 sm:w-auto justify-center items-center gap-x-2 rounded-full sm:hover:bg-black/10 sm:hover:bg-transparent will-change-transform sm:hover:-translate-y-0.5"
            >
              <div class="size-6 bg-[url(/icons/favorite.svg)]"></div>
              <span class="mt-0.5 hidden sm:block">{{ $t("favorites") }}</span>
            </NuxtLinkLocale>
            <NuxtLinkLocale
              to="cart"
              class="transition duration-200 ease-in-out flex h-10 w-10 sm:w-auto justify-center items-center gap-x-2 rounded-full sm:hover:bg-black/10 sm:hover:bg-transparent will-change-transform sm:hover:-translate-y-0.5"
            >
              <div class="size-6 bg-[url(/icons/cart.svg)]"></div>
              <span class="mt-0.5 hidden sm:block">{{ $t("cart") }}</span>
            </NuxtLinkLocale>
          </div>
        </div>
      </div>
      <div class="hidden sm:block group mr-4">
        <div
          class="size-12 border rounded-full cursor-pointer bg-[url(/icons/translate.svg)] bg-no-repeat bg-center"
        ></div>
        <div
          class="transition-visibility duration-300 absolute right-4 flex flex-col items-center gap-2 p-2 w-28 absolute group-hover:mt-3 group-hover:visible group-hover:opacity-100 invisible opacity-0 bg-white border rounded-md shadow-xl"
        >
          <NuxtLink
            v-for="locale in locales"
            :key="locale.code"
            :to="switchLocalePath(locale.code)"
            class="p-2"
          >
            {{ locale.name }}
          </NuxtLink>
        </div>
      </div>
    </nav>
    <slot />
    <footer></footer>
  </div>
</template>
<script setup lang="ts">
const config = useRuntimeConfig().public;
const role = useCookie<number | null>("role");
const { locales } = useI18n();
const switchLocalePath = useSwitchLocalePath();
const localePath = useLocalePath();
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
        navigateTo(localePath("/"));
      }
    },
  });
};
const search = ref("");
const mobile_search = ref(false);
const products = ref<Product[]>([]);
watch(search, async () => {
  if (search.value.length > 1) {
    await useFetch(config.apiBase + "/search", {
      query: { q: search.value },
      onResponse({ response }) {
        products.value = response._data;
      },
    });
  }
});
watch(products, () => {
  const regex = new RegExp(search.value.split(" ").join("|"), "gi");
  products.value.map((product: Product) => {
    let out = "";
    product.title.split(" ").map((word) => {
      word += " ";
      const match = word.match(regex);
      if (match) {
        word = match.reduce((acc, m) => {
          return acc.replace(m, "<strong>" + m + "</strong>");
        }, word);
      }
      out += word;
    });
    product.title = out;
  });
});
</script>
