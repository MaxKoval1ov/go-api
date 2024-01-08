package main

import (
	lib "github.com/MaxKoval1ov/go-api/lib"
	"github.com/MaxKoval1ov/go-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	db := lib.CreateConnection(lib.Config())

	if !db.Migrator().HasTable(&models.Test{}) {
		db.Migrator().AutoMigrate(&models.Test{})
	}

	r.GET("/test", func(c *gin.Context) {
		var tests []models.Test
        db.Find(&tests)
        c.JSON(http.StatusOK, tests)
	})

	r.PUT("/test", func(c *gin.Context) {
		var tests []models.Test
        db.Find(&tests)
        c.JSON(http.StatusOK, tests)
	})
	
	r.POST("/test", func(c *gin.Context) {
        var inputTest models.Test

        if err := c.ShouldBindJSON(&inputTest); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        db.Create(&inputTest)

        c.JSON(http.StatusCreated, gin.H{"message": "Test created successfully", "test": inputTest})
    })
	
	// CreateConnection()
	r.Run() // listen
}

// func multiply(perem *int) int{
// 	perem = 125
// 	return perem * perem
// }

// numberHAsTobe := 123 

// multiply(&numberHAsTobe)

// print(numberHAsTobe)

//nunber changed 125

// func run() error {
// 	// read config from env
// 	cfg := config.Read()

// 	// create port repository
// 	portStoreRepo := inmem.NewPortStore()

// 	// create port service
// 	portService := services.NewPortService(portStoreRepo)

// 	// create http server with application injected
// 	httpServer := transport.NewHttpServer(portService)

// 	// create http router
// 	router := mux.NewRouter()
// 	router.HandleFunc("/port", httpServer.GetPort).Methods("GET")
// 	router.HandleFunc("/count", httpServer.CountPorts).Methods("GET")
// 	router.HandleFunc("/ports", httpServer.UploadPorts).Methods("POST")

// 	srv := &http.Server{
// 		Addr:    cfg.HTTPAddr,
// 		Handler: router,
// 	}

// 	// listen to OS signals and gracefully shutdown HTTP server
// 	stopped := make(chan struct{})
// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
// 		<-sigint
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()
// 		if err := srv.Shutdown(ctx); err != nil {
// 			log.Printf("HTTP Server Shutdown Error: %v", err)
// 		}
// 		close(stopped)
// 	}()

// 	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

// 	// start HTTP server
// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
// 		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
// 	}

// 	<-stopped

// 	log.Printf("Have a nice day!")

// 	return nil
// }
