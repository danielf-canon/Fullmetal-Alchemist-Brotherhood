package server

import (
	"github.com/gorilla/mux"
)

func (s *Server) routes() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/register", s.HandleRegister).Methods("POST")
	router.HandleFunc("/auth/login", s.HandleLogin).Methods("POST")

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(s.AuthMiddleware)

	// Alquimistas
	protected.HandleFunc("/alquimistas", s.HandleAlquimista).Methods("GET", "POST")
	protected.HandleFunc("/alquimistas/{id}", s.HandleAlquimistaWithId).Methods("GET", "PUT", "DELETE")

	// Misiones
	protected.HandleFunc("/missions", s.HandleMission).Methods("GET", "POST")
	protected.HandleFunc("/missions/{id}", s.HandleMissionWithId).Methods("GET", "PUT", "DELETE")

	// Materiales
	protected.HandleFunc("/materials", s.HandleMaterial).Methods("GET", "POST")
	protected.HandleFunc("/materials/{id}", s.HandleMaterialWithId).Methods("GET", "PUT", "DELETE")

	// Transmutaciones
	protected.HandleFunc("/transmutations", s.HandleTransmutation).Methods("GET", "POST")
	protected.HandleFunc("/transmutations/{id}", s.HandleTransmutationWithId).Methods("GET", "PUT", "DELETE")

	// Auditorias
	protected.HandleFunc("/auditorias", s.handleGetAllAuditorias).Methods("GET")

	s.Router = router
}
