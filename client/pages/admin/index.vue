<script setup>
definePageMeta({
  middleware: "auth",
});
navigateTo("/admin/products");
</script>