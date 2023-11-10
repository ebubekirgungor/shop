import { useUser } from "~/store/user";
export default defineNuxtRouteMiddleware((to, from) => {
  const { user } = useUser();
  if (!user.token && to.path === "/account") {
    return navigateTo("/login");
  }
  else if (user.token && (to.path === "/login" || to.path === "/register")) {
    return navigateTo("/");
  }
});
