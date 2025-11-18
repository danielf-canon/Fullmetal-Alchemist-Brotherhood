// src/components/shared/Input.tsx
import React from "react";

export default function Input(props: React.InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input {...props} style={{
      padding: "8px 10px",
      borderRadius: 6,
      border: "1px solid #ccc",
      width: "100%",
      boxSizing: "border-box"
    }} />
  );
}
