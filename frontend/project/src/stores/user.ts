import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { User } from '../types/user';
import { authApi } from '../api/auth';

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null);
  const token = ref<string | null>(null);

  const setUser = (userData: User) => {
    user.value = userData;
  };

  const setToken = (newToken: string) => {
    token.value = newToken;
    localStorage.setItem('token', newToken);
  };

  const login = async (username: string, password: string) => {
    const response = await authApi.login({ username, password });
    setUser(response.data.user);
    setToken(response.data.token);
  };

  const register = async (formData: { username: string; email: string; password: string }) => {
    const response = await authApi.register(formData);
    setUser(response.data.user);
    setToken(response.data.token);
  };

  const logout = () => {
    user.value = null;
    token.value = null;
    localStorage.removeItem('token');
  };

  return {
    user,
    token,
    login,
    register,
    logout,
    setUser,
    setToken
  };
});