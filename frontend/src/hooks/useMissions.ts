import { useEffect, useState } from "react";
import {
  getMissions,
  getMissionById,
  createMission,
  updateMission,
  deleteMission,
  type MissionRequestDto,
  type MissionResponseDto
} from "../api";

export const useMissions = () => {
  const [missions, setMissions] = useState<MissionResponseDto[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchAll = async () => {
    setLoading(true);
    const res = await getMissions();
    setMissions(res.data);
    setLoading(false);
  };

  const fetchById = async (id: number) => {
    const res = await getMissionById(id);
    return res.data;
  };

  const create = async (data: MissionRequestDto) => {
    const res = await createMission(data);
    await fetchAll();
    return res.data;
  };

  const update = async (id: number, data: MissionRequestDto) => {
    const res = await updateMission(id, data);
    await fetchAll();
    return res.data;
  };

  const remove = async (id: number) => {
    await deleteMission(id);
    await fetchAll();
  };

  useEffect(() => {
    fetchAll();
  }, []);

  return { missions, loading, fetchAll, fetchById, create, update, remove };
};
