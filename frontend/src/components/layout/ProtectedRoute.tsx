// src/components/layout/ProtectedRoute.tsx
import { Navigate } from "react-router-dom";
import type { ReactNode } from "react";
import { useAuth } from "../../context/AuthContext";

interface Props {
  children: ReactNode;
}

export default function ProtectedRoute({ children }: Props) {
  const { token } = useAuth(); // ahora no es unknown

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  return children;
}
