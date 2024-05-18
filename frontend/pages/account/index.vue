<script setup>
const localePath = useLocalePath();
definePageMeta({
  middleware: "auth",
});
navigateTo(localePath("account-orders"));
</script>
