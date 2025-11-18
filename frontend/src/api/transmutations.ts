import api from "./http";

export interface TransmutationRequestDto {
  alquimista_id: number;
  material_id: number;
  costo: number;
  resultado?: string;
  estado?: string;
}

export interface TransmutationResponseDto {
  id: number;
  alquimista_id: number;
  material_id: number;
  costo: number;
  resultado: string;
  estado: string;
  fecha_creacion: string;
}

export const getTransmutations = () =>
  api.get<TransmutationResponseDto[]>("/transmutations");

export const getTransmutationById = (id: number) =>
  api.get<TransmutationResponseDto>(`/transmutations/${id}`);

export const createTransmutation = (payload: TransmutationRequestDto) =>
  api.post("/transmutations", payload);

export const updateTransmutation = (id: number, payload: TransmutationRequestDto) =>
  api.put<TransmutationResponseDto>(`/transmutations/${id}`, payload);

export const deleteTransmutation = (id: number) =>
  api.delete<void>(`/transmutations/${id}`);
