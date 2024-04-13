<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div
      class="flex items-center gap-x-4 p-6 text-xl h-auto bg-white rounded-xl shadow-md"
    >
      Orders
      <button
        @click="status_filter = null"
        class="transition duration-200 ease-in-out px-4 text-[15px] border-2 rounded-xl hover:bg-black/10"
        :class="{ 'bg-black/10': status_filter == null }"
      >
        All
      </button>
      <button
        @click="status_filter = DeliveryStatus.InProgress"
        class="transition duration-200 ease-in-out px-4 text-[15px] border-2 rounded-xl hover:bg-black/10 whitespace-nowrap"
        :class="{ 'bg-black/10': status_filter == DeliveryStatus.InProgress }"
      >
        In Progress
      </button>
      <button
        @click="status_filter = DeliveryStatus.Returned"
        class="transition duration-200 ease-in-out px-4 text-[15px] border-2 rounded-xl hover:bg-black/10"
        :class="{ 'bg-black/10': status_filter == DeliveryStatus.Returned }"
      >
        Returned
      </button>
      <button
        @click="status_filter = DeliveryStatus.Canceled"
        class="transition duration-200 ease-in-out px-4 text-[15px] border-2 rounded-xl hover:bg-black/10"
        :class="{ 'bg-black/10': status_filter == DeliveryStatus.Canceled }"
      >
        Canceled
      </button>
    </div>

    <div class="flex p-6 min-w-[35rem] bg-white rounded-xl shadow-md">
      <div class="w-full flex flex-col gap-y-6">
        <input
          class="transition duration-300 ease-in-out bg-[url(/icons/search.svg)] bg-no-repeat bg-[position:98%_60%] w-64 rounded-md border-0 bg-gray-100 text-sm focus:ring-2 focus:ring-slate-300"
          v-model="search_filter"
          type="text"
          placeholder="Search"
        />
        <div
          v-if="filtered_orders.length == 0"
          class="flex flex-col justify-center items-center gap-y-8 h-auto"
        >
          <div
            class="size-36 bg-no-repeat bg-center bg-contain bg-[url(/icons/order.svg)] contrast-0"
          ></div>
          <div class="text-xl text-gray-500">You don't have an order yet</div>
        </div>
        <div
          v-else
          v-for="order in filtered_orders"
          class="flex flex-col gap-x-8 bg-white rounded-xl border"
        >
          <div class="flex justify-between p-5 bg-black/5 rounded-t-xl">
            <div class="flex flex-col">
              <div class="text-gray-600">Order date</div>
              <div class="text-lg">
                {{ new Date(order.date).toLocaleString("tr-TR", options) }}
              </div>
            </div>
            <div class="flex flex-col">
              <div class="text-gray-600">Customer</div>
              <div class="text-lg">
                {{ order.customer_name }}
              </div>
            </div>
            <div class="flex flex-col">
              <div class="text-gray-600">Total amount</div>
              <div class="text-lg">
                {{ order.price.toLocaleString("tr-TR") }}
                TL
              </div>
            </div>
            <div class="flex flex-col">
              <div class="text-gray-600">Total products</div>
              <div class="text-lg">
                {{
                  order.products.reduce((total: number, product: Product) => {
                    return total + product.quantity;
                  }, 0)
                }}
              </div>
            </div>
            <NuxtLink
              :to="'orders/' + order.id"
              class="transition duration-300 ease-in-out flex justify-center items-center w-20 h-12 rounded-xl bg-black text-white hover:bg-black/80 font-medium"
              >Details</NuxtLink
            >
          </div>
          <div class="flex items-center px-5 pt-4 gap-x-4">
            <div :class="status_names[order.delivery_status].icon"></div>
            {{ status_names[order.delivery_status].name }}
          </div>
          <div class="grid grid-cols-auto_box_orders items-center p-5 gap-4">
            <NuxtLink
              v-for="product in order.products"
              :to="'/' + product.url"
              class="flex justify-end items-end min-w-24 size-24 border border-gray-300 rounded-lg bg-center bg-contain bg-no-repeat"
              :style="
                'background-image: url(/images/products/' + product.image + ');'
              "
              ><span
                class="-mb-3.5 -mr-3.5 flex justify-center items-center size-7 rounded-full bg-gray-300 font-medium"
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
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
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
    name: "Delivered",
    icon: "bg-center bg-no-repeat size-6 bg-[url(/icons/check.svg)]",
  },
  {
    name: "In Progress",
    icon: "bg-center bg-no-repeat size-6 size-6 bg-[url(/icons/progress.svg)]",
  },
  {
    name: "Returned",
    icon: "bg-center bg-no-repeat size-6 size-6 bg-[url(/icons/return.svg)]",
  },
  {
    name: "Canceled",
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
const orders = ref<Order[]>([]);
const status_filter = ref<DeliveryStatus | null>(null);
const search_filter = ref("");
onMounted(() => {
  nextTick(async () => {
    await useFetch<Order[]>(config.apiBase + "/orders", {
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
  if (status_filter.value) {
    return orders.value
      .filter((order: Order) => {
        const filtered_products = order.products.filter((product) => {
          return product.title
            .toLowerCase()
            .includes(search_filter.value.toLowerCase());
        });
        return filtered_products.length > 0;
      })
      .filter((order: Order) => order.delivery_status == status_filter.value);
  } else {
    return orders.value.filter((order: Order) => {
      const filtered_products = order.products.filter((product) => {
        return product.title
          .toLowerCase()
          .includes(search_filter.value.toLowerCase());
      });
      return filtered_products.length > 0;
    });
  }
});
</script>
