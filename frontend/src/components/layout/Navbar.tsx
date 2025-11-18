// src/components/layout/Navbar.tsx
import { Link } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";

export default function Navbar() {
  const { token, logout } = useAuth();

  return (
    <nav
      style={{
        padding: "10px",
        background: "#111",
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
      }}
    >
      <div style={{ display: "flex", gap: "15px" }}>
        {token && (
          <>
            <Link to="/alquimistas" style={{ color: "#fff" }}>
              Alquimistas
            </Link>
            <Link to="/materials" style={{ color: "#fff" }}>
              Materiales
            </Link>
            <Link to="/missions" style={{ color: "#fff" }}>
              Misiones
            </Link>
            <Link to="/transmutations" style={{ color: "#fff" }}>
              Transmutaciones
            </Link>
            <Link to="/auditorias" style={{ color: "#fff" }}>
              Auditorías
            </Link>
          </>
        )}
      </div>

      <div>
        {token ? (
          <button
            onClick={logout}
            style={{
              padding: "5px 10px",
              borderRadius: "5px",
              background: "#444",
              color: "#fff",
              cursor: "pointer",
              border: "none",
            }}
          >
            Cerrar sesión
          </button>
        ) : (
          <Link to="/login" style={{ color: "#fff" }}>
            Iniciar sesión
          </Link>
        )}
      </div>
    </nav>
  );
}
