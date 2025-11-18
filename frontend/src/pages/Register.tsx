import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { registerRequest } from "../api/auth";
import Input from "../components/shared/Input";
import Button from "../components/shared/Button";

export default function Register() {
  const navigate = useNavigate();

  const [loading, setLoading] = useState(false);

  const [form, setForm] = useState({
    name: "",
    email: "",
    password: ""
  });

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const res = await registerRequest(form);

      if (res.status === 200 || res.status === 201) {
        alert("Usuario registrado con éxito");
        navigate("/login");
      }
    } catch {
      alert("Error registrando usuario");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: 360, margin: "60px auto", padding: 16 }}>
      <h2>Registrar nuevo usuario</h2>

      <form onSubmit={submit} style={{ display: "grid", gap: 12 }}>
        <Input
          placeholder="Nombre"
          value={form.name}
          onChange={(e) => setForm({ ...form, name: e.target.value })}
        />
        <Input
          placeholder="Email"
          value={form.email}
          onChange={(e) => setForm({ ...form, email: e.target.value })}
        />
        <Input
          type="password"
          placeholder="Contraseña"
          value={form.password}
          onChange={(e) => setForm({ ...form, password: e.target.value })}
        />

        <Button type="submit" disabled={loading}>
          {loading ? "Registrando..." : "Registrarse"}
        </Button>
      </form>

      <p style={{ marginTop: 12 }}>
        ¿Ya tienes cuenta? <Link to="/login">Inicia sesión</Link>
      </p>
    </div>
  );
}
