// src/components/alquimistas/AlquimistasList.tsx
import { useState } from "react";
import { useAlquimistas } from "../../hooks/useAlquimistas";
import Table from "../shared/Table";
import Button from "../shared/Button";
import AlquimistaForm from "./AlquimistaForm";
import ConfirmDialog from "../shared/ConfirmDialog";

export default function AlquimistasList() {
  const { alquimistas, loading, create, update, remove } = useAlquimistas();
  const [editing, setEditing] = useState<any | null>(null);
  const [modalOpen, setModalOpen] = useState(false);
  const [confirmOpen, setConfirmOpen] = useState(false);
  const [toDelete, setToDelete] = useState<number | null>(null);

  const columns = [
    { key: "nombre", header: "Nombre" },
    { key: "edad", header: "Edad" },
    { key: "especialidad", header: "Especialidad" },
    { key: "rango", header: "Rango" },
    {
      key: "actions", header: "Acciones", render: (row: any) => (
        <div style={{display: "flex", gap: 8}}>
          <Button onClick={() => { setEditing(row); setModalOpen(true); }}>Editar</Button>
          <Button onClick={() => { setToDelete(row.id); setConfirmOpen(true); }}>Eliminar</Button>
        </div>
      )
    }
  ];

  return (
    <div>
      <div style={{display: "flex", justifyContent: "space-between", marginBottom: 12}}>
        <h3>Alquimistas</h3>
        <Button onClick={() => { setEditing(null); setModalOpen(true); }}>Nuevo</Button>
      </div>

      {loading ? <p>Cargando...</p> : <Table columns={columns} data={alquimistas as any[]} />}

      <AlquimistaForm open={modalOpen} onClose={() => setModalOpen(false)} initial={editing} onSave={async (payload) => {
        if (editing) await update(editing.id, payload);
        else await create(payload);
        setModalOpen(false);
      }} />

      <ConfirmDialog open={confirmOpen} onCancel={() => setConfirmOpen(false)} onConfirm={async () => {
        if (toDelete) await remove(toDelete);
        setConfirmOpen(false);
      }} message="Eliminar alquimista?" />
    </div>
  );
}
