<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div
      class="mt-12 sm:mt-0 flex items-center sm:p-6 text-xl h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      Payment Methods
    </div>
    <div
      class="flex items-center sm:p-6 sm:min-w-[40rem] h-auto bg-white sm:rounded-xl sm:shadow-md"
    >
      Payment Methods
    </div>
  </main>
</template>
<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: "account",
});
</script>
