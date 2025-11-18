import { useEffect, useState } from "react";
import {
  getTransmutations,
  getTransmutationById,
  createTransmutation,
  updateTransmutation,
  deleteTransmutation,
  type TransmutationRequestDto,
  type TransmutationResponseDto
} from "../api";

export const useTransmutations = () => {
  const [transmutations, setTransmutations] = useState<TransmutationResponseDto[]>([]);
  const [loading, setLoading] = useState(false);
  const [processing, setProcessing] = useState(false);

  const fetchAll = async () => {
    setLoading(true);
    const res = await getTransmutations();
    setTransmutations(res.data);
    setLoading(false);
  };

  const fetchById = async (id: number) => {
    const res = await getTransmutationById(id);
    return res.data;
  };

  const create = async (data: TransmutationRequestDto) => {
    setProcessing(true);
    const res = await createTransmutation(data);
    setProcessing(false);
    await fetchAll();
    return res.data;
  };

  const update = async (id: number, data: TransmutationRequestDto) => {
    const res = await updateTransmutation(id, data);
    await fetchAll();
    return res.data;
  };

  const remove = async (id: number) => {
    await deleteTransmutation(id);
    await fetchAll();
  };

  useEffect(() => {
    fetchAll();
  }, []);

  return {
    transmutations,
    loading,
    processing,
    fetchAll,
    fetchById,
    create,
    update,
    remove
  };
};
