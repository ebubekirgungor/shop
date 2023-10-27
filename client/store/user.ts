import { defineStore } from "pinia";
export const useUser = defineStore("user", {
  state: () => ({
    user: {
      id: null,
      email: null,
      name: null,
      token: null,
    },
  }),
  persist: true,
});
