<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div
      class="flex items-center mt-12 sm:mt-0 gap-x-4 sm:p-6 text-lg sm:text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <button
        @click="navigateTo('/account/orders')"
        class="transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/previous.svg)] ml-4 sm:ml-0 size-10 bg-black/5 rounded-full hover:bg-black/10"
      ></button>
      {{ $t("order_detail") }}
      <div v-if="fetch_complete" class="mr-2 sm:mr-0 text-[16px] text-gray-500">
        {{ new Date(order?.date!).toLocaleString("tr-TR", options) }}
      </div>
    </div>
    <div
      class="sm:min-w-[35rem] sm:min-h-[43.75rem] h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <div v-if="fetch_complete" class="flex flex-col gap-y-4 sm:p-6">
        <div class="w-screen sm:w-full flex flex-col sm:border sm:rounded-xl">
          <div class="bg-black/5 text-lg p-4 sm:rounded-t-xl">
            {{ $t("items") }}
          </div>
          <div class="flex items-center px-5 pt-4 gap-x-4">
            <div :class="status_names[order?.delivery_status!].icon"></div>
            {{ $t(status_names[order?.delivery_status!].name + "_status") }}
          </div>
          <div
            class="w-full flex flex-wrap gap-2 sm:gap-4 p-2 sm:p-4 items-center"
          >
            <div
              v-for="product in order?.products"
              class="flex flex-col w-[calc(50vw-12px)] sm:w-56 sm:h-[21rem] bg-white rounded-xl border"
            >
              <NuxtLink :to="'/product/' + product.url">
                <img
                  class="w-[calc(50vw-14px)] h-[calc(50vw-14px)] sm:size-56 object-contain"
                  :src="'/images/products/' + product.image"
                />
              </NuxtLink>
              <div class="flex flex-col gap-y-6 p-2 sm:p-4">
                <NuxtLink :to="'/product/' + product.url" class="text-md">{{
                  product.title
                }}</NuxtLink>
                <div class="flex items-center justify-between">
                  <div class="text-lg sm:text-2xl">
                    {{ product.list_price.toLocaleString("tr-TR") }}
                    TL
                  </div>
                  <div class="text-gray-500">
                    {{ $t("quantity") }}: {{ product.quantity }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex flex-col sm:flex-row gap-x-4">
          <div class="w-full h-min flex flex-col sm:border sm:rounded-xl">
            <div class="bg-black/5 text-lg p-4 sm:rounded-t-xl">
              {{ $t("shipment") }}
            </div>
            <div class="flex flex-col p-4 gap-y-6">
              <div class="flex flex-col gap-y-1">
                <div class="text-gray-500">{{ $t("customer") }}</div>
                {{ order?.customer_name }}
              </div>
              <div class="flex flex-col gap-y-1">
                <div class="text-gray-500">{{ $t("address") }}</div>
                <div class="whitespace-pre">{{ order?.delivery_address }}</div>
              </div>
            </div>
          </div>
          <div class="w-full h-min flex flex-col sm:border sm:rounded-xl">
            <div class="bg-black/5 text-lg p-4 sm:rounded-t-xl">
              {{ $t("payment") }}
            </div>
            <div class="flex flex-col p-4 divide-y">
              <div class="flex justify-between items-center h-8">
                <div class="text-gray-500">{{ $t("subtotal") }}</div>
                {{
                  order?.products
                    .reduce((total: number, product: Product) => {
                      return total + product.list_price * product.quantity;
                    }, 0)
                    .toLocaleString("tr-TR", {
                      maximumFractionDigits: 2,
                    })
                }}
                TL
              </div>
              <div class="flex justify-between items-center h-8">
                <div class="text-gray-500">{{ $t("shipping_cost") }}</div>
                {{ shipping }} TL
              </div>
              <div class="flex justify-between items-center h-8">
                <div class="font-medium">{{ $t("total_charge") }}</div>
                {{ order?.price.toLocaleString("tr-TR") }} TL
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-else class="flex items-center justify-center w-full h-full">
        <div
          class="animate-spin size-24 bg-[url(/icons/loading.svg)] bg-no-repeat bg-cover"
        ></div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
const route = useRoute();
const config = useRuntimeConfig().public;
definePageMeta({
  middleware: "auth",
  layout: "account",
});
interface Product {
  title: string;
  url: string;
  image: string;
  list_price: number;
  quantity: number;
}
interface Order {
  id: number;
  date: Date;
  customer_name: string;
  delivery_address: string;
  price: number;
  products: Product[];
  delivery_status: DeliveryStatus;
}
enum DeliveryStatus {
  Delivered,
  InProgress,
  Returned,
  Canceled,
}
const status_names = [
  {
    name: "delivered",
    icon: "bg-center bg-no-repeat size-6 bg-[url(/icons/check.svg)]",
  },
  {
    name: "in_progress",
    icon: "bg-center bg-no-repeat size-6 size-6 bg-[url(/icons/progress.svg)]",
  },
  {
    name: "returned",
    icon: "bg-center bg-no-repeat size-6 size-6 bg-[url(/icons/return.svg)]",
  },
  {
    name: "canceled",
    icon: "bg-center bg-no-repeat size-6 size-6 bg-[url(/icons/close.svg)]",
  },
];
const options = <Object>{
  day: "numeric",
  month: "long",
  year: "numeric",
  hour: "numeric",
  minute: "numeric",
};
const order = ref<Order>();
const shipping = ref(50);
const fetch_complete = ref(false);
onMounted(() => {
  nextTick(async () => {
    await useFetch(config.apiBase + "/orders/" + route.params.order, {
      headers: {
        Authorization: config.apiKey,
      },
      onResponse({ response }) {
        if (response._data) {
          order.value = response._data;
          fetch_complete.value = true;
        }
      },
    });
  });
});
</script>
