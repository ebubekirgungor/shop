import { defineStore } from "pinia";
export const useUser = defineStore("user", {
  state: () => ({
    user: {
      id: null,
      token: null,
      role: null,
    },
  }),
  persist: true,
});
