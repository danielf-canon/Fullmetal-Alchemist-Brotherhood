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
	"github.com/rabbitmq/amqp091-go"
)


func (s *Server) HandleTransmutation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllTransmutations(w, r)
		return
	case http.MethodPost:
		s.handleCreateTransmutation(w, r) 
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleTransmutationWithId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetTransmutationById(w, r)
		return
	case http.MethodPut:
		s.handleEditTransmutation(w, r)
		return
	case http.MethodDelete:
		s.handleDeleteTransmutation(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}


func (s *Server) handleGetAllTransmutations(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	transmutations, err := s.TransmutationRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	result := []*api.TransmutationResponseDto{}
	for _, t := range transmutations {
		result = append(result, t.ToTransmutationResponseDto())
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


func (s *Server) handleGetTransmutationById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(int(id))
	if transmutation == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path,
			fmt.Errorf("transmutation with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response, err := json.Marshal(transmutation.ToTransmutationResponseDto())
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	s.logger.Info(http.StatusOK, r.URL.Path, start)
}




func (s *Server) handleCreateTransmutation(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var dto api.TransmutationRequestDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	transmutation := &models.Transmutation{
		AlquimistaID: dto.AlquimistaID,
		MaterialID:   dto.MaterialID,
		Costo:        dto.Costo,
		Resultado:    dto.Resultado,
		Estado:       "Pendiente",
	}

	transmutation, err = s.TransmutationRepository.Save(transmutation)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	body, _ := json.Marshal(transmutation)

	err = s.MQ.Channel.Publish(
		"",
		s.MQ.Queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path,
			fmt.Errorf("error enviando tarea a RabbitMQ: %v", err))
		return
	}

	s.createAuditoria("system", "CREATE_ASYNC", "Transmutación",
		fmt.Sprintf("Transmutación %d enviada al worker", transmutation.ID))

	response := map[string]interface{}{
		"status": "processing",
		"id":     transmutation.ID,
	}

	out, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

	s.logger.Info(http.StatusAccepted, r.URL.Path, start)
}




func (s *Server) handleEditTransmutation(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var dto api.TransmutationRequestDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(int(id))
	if transmutation == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path,
			fmt.Errorf("transmutation with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	transmutation.AlquimistaID = dto.AlquimistaID
	transmutation.MaterialID = dto.MaterialID
	transmutation.Costo = dto.Costo
	transmutation.Resultado = dto.Resultado
	transmutation.Estado = dto.Estado

	transmutation, err = s.TransmutationRepository.Save(transmutation)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	result, err := json.Marshal(transmutation.ToTransmutationResponseDto())
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	s.createAuditoria("system", "UPDATE", "Transmutación",
		fmt.Sprintf("Se actualizó la transmutación ID %d", transmutation.ID))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
	s.logger.Info(http.StatusAccepted, r.URL.Path, start)
}



func (s *Server) handleDeleteTransmutation(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		s.HandleError(w, http.StatusBadRequest, r.URL.Path, err)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(int(id))
	if transmutation == nil && err == nil {
		s.HandleError(w, http.StatusNotFound, r.URL.Path,
			fmt.Errorf("transmutation with id %d not found", id))
		return
	}
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	err = s.TransmutationRepository.Delete(transmutation)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	s.createAuditoria("system", "DELETE", "Transmutación",
		fmt.Sprintf("Se eliminó la transmutación ID %d", id))

	w.WriteHeader(http.StatusNoContent)
	s.logger.Info(http.StatusNoContent, r.URL.Path, start)
}
