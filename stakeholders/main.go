package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"stakeholders-service/handler"
	"stakeholders-service/model"
	"stakeholders-service/proto/stakeholder"
	"stakeholders-service/repository"
	"stakeholders-service/service"
	"syscall"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initDB() *gorm.DB {

	connectionUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", "postgresql", "postgres", "super", "localhost", "5432", "explorer-v1")

	database, databaseErr := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		NoLowerCase: true,
	}})
	if databaseErr != nil {
		log.Fatalf(databaseErr.Error())
		return nil
	}
	if err := database.AutoMigrate(&model.User{}, &model.Person{}, &model.Profile{}); err != nil {
		log.Fatalf("error migrating Tour model: %v", err)
		return nil
	}
	return database
}
func startServer(authHandler *handler.AuthenticationHandler) {
	//router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/", authHandler.RegisterTourist).Methods("POST")
	//router.HandleFunc("/login", authHandler.Login).Methods("POST")

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	stakeholder.RegisterStakeholderServiceServer(grpcServer, authHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
	/*
		allowedOrigins := handlers.AllowedOrigins([]string{"*"})
		allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
		allowedHeaders := handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"X-Custom-Header",
		})*/

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	//corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)

	println("Server starting")
	//log.Fatal(http.ListenAndServe(":8085", corsRouter))

}

func main() {

	log.Println("Stakeholders-service staring")
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	userRepository := &repository.UserRepository{DatabaseConnection: database}
	profileRepository := &repository.ProfileRepository{DatabaseConnection: database}
	jwtGenerator := repository.NewJwtGenerator()

	authService := service.NewAuthenticationService(userRepository, jwtGenerator, profileRepository)
	authHandler := &handler.AuthenticationHandler{AuthenticationService: authService}

	startServer(authHandler)

}
