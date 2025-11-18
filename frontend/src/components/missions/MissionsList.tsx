// src/components/missions/MissionsList.tsx
import { useState } from "react";
import { useMissions } from "../../hooks/useMissions";
import Table from "../shared/Table";
import Button from "../shared/Button";
import MissionForm from "./MissionForm";
import ConfirmDialog from "../shared/ConfirmDialog";

export default function MissionsList() {
  const { missions, loading, create, update, remove } = useMissions();
  const [open, setOpen] = useState(false);
  const [editing, setEditing] = useState<any | null>(null);
  const [confirmOpen, setConfirmOpen] = useState(false);
  const [toDelete, setToDelete] = useState<number | null>(null);

  const columns = [
    { key: "title", header: "Título" },
    { key: "description", header: "Descripción" },
    { key: "status", header: "Estado" },
    { key: "assigned_to", header: "Asignado a" },
    { key: "actions", header: "Acciones", render: (row: any) => (
      <div style={{display: "flex", gap: 8}}>
        <Button onClick={() => { setEditing(row); setOpen(true); }}>Editar</Button>
        <Button onClick={() => { setToDelete(Number(row.mission_id || row.id)); setConfirmOpen(true); }}>Eliminar</Button>
      </div>
    ) }
  ];

  return (
    <div>
      <div style={{display: "flex", justifyContent: "space-between", marginBottom: 12}}>
        <h3>Misiones</h3>
        <Button onClick={() => { setEditing(null); setOpen(true); }}>Nueva</Button>
      </div>

      {loading ? <p>Cargando...</p> : <Table columns={columns} data={missions as any[]} />}

      <MissionForm open={open} onClose={() => setOpen(false)} initial={editing} onSave={async (payload) => {
        if (editing) await update(Number(editing.mission_id || editing.id), payload);
        else await create(payload);
        setOpen(false);
      }} />

      <ConfirmDialog open={confirmOpen} onCancel={() => setConfirmOpen(false)} onConfirm={async () => {
        if (toDelete) await remove(toDelete);
        setConfirmOpen(false);
      }} message="Eliminar misión?" />
    </div>
  );
}
