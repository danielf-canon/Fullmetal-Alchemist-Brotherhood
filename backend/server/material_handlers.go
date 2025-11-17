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

// Manejo de rutas
func (s *Server) HandleMaterial(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllMaterials(w, r)
	case http.MethodPost:
		s.handleCreateMaterial(w, r)
	}
}

func (s *Server) HandleMaterialWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetMaterialById(w, r)
	case http.MethodPut:
		s.handleEditMaterial(w, r)
	case http.MethodDelete:
		s.handleDeleteMaterial(w, r)
	}
}

// GET /materials
func (s *Server) handleGetAllMaterials(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	materials, err := s.MaterialRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	result := []*api.MaterialResponseDto{}
	for _, m := range materials {
		result = append(result, m.ToMaterialResponseDto())
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

// GET /materials/{id}
func (s *Server) handleGetMaterialById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	material, err := s.MaterialRepository.FindById(int(id))
	if material == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path,
			fmt.Errorf("material with id %d not found", id))
		return
	}

	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response, err := json.Marshal(material.ToMaterialResponseDto())
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

// POST /materials
func (s *Server) handleCreateMaterial(w http.ResponseWriter, r *http.Request) {
	var dto api.MaterialRequestDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	material := &models.Material{
		NombreMaterial: dto.NombreMaterial,
	}

	material, err = s.MaterialRepository.Save(material)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(material.ToMaterialResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "CREATE", "Material",
	fmt.Sprintf("Se creó el material %s", material.NombreMaterial))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

// PUT /materials/{id}
func (s *Server) handleEditMaterial(w http.ResponseWriter, r *http.Request) {
	var dto api.MaterialRequestDto
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

	material, err := s.MaterialRepository.FindById(int(id))
	if material == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	material.NombreMaterial = dto.NombreMaterial

	material, err = s.MaterialRepository.Save(material)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(material.ToMaterialResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "UPDATE", "Material",
	fmt.Sprintf("Se editó el material %s (ID %d)", material.NombreMaterial, material.ID))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

// DELETE /materials/{id}
func (s *Server) handleDeleteMaterial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	material, err := s.MaterialRepository.FindById(int(id))
	if material == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.MaterialRepository.Delete(material)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "DELETE", "Material",
	fmt.Sprintf("Se eliminó el material con ID %d", id))
	w.WriteHeader(http.StatusNoContent)
}
