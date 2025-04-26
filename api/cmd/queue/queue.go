package queue

import "github.com/hibiken/asynq"

var Client *asynq.Client
const redisAddr = "localhost:6379"

func Init() {
	Client = asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})
}

func Close() {
	if Client != nil {
		defer Client.Close()
	}
}
