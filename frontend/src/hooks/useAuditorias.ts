import { useEffect, useState } from "react";
import {
  getAuditorias,
  type AuditoriaResponseDto
} from "../api";

export const useAuditorias = () => {
  const [auditorias, setAuditorias] = useState<AuditoriaResponseDto[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchAll = async () => {
    setLoading(true);
    const res = await getAuditorias();
    setAuditorias(res.data);
    setLoading(false);
  };

  useEffect(() => {
    fetchAll();
  }, []);

  return { auditorias, loading, fetchAll };
};
