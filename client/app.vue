<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
<style>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s;
}
.modal-enter-from,
.modal-leave-to {
  transform: translateY(30px);
  opacity: 0;
}
.background-enter-active,
.background-leave-active {
  transition: all 0.3s;
}
.background-enter-from,
.background-leave-to {
  opacity: 0;
}
</style>
