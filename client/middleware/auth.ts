import { useUser } from "~/store/user";
export default defineNuxtRouteMiddleware((to, from) => {
  const { user } = useUser();
  if (!user.token && (to.path.startsWith("/account") || to.path.startsWith("/admin"))) {
    return navigateTo("/login");
  }
  else if (user.token && (to.path === "/login" || to.path === "/register")) {
    return navigateTo("/");
  }
  /*else if (to.path.startsWith("/admin") && user.role == 0) {
    return navigateTo("/");
  }*/
});
