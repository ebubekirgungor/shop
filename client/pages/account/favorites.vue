<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <transition name="background" mode="out-in">
      <div
        v-if="delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 size-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="delete_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed z-10"
      >
        <div
          class="flex flex-col p-2 bg-white sm:-mt-48 w-80 sm:w-96 h-auto rounded-xl"
        >
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">
            {{ $t("remove_product_from_favorites") }}
          </h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_product_id" required />
            <div class="text-center">{{ $t("are_you_sure") }}</div>
            <button
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              {{ $t("delete") }}
            </button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="mt-12 sm:mt-0 flex justify-center sm:justify-start items-center sm:p-6 text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      {{ $t("favorites") }}
    </div>
    <div
      v-if="favorites.length == 0"
      class="flex flex-col justify-center items-center gap-y-8 sm:min-w-[35rem] min-h-96 h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <div
        class="size-36 bg-no-repeat bg-center bg-contain bg-[url(/icons/favorite.svg)] contrast-0"
      ></div>
      <div class="text-xl text-gray-500">
        {{ $t("your_favorites_list_is_empty") }}
      </div>
    </div>
    <div
      v-else
      class="flex p-2 sm:p-6 w-screen sm:w-auto sm:min-w-[35rem] sm:min-h-96 h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <div class="w-full flex flex-wrap gap-2 sm:gap-6 items-center">
        <div
          v-for="product in favorites"
          class="flex flex-col w-[calc(50vw-12px)] sm:w-56 sm:h-[21rem] bg-white rounded-xl border"
        >
          <button
            @click="
              delete_dialog = true;
              delete_product_id = product.ID;
            "
            class="absolute m-2 self-end transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/delete.svg)] size-10 bg-black/10 rounded-full sm:hover:bg-black/25"
          ></button>
          <NuxtLink :to="'/product/' + product.url">
            <img
              v-if="product.images.length > 0"
              class="w-[calc(50vw-14px)] h-[calc(50vw-14px)] sm:size-56 object-contain rounded-t-xl"
              :src="'/images/products/' + product.images[0].name"
            />
            <div
              v-else
              class="sm:size-56 bg-no-repeat bg-center bg-contain bg-[url(/images/products/product.png)]"
            ></div>
          </NuxtLink>
          <div class="flex flex-col gap-y-6 p-4">
            <NuxtLink :to="'/product/' + product.url" class="text-md">{{
              product.title
            }}</NuxtLink>
            <div class="text-2xl">
              {{ product.list_price.toLocaleString("tr-TR") }}
              TL
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const { t } = useI18n();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
interface Image {
  order: number;
  name: string;
  url: string;
}
interface Product {
  ID: number | null;
  title: string;
  url: string;
  category_id: string;
  list_price: number;
  stock_quantity: string;
  images: Image[];
}
const delete_dialog = ref(false);
const delete_product_id = ref<number | null>(null);
const favorites = ref<Product[]>([]);
onMounted(() => {
  nextTick(async () => {
    await useFetch(config.apiBase + "/favorites", {
      headers: {
        Authorization: config.apiKey,
      },
      onResponse({ response }) {
        if (response._data) {
          favorites.value = response._data.favorites;
        }
      },
    });
  });
});
const remove = async () => {
  await useFetch(config.apiBase + "/favorites/" + delete_product_id.value, {
    headers: {
      Authorization: config.apiKey,
    },
    method: "delete",
    onResponse({ response }) {
      if (response._data == "ok") {
        favorites.value.splice(
          favorites.value
            .map((product: Product) => product.ID)
            .indexOf(delete_product_id.value),
          1
        );
        delete_dialog.value = false;
        delete_product_id.value = null;
        toast.success(t("product_removed"), {
          bodyClassName: "toast-font",
        });
      } else {
        toast.warning(response._data.error, {
          bodyClassName: "toast-font",
        });
      }
    },
  });
};
</script>
