import Sidebar from "../components/layout/Sidebar";
import Navbar from "../components/layout/Navbar";
import AlquimistasList from "../components/alquimistas/AlquimistasList";

export default function AlquimistasPage() {
  return (
    <div style={{display: "flex"}}>
      <Sidebar />

      <div style={{flex: 1}}>
        <Navbar />

        <main style={{padding: 16}}>
          <AlquimistasList />
        </main>
      </div>
    </div>
  );
}
