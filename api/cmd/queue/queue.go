package queue

import "github.com/hibiken/asynq"

var Client *asynq.Client

func Init() {
	Client = asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
}

func Close() {
	if Client != nil {
		defer Client.Close()
	}
}
