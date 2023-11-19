package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"ms-go-initial/config"
	"ms-go-initial/db"
	"ms-go-initial/handler"
	"ms-go-initial/model"
	"ms-go-initial/service"
	"net/http"
	_ "os"
	"strconv"
	"time"
)

var profile *string

func init() {
	profile = flag.String("profile", "default", "Application run profile")
}

func main() {
	flag.Parse()
	initLogger()

	log.Info("Application is starting with profile: ", *profile)

	initEnvVars()

	config.LoadConfig()

	db.InitDb()

	Migrate()

	repo := db.NewRepository(db.GetDb())

	r := chi.NewRouter()
	s := service.NewService(repo)

	handler.BaseHandler(r)
	handler.Handler(r, s)

	port := strconv.Itoa(config.Conf.Port)
	log.Info("Starting server at port ", port)
	log.Fatal((&http.Server{
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
		Handler:      r,
		Addr:         ":" + port,
	}).ListenAndServe())
}

func initEnvVars() {
	_ = godotenv.Load("profiles/" + *profile + ".env")
}

func initLogger() {
	log.SetLevel(log.InfoLevel)
	if *profile == "default" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	logger := log.New()
	logger.Level = log.InfoLevel

}

func Migrate() {
	err := db.GetDb().AutoMigrate(&model.Contract{})
	err = db.GetDb().AutoMigrate(&model.Users{})
	if err != nil {
		log.Error("ActionLog.Migrate.error ", err.Error())
		return
	}
}
