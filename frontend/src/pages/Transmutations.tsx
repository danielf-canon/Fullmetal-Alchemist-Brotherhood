import Sidebar from "../components/layout/Sidebar";
import Navbar from "../components/layout/Navbar";
import TransmutationsList from "../components/transmutations/TransmutationsList";

export default function TransmutationsPage() {
  return (
    <div style={{display: "flex"}}>
      <Sidebar />
      <div style={{flex: 1}}>
        <Navbar />
        <main style={{padding: 16}}>
          <TransmutationsList />
        </main>
      </div>
    </div>
  );
}
