export default defineNuxtRouteMiddleware((to, from) => {
  const userid = useCookie("userid");
  const role = useCookie<string>("role");
  if (
    !userid.value &&
    (to.path.startsWith("/account") || to.path.startsWith("/admin"))
  ) {
    return navigateTo("/login");
  } else if (
    userid.value &&
    (to.path === "/login" || to.path === "/register")
  ) {
    return navigateTo("/");
  } else if (to.path.startsWith("/admin") && parseInt(role.value) === 0) {
    return navigateTo("/");
  }
});
