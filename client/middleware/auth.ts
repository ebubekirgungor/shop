import { useUser } from "~/store/user";
export default defineNuxtRouteMiddleware((to, from) => {
  const { user } = useUser();
  if (!user.token && to.path === "/account/orders") {
    return navigateTo("/login");
  }
  else if (user.token && (to.path === "/login" || to.path === "/register")) {
    return navigateTo("/");
  }
});
