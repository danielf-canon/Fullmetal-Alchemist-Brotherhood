import Sidebar from "../components/layout/Sidebar";
import Navbar from "../components/layout/Navbar";
import AuditoriasList from "../components/auditorias/AuditoriasList";

export default function AuditoriasPage() {
  return (
    <div style={{display: "flex"}}>
      <Sidebar />
      <div style={{flex: 1}}>
        <Navbar />
        <main style={{padding: 16}}>
          <AuditoriasList />
        </main>
      </div>
    </div>
  );
}
