<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <transition name="background" mode="out-in">
      <div
        v-if="add_dialog || edit_dialog || delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 size-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="add_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed"
      >
        <div
          class="flex flex-col justify-center sm:justify-start sm:p-3 bg-white sm:-mt-[10vh] w-screen sm:w-[25vw] sm:min-w-[28rem] h-screen sm:h-auto rounded-xl"
        >
          <button
            @click="add_dialog = false"
            class="absolute sm:relative top-0 mt-4 mr-4 sm:m-0 self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-cover sm:bg-auto bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl sm:grow">
            {{ $t("add_new_address") }}
          </h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-4 sm:mx-8"
            @submit.prevent="create"
          >
            <input
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              :placeholder="$t('title')"
              v-model="new_address.title"
              required
            />
            <input
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              :placeholder="$t('customer_name')"
              v-model="new_address.customer_name"
              required
            />
            <textarea
              rows="10"
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              :placeholder="$t('address')"
              v-model="new_address.address"
              required
            ></textarea>
            <button
              :disabled="!new_address.title || !new_address.address"
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              {{ $t("create") }}
            </button>
          </form>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="edit_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed"
      >
        <div
          class="flex flex-col justify-center sm:justify-start sm:p-3 bg-white sm:-mt-[10vh] w-screen sm:w-[25vw] sm:min-w-[28rem] h-screen sm:h-auto rounded-xl"
        >
          <button
            @click="edit_dialog = false"
            class="absolute sm:relative top-0 mt-4 mr-4 sm:m-0 self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-cover sm:bg-auto bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl sm:grow">{{ $t("edit_address") }}</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-4 sm:mx-8"
            @submit.prevent="update"
          >
            <input type="hidden" v-model="edit_address.ID" required />
            <input
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              :placeholder="$t('title')"
              v-model="edit_address.title"
              required
            />
            <input
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              :placeholder="$t('customer_name')"
              v-model="edit_address.customer_name"
              required
            />
            <textarea
              rows="10"
              class="transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300"
              type="text"
              :placeholder="$t('address')"
              v-model="edit_address.address"
              required
            ></textarea>
            <button
              :disabled="!edit_address.title || !edit_address.address"
              type="submit"
              class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
            >
              {{ $t("update") }}
            </button>
          </form>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="delete_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 size-full fixed"
      >
        <div
          class="flex flex-col p-2 bg-white sm:-mt-48 w-80 h-auto rounded-xl z-3"
        >
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out size-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl">{{ $t("delete_address") }}</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_address_id" required />
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
      {{ $t("addresses") }}
    </div>
    <div
      class="flex sm:p-6 sm:min-w-[35rem] sm:min-h-[17rem] bg-white sm:rounded-xl sm:shadow-md"
    >
      <div
        v-if="fetch_complete"
        class="w-full flex flex-col sm:flex-row flex-wrap gap-6"
      >
        <div
          @click="add_dialog = true"
          class="flex flex-col p-4 shadow-md text-xl sm:size-56 rounded-xl transition duration-200 ease-in-out items-center justify-center cursor-pointer bg-gray-200 hover:bg-gray-300"
        >
          <div
            class="size-24 bg-[url(/icons/add.svg)] bg-no-repeat bg-cover"
          ></div>
        </div>
        <div
          v-for="address in addresses"
          class="flex flex-col p-4 shadow-md text-xl sm:size-56 rounded-xl bg-gray-50 gap-y-2"
        >
          <div class="flex">
            <div class="grow w-0 break-words font-medium">
              {{ address.title }}
            </div>
            <button
              @click="
                open_edit_dialog(
                  address.ID,
                  address.title,
                  address.customer_name,
                  address.address
                )
              "
              class="mr-1 transition duration-300 ease-in-out size-9 bg-[url(/icons/edit.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
            ></button>
            <button
              @click="open_delete_dialog(address.ID)"
              class="transition duration-300 ease-in-out size-9 bg-[url(/icons/delete.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
            ></button>
          </div>
          <div class="text-sm text-gray-500">{{ address.customer_name }}</div>
          <div class="text-sm select-text">{{ address.address }}</div>
        </div>
      </div>
      <div v-else class="flex items-center justify-center w-full">
        <div
          class="animate-spin size-24 bg-[url(/icons/loading.svg)] bg-no-repeat bg-cover"
        ></div>
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
interface Address {
  ID: number | null;
  title: string;
  customer_name: string;
  address: string;
}
interface NewAddress {
  title: string;
  customer_name: string;
  address: string;
}
const fetch_complete = ref(false);
const add_dialog = ref(false);
const edit_dialog = ref(false);
const delete_dialog = ref(false);
const open_edit_dialog = (
  id: number | null,
  title: string,
  customer_name: string,
  address: string
) => {
  edit_address.value = {
    ID: id,
    title: title,
    customer_name: customer_name,
    address: address,
  };
  edit_dialog.value = true;
};
const open_delete_dialog = (id: number | null) => {
  delete_address_id.value = id;
  delete_dialog.value = true;
};
const addresses = ref<Address[]>([]);
const new_address = ref<NewAddress>({
  title: "",
  customer_name: "",
  address: "",
});
const edit_address = ref<Address>({
  ID: null,
  title: "",
  customer_name: "",
  address: "",
});
const delete_address_id = ref<number | null>(null);
onMounted(() => {
  nextTick(async () => {
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
const create = async () => {
  await useFetch(config.apiBase + "/addresses", {
    headers: {
      Authorization: config.apiKey,
    },
    method: "post",
    body: {
      title: new_address.value.title,
      customer_name: new_address.value.customer_name,
      address: new_address.value.address,
    },
    onResponse({ response }) {
      if (response._data.title) {
        add_dialog.value = false;
        addresses.value.unshift(response._data);
        new_address.value = {
          title: "",
          customer_name: "",
          address: "",
        };
        toast.success(t("address_created"), {
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
const update = async () => {
  await useFetch(config.apiBase + "/addresses/" + edit_address.value.ID, {
    headers: {
      Authorization: config.apiKey,
    },
    method: "patch",
    body: {
      title: edit_address.value.title,
      customer_name: edit_address.value.customer_name,
      address: edit_address.value.address,
    },
    onResponse({ response }) {
      if (response._data.title) {
        addresses.value[
          addresses.value
            .map((address: Address) => address.ID)
            .indexOf(edit_address.value.ID)
        ] = edit_address.value;
        edit_dialog.value = false;
        edit_address.value = {
          ID: null,
          title: "",
          customer_name: "",
          address: "",
        };
        toast.success(t("address_updated"), {
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
const remove = async () => {
  await useFetch(config.apiBase + "/addresses/" + delete_address_id.value, {
    headers: {
      Authorization: config.apiKey,
    },
    method: "delete",
    onResponse({ response }) {
      if (response._data) {
        addresses.value.splice(
          addresses.value
            .map((address: Address) => address.ID)
            .indexOf(delete_address_id.value),
          1
        );
        delete_dialog.value = false;
        delete_address_id.value = null;
        toast.success(t("address_removed"), {
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
