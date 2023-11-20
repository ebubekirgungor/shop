<template>
  <main class="flex flex-col gap-y-4">
    <transition name="background" mode="out-in">
      <div
        v-if="category_dialog"
        class="bg-black/40 inset-x-0 inset-y-0 w-full h-full fixed"
      ></div>
    </transition>
    <transition name="modal" mode="out-in">
      <div
        v-if="category_dialog"
        class="flex justify-center items-center inset-x-0 inset-y-0 w-full h-full fixed"
      >
        <div
          class="flex flex-col p-3 bg-white -mt-48 w-[25vw] h-auto rounded-xl z-3"
        >
          <button
            @click="category_dialog = false"
            class="self-end transition duration-300 ease-in-out w-8 h-8 bg-[url(/icons/close.svg)] bg-no-repeat bg-center rounded-full hover:bg-black/10"
          ></button>
          <h1 class="text-center text-xl grow">Create new category</h1>
          <form
            class="flex flex-col gap-y-6 mt-8 mb-4 mx-8"
            @submit.prevent="create_category"
          >
            <input
              :class="input"
              type="text"
              placeholder="Category title"
              v-model="category_new"
              required
            />
            <button
              :disabled="!category_new"
              type="submit"
              :class="create_button"
            >
              Create
            </button>
          </form>
        </div>
      </div>
    </transition>
    <div
      class="flex items-center gap-x-4 p-6 text-xl h-20 bg-white rounded-xl shadow-md"
    >
      Products
    </div>
    <form @submit.prevent="create" :class="form_div">
      <label :class="label"
        >Title<input :class="input" type="text" v-model="form.title"
      /></label>
      <label :class="label"
        >Category
        <div class="flex gap-x-4">
          <select :class="input" v-model="form.category_id">
            <option value="" selected>Select category</option>
            <option v-for="c in categories" :value="c.ID">{{ c.title }}</option>
          </select>
          <button
            type="button"
            @click="category_dialog = true"
            :class="add_category_button"
          >
            Add New
          </button>
        </div>
      </label>
      <label :class="label"
        >List Price<input
          :class="input"
          type="number"
          min="1"
          v-model="form.list_price"
      /></label>
      <label :class="label"
        >Stock Quantity<input
          :class="input"
          type="number"
          min="0"
          v-model="form.stock_quantity"
      /></label>
      <button
        :disabled="
          form.title == '' ||
          form.category_id == '' ||
          !form.list_price ||
          !form.stock_quantity
        "
        type="submit"
        :class="create_button"
      >
        Create
      </button>
    </form>
  </main>
</template>
<script setup lang="ts">
import { nextTick } from "vue";
import { useUser } from "@/store/user";
const { user } = useUser();
definePageMeta({
  middleware: "auth",
  layout: "admin",
});
const categories = ref<any>([]);
const category_dialog = ref(false);
const category_new = ref("");
const form = ref({
  title: "",
  category_id: "",
  list_price: null,
  stock_quantity: null,
});
const form_div =
  "grid grid-cols-2 gap-x-[7%] gap-y-8 items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md";
const input =
  "transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const add_category_button =
  "transition duration-300 ease-in-out w-32 rounded-md bg-black/5 text-sm hover:bg-black/10";
const create_button =
  "transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none";
const label = "flex flex-col gap-y-2";
onMounted(() => {
  nextTick(async () => {
    const { data } = await useFetch<any>("/api/categories");
    categories.value = data.value;
  });
});
const create = async () => {
  await useFetch("/api/products", {
    method: "post",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { id: user.id },
    body: {
      title: form.value.title,
      category_id: form.value.category_id,
      list_price: form.value.list_price,
      stock_quantity: form.value.stock_quantity,
    },
  });
};
const create_category = async () => {
  const { data } = await useFetch("/api/categories", {
    method: "post",
    headers: {
      Authorization: `Bearer ${user.token}`,
    },
    query: { id: user.id },
    body: {
      title: category_new.value,
    },
  });
  console.log((data as any).value);
  console.log(categories.value);
  if ((data as any).value.title) {
    category_dialog.value = false;
    categories.value.push((data as any).value);
  }
};
</script>
