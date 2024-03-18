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
        <div class="flex flex-col p-2 bg-white -mt-48 w-96 h-auto rounded-xl">
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">Delete product from favorites</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_product_id" required />
            <div class="text-center">Are you sure?</div>
            <button
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              Delete
            </button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center p-6 text-xl h-auto bg-white rounded-xl shadow-md"
    >
      Favorites
    </div>
    <div
      v-if="favorites.length == 0"
      class="flex flex-col justify-center items-center gap-y-8 min-w-[35rem] min-h-96 h-auto bg-white rounded-xl shadow-md"
    >
      <div
        class="size-36 bg-no-repeat bg-center bg-contain bg-[url(/icons/favorite.svg)] contrast-0"
      ></div>
      <div class="text-xl text-gray-500">Your favorites list is empty</div>
    </div>
    <div
      v-else
      class="flex p-6 min-w-[35rem] min-h-96 h-auto bg-white rounded-xl shadow-md"
    >
      <div
        class="w-full grid grid-cols-4 grid-cols-auto_box gap-6 items-center"
      >
        <div
          v-for="product in favorites"
          class="flex flex-col w-56 h-[21rem] bg-white rounded-xl border"
        >
          <button
            @click="
              delete_dialog = true;
              delete_product_id = product.ID;
            "
            class="absolute mt-2 ml-44 transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/delete.svg)] size-10 bg-black/10 rounded-full hover:bg-black/25"
          ></button>
          <img
            v-if="product.images.length > 0"
            class="size-56 object-contain"
            :src="'/images/products/' + product.images[0].name"
          />
          <div v-else class="size-56 border rounded-lg">
            <div
              class="size-56 bg-no-repeat bg-center bg-contain bg-[url(/icons/order.svg)] contrast-75 invert"
            ></div>
          </div>
          <div class="flex flex-col gap-y-4 p-4">
            <div class="text-lg">{{ product.title }}</div>
            <div class="text-2xl">
              {{
                product.list_price.toLocaleString("tr-TR", {
                  minimumFractionDigits: 2,
                })
              }}
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
    await useFetch<Product[]>(config.apiBase + "/favorites", {
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
        toast.success("Product removed", {
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
