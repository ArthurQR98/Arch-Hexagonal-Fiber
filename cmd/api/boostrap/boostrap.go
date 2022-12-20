package boostrap

import (
	"database/sql"
	"fmt"

	"github.com/ArthurQR98/challenge_fiber/internal/platform/server"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql" //important
)

const (
	host = "localhost"
	port = 3000

	dbUser = "arthur"
	dbPass = "020398"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "test"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}
	courseRepository := mysql.NewCourseRepository(db)
	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
