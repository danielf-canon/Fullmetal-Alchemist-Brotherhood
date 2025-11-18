import { useEffect, useState } from "react";
import {
  getAlquimistas,
  getAlquimistaById,
  createAlquimista,
  updateAlquimista,
  deleteAlquimista,
  type AlquimistaRequestDto,
  type AlquimistaResponseDto
} from "../api";

export const useAlquimistas = () => {
  const [alquimistas, setAlquimistas] = useState<AlquimistaResponseDto[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchAll = async () => {
    setLoading(true);
    const res = await getAlquimistas();
    setAlquimistas(res.data);
    setLoading(false);
  };

  const fetchById = async (id: number) => {
    const res = await getAlquimistaById(id);
    return res.data;
  };

  const create = async (data: AlquimistaRequestDto) => {
    const res = await createAlquimista(data);
    await fetchAll();
    return res.data;
  };

  const update = async (id: number, data: AlquimistaRequestDto) => {
    const res = await updateAlquimista(id, data);
    await fetchAll();
    return res.data;
  };

  const remove = async (id: number) => {
    await deleteAlquimista(id);
    await fetchAll();
  };

  useEffect(() => {
    fetchAll();
  }, []);

  return { alquimistas, loading, fetchAll, fetchById, create, update, remove };
};
