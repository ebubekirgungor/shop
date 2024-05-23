<template>
  <main class="flex justify-center gap-x-4 mb-20 m-4">
    <div
      class="flex flex-col gap-y-4 w-[clamp(48rem,65rem,65rem)] sm:min-w-[48rem]"
    >
      <div
        class="flex justify-center sm:justify-start items-center gap-x-4 p-6 text-xl h-20 bg-white sm:rounded-xl sm:shadow-md"
      >
        <button
          @click="step == 1 ? navigateTo(localePath('cart')) : (step = 1)"
          class="absolute sm:relative left-4 sm:left-0 transition duration-200 ease-in-out bg-no-repeat bg-center bg-[url(/icons/previous.svg)] size-10 bg-black/5 rounded-full hover:bg-black/10"
        ></button>
        {{ step == 1 ? $t("select_address") : $t("payment") }}
      </div>
      <div class="flex sm:p-6 h-auto bg-white sm:rounded-xl sm:shadow-md">
        <div v-if="step == 1" class="w-full">
          <div v-if="fetch_complete" class="flex flex-wrap gap-6">
            <div
              v-for="address in addresses"
              @click="selected_address = address"
              class="transition duration-300 ease-in-out flex flex-col p-4 text-xl size-full sm:size-52 ring-1 ring-slate-300 rounded-xl gap-y-2 cursor-pointer"
              :class="{ 'ring-4': selected_address?.ID == address.ID }"
            >
              <div class="flex">
                <div class="grow w-0 break-words font-medium">
                  {{ address.title }}
                </div>
              </div>
              <div class="text-sm text-gray-500">
                {{ address.customer_name }}
              </div>
              <div class="text-sm select-text">{{ address.address }}</div>
            </div>
          </div>
          <div v-else class="flex items-center justify-center w-full">
            <div
              class="animate-spin size-24 bg-[url(/icons/loading.svg)] bg-no-repeat bg-cover"
            ></div>
          </div>
        </div>
        <div v-else class="flex flex-col sm:flex-row gap-6 w-full">
          <div
            class="flex flex-col justify-between p-6 bg-sky-950/95 rounded-xl sm:w-[27rem] sm:min-w-[27rem] h-48 sm:h-60 text-white text-2xl sm:text-3xl font-[OCRA]"
          >
            <div
              class="-mt-1 size-10 sm:size-16 bg-[url(/icons/chip.png)] bg-no-repeat bg-cover"
            ></div>
            <div class="flex flex-col gap-y-4">
              <span v-if="!card.number">**** **** **** ****</span
              >{{ card.number }}
              <div class="flex justify-between text-[19px]">
                <div class="font-[Consolas]">{{ card.name }}</div>
                <div>
                  <span v-if="!card.month">{{ $t("MM") }}</span
                  >{{ card.month }}/<span v-if="!card.year">YY</span
                  >{{ card.year }}
                </div>
              </div>
            </div>
          </div>
          <div class="flex flex-col gap-y-3 w-full">
            <div class="flex flex-col gap-y-2">
              <label for="number">{{ $t("card_number") }}</label>
              <input
                id="number"
                class="transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300 text-center w-12 h-10 !w-auto !text-left"
                type="text"
                v-model="card.number"
                @input="card_number_format"
                maxlength="19"
                placeholder="0123 4567 8910 0000"
              />
            </div>
            <div class="flex flex-col gap-y-2">
              <label for="holder">{{ $t("card_holder") }}</label>
              <input
                id="holder"
                class="transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300 text-center w-12 h-10 !w-auto !text-left"
                type="text"
                v-model="card.name"
                @input="card_name_format"
              />
            </div>
            <div class="flex justify-between">
              <div class="flex flex-col gap-y-2">
                <label>{{ $t("expiration_date") }}</label>
                <div class="flex gap-x-4">
                  <input
                    class="transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300 text-center w-12 h-10"
                    type="text"
                    v-model="card.month"
                    maxlength="2"
                    :placeholder="$t('MM')"
                    @input="card.month = exp_cvv_format(card.month)"
                  />
                  <input
                    class="transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300 text-center w-12 h-10"
                    type="text"
                    v-model="card.year"
                    maxlength="2"
                    placeholder="YY"
                    @input="card.year = exp_cvv_format(card.year)"
                  />
                </div>
              </div>
              <div class="flex flex-col gap-y-2">
                <label for="cvv">CVV</label>
                <input
                  id="cvv"
                  class="transition duration-300 ease-in-out rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300 text-center w-12 h-10 !w-14"
                  type="text"
                  v-model="card.cvv"
                  maxlength="3"
                  placeholder="123"
                  @input="card.cvv = exp_cvv_format(card.cvv)"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <button
      class="transition-margin duration-300 ease-in-out z-10 peer absolute left-1 bottom-0 focus:mb-20 sm:hidden min-w-7 h-20 bg-[url(/icons/expand_less.svg)] focus:bg-[url(/icons/expand_more.svg)] bg-center bg-contain bg-no-repeat"
    ></button>
    <div
      class="transition-height duration-300 ease-in-out peer-focus:h-40 fixed sm:relative bottom-0 w-full sm:w-auto h-20 sm:!h-fit flex sm:flex-col gap-x-2 gap-y-5 pl-2 p-4 sm:p-6 sm:min-w-64 bg-white sm:rounded-xl border-t sm:border-t-0 shadow-[0_0_15px_rgba(0,0,0,0.2)] sm:shadow-md"
    >
      <div class="pl-8 sm:pl-0 w-full flex flex-col sm:gap-y-5">
        <div class="sm:text-[17px]">{{ $t("total_amount") }}</div>
        <div class="flex gap-x-1 text-xl sm:text-3xl">
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
          <div class="sm:text-[20px] sm:mt-0.5">TL</div>
        </div>
      </div>
      <div class="flex flex-col gap-y-5">
        <div
          class="w-full absolute sm:relative left-0 top-20 sm:top-0 bottom-0 p-4 sm:p-0 order-1 sm:order-none flex flex-col gap-y-1"
        >
          <div class="flex justify-between">
            <div>{{ $t("subtotal") }}</div>
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
            <div>{{ $t("shipping") }}</div>
            <div>{{ shipping }} TL</div>
          </div>
        </div>
      </div>
      <button
        v-if="step == 1"
        @click="step = 2"
        :disabled="!selected_address"
        class="transition duration-300 ease-in-out min-w-36 h-12 col-span-2 rounded-xl bg-black text-white sm:hover:bg-black/80 font-medium sm:mt-4 disabled:bg-black/60 disabled:pointer-events-none"
      >
        {{ $t("continue") }}
      </button>
      <button
        v-else
        @click="order"
        :disabled="
          card.number.length < 19 ||
          !card.name ||
          !card.month ||
          card.month.length < 2 ||
          parseInt(card.month) > 12 ||
          parseInt(card.month) < 1 ||
          !card.year ||
          card.year.length < 2 ||
          parseInt(card.year) <
            parseInt(new Date().getFullYear().toString().substring(2, 4)) ||
          !card.cvv ||
          card.cvv.length < 3
        "
        class="transition duration-300 ease-in-out min-w-36 h-12 col-span-2 rounded-xl bg-black text-white sm:hover:bg-black/80 font-medium sm:mt-4 disabled:bg-black/60 disabled:pointer-events-none"
      >
        {{ $t("place_order") }}
      </button>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useToast } from "vue-toastification";
const toast = useToast();
const config = useRuntimeConfig().public;
const localePath = useLocalePath();
const { t } = useI18n();
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

interface Address {
  ID: number | null;
  title: string;
  customer_name: string;
  address: string;
}

interface Card {
  number: string;
  name: string;
  month: string;
  year: string;
  cvv: string;
}

const cart = ref<Cart[]>([]);
const shipping = ref(50);
const addresses = ref<Address[]>([]);
const selected_address = ref<Address>();
const fetch_complete = ref(false);
const step = ref(1);

const card = ref<Card>({
  number: "",
  name: "",
  month: "",
  year: "",
  cvv: "",
});

onMounted(() => {
  nextTick(async () => {
    await useFetch(
      role.value != undefined
        ? config.apiBase + "/users/cart"
        : config.apiBase + "/cart",
      {
        headers: {
          Authorization: config.apiKey,
        },
        onResponse({ response }) {
          cart.value = response._data;
        },
      }
    );
    await useFetch(config.apiBase + "/addresses", {
      headers: {
        Authorization: config.apiKey,
      },
      onResponse({ response }) {
        if (response._data.addresses) {
          addresses.value = response._data.addresses;
          fetch_complete.value = true;
        }
      },
    });
  });
});

const order = async () => {
  await useFetch(config.apiBase + "/orders", {
    headers: {
      Authorization: config.apiKey,
    },
    method: "post",
    body: {
      delivery_address: selected_address.value?.address,
      customer_name: selected_address.value?.customer_name,
    },
    onResponse({ response }) {
      if (response._data.ID) {
        toast.success(t("order_placed"), {
          bodyClassName: "toast-font",
        });
      }
    },
  });
};

const card_number_format = () => {
  card.value.number = card.value.number
    .replace(/[^0-9]/gi, "")
    .replace(/(\d{4})(?=\d)/g, "$1 ");
};

const card_name_format = () => {
  card.value.name = card.value.name
    .toLocaleUpperCase("tr-TR")
    .replace(/[^\p{L}\s-]/u, "");
};

const exp_cvv_format = (value: string) => {
  return value.replace(/[^0-9]/gi, "");
};
</script>
<style>
@font-face {
  font-family: OCRA;
  src: url(/fonts/ocr.ttf);
}
</style>
