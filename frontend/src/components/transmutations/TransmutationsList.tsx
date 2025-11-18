// src/components/transmutations/TransmutationsList.tsx
import { useState } from "react";
import { useTransmutations } from "../../hooks/useTransmutations";
import Table from "../shared/Table";
import Button from "../shared/Button";
import TransmutationForm from "./TransmutationForm";
import ConfirmDialog from "../shared/ConfirmDialog";

export default function TransmutationsList() {
  const { transmutations, loading, processing, create, update, remove } = useTransmutations();
  const [open, setOpen] = useState(false);
  const [editing, setEditing] = useState<any | null>(null);
  const [confirmOpen, setConfirmOpen] = useState(false);
  const [toDelete, setToDelete] = useState<number | null>(null);

  const columns = [
    { key: "id", header: "ID" },
    { key: "alquimista_id", header: "Alquimista ID" },
    { key: "material_id", header: "Material ID" },
    { key: "costo", header: "Costo" },
    { key: "estado", header: "Estado" },
    { key: "resultado", header: "Resultado" },
    { key: "fecha_creacion", header: "Creación" },
    { key: "actions", header: "Acciones", render: (row: any) => (
      <div style={{display: "flex", gap: 8}}>
        <Button onClick={() => { setEditing(row); setOpen(true); }}>Editar</Button>
        <Button onClick={() => { setToDelete(row.id); setConfirmOpen(true); }}>Eliminar</Button>
      </div>
    ) }
  ];

  return (
    <div>
      <div style={{display: "flex", justifyContent: "space-between", marginBottom: 12}}>
        <h3>Transmutaciones</h3>
        <div style={{display: "flex", gap: 8, alignItems: "center"}}>
          {processing && <small>Procesando...</small>}
          <Button onClick={() => { setEditing(null); setOpen(true); }}>Nueva</Button>
        </div>
      </div>

      {loading ? <p>Cargando...</p> : <Table columns={columns} data={transmutations as any[]} />}

      <TransmutationForm open={open} onClose={() => setOpen(false)} initial={editing} onSave={async (payload) => {
        if (editing) await update(editing.id, payload);
        else await create(payload);
        setOpen(false);
      }} />

      <ConfirmDialog open={confirmOpen} onCancel={() => setConfirmOpen(false)} onConfirm={async () => {
        if (toDelete) await remove(toDelete);
        setConfirmOpen(false);
      }} message="Eliminar transmutación?" />
    </div>
  );
}
