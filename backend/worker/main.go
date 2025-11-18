package main

import (
    "backend-alquimia/models"
    "backend-alquimia/repository"
    "encoding/json"
    "fmt"
    "log"

    "github.com/rabbitmq/amqp091-go"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {

	dsn := "host=backend-alquimia-db user=postgres password=postgres dbname=backend_alquimia sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a BD:", err)
	}

	transRepo := repository.NewTransmutationRepository(db)

	conn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Error conectando a RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error creando canal RabbitMQ:", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"transmutaciones_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error declarando cola:", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true, // autoAck
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error iniciando consumidor:", err)
	}

	fmt.Println("Worker escuchando tareas...")

	for msg := range msgs {
		var t models.Transmutation
		json.Unmarshal(msg.Body, &t)

		fmt.Printf("⏳ Procesando transmutación %d...\n", t.ID)

		t.Estado = "Completado"
		t.Resultado = "Éxito"

		_, err := transRepo.Save(&t)
		if err != nil {
			log.Println("❌ Error actualizando BD:", err)
		}

		fmt.Printf("✅ Transmutación %d completada.\n", t.ID)
	}
}
