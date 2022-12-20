package boostrap

import "github.com/ArthurQR98/challenge_fiber/internal/platform/server"

const (
	host = "localhost"
	port = 3000
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
