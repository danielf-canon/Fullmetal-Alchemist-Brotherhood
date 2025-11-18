import { useState } from "react";
import { loginRequest, registerRequest, type LoginRequest, type RegisterRequest } from "../api";

export const useAuth = () => {
  const [loading, setLoading] = useState(false);

  const login = async (payload: LoginRequest) => {
    setLoading(true);
    try {
      const res = await loginRequest(payload);
      localStorage.setItem("token", res.data.token);
      return true;
    } catch {
      return false;
    } finally {
      setLoading(false);
    }
  };

  const register = async (payload: RegisterRequest) => {
    setLoading(true);
    try {
      const res = await registerRequest(payload);
      localStorage.setItem("token", res.data.token);
      return true;
    } catch {
      return false;
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    localStorage.removeItem("token");
  };

  return { login, register, logout, loading };
};
