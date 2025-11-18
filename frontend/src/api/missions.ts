import api from "./http";

export interface MissionRequestDto {
  title: string;
  description: string;
  assigned_to: number; 
}

export interface AlquimistaEmbedded {
  id: number;
  nombre: string;
  edad: number;
  especialidad: string;
  rango: string;
  fecha_creacion?: string;
}

export interface MissionResponseDto {
  mission_id: number | string; 
  title: string;
  description: string;
  status: string;
  assigned_to: number;
  alchemist?: AlquimistaEmbedded | null;
  created_at?: string;
}

export const getMissions = () =>
  api.get<MissionResponseDto[]>("/missions");

export const getMissionById = (id: number) =>
  api.get<MissionResponseDto>(`/missions/${id}`);

export const createMission = (payload: MissionRequestDto) =>
  api.post<MissionResponseDto>("/missions", payload);

export const updateMission = (id: number, payload: MissionRequestDto) =>
  api.put<MissionResponseDto>(`/missions/${id}`, payload);

export const deleteMission = (id: number) =>
  api.delete<void>(`/missions/${id}`);
