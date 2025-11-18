// src/components/materials/MaterialForm.tsx
import React, { useEffect, useState } from "react";
import Modal from "../shared/Modal";
import Input from "../shared/Input";
import Button from "../shared/Button";

export type MaterialPayload = { nombre_material: string };

export default function MaterialForm({ open, onClose, initial, onSave }: { open: boolean; onClose: () => void; initial?: any | null; onSave: (payload: MaterialPayload) => Promise<void> }) {
  const [form, setForm] = useState<MaterialPayload>({ nombre_material: "" });

  useEffect(() => {
    if (initial) setForm({ nombre_material: initial.nombre_material || "" });
    else setForm({ nombre_material: "" });
  }, [initial, open]);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSave(form);
  };

  return (
    <Modal open={open} onClose={onClose} title={initial ? "Editar material" : "Crear material"}>
      <form onSubmit={submit} style={{display: "grid", gap: 8}}>
        <label>Nombre <Input value={form.nombre_material} onChange={e => setForm({...form, nombre_material: e.target.value})} /></label>
        <div style={{display: "flex", gap: 8, justifyContent: "flex-end"}}>
          <Button onClick={onClose}>Cancelar</Button>
          <Button type="submit">Guardar</Button>
        </div>
      </form>
    </Modal>
  );
}
