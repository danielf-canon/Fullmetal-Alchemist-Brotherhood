import Navbar from "../components/layout/Navbar";
import Sidebar from "../components/layout/Sidebar";

export default function Dashboard() {
  return (
    <div style={{display: "flex"}}>
      <Sidebar />
      <div style={{flex: 1}}>
        <Navbar />
        <main style={{padding: 16}}>
          <h2>Panel principal</h2>
          <p>Bienvenido al sistema de gestión alquímica.</p>
          <ul>
            <li>Gestiona alquimistas</li>
            <li>Administra materiales</li>
            <li>Asigna misiones</li>
            <li>Realiza transmutaciones</li>
            <li>Consulta auditorías</li>
          </ul>
        </main>
      </div>
    </div>
  );
}
