import api from "./http";

export interface AlquimistaRequestDto {
  nombre: string;
  edad: number;
  especialidad: string;
  rango: string;
}

export interface AlquimistaResponseDto {
  id: number;
  nombre: string;
  edad: number;
  especialidad: string;
  rango: string;
  fecha_creacion: string;
}

export const getAlquimistas = () =>
  api.get<AlquimistaResponseDto[]>("/alquimistas");

export const getAlquimistaById = (id: number) =>
  api.get<AlquimistaResponseDto>(`/alquimistas/${id}`);

export const createAlquimista = (payload: AlquimistaRequestDto) =>
  api.post<AlquimistaResponseDto>("/alquimistas", payload);

export const updateAlquimista = (id: number, payload: AlquimistaRequestDto) =>
  api.put<AlquimistaResponseDto>(`/alquimistas/${id}`, payload);

export const deleteAlquimista = (id: number) =>
  api.delete<void>(`/alquimistas/${id}`);
