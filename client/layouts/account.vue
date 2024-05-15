<template>
  <NuxtLayout name="default">
    <main class="flex justify-center m-4 gap-x-4">
      <button
        @click="nav_menu = true"
        class="sm:hidden absolute left-0 ml-5 size-8 bg-[url(/icons/menu.svg)] bg-cover bg-no-repeat"
      ></button>
      <button
        v-if="role != 0"
        @click="navigateTo(localePath('admin'))"
        class="sm:hidden absolute right-0 mr-20 size-7 bg-[url(/icons/admin.svg)] bg-cover bg-no-repeat"
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
          :class="$route.name!.toString().startsWith('account-orders') ? link + active : link"
          to="account-orders"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/order.svg)]"></div>
          <span>{{ $t("orders") }}</span>
        </NuxtLinkLocale>
        <NuxtLinkLocale
          :class="
            $route.name!.toString().startsWith('account-personal-details') ? link + active : link
          "
          to="account-personal-details"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/account.svg)]"></div>
          <span>{{ $t("personal_details") }}</span>
        </NuxtLinkLocale>
        <NuxtLinkLocale
          :class="
            $route.name!.toString().startsWith('account-payment-methods') ? link + active : link
          "
          to="account-payment-methods"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/payment.svg)]"></div>
          <span>{{ $t("payment_methods") }}</span>
        </NuxtLinkLocale>
        <NuxtLinkLocale
          :class="$route.name!.toString().startsWith('account-addresses') ? link + active : link"
          to="account-addresses"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/address.svg)]"></div>
          <span>{{ $t("addresses") }}</span>
        </NuxtLinkLocale>
        <NuxtLinkLocale
          :class="
            $route.name!.toString().startsWith('account-change-password') ? link + active : link
          "
          to="account-change-password"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/password.svg)]"></div>
          <span>{{ $t("change_password") }}</span>
        </NuxtLinkLocale>
        <NuxtLinkLocale
          :class="$route.name!.toString().startsWith('account-favorites') ? link + active : link"
          to="account-favorites"
          @click="nav_menu = false"
        >
          <div class="size-6 bg-[url(/icons/favorite.svg)]"></div>
          <span>{{ $t("favorites") }}</span>
        </NuxtLinkLocale>
        <div class="sm:hidden group absolute right-0 top-0 w-36 m-4">
          <div
            class="w-full flex justify-between items-center border rounded-md p-2 cursor-pointer"
          >
            {{ $t("language") }}
            <div class="size-6 bg-[url(/icons/translate.svg)]"></div>
          </div>
          <div
            class="transition-visibility duration-300 flex flex-col gap-2 p-2 w-full group-hover:mt-4 group-hover:visible group-hover:opacity-100 invisible opacity-0 bg-white border rounded-md shadow-xl"
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
    </main>
  </NuxtLayout>
</template>
<script setup lang="ts">
const config = useRuntimeConfig().public;
const { locales } = useI18n();
const switchLocalePath = useSwitchLocalePath();
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
