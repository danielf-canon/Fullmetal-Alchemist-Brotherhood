import Sidebar from "../components/layout/Sidebar";
import Navbar from "../components/layout/Navbar";
import MissionsList from "../components/missions/MissionsList";

export default function MissionsPage() {
  return (
    <div style={{display: "flex"}}>
      <Sidebar />
      <div style={{flex: 1}}>
        <Navbar />
        <main style={{padding: 16}}>
          <MissionsList />
        </main>
      </div>
    </div>
  );
}
