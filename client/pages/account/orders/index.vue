<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div
      class="flex flex-col sm:flex-row justify-center sm:justify-start items-center mt-12 sm:mt-0 gap-4 sm:p-6 text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      {{ $t("orders") }}
      <div
        class="w-screen sm:w-auto flex gap-x-4 px-4 overflow-x-auto"
        style="scrollbar-width: none"
      >
        <button
          v-for="status in status_names.filter(
            (status) => status.status != DeliveryStatus.Delivered
          )"
          @click="status_filter = status.status"
          class="transition duration-200 ease-in-out h-8 px-4 text-[15px] border-2 rounded-xl sm:hover:bg-black/10 whitespace-nowrap"
          :class="{ 'bg-black/10': status_filter == status.status }"
        >
          {{ $t(status.name) }}
        </button>
      </div>
    </div>
    <div
      class="flex sm:p-6 sm:min-w-[35rem] bg-white sm:rounded-xl sm:shadow-md"
    >
      <div class="w-full flex flex-col gap-y-6 p-4 sm:p-0">
        <input
          class="transition duration-300 ease-in-out bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:98%_60%] sm:w-64 rounded-md border-0 bg-gray-100 text-sm focus:ring-2 focus:ring-slate-300"
          v-model="search_filter"
          type="text"
          :placeholder="$t('search')"
        />
        <div
          v-if="filtered_orders.length == 0"
          class="flex flex-col justify-center items-center gap-y-8 h-auto"
        >
          <div
            class="size-36 bg-no-repeat bg-center bg-contain bg-[url(/icons/order.svg)] contrast-0"
          ></div>
          <div class="text-xl text-gray-500">
            {{ $t("dont_have_order_yet") }}
          </div>
        </div>
        <div
          v-else
          v-for="order in filtered_orders"
          class="flex flex-col sm:gap-x-8 bg-white rounded-xl border"
        >
          <div
            class="flex flex-col sm:flex-row justify-between p-5 bg-black/5 rounded-t-xl"
          >
            <div class="flex flex-col">
              <div class="hidden sm:block text-gray-600">
                {{ $t("order_date") }}
              </div>
              <div class="sm:text-lg">
                {{ new Date(order.date).toLocaleString("tr-TR", options) }}
              </div>
            </div>
            <div class="hidden sm:flex flex-col">
              <div class="text-gray-600">{{ $t("customer") }}</div>
              <div class="text-lg">
                {{ order.customer_name }}
              </div>
            </div>
            <div class="flex flex-col">
              <div class="hidden sm:block text-gray-600">
                {{ $t("total_amount") }}
              </div>
              <div class="text-lg">
                {{ order.price.toLocaleString("tr-TR") }}
                TL
              </div>
            </div>
            <div class="hidden sm:flex flex-col">
              <div class="text-gray-600">{{ $t("total_products") }}</div>
              <div class="text-lg">
                {{
                  order.products.reduce((total: number, product: Product) => {
                    return total + product.quantity;
                  }, 0)
                }}
              </div>
            </div>
            <NuxtLink
              :to="localePath('account-orders') + '/' + order.id"
              class="absolute sm:relative right-0 mr-8 sm:mr-0 transition duration-300 ease-in-out flex justify-center items-center w-20 h-12 rounded-xl bg-black text-white hover:bg-black/80 font-medium"
              >{{ $t("details") }}</NuxtLink
            >
          </div>
          <div class="flex items-center px-5 pt-4 gap-x-4">
            <div :class="status_names[order.delivery_status + 1].icon"></div>
            {{ $t(status_names[order.delivery_status + 1].name + "_status") }}
          </div>
          <div
            class="max-w-[calc(100vw-2rem)] flex overflow-x-auto sm:flex-wrap items-center p-5 gap-x-3 gap-y-6 sm:gap-6"
          >
            <NuxtLink
              v-for="product in order.products"
              :to="'/product/' + product.url"
              class="flex justify-center min-w-12 size-12 sm:min-w-24 sm:size-24 border border-gray-300 rounded-lg"
              ><img
                class="object-center object-contain rounded-lg"
                :src="'/images/products/' + product.image"
              />
              <span
                class="hidden absolute mt-10 ml-12 sm:mt-20 sm:ml-24 sm:flex justify-center items-center size-6 sm:size-7 rounded-full bg-gray-300 font-medium"
                >{{ product.quantity }}</span
              ></NuxtLink
            >
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
const config = useRuntimeConfig().public;
const localePath = useLocalePath();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
interface Product {
  url: string;
  title: string;
  image: string;
  quantity: number;
}
interface Order {
  id: number;
  date: Date;
  customer_name: string;
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
    name: "all",
    icon: "",
    status: null,
  },
  {
    name: "delivered",
    icon: "bg-center bg-no-repeat size-6 bg-[url(/icons/check.svg)]",
    status: DeliveryStatus.Delivered,
  },
  {
    name: "in_progress",
    icon: "bg-center bg-no-repeat size-6 bg-[url(/icons/progress.svg)]",
    status: DeliveryStatus.InProgress,
  },
  {
    name: "returned",
    icon: "bg-center bg-no-repeat size-6 bg-[url(/icons/return.svg)]",
    status: DeliveryStatus.Returned,
  },
  {
    name: "canceled",
    icon: "bg-center bg-no-repeat size-6 bg-[url(/icons/close.svg)]",
    status: DeliveryStatus.Canceled,
  },
];
const options = <Object>{
  day: "numeric",
  month: "long",
  year: "numeric",
  hour: "numeric",
  minute: "numeric",
};
const orders = ref<Order[]>([]);
const status_filter = ref<DeliveryStatus | null>(null);
const search_filter = ref("");
onMounted(() => {
  nextTick(async () => {
    await useFetch(config.apiBase + "/orders", {
      headers: {
        Authorization: config.apiKey,
      },
      onResponse({ response }) {
        if (response._data) {
          orders.value = response._data;
        }
      },
    });
  });
});
const filtered_orders = computed(() => {
  const filtered_orders = orders.value.filter((order: Order) => {
    const filtered_products = order.products.filter((product) => {
      return product.title
        .toLowerCase()
        .includes(search_filter.value.toLowerCase());
    });
    return filtered_products.length > 0;
  });
  if (status_filter.value) {
    return filtered_orders.filter(
      (order: Order) => order.delivery_status == status_filter.value
    );
  } else {
    return filtered_orders;
  }
});
</script>
