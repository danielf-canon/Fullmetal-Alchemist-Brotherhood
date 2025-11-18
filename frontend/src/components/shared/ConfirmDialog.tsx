// src/components/shared/ConfirmDialog.tsx
import Modal from "./Modal";
import Button from "./Button";

export default function ConfirmDialog({ open, onCancel, onConfirm, message }: { open: boolean; onCancel: () => void; onConfirm: () => void; message?: string }) {
  return (
    <Modal open={open} onClose={onCancel} title="Confirmar">
      <p>{message ?? "¿Estás seguro?"}</p>
      <div style={{display: "flex", gap: 8, justifyContent: "flex-end"}}>
        <Button onClick={onCancel}>Cancelar</Button>
        <Button onClick={onConfirm}>Confirmar</Button>
      </div>
    </Modal>
  );
}
