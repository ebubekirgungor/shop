<script setup>
const localePath = useLocalePath();
definePageMeta({
  middleware: "auth",
});
navigateTo(localePath("admin-products"));
</script>
