package main

import (
	"log"

	"github.com/ArthurQR98/challenge_fiber/cmd/api/boostrap"
	_ "github.com/ArthurQR98/challenge_fiber/cmd/api/docs"
)

//go:generate swag init --dir ./,../../internal/platform/server

// @title			Challenge Fiber API
// @version		1.0
// @description	Documentation with Swagger of the API
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	arthur.quezada98@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {
	if err := boostrap.Run(); err != nil {
		log.Fatal(err)
	}
}
