import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import { useNavigate, Link } from "react-router-dom";
import Input from "../components/shared/Input";
import Button from "../components/shared/Button";

export default function Login() {
  const { login, loading } = useAuth();
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    const ok = await login({ email, password });
    if (ok) navigate("/");
    else alert("Credenciales incorrectas");
  };

  return (
    <div style={{ maxWidth: 360, margin: "80px auto", padding: 16 }}>
      <h2>Iniciar sesión</h2>

      <form onSubmit={submit} style={{ display: "grid", gap: 12 }}>
        <Input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
        />
        <Input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Contraseña"
        />

        <Button type="submit" disabled={loading}>
          {loading ? "Cargando..." : "Ingresar"}
        </Button>
      </form>

      <p style={{ marginTop: 12 }}>
        ¿No tienes cuenta? <Link to="/register">Registrarse</Link>
      </p>
    </div>
  );
}
