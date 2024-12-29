export interface User {
  id?: number;
  username: string;
  email: string;
  token?: string;
}

export interface LoginForm {
  username: string;
  password: string;
}

export interface RegisterForm extends LoginForm {
  email: string;
  confirmPassword: string;
}