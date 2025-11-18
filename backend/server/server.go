package server

import (
	"backend-alquimia/logger"
	"backend-alquimia/models"
	"backend-alquimia/repository"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router 
	logger *logger.Logger

	AlquimistaRepository    repository.Repository[models.Alquimista]
	MissionRepository       repository.Repository[models.Mission]
	TransmutationRepository repository.Repository[models.Transmutation]
	MaterialRepository      repository.Repository[models.Material]
	AuditoriaRepository     repository.Repository[models.Auditoria]
	UsuarioRepository       *repository.UsuarioRepository

	MQ struct {
        Conn    *amqp091.Connection
        Channel *amqp091.Channel
        Queue   amqp091.Queue
    }
}

func NewServer() *Server {
	s := &Server{
		logger: logger.NewLogger(),
	}
	return s
}

func (s *Server) StartServer() {
	fmt.Println("Inicializando base de datos...")
	s.initDB()
	fmt.Println("Inicializando rutas...")
	s.routes()
	fmt.Println("Inicializando RabbitMQ...")
	s.initRabbitMQ()


	corsObj := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
	)

	fmt.Println("Inicializando mux...")
	srv := &http.Server{
		Addr:    ":8000",
		Handler: corsObj(s.Router),
	}

	if err := srv.ListenAndServe(); err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) initDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		s.logger.Fatal(err)
	}

	s.DB = db

	fmt.Println("Aplicando migraciones...")
	s.DB.AutoMigrate(
		&models.Alquimista{},
		&models.Mission{},
		&models.Transmutation{},
		&models.Material{},
		&models.Auditoria{},
		&models.Usuario{},
	)

	s.AlquimistaRepository = repository.NewAlquimistaRepository(s.DB)
	s.MissionRepository = repository.NewMissionRepository(s.DB)
	s.TransmutationRepository = repository.NewTransmutationRepository(s.DB)
	s.MaterialRepository = repository.NewMaterialRepository(s.DB)
	s.AuditoriaRepository = repository.NewAuditoriaRepository(s.DB)
	s.UsuarioRepository = repository.NewUsuarioRepository(s.DB)
}

func (s *Server) initRabbitMQ() {
    conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
    if err != nil {
        s.logger.Fatal(fmt.Errorf("no se pudo conectar a RabbitMQ: %v", err))
    }

    ch, err := conn.Channel()
    if err != nil {
        s.logger.Fatal(fmt.Errorf("no se pudo abrir canal RabbitMQ: %v", err))
    }

    queue, err := ch.QueueDeclare(
        "transmutaciones_queue",
        true, 
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        s.logger.Fatal(fmt.Errorf("no se pudo declarar la queue: %v", err))
    }

    s.MQ.Conn = conn
    s.MQ.Channel = ch
    s.MQ.Queue = queue

    fmt.Println("RabbitMQ inicializado y cola declarada.")
}

func (s *Server) createAuditoria(user, action, entity string, details string) {
	a := &models.Auditoria{
		User:        user,
		Accion:      action,
		Entidad:     entity,
		Descripcion: details,
	}

	_, err := s.AuditoriaRepository.Save(a)
	if err != nil {
		s.logger.Error(500, "AUDIT", err)
	}
}
