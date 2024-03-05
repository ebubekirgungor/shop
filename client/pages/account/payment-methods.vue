<template>
  <main class="flex flex-col gap-y-4 w-[clamp(30rem,65rem,65rem)]">
    <div class="flex items-center p-6 text-xl h-auto bg-white rounded-xl shadow-md">
      Payment Methods
    </div>
    <div class="flex items-center p-6 min-w-[40rem] h-auto bg-white rounded-xl shadow-md">
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
