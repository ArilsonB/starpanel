package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/arilsonb/starpanel/cmd/queue"
	"github.com/arilsonb/starpanel/internal/api/auth"
	"github.com/arilsonb/starpanel/internal/api/health"
	v1 "github.com/arilsonb/starpanel/internal/api/v1"
	"github.com/gofiber/fiber/v2"
)

func main() {
	queue.Init()
	defer queue.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	workerDone := make(chan struct{})
	go func() {
		startWorker(ctx)
		close(workerDone)
	}()

	app := fiber.New()

	api := app.Group("/api")

	authGroup := api.Group("/auth")
	auth.RegisterRoutes(authGroup)

	healthGroup := api.Group("/health")
	health.RegisterRoutes(healthGroup)

	v1Group := api.Group("/v1")
	v1.RegisterRoutes(v1Group)

	app.Static("/", "./public")

	go func() {
		log.Println("🚀 API rodando na porta 8080")
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("erro ao iniciar o servidor: %v", err)
		}
	}()

	<-ctx.Done() // Espera sinal de parada

	log.Println("🛑 Encerrando servidor...")

	_ = app.Shutdown()

	// Espera o worker finalizar
	<-workerDone

	log.Println("✅ Encerrado com sucesso.")
}