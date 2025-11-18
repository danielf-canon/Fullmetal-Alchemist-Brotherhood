// src/components/transmutations/TransmutationForm.tsx
import React, { useEffect, useState } from "react";
import Modal from "../shared/Modal";
import Input from "../shared/Input";
import Button from "../shared/Button";
import { useAlquimistas } from "../../hooks/useAlquimistas";
import { useMaterials } from "../../hooks/useMaterials";

export type TransmutationPayload = {
  alquimista_id: number;
  material_id: number;
  costo: number;
  resultado?: string;
  estado?: string;
};

export default function TransmutationForm({ open, onClose, initial, onSave }: { open: boolean; onClose: () => void; initial?: any | null; onSave: (payload: TransmutationPayload) => Promise<void> }) {
  const { alquimistas } = useAlquimistas();
  const { materials } = useMaterials();

  const [form, setForm] = useState<TransmutationPayload>({ alquimista_id: 0, material_id: 0, costo: 0, resultado: "", estado: "Pendiente" });

  useEffect(() => {
    if (initial) {
      setForm({
        alquimista_id: initial.alquimista_id ?? 0,
        material_id: initial.material_id ?? 0,
        costo: initial.costo ?? 0,
        resultado: initial.resultado ?? "",
        estado: initial.estado ?? "Pendiente"
      });
    } else {
      setForm({
        alquimista_id: alquimistas?.[0]?.id ?? 0,
        material_id: materials?.[0]?.id ?? 0,
        costo: 0,
        resultado: "",
        estado: "Pendiente"
      });
    }
  }, [initial, open, alquimistas, materials]);

  const submit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSave(form);
  };

  return (
    <Modal open={open} onClose={onClose} title={initial ? "Editar transmutación" : "Crear transmutación"}>
      <form onSubmit={submit} style={{display: "grid", gap: 8}}>
        <label>Alquimista
          <select value={form.alquimista_id} onChange={e => setForm({...form, alquimista_id: Number(e.target.value)})} style={{width: "100%", padding: 8, borderRadius: 6}}>
            <option value={0}>-- Seleccionar --</option>
            {alquimistas?.map(a => <option key={a.id} value={a.id}>{a.nombre}</option>)}
          </select>
        </label>

        <label>Material
          <select value={form.material_id} onChange={e => setForm({...form, material_id: Number(e.target.value)})} style={{width: "100%", padding: 8, borderRadius: 6}}>
            <option value={0}>-- Seleccionar --</option>
            {materials?.map(m => <option key={(m as any).id} value={(m as any).id}>{(m as any).nombre_material}</option>)}
          </select>
        </label>

        <label>Costo <Input type="number" value={String(form.costo)} onChange={e => setForm({...form, costo: Number(e.target.value)})} /></label>

        <label>Resultado <Input value={form.resultado} onChange={e => setForm({...form, resultado: e.target.value})} /></label>

        <label>Estado
          <select value={form.estado} onChange={e => setForm({...form, estado: e.target.value})} style={{width: "100%", padding: 8, borderRadius: 6}}>
            <option>Pendiente</option>
            <option>Completado</option>
            <option>Fallido</option>
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
