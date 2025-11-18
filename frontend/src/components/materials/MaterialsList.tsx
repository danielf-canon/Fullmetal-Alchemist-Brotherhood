// src/components/materials/MaterialsList.tsx
import { useState } from "react";
import { useMaterials } from "../../hooks/useMaterials";
import Table from "../shared/Table";
import Button from "../shared/Button";
import MaterialForm from "./MaterialForm";
import ConfirmDialog from "../shared/ConfirmDialog";

export default function MaterialsList() {
  const { materials, loading, create, update, remove } = useMaterials();
  const [open, setOpen] = useState(false);
  const [editing, setEditing] = useState<any | null>(null);
  const [confirmOpen, setConfirmOpen] = useState(false);
  const [toDelete, setToDelete] = useState<number | null>(null);

  const columns = [
    { key: "nombre_material", header: "Nombre" },
    { key: "fecha_creacion", header: "Creado" },
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
        <h3>Materiales</h3>
        <Button onClick={() => { setEditing(null); setOpen(true); }}>Nuevo</Button>
      </div>

      {loading ? <p>Cargando...</p> : <Table columns={columns} data={materials as any[]} />}

      <MaterialForm open={open} onClose={() => setOpen(false)} initial={editing} onSave={async (payload) => {
        if (editing) await update(editing.id, payload);
        else await create(payload);
        setOpen(false);
      }} />

      <ConfirmDialog open={confirmOpen} onCancel={() => setConfirmOpen(false)} onConfirm={async () => {
        if (toDelete) await remove(toDelete);
        setConfirmOpen(false);
      }} message="Eliminar material?" />
    </div>
  );
}
