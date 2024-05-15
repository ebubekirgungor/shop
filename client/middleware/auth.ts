export default defineNuxtRouteMiddleware((to) => {
  const role = useCookie<number>("role");
  const localePath = useLocalePath();
  if (
    role.value == undefined &&
    (to.name!.toString().startsWith("account") ||
      to.name!.toString().startsWith("admin"))
  ) {
    return navigateTo(localePath("login"));
  } else if (
    role.value != undefined &&
    (to.name!.toString().startsWith("login") ||
      to.name!.toString().startsWith("register"))
  ) {
    return navigateTo(localePath("/"));
  } else if (to.name!.toString().startsWith("admin") && role.value === 0) {
    return navigateTo(localePath("/"));
  }
});
