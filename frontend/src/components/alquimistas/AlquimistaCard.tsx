
export default function AlquimistaCard({ a }: { a: any }) {
  return (
    <div style={{border: "1px solid #eee", padding: 12, borderRadius: 8}}>
      <h4 style={{margin: 0}}>{a.nombre}</h4>
      <p style={{margin: 0}}>Edad: {a.edad} — {a.especialidad}</p>
      <small>{a.rango}</small>
    </div>
  );
}
