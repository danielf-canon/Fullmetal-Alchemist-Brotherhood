// src/components/missions/MissionForm.tsx
import React, { useEffect, useState } from "react";
import Modal from "../shared/Modal";
import Input from "../shared/Input";
import Button from "../shared/Button";
import { useAlquimistas } from "../../hooks/useAlquimistas";

export type MissionPayload = { title: string; description: string; assigned_to: number };

export default function MissionForm({ open, onClose, initial, onSave }: { open: boolean; onClose: () => void; initial?: any | null; onSave: (payload: MissionPayload) => Promise<void> }) {
  const { alquimistas } = useAlquimistas();
  const [form, setForm] = useState<MissionPayload>({ title: "", description: "", assigned_to: 0 });

  useEffect(() => {
    if (initial) setForm({ title: initial.title || "", description: initial.description || "", assigned_to: initial.assigned_to || 0 });
    else setForm({ title: "", description: "", assigned_to: alquimistas?.[0]?.id ?? 0 });
  }, [initial, open, alquimistas]);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSave(form);
  };

  return (
    <Modal open={open} onClose={onClose} title={initial ? "Editar misión" : "Crear misión"}>
      <form onSubmit={submit} style={{display: "grid", gap: 8}}>
        <label>Título <Input value={form.title} onChange={e => setForm({...form, title: e.target.value})} /></label>
        <label>Descripción <Input value={form.description} onChange={e => setForm({...form, description: e.target.value})} /></label>
        <label>Asignado a
          <select value={form.assigned_to} onChange={e => setForm({...form, assigned_to: Number(e.target.value)})} style={{width: "100%", padding: 8, borderRadius: 6, border: "1px solid #ccc"}}>
            <option value={0}>-- Seleccionar --</option>
            {alquimistas?.map(a => <option key={a.id} value={a.id}>{a.nombre}</option>)}
          </select>
        </label>
        <div style={{display: "flex", gap: 8, justifyContent: "flex-end"}}>
          <Button onClick={onClose}>Cancelar</Button>
          <Button type="submit">Guardar</Button>
        </div>
      </form>
    </Modal>
  );
}
