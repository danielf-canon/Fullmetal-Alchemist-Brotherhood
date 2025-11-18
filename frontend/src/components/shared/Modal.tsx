// src/components/shared/Modal.tsx
import React from "react";

export default function Modal({ open, onClose, title, children }: { open: boolean; onClose: () => void; title?: string; children: React.ReactNode }) {
  if (!open) return null;
  return (
    <div style={{
      position: "fixed", top: 0, left: 0, right:0, bottom:0,
      background: "rgba(0,0,0,0.4)", display: "flex", justifyContent: "center", alignItems: "center", zIndex: 60
    }}>
      <div style={{width: 720, background: "#fff", borderRadius: 8, padding: 16}}>
        <div style={{display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 12}}>
          <h3 style={{margin: 0}}>{title}</h3>
          <button onClick={onClose}>X</button>
        </div>
        <div>{children}</div>
      </div>
    </div>
  );
}
