package micro

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	serverPort = "SERVER_PORT"
)

type service struct {
	name string
}

func NewService(name string) *service {
	return &service{
		name: name,
	}
}

func (s *service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	file, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		log.Fatal("Fail to load GOIF banner logo.")
	}
	log.Print(string(file))

	log.Printf("%v starting...", s.name)

	port, err := strconv.Atoi(os.Getenv(serverPort))
	if err != nil {
		log.Fatal("Failed to load server port.")
	}

	log.Printf("Server running on port: %v", port)

	<-ctx.Done()

	log.Printf("%v gracefull shutdown...", s.name)

	stop()
}
