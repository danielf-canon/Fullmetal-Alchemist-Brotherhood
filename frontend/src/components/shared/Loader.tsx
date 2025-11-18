// src/components/shared/Loader.tsx

export default function Loader({ size = 24 }: { size?: number }) {
  return (
    <div style={{display: "inline-block", width: size, height: size, border: "3px solid #ddd", borderTopColor: "#333", borderRadius: "50%", animation: "spin 1s linear infinite"}} />
  );
}

