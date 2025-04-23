package main

import (
	"log"

	"github.com/arilsonb/starpanel/internal/api/v1/nginx/tasks"
	"github.com/hibiken/asynq"
)

func startWorker() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeInstallNginx, tasks.HandleinstallNginxTask)

	log.Println("🧠 Worker rodando...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("erro no worker: %v", err)
	}
}
