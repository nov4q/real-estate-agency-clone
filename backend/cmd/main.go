package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
	"net/http"
	"polsl/tab/estate-agency/internal/api"
	"polsl/tab/estate-agency/internal/database"
	"polsl/tab/estate-agency/internal/logging"
)

// Configuration struct containing environmental variables
type Configuration struct {
	MySQL struct {
		User     string
		Password string
		DbName   string
	}
}

// initEnvConfiguration initialize environmental variables
func initEnvConfiguration() (Configuration, error) {
	configuration := Configuration{}
	if err := envconfig.InitWithPrefix(&configuration, "APP"); err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}

func initAPIHandler(conf Configuration, logger *log.Logger) (api.Handler, error) {

	db := database.NewDatabaseConnection(conf.MySQL.User, conf.MySQL.Password, conf.MySQL.DbName, logger)
	err := db.Connect()
	if err != nil {
		logger.Fatal("Error during connection to DB")
		return api.Handler{}, err
	}
	apiHandler := api.NewHandler(db, logger)
	return apiHandler, nil
}

func main() {
	conf, err := initEnvConfiguration()
	if err != nil {
		log.Fatal("Error when getting environmental variables: " + err.Error())
	}

	logger := logging.InitLogger(log.WarnLevel)

	apiHanlder, err := initAPIHandler(conf, logger)
	if err != nil {
		log.Fatal("Cannot initialize API handler")
		return
	}
	router := mux.NewRouter()
	apiHanlder.InitializeEndpoints(router)
	handler := cors.Default().Handler(router)
	port := "8081"
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Starting server at port %s failed!", port)
	}
}
