<template>
  <main class="flex justify-center m-4">
    <div
      class="flex flex-col sm:flex-row sm:w-[clamp(60rem,70rem,70rem)] sm:min-w-[60rem] h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      <div
        v-if="images.length > 0"
        class="flex flex-col gap-y-4 pb-4 min-w-screen sm:min-w-[35rem] rounded-tl-xl relative overflow-hidden touch-pan-y"
      >
        <div
          class="transition-transform duration-500 ease-in-out flex w-screen h-[100vw] sm:size-[35rem]"
          ref="slider"
          @touchstart="touch_start"
          @touchmove="touch_move"
        >
          <div v-for="image in images" style="flex: 0 0 100%">
            <img
              :src="image.url"
              class="object-contain min-w-full min-h-[100vw] max-w-full max-h-[100vw] sm:min-w-[35rem] sm:min-h-[35rem] sm:max-w-[35rem] sm:max-h-[35rem]"
              :alt="image.name"
            />
          </div>
        </div>
        <button
          v-if="images.length > 1"
          @click="goto_slide(slide_index - 1)"
          class="transition duration-300 ease-in-out absolute top-[50vw] sm:top-[17.5rem] -translate-y-1/2 bg-no-repeat bg-center size-10 bg-gray-200 sm:bg-white rounded-full sm:hover:bg-gray-200 disabled:pointer-events-none disabled:opacity-0 left-3 bg-[url(/icons/previous.svg)]"
        ></button>
        <button
          v-if="images.length > 1"
          @click="goto_slide(slide_index + 1)"
          class="transition duration-300 ease-in-out absolute top-[50vw] sm:top-[17.5rem] -translate-y-1/2 bg-no-repeat bg-center size-10 bg-gray-200 sm:bg-white rounded-full sm:hover:bg-gray-200 disabled:pointer-events-none disabled:opacity-0 right-3 bg-[url(/icons/next.svg)]"
        ></button>
        <div
          v-if="images.length > 1"
          class="flex justify-center gap-x-2 sm:gap-x-4 mx-4"
        >
          <button
            v-for="(image, index) in images"
            class="flex justify-center transition duration-200 ease-in-out size-3 sm:size-12 border border-gray-400 sm:border-0 bg-white rounded-full sm:rounded-md"
            :class="{
              '!bg-gray-400 sm:!bg-white sm:ring sm:ring-3 sm:ring-black':
                index == slide_index,
            }"
            @click="goto_slide(index)"
          >
            <img
              class="hidden sm:block h-full object-center object-cover rounded-md"
              :src="image.url"
            />
          </button>
        </div>
      </div>
      <div v-else class="bg-gray-200 sm:rounded-l-xl">
        <div
          class="size-[100vw] sm:size-[35rem] bg-no-repeat bg-center bg-contain bg-[url(/icons/order.svg)] contrast-0"
        ></div>
      </div>
      <div class="hidden sm:flex flex-col w-full gap-y-6 p-6">
        <div class="flex justify-between items-center">
          <div class="text-xl select-text">{{ product.title }}</div>
          <button
            v-if="role != undefined"
            @click="toggle_favorite"
            class="transition duration-200 ease-in-out flex justify-center items-center size-12 bg-black/5 rounded-full hover:bg-black/10"
          >
            <div
              class="bg-no-repeat bg-center size-7 bg-contain"
              :class="
                product.is_favorite
                  ? 'bg-[url(/icons/favorite_filled.svg)]'
                  : 'bg-[url(/icons/favorite.svg)] contrast-0'
              "
            ></div>
            <div
              v-if="animate"
              class="absolute bg-no-repeat bg-center size-7 bg-contain animate-ping"
              :class="
                product.is_favorite
                  ? 'bg-[url(/icons/favorite_filled.svg)]'
                  : 'bg-[url(/icons/favorite.svg)] contrast-0'
              "
            ></div>
          </button>
        </div>
        <div class="flex gap-x-2 text-5xl">
          {{ product.list_price.toLocaleString("tr-TR") }}
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
                parseInt(product.stock_quantity) == 0 ||
                quantity == parseInt(product.stock_quantity) ||
                product_cart_quantity + quantity >= product.stock_quantity
              "
              class="transition duration-300 ease-in-out bg-black size-12 rounded-r-full pr-1 hover:bg-black/80 disabled:pointer-events-none"
            >
              +
            </button>
          </div>
          <button
            class="transition duration-300 ease-in-out w-full h-full sm:h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 font-medium disabled:bg-black/60 disabled:pointer-events-none"
            :disabled="
              parseInt(product.stock_quantity) == 0 ||
              product_cart_quantity >= product.stock_quantity
            "
            @click="add_to_cart()"
          >
            {{ $t("add_to_cart_capital") }}
          </button>
        </div>
      </div>
      <div class="sm:hidden mb-20 p-4 flex justify-between items-center">
        <div class="text-xl select-text">{{ product.title }}</div>
        <button
          v-if="role != undefined"
          @click="toggle_favorite"
          class="transition duration-200 ease-in-out flex justify-center items-center size-12 bg-black/5 rounded-full hover:bg-black/10"
        >
          <div
            class="bg-no-repeat bg-center size-7 bg-contain"
            :class="
              product.is_favorite
                ? 'bg-[url(/icons/favorite_filled.svg)]'
                : 'bg-[url(/icons/favorite.svg)] contrast-0'
            "
          ></div>
          <div
            v-if="animate"
            class="absolute bg-no-repeat bg-center size-7 bg-contain animate-ping"
            :class="
              product.is_favorite
                ? 'bg-[url(/icons/favorite_filled.svg)]'
                : 'bg-[url(/icons/favorite.svg)] contrast-0'
            "
          ></div>
        </button>
      </div>
      <div
        class="sm:hidden flex justify-between items-center gap-x-16 fixed bottom-0 w-full h-20 p-4 bg-white border-t shadow-[0_0_15px_rgba(0,0,0,0.2)]"
      >
        <div class="flex gap-x-2 text-xl">
          {{ product.list_price.toLocaleString("tr-TR") }}
          <div>TL</div>
        </div>
        <button
          class="transition duration-300 ease-in-out w-full h-full sm:h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 font-medium disabled:bg-black/60 disabled:pointer-events-none"
          :disabled="
            parseInt(product.stock_quantity) == 0 ||
            product_cart_quantity >= product.stock_quantity
          "
          @click="add_to_cart()"
        >
          {{ $t("add_to_cart_capital") }}
        </button>
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
const route = useRoute();
const role = useCookie<number>("role");

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
  is_favorite: boolean;
}

interface Cart {
  id: number | null;
  quantity: number;
  selected: boolean;
}

const product = ref<Product>({
  ID: null,
  title: "",
  category_id: "",
  list_price: 0,
  stock_quantity: "",
  is_favorite: false,
});

const images = ref<Image[]>([]);
const startX = ref(0);
const dragging = ref(false);
const slide_index = ref(0);
const slider = ref<any>();
const quantity = ref(1);
const cart_unregistered = useCookie<Cart[]>("cart");
if (!cart_unregistered.value) cart_unregistered.value = [];
const cart = ref<Cart[]>([]);
const product_cart_quantity = ref();
const animate = ref(false);

onMounted(() => {
  nextTick(async () => {
    await useFetch(
      config.apiBase +
        "/products/" +
        route.params.product.toString().match(/-(\d+)$/)![1],
      {
        onResponse({ response }) {
          history.replaceState({}, "", response._data.url);
          product.value = response._data;
          if (parseInt(product.value.stock_quantity) == 0) quantity.value = 0;
          const images_array: Image[] = [];
          response._data.images.forEach((image: Image) => {
            images_array.push({
              order: image.order,
              name: image.name,
              url: "/images/products/" + image.name,
            });
          });
          images.value = images_array;
        },
      }
    );
    if (role.value != undefined) {
      await useFetch(config.apiBase + "/users", {
        headers: {
          Authorization: config.apiKey,
        },
        onResponse({ response }) {
          cart.value = response._data.cart;
          product_cart_quantity.value =
            cart.value[
              cart.value.findIndex((item: Cart) => item.id == product.value.ID)
            ].quantity;
          if (product_cart_quantity.value >= product.value.stock_quantity)
            quantity.value = 0;
        },
      });
    } else {
      cart.value = cart_unregistered.value;
      if (cart_unregistered.value.length > 0) {
        product_cart_quantity.value =
          cart.value[
            cart.value.findIndex((item: Cart) => item.id == product.value.ID)
          ].quantity;
        if (product_cart_quantity.value >= product.value.stock_quantity)
          quantity.value = 0;
      }
    }
  });
});

const goto_slide = async (index: number) => {
  if (index < 0) {
    slide_index.value = images.value.length - 1;
  } else if (index >= images.value.length) {
    slide_index.value = 0;
  } else {
    slide_index.value = index;
  }
};

const touch_start = async (event: any) => {
  dragging.value = true;
  startX.value = event.touches[0].clientX;
};

const touch_move = async (event: any) => {
  if (!dragging.value) return;
  const currentX = event.touches[0].clientX;
  const diffX = currentX - startX.value;
  if (Math.abs(diffX) > 50) {
    goto_slide(slide_index.value + (diffX > 0 ? -1 : 1));
    dragging.value = false;
  }
};

watch(slide_index, () => {
  const transformValue = `translateX(-${slide_index.value * 100}%)`;
  slider.value.style.transform = transformValue;
});

const add_to_cart = async () => {
  const existing_index = cart.value.findIndex(
    (item: Cart) => item.id == product.value.ID
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
    await useFetch(config.apiBase + "/users/cart", {
      headers: {
        Authorization: config.apiKey,
      },
      method: "patch",
      body: {
        cart: JSON.parse(JSON.stringify(cart.value)),
      },
      onResponse({ response }) {
        if (response._data == "ok") {
          toast.success(t("added_to_cart"), {
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
    toast.success(t("added_to_cart"), {
      bodyClassName: "toast-font",
    });
    product_cart_quantity.value >= product.value.stock_quantity
      ? (quantity.value = 0)
      : (quantity.value = 1);
  }
};

const toggle_favorite = async () => {
  await useFetch(config.apiBase + "/favorites/" + product.value.ID, {
    headers: {
      Authorization: config.apiKey,
    },
    method: product.value.is_favorite ? "delete" : "post",
    onResponse({ response }) {
      if (response._data == "ok") {
        product.value.is_favorite = !product.value.is_favorite;
        if (product.value.is_favorite) {
          animate.value = true;
          setTimeout(() => {
            animate.value = false;
          }, 500);
        }
      }
    },
  });
};
</script>
