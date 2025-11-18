import { useEffect, useState } from "react";
import {
  getMaterials,
  getMaterialById,
  createMaterial,
  updateMaterial,
  deleteMaterial,
  type MaterialRequestDto,
  type MaterialResponseDto
} from "../api";

export const useMaterials = () => {
  const [materials, setMaterials] = useState<MaterialResponseDto[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchAll = async () => {
    setLoading(true);
    const res = await getMaterials();
    setMaterials(res.data);
    setLoading(false);
  };

  const fetchById = async (id: number) => {
    const res = await getMaterialById(id);
    return res.data;
  };

  const create = async (data: MaterialRequestDto) => {
    const res = await createMaterial(data);
    await fetchAll();
    return res.data;
  };

  const update = async (id: number, data: MaterialRequestDto) => {
    const res = await updateMaterial(id, data);
    await fetchAll();
    return res.data;
  };

  const remove = async (id: number) => {
    await deleteMaterial(id);
    await fetchAll();
  };

  useEffect(() => {
    fetchAll();
  }, []);

  return { materials, loading, fetchAll, fetchById, create, update, remove };
};
