import { useAuditorias } from "../../hooks/useAuditorias";
import Table from "../shared/Table";

export default function AuditoriasList() {
  const { auditorias, loading } = useAuditorias();

  const columns = [
    { key: "id", header: "ID" },
    { key: "user", header: "Usuario" },
    { key: "accion", header: "Acción" },
    { key: "entidad", header: "Entidad" },
    { key: "descripcion", header: "Descripción" },
    { key: "fecha_creacion", header: "Fecha" }
  ];

  return (
    <div>
      <h3>Auditorías</h3>
      {loading ? <p>Cargando...</p> : <Table columns={columns} data={auditorias as any[]} />}
    </div>
  );
}
