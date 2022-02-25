package worker

import (
	"log"
	"os"
	"skate/internal/queue"
)

func Start() {
	log.Println("Worker Running")

	var (
		queueURL = os.Getenv("QUEUE_URL")
	)

	var (
		q = queue.NewAMQPConsumer(queueURL, "skate")
	)

	if err := q.Open(func(err error) {
		log.Fatalln(err)
	}); err != nil {
		log.Fatalln(err)
	}

	if err := q.Listen(func(buff []byte) {
		log.Println(string(buff))
	}); err != nil {
		log.Fatalln(err)
	}

	// wait forever
	<-make(chan bool)
}
