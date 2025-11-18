import Sidebar from "../components/layout/Sidebar";
import Navbar from "../components/layout/Navbar";
import MaterialsList from "../components/materials/MaterialsList";

export default function MaterialsPage() {
  return (
    <div style={{display: "flex"}}>
      <Sidebar />
      <div style={{flex: 1}}>
        <Navbar />
        <main style={{padding: 16}}>
          <MaterialsList />
        </main>
      </div>
    </div>
  );
}
