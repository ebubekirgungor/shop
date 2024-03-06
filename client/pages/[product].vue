<template>
  <main class="flex justify-center m-4">
    <div
      class="flex w-[clamp(60rem,70rem,70rem)] min-w-[60rem] h-auto bg-white rounded-xl shadow-md"
    >
      <div v-if="images.length > 0" class="flex flex-col gap-y-4 pb-4">
        <div
          class="overflow-hidden size-[35rem] relative mx-auto rounded-tl-xl"
        >
          <div
            v-for="(image, index) in images"
            class="transition duration-300 ease-in-out absolute inset-0"
            :class="{
              '-translate-x-full': index < current_slide,
              'translate-x-full': index > current_slide,
            }"
          >
            <img
              :src="image.url"
              class="object-contain relative size-full"
              alt="Product Image"
            />
          </div>
          <div class="flex justify-between items-center size-full p-4 absolute">
            <button
              :class="icon + 'bg-[url(/icons/previous.svg)]'"
              @click="current_slide--"
              :disabled="current_slide == 0"
            ></button>
            <button
              :class="icon + 'bg-[url(/icons/next.svg)]'"
              @click="current_slide++"
              :disabled="current_slide == images.length - 1"
            ></button>
          </div>
        </div>
        <div class="flex justify-center gap-x-4 mx-4">
          <button
            v-for="(image, index) in images"
            class="transition duration-200 ease-in-out size-12 rounded-md bg-no-repeat bg-cover"
            :class="{ 'ring ring-3 ring-black': index == current_slide }"
            :style="'background-image: url(' + image.url + ');'"
            @click="current_slide = index"
          ></button>
        </div>
      </div>
      <div v-else class="bg-gray-200 rounded-l-xl">
        <div
          class="size-[35rem] bg-no-repeat bg-center bg-contain bg-[url(/icons/order.svg)] contrast-0"
        ></div>
      </div>
      <div class="flex flex-col w-full gap-y-6 p-6">
        <div class="text-xl select-text">{{ product.title }}</div>
        <div class="flex gap-x-2 text-5xl">
          {{
            product.list_price.toLocaleString("tr-TR", {
              minimumFractionDigits: 2,
            })
          }}
          <div class="self-end text-2xl">TL</div>
        </div>
        <div class="flex gap-x-2">
          <div class="flex text-white text-2xl">
            <button
              @click="quantity--"
              :disabled="quantity <= 1"
              class="transition duration-300 ease-in-out bg-black size-12 rounded-l-full pl-2 hover:bg-black/80 disabled:pointer-events-none"
            >
              -
            </button>
            <div
              class="flex justify-center items-center text-black text-lg size-12 border-x border-4 border-black"
            >
              {{ quantity }}
            </div>
            <button
              @click="quantity++"
              :disabled="
                quantity == parseInt(product.stock_quantity) ||
                product_cart_quantity + quantity >= product.stock_quantity
              "
              class="transition duration-300 ease-in-out bg-black size-12 rounded-r-full pr-1 hover:bg-black/80 disabled:pointer-events-none"
            >
              +
            </button>
          </div>
          <button
            :class="button"
            :disabled="product_cart_quantity >= product.stock_quantity"
            @click="add_to_cart()"
          >
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
const route = useRoute();
import { useToast } from "vue-toastification";
const role = useCookie<number>("role");
const toast = useToast();
const product = ref<any>({
  ID: "",
  title: "",
  category_id: "",
  list_price: 0,
  stock_quantity: "",
});
const images = ref<any>([]);
const current_slide = ref(0);
const quantity = ref(1);
const cart_unregistered = useCookie<any>("cart");
if (!cart_unregistered.value) cart_unregistered.value = [];
const cart = ref<any>([]);
const product_cart_quantity = ref();
onMounted(() => {
  nextTick(async () => {
    await useFetch<any>("/api/products/" + route.params.product, {
      onResponse({ response }) {
        product.value = response._data;
        const images_array: any = [];
        response._data.images.forEach((image: any) => {
          images_array.push({
            order: image.order,
            name: image.name,
            url: "/images/products/" + image.name,
          });
        });
        images.value = images_array;
      },
    });
    if (role.value != undefined) {
      await useFetch<any>("/api/users", {
        onResponse({ response }) {
          cart.value = response._data.cart;
          product_cart_quantity.value =
            cart.value[
              cart.value.findIndex((item: any) => item.id == product.value.ID)
            ].quantity;
          if (product_cart_quantity.value >= product.value.stock_quantity)
            quantity.value = 0;
        },
      });
    } else {
      console.log(cart_unregistered.value);
      cart.value = cart_unregistered.value;
      if (cart_unregistered.value.length > 0) {
        product_cart_quantity.value =
          cart.value[
            cart.value.findIndex((item: any) => item.id == product.value.ID)
          ].quantity;
        if (product_cart_quantity.value >= product.value.stock_quantity)
          quantity.value = 0;
      }
    }
  });
});
const add_to_cart = async () => {
  const existing_index = cart.value.findIndex(
    (item: any) => item.id == product.value.ID
  );
  if (existing_index === -1) {
    cart.value.push({
      id: product.value.ID,
      quantity: quantity.value,
      selected: true,
    });
    product_cart_quantity.value = quantity.value;
  } else {
    cart.value[existing_index].quantity += quantity.value;
    product_cart_quantity.value += quantity.value;
  }
  if (role.value != undefined) {
    await useFetch("/api/users/cart", {
      method: "patch",
      body: {
        cart: JSON.parse(JSON.stringify(cart.value)),
      },
      onResponse({ response }) {
        if (response._data == "ok") {
          toast.success("Added to cart", {
            bodyClassName: "toast-font",
          });
          product_cart_quantity.value >= product.value.stock_quantity
            ? (quantity.value = 0)
            : (quantity.value = 1);
        } else {
          toast.warning(response._data.error, {
            bodyClassName: "toast-font",
          });
        }
      },
    });
  } else {
    cart_unregistered.value = cart.value;
    toast.success("Added to cart", {
      bodyClassName: "toast-font",
    });
    product_cart_quantity.value >= product.value.stock_quantity
      ? (quantity.value = 0)
      : (quantity.value = 1);
  }
};
const icon =
  "transition duration-300 ease-in-out bg-no-repeat bg-center size-10 bg-white rounded-full hover:bg-gray-200 disabled:pointer-events-none disabled:opacity-0 ";
const button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 font-medium disabled:bg-black/60 disabled:pointer-events-none";
</script>
