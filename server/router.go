package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.Use(s.logger.RequestLogger)
	router.HandleFunc("/alquimistas", s.HandleAlquimista).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/alquimista/{id}", s.HandleAlquimistaWithId).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	router.HandleFunc("/missions", s.HandleMission).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/mission/{id}", s.HandleMissionWithId).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	router.HandleFunc("/transmutations", s.HandleTransmutation).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/transmutation/{id}", s.HandleTransmutationWithId).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	router.HandleFunc("/materials", s.HandleMaterial).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/material/{id}", s.HandleMaterialWithId).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	router.HandleFunc("/auditorias", s.handleGetAllAuditorias).Methods(http.MethodGet)
	return router
}
