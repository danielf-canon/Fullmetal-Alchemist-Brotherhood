import api from "./http";

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
}

export const registerRequest = (payload: RegisterRequest) =>
  api.post<AuthResponse>("/auth/register", payload);

export const loginRequest = (payload: LoginRequest) =>
  api.post<AuthResponse>("/auth/login", payload);
