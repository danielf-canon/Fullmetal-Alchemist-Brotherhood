import api from "./http";

export interface MaterialRequestDto {
  nombre_material: string;
}

export interface MaterialResponseDto {
  id: number;
  nombre_material: string;
  fecha_creacion: string;
}

export const getMaterials = () =>
  api.get<MaterialResponseDto[]>("/materials");

export const getMaterialById = (id: number) =>
  api.get<MaterialResponseDto>(`/materials/${id}`);

export const createMaterial = (payload: MaterialRequestDto) =>
  api.post<MaterialResponseDto>("/materials", payload);

export const updateMaterial = (id: number, payload: MaterialRequestDto) =>
  api.put<MaterialResponseDto>(`/materials/${id}`, payload);

export const deleteMaterial = (id: number) =>
  api.delete<void>(`/materials/${id}`);
