package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lgdev07/superapi/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf(`host=%s port=%s user=%s dbname=%s sslmode=disable 
	password=%s`, DbHost, DbPort, DbUser, DbName, DbPassword)

	s.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("Cannot conect to database %s\n", DbName)
		log.Fatal(err)
	} else {
		fmt.Printf("We are connected to database %s\n", DbName)
	}

	s.DB.AutoMigrate(
		&models.Super{},
	)

	s.Router = mux.NewRouter().StrictSlash(true)
	s.InitializeRoutes()

}

func (s *Server) Run() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	log.Printf("\nServer starting on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Router))

}
