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

func (s *Server) HandleAlquimista(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllAlquimistas(w, r)
		return
	case http.MethodPost:
		s.handleCreateAlquimista(w, r)
		return
	}
}

func (s *Server) HandleAlquimistaWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAlquimistaById(w, r)
		return
	case http.MethodPut:
		s.handleEditAlquimista(w, r)
		return
	case http.MethodDelete:
		s.handleDeleteAlquimista(w, r)
		return
	}
}

func (s *Server) handleGetAllAlquimistas(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	result := []*api.AlquimistaResponseDto{}
	alquimistas, err := s.AlquimistaRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	for _, v := range alquimistas {

		result = append(result, v.ToAlquimistaResponseDto())
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

func (s *Server) handleGetAlquimistaById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}
	p, err := s.AlquimistaRepository.FindById(int(id))
	if p == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path, fmt.Errorf("person with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	response, err := json.Marshal(p.ToAlquimistaResponseDto())
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}

func (s *Server) handleCreateAlquimista(w http.ResponseWriter, r *http.Request) {
	var p api.AlquimistaRequestDto
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	alquimista := &models.Alquimista{
		Nombre:       p.Nombre,
		Edad:         int(p.Edad),
		Especialidad: p.Especialidad,
		Rango:        p.Rango,
	}
	alquimista, err = s.AlquimistaRepository.Save(alquimista)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "CREATE", "Alquimista",
	fmt.Sprintf("Se creó el alquimista %s", alquimista.Nombre))

	result, err := json.Marshal(alquimista.ToAlquimistaResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (s *Server) handleEditAlquimista(w http.ResponseWriter, r *http.Request) {
	var p api.AlquimistaRequestDto
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	alquimista, err := s.AlquimistaRepository.FindById(int(id))
	if alquimista == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	alquimista.Nombre = p.Nombre
	alquimista.Edad = int(p.Edad)
	alquimista.Especialidad = p.Especialidad
	alquimista.Rango = p.Rango
	alquimista, err = s.AlquimistaRepository.Save(alquimista)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(alquimista.ToAlquimistaResponseDto())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	s.createAuditoria("system", "UPDATE", "Alquimista",
	fmt.Sprintf("Se editó el alquimista %s (ID %d)", alquimista.Nombre, alquimista.ID))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

func (s *Server) handleDeleteAlquimista(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	person, err := s.AlquimistaRepository.FindById(int(id))
	if person == nil && err == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = s.AlquimistaRepository.Delete(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	s.createAuditoria("system", "DELETE", "Alquimista",
	fmt.Sprintf("Se eliminó el alquimista con ID %d", id))

	
}
