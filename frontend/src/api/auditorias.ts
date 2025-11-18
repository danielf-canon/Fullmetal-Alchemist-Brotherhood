import api from "./http";

export interface AuditoriaResponseDto {
  id: number;
  user: string;
  accion: string;
  entidad: string;
  descripcion: string;
  fecha_creacion: string;
}

export const getAuditorias = () =>
  api.get<AuditoriaResponseDto[]>("/auditorias");
