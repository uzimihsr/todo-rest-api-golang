package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/uzimihsr/todo-rest-api-golang/config"
	"github.com/uzimihsr/todo-rest-api-golang/infrastructure/database"
	"github.com/uzimihsr/todo-rest-api-golang/presentation/handler"
	"github.com/uzimihsr/todo-rest-api-golang/presentation/router"
	"github.com/uzimihsr/todo-rest-api-golang/usecase/service"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Work in Progress")

	configFile := flag.String("config", "config/config.yaml", "config file")
	flag.Parse()

	buffer, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config config.Config
	err = yaml.Unmarshal([]byte(buffer), &config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(
		"mysql",
		config.Database.User+":"+config.Database.Password+"@tcp("+config.Database.Host+":"+config.Database.Port+")/"+config.Database.DatabaseName+"?charset=utf8mb4&parseTime=true",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println(config.Database.User + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.DatabaseName + "?charset=utf8mb4&parseTime=true")

	repository := database.NewToDoRepositoryMySQL(db)
	service := service.NewToDoService(repository)
	handler := handler.NewToDoHandler(service)
	router := router.NewToDoRouter(handler)
	server := &http.Server{
		Addr:    ":" + string(config.Server.Port),
		Handler: router.GetRouter(),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
