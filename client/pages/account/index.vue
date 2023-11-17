<script setup>
definePageMeta({
  middleware: "auth",
});
navigateTo("/account/orders");
</script>