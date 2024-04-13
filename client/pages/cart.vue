<template>
  <div v-if="cart.length == 0" class="flex flex-col items-center mt-8 gap-y-12">
    <div
      class="size-36 bg-no-repeat bg-center bg-contain bg-[url(/icons/cart.svg)] contrast-0"
    ></div>
    <div class="text-3xl text-gray-600">Your cart is empty</div>
    <NuxtLink
      to="/"
      class="transition duration-200 ease-in-out flex justify-center items-center w-52 h-10 rounded-full text-lg border-2 border-black hover:bg-black/10"
      >Return to shopping</NuxtLink
    >
  </div>
  <main v-else class="flex justify-center gap-x-4 m-4">
    <div
      class="flex flex-col gap-y-4 w-[clamp(40rem,65rem,65rem)] min-w-[40rem]"
    >
      <div
        class="flex gap-x-2 p-6 text-xl h-auto bg-white rounded-xl shadow-md"
      >
        Cart
        <div class="text-[16px] text-gray-500 mt-0.5">
          ({{ items_count + (items_count == 1 ? " item" : " items") }})
        </div>
      </div>
      <div class="flex h-auto bg-white rounded-xl shadow-md">
        <div class="flex flex-col w-full min-w-[40rem] divide-y">
          <div v-for="product in cart" class="flex justify-between p-6">
            <div class="flex items-center gap-x-4">
              <input
                v-model="product.cart.selected"
                @click="update_cart_unregistered"
                type="checkbox"
                class="transition duration-200 ease-in-out size-5 cursor-pointer rounded-md border-gray-300 text-gray-800 hover:border-gray-500 focus:ring-0 focus:ring-offset-0"
              />
              <img
                v-if="product.images.length > 0"
                class="min-w-24 size-24 border border-gray-300 rounded-lg object-contain"
                :src="'/images/products/' + product.images[0].name"
              />
              <div v-else class="size-24 border border-gray-300 rounded-lg">
                <div
                  class="size-24 bg-no-repeat bg-center bg-contain bg-[url(/icons/order.svg)] contrast-75 invert"
                ></div>
              </div>
              <NuxtLink :to="product.url" class="font-medium w-48">
                {{ product.title }}
              </NuxtLink>
            </div>
            <div class="flex justify-between items-center w-96">
              <div class="flex text-xl border border-gray-300 rounded-full">
                <button
                  @click="
                    product.cart.quantity--;
                    update_cart_unregistered();
                  "
                  :disabled="product.cart.quantity == 1"
                  class="transition duration-300 ease-in-out size-8 rounded-l-full pl-1 hover:bg-black/10 disabled:pointer-events-none"
                >
                  -
                </button>
                <div
                  class="flex justify-center items-center text-black text-sm size-8"
                >
                  {{ product.cart.quantity }}
                </div>
                <button
                  @click="
                    product.cart.quantity++;
                    update_cart_unregistered();
                  "
                  :disabled="product.cart.quantity == product.stock_quantity"
                  class="transition duration-300 ease-in-out size-8 rounded-r-full hover:bg-black/10 disabled:pointer-events-none"
                >
                  +
                </button>
              </div>
              <div class="font-medium text-lg">
                {{
                  (product.list_price * product.cart.quantity).toLocaleString(
                    "tr-TR",
                    {
                      maximumFractionDigits: 2,
                    }
                  )
                }}
                TL
              </div>
              <button
                @click="remove_from_cart(product.id)"
                class="transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/delete.svg)] size-10 bg-white rounded-full hover:bg-black/10"
              ></button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div
      class="flex flex-col gap-y-5 p-6 min-w-64 h-fit bg-white rounded-xl shadow-md"
    >
      <div class="text-[17px]">
        Selected items ({{
          cart.reduce((total: number, product: Cart) => {
            return total + (product.cart.selected ? product.cart.quantity : 0);
          }, 0)
        }})
      </div>
      <div class="flex gap-x-1 text-3xl">
        {{
          (
            cart.reduce((total: number, product: Cart) => {
              return (
                total +
                product.list_price *
                  (product.cart.selected ? product.cart.quantity : 0)
              );
            }, 0) + shipping
          ).toLocaleString("tr-TR", {
            maximumFractionDigits: 2,
          })
        }}
        <div class="text-[20px] mt-0.5">TL</div>
      </div>
      <div class="flex flex-col gap-y-1">
        <div class="flex justify-between">
          <div>Subtotal</div>
          <div>
            {{
              cart
                .reduce((total: number, product: Cart) => {
                  return (
                    total +
                    product.list_price *
                      (product.cart.selected ? product.cart.quantity : 0)
                  );
                }, 0)
                .toLocaleString("tr-TR", {
                  maximumFractionDigits: 2,
                })
            }}
            TL
          </div>
        </div>
        <div class="flex justify-between">
          <div>Shipping</div>
          <div>{{ shipping }} TL</div>
        </div>
      </div>
      <button
        @click="checkout"
        :disabled="
          cart.reduce((total: number, product: Cart) => {
            return total + (product.cart.selected ? 1 : 0);
          }, 0) == 0"
        class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-xl bg-black text-white hover:bg-black/80 font-medium mt-4 disabled:bg-black/60 disabled:pointer-events-none"
      >
        Checkout
      </button>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const role = useCookie("role");
interface Image {
  order: number;
  name: string;
  url: string;
}
interface Cart {
  id: number;
  title: string;
  url: string;
  list_price: number;
  stock_quantity: number;
  images: Image[];
  cart: {
    id: number;
    quantity: number;
    selected: boolean;
  };
}
const cart_unregistered = useCookie<Object[]>("cart");
if (!cart_unregistered.value) cart_unregistered.value = [];
const cart = ref<Cart[]>([]);
let cart_copy = "";
const items_count = computed(() => {
  return cart.value.reduce((total: number, product: Cart) => {
    return total + product.cart.quantity;
  }, 0);
});
const shipping = ref(50);
onMounted(() => {
  nextTick(async () => {
    await useFetch<Cart[]>(
      role.value != undefined
        ? config.apiBase + "/users/cart"
        : config.apiBase + "/cart",
      {
        headers: {
          Authorization: config.apiKey,
        },
        onResponse({ response }) {
          cart.value = response._data;
          cart_copy = JSON.stringify(response._data);
        },
      }
    );
  });
  if (role.value != undefined) {
    setInterval(async () => {
      if (JSON.stringify(cart.value) != cart_copy) {
        if (await update_cart()) cart_copy = JSON.stringify(cart.value);
      }
    }, 10000);
  }
});
const update_cart = async () => {
  const { data } = await useFetch<string>(config.apiBase + "/users/cart", {
    headers: {
      Authorization: config.apiKey,
    },
    method: "patch",
    body: {
      cart: JSON.parse(
        JSON.stringify(cart.value.map((item: Cart) => item.cart))
      ),
    },
  });
  return data.value == "ok";
};
const update_cart_unregistered = async () => {
  if (role.value == undefined) {
    setTimeout(() => {
      cart_unregistered.value = JSON.parse(
        JSON.stringify(cart.value.map((item: Cart) => item.cart))
      );
    }, 500);
  }
};
const remove_from_cart = async (product_id: number) => {
  cart.value.splice(
    cart.value.map((item: Cart) => item.id).indexOf(product_id),
    1
  );
  update_cart_unregistered();
};
const checkout = async () => {
  if (await update_cart()) navigateTo("/checkout");
};
</script>
