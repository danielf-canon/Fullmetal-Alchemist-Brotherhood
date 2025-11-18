// src/context/AuthContext.tsx
import { createContext, useState, useContext, type ReactNode } from "react";
import { loginRequest } from "../api/auth";
import type { LoginRequest } from "../api/auth";

interface AuthContextProps {
  token: string | null;
  loading: boolean;
  login: (data: LoginRequest) => Promise<boolean>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextProps | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token")
  );
  const [loading, setLoading] = useState(false);

  const login = async (data: LoginRequest) => {
    setLoading(true);
    try {
      const res = await loginRequest(data);
      const t = res.data.token;
      localStorage.setItem("token", t);
      setToken(t);
      return true;
    } catch {
      return false;
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    localStorage.removeItem("token");
    setToken(null);
  };

  return (
    <AuthContext.Provider value={{ token, loading, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error("useAuth must be used inside AuthProvider");
  return ctx;
};
