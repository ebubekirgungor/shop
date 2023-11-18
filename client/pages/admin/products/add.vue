<template>
  <main class="flex flex-col gap-y-4">
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
            <option value="">Select category</option>
            <option v-for="c in categories" :value="c.id">{{ c.title }}</option>
          </select>
          <button :class="button">Add New</button>
        </div>
      </label>
      <label :class="label"
        >List Price<input
          :class="input"
          type="number"
          v-model="form.list_price"
      /></label>
      <label :class="label"
        >Stock Quantity<input
          :class="input"
          type="number"
          v-model="form.stock_quantity"
      /></label>
      <button
        :disabled="!form.title"
        type="submit"
        class="transition duration-300 ease-in-out w-full h-12 col-span-2 rounded-full bg-black text-white hover:bg-black/80 disabled:bg-black/60 disabled:pointer-events-none"
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
const form = ref({
  title: "",
  category_id: "",
  list_price: 0,
  stock_quantity: 0,
});
const form_div =
  "grid grid-cols-2 gap-x-[7%] gap-y-8 items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md";
const input =
  "transition duration-300 ease-in-out w-full rounded-md border-0 bg-black/5 text-sm focus:ring-2 focus:ring-slate-300";
const button =
  "transition duration-300 ease-in-out w-32 rounded-md bg-black/5 text-sm hover:bg-black/10";
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
    query: { id: 1 },
    body: {
      title: form.value.title,
      category_id: form.value.category_id,
      list_price: form.value.list_price,
      stock_quantity: form.value.stock_quantity,
    },
  });
};
</script>