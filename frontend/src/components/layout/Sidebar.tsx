import { Link, useLocation } from "react-router-dom";

export default function Sidebar() {
  const { pathname } = useLocation();

  const linkClasses = (path: string) =>
    `block px-4 py-2 rounded-md transition ${
      pathname === path
        ? "bg-blue-600 text-white font-semibold"
        : "text-gray-700 hover:bg-gray-200"
    }`;

  return (
    <aside className="w-60 bg-white shadow-md px-4 py-6">
      <h2 className="text-xl font-bold mb-6 text-blue-700">Alquimia</h2>

      <nav className="space-y-2">
        <Link className={linkClasses("/")} to="/">Dashboard</Link>
        <Link className={linkClasses("/alquimistas")} to="/alquimistas">Alquimistas</Link>
        <Link className={linkClasses("/materials")} to="/materials">Materiales</Link>
        <Link className={linkClasses("/missions")} to="/missions">Misiones</Link>
        <Link className={linkClasses("/transmutations")} to="/transmutations">Transmutaciones</Link>
        <Link className={linkClasses("/auditorias")} to="/auditorias">Auditorías</Link>
      </nav>
    </aside>
  );
}
