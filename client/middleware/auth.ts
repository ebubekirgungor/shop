import { useUser } from "~/store/user";
export default defineNuxtRouteMiddleware((to, from) => {
  const { user } = useUser();
  if (!user.token) {
    return navigateTo("/?login");
  }
});
