<template>
  <main class="flex flex-col gap-y-4">
    <div class="flex items-center p-6 text-xl h-auto bg-white rounded-xl shadow-md">
      Addresses
    </div>
    <div class="flex items-center p-6 w-[50vw] h-auto bg-white rounded-xl shadow-md">
      Addresses
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: "account",
});
</script>
