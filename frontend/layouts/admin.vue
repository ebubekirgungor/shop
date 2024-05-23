<template>
  <NuxtLayout name="default">
    <main class="flex justify-center m-4 gap-x-4">
      <button
        @click="nav_menu = true"
        class="sm:hidden absolute left-0 ml-5 size-8 bg-[url(/icons/menu.svg)] bg-cover bg-no-repeat"
      ></button>
      <button
        @click="logout()"
        class="sm:hidden absolute right-0 mr-5 size-7 bg-[url(/icons/logout.svg)] bg-cover bg-no-repeat"
      ></button>
      <nav
        class="transition-all duration-500 ease-in-out right-[100vw] sm:right-0 z-10 sm:z-0 top-0 fixed sm:relative flex flex-col bg-white sm:rounded-xl sm:shadow-md w-screen sm:w-auto min-w-64 h-screen sm:h-min gap-y-4 p-4"
        :class="{ '!right-[0vw]': nav_menu }"
      >
        <button
          @click="nav_menu = false"
          class="sm:hidden ml-3 mt-3 size-7 bg-[url(/icons/previous.svg)] bg-cover bg-no-repeat"
        ></button>
        <NuxtLinkLocale
          :class="$route.name!.toString().startsWith('admin-products') ? link + active : link"
          to="admin-products"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/product.svg)]"></div>
          <span>{{ $t("products") }}</span>
        </NuxtLinkLocale>
        <NuxtLinkLocale
          :class="$route.name!.toString().startsWith('admin-categories') ? link + active : link"
          to="admin-categories"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/category.svg)]"></div>
          <span>{{ $t("categories") }}</span>
        </NuxtLinkLocale>
      </nav>
      <slot />
    </main>
  </NuxtLayout>
</template>
<script setup lang="ts">
const config = useRuntimeConfig().public;
const localePath = useLocalePath();
const role = useCookie<number | null>("role");

const nav_menu = ref(false);

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

const link =
  "transition duration-300 ease-in-out flex items-center gap-x-4 p-4 h-12 hover:bg-black/5 rounded-full ";
const active = "bg-black/10 pointer-events-none font-medium";
</script>
