import api from './axios';
import { API_ENDPOINTS } from '../config/api';
import type { LoginForm, RegisterForm, User } from '../types/user';

export const authApi = {
  login: (data: LoginForm) => 
    api.post<{ user: User; token: string }>(API_ENDPOINTS.AUTH.LOGIN, data),
  
  register: (data: RegisterForm) =>
    api.post<{ user: User; token: string }>(API_ENDPOINTS.AUTH.REGISTER, data),
    
  getCurrentUser: () =>
    api.get<{ user: User }>(API_ENDPOINTS.AUTH.ME)
};