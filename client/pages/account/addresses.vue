<template>
  <main class="flex flex-col gap-y-4">
    <transition name="background" mode="out-in">
      <div
        v-if="add_dialog || edit_dialog || delete_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 w-full h-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="add_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 w-full h-full fixed"
      >
        <div
          class="flex flex-col p-3 bg-white -mt-48 w-[25vw] h-auto rounded-xl z-3"
        >
          <button
            @click="add_dialog = false"
            class="self-end transition duration-300 ease-in-out w-8 h-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl grow">Add new address</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="create"
          >
            <input
              :class="input"
              type="text"
              placeholder="Title"
              v-model="new_address.title"
              required
            />
            <textarea
              rows="10"
              :class="input"
              type="text"
              placeholder="Address"
              v-model="new_address.address"
              required
            ></textarea>
            <button
              :disabled="!new_address.title || !new_address.address"
              type="submit"
              :class="button"
            >
              Create
            </button>
          </form>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="edit_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 w-full h-full fixed"
      >
        <div
          class="flex flex-col p-3 bg-white -mt-48 w-[25vw] h-auto rounded-xl z-3"
        >
          <button
            @click="edit_dialog = false"
            class="self-end transition duration-300 ease-in-out w-8 h-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl grow">Edit address</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="update"
          >
            <input type="hidden" v-model="edit_address.ID" required />
            <input
              :class="input"
              type="text"
              placeholder="Title"
              v-model="edit_address.title"
              required
            />
            <textarea
              rows="10"
              :class="input"
              type="text"
              placeholder="Address"
              v-model="edit_address.address"
              required
            ></textarea>
            <button
              :disabled="!edit_address.title || !edit_address.address"
              type="submit"
              :class="button"
            >
              Update
            </button>
          </form>
        </div>
      </div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="delete_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 w-full h-full fixed"
      >
        <div
          class="flex flex-col p-3 bg-white -mt-48 w-[25vw] h-auto rounded-xl z-3"
        >
          <button
            @click="delete_dialog = false"
            class="self-end transition duration-300 ease-in-out w-8 h-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl grow">Delete address</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="remove"
          >
            <input type="hidden" v-model="delete_address_id" required />
            <div>Are you sure?</div>
            <button type="submit" :class="button">Delete</button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center p-6 text-xl h-auto bg-white rounded-xl shadow-md"
    >
      Addresses
    </div>
    <div
      class="flex justify-center items-center p-6 w-[50vw] h-full bg-white rounded-xl shadow-md"
    >
      <div v-if="fetch_complete" class="w-full grid grid-cols-auto_box gap-6">
        <div v-for="address in addresses" :class="box + ' bg-gray-50'">
          <div class="flex">
            <div class="grow">{{ address.title }}</div>
            <button
              @click="
                open_edit_dialog(address.ID, address.title, address.address)
              "
              class="mr-1 transition duration-300 ease-in-out w-9 h-9 bg-[url(/icons/edit.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
            ></button>
            <button
              @click="open_delete_dialog(address.ID)"
              class="transition duration-300 ease-in-out w-9 h-9 bg-[url(/icons/delete.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
            ></button>
          </div>
          <div class="text-sm select-text">{{ address.address }}</div>
        </div>
        <div
          @click="add_dialog = true"
          :class="
            box +
            ' transition duration-200 ease-in-out flex items-center justify-center cursor-pointer bg-gray-200 hover:bg-gray-300'
          "
        >
          <div
            class="w-24 h-24 bg-[url(/icons/add.svg)] bg-no-repeat bg-cover"
          ></div>
        </div>
      </div>
      <div
        v-else
        class="animate-spin w-24 h-24 bg-[url(/icons/loading.svg)] bg-no-repeat bg-cover"
      ></div>
    </div>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  middleware: "auth",
  layout: "account",
});
const fetch_complete = ref(false);
const add_dialog = ref(false);
const edit_dialog = ref(false);
const delete_dialog = ref(false);
const open_edit_dialog = (id: any, title: string, address: string) => {
  edit_address.value = {
    ID: id,
    title: title,
    address: address,
  };
  edit_dialog.value = true;
};
const open_delete_dialog = (id: string) => {
  delete_address_id.value = id;
  delete_dialog.value = true;
};
const addresses = ref<any>([]);
const new_address = ref({
  title: "",
  address: "",
});
const edit_address = ref({
  ID: null,
  title: "",
  address: "",
});
const delete_address_id = ref("");
onMounted(() => {
  nextTick(async () => {
    const { data } = await useFetch<any>(`/api/addresses/${user.id}`, {
      headers: {
        Authorization: `Bearer ${user.token}`,
      },
    });
    addresses.value = data.value.addresses;
    fetch_complete.value = true;
  });
});
const create = async () => {
  const { data } = await useFetch("/api/addresses", {
    method: "post",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { id: user.id },
    body: {
      title: new_address.value.title,
      address: new_address.value.address,
    },
  });
  if ((data as any).value.title) {
    add_dialog.value = false;
    addresses.value.push((data as any).value);
    new_address.value = {
      title: "",
      address: "",
    };
  }
};
const update = async () => {
  const { data } = await useFetch(`/api/addresses/${edit_address.value.ID}`, {
    method: "patch",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { userid: user.id },
    body: {
      title: edit_address.value.title,
      address: edit_address.value.address,
    },
  });
  if ((data as any).value.title) {
    addresses.value[
      addresses.value
        .map((address: any) => address.ID)
        .indexOf(edit_address.value.ID)
    ] = edit_address.value;
    edit_dialog.value = false;
    edit_address.value = {
      ID: null,
      title: "",
      address: "",
    };
  }
};
const remove = async () => {
  const { data } = await useFetch(`/api/addresses/${delete_address_id.value}`, {
    method: "delete",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { userid: user.id },
  });
  if ((data as any).value) {
    addresses.value.splice(
      addresses.value
        .map((address: any) => address.ID)
        .indexOf(delete_address_id.value),
      1
    );
    delete_dialog.value = false;
    delete_address_id.value = "";
  }
};
const box = "flex flex-col p-4 shadow-md text-xl w-60 h-60 rounded-xl";
const input =
  "transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none";
</script>
