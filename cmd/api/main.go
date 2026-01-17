package main

// "context"
// "os"
// "os/signal"
// "syscall"

// messagingRouter "github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/messaging/router"
// "github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/middleware"
// webRouter "github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/router"
// "github.com/7yrionLannister/golang-technical-assesment-v2/pkg/config"
// "github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
// "github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"

func main() {
	// config.Setup()

	// app := fiber.New()
	// err := webRouter.Setup(app)
	// if err != nil {
	// 	log.L.Error("Error initializing web router", "error", err)
	// 	os.Exit(1)
	// }
	// middleware.Setup(app)

	// router, err := messagingRouter.Setup()
	// if err != nil {
	// 	log.L.Error("Error initializing messaging router", "error", err)
	// 	os.Exit(1)
	// }
	// shutdown := make(chan os.Signal, 1)
	// signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// go func() {
	// 	go util.StartServerWithGracefulShutdown(app, config.Env.ServerHost, config.Env.ServerPort)
	// 	router.Run(context.Background())
	// }()
	// <-shutdown
}
