package server

import (
	"backend-alquimia/api"
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) handleGetAllAuditorias(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	auditorias, err := s.AuditoriaRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	result := []*api.AuditoriaResponseDto{}
	for _, a := range auditorias {
		result = append(result, a.ToAuditoriaResponseDto())
	}

	response, err := json.Marshal(result)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}
