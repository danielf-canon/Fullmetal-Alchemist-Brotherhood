// src/components/alquimistas/AlquimistaForm.tsx
import React, { useEffect, useState } from "react";
import Modal from "../shared/Modal";
import Input from "../shared/Input";
import Button from "../shared/Button";

export type AlquimistaPayload = {
  nombre: string;
  edad: number;
  especialidad: string;
  rango: string;
};

export default function AlquimistaForm({ open, onClose, initial, onSave }: { open: boolean; onClose: () => void; initial?: any | null; onSave: (payload: AlquimistaPayload) => Promise<void> }) {
  const [form, setForm] = useState<AlquimistaPayload>({ nombre: "", edad: 0, especialidad: "", rango: "" });

  useEffect(() => {
    if (initial) {
      setForm({
        nombre: initial.nombre || "",
        edad: initial.edad ?? 0,
        especialidad: initial.especialidad || "",
        rango: initial.rango || ""
      });
    } else {
      setForm({ nombre: "", edad: 0, especialidad: "", rango: "" });
    }
  }, [initial, open]);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSave(form);
  };

  return (
    <Modal open={open} onClose={onClose} title={initial ? "Editar alquimista" : "Crear alquimista"}>
      <form onSubmit={submit} style={{display: "grid", gap: 8}}>
        <label>Nombre <Input value={form.nombre} onChange={e => setForm({...form, nombre: e.target.value})} /></label>
        <label>Edad <Input type="number" value={String(form.edad)} onChange={e => setForm({...form, edad: Number(e.target.value)})} /></label>
        <label>Especialidad <Input value={form.especialidad} onChange={e => setForm({...form, especialidad: e.target.value})} /></label>
        <label>Rango <Input value={form.rango} onChange={e => setForm({...form, rango: e.target.value})} /></label>
        <div style={{display: "flex", gap: 8, justifyContent: "flex-end"}}>
          <Button onClick={onClose}>Cancelar</Button>
          <Button type="submit">Guardar</Button>
        </div>
      </form>
    </Modal>
  );
}
