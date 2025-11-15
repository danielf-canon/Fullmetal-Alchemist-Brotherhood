package server

import (
	"backend-alquimia/api"
	"backend-alquimia/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


func (s *Server) HandleMission(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllMissions(w, r)
		return
	case http.MethodPost:
		s.handleCreateMission(w, r)
		return
	}
}

func (s *Server) HandleMissionWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetMissionById(w, r)
		return
	case http.MethodPut:
		s.handleEditMission(w, r)
		return
	case http.MethodDelete:
		s.handleDeleteMission(w, r)
		return
	}
}


func (s *Server) handleGetAllMissions(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	missions, err := s.MissionRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	result := []*api.MissionResponseDto{}
	for _, m := range missions {
		result = append(result, m.ToMissionResponseDto())
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

func (s *Server) handleGetMissionById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	mission, err := s.MissionRepository.FindById(int(id))
	if mission == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path,
			fmt.Errorf("mission with id %d not found", id))
		return
	}

	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response, err := json.Marshal(mission.ToMissionResponseDto())
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleCreateMission(w http.ResponseWriter, r *http.Request) {
	var dto api.MissionRequestDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mission := &models.Mission{
		Title:       dto.Title,
		Description: dto.Description,
		Status:      "assigned",
		AssignedTo:  dto.AssignedTo,
	}

	mission, err = s.MissionRepository.Save(mission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(mission.ToMissionResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "CREATE", "Mission",
	fmt.Sprintf("Se creó la misión %s", mission.Title))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (s *Server) handleEditMission(w http.ResponseWriter, r *http.Request) {
	var dto api.MissionRequestDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mission, err := s.MissionRepository.FindById(int(id))
	if mission == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mission.Title = dto.Title
	mission.Description = dto.Description
	mission.AssignedTo = dto.AssignedTo

	mission, err = s.MissionRepository.Save(mission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(mission.ToMissionResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "UPDATE", "Mission",
	fmt.Sprintf("Se editó la misión %s (ID %d)", mission.Title, mission.ID))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

func (s *Server) handleDeleteMission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mission, err := s.MissionRepository.FindById(int(id))
	if mission == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.MissionRepository.Delete(mission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "DELETE", "Mission",
	fmt.Sprintf("Se eliminó la misión con ID %d", id))
	w.WriteHeader(http.StatusNoContent)
}
