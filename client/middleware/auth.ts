export default defineNuxtRouteMiddleware((to) => {
  const role = useCookie<number>("role");
  if (
    role.value == undefined &&
    (to.path.startsWith("/account") || to.path.startsWith("/admin"))
  ) {
    return navigateTo("/login");
  } else if (
    role.value != undefined &&
    (to.path === "/login" || to.path === "/register")
  ) {
    return navigateTo("/");
  } else if (to.path.startsWith("/admin") && role.value === 0) {
    return navigateTo("/");
  }
});
