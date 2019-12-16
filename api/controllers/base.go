package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/febrielven/go_backend_webcompro/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	config.Connect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName)
	server.DB = config.GetDB()
	server.Router = mux.NewRouter()
	server.initallizeRoutes()

}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 4040")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
