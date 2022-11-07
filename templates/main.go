package templates

// TODO mod file name check and mod file name with import
var Main = `
	/*
	{{.Copyright}}
	{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
	*/
	package main

	import (
		"context"
		"os/signal"
		"syscall"
		"time"
	)

	// @title Swagger Example API
	// @version 1.0
	// description This is a sample client server
	// @contact.name API support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io
	// @license.name Apache 2.0
	// @license.url http://www.apache.org/license/LICENSE-2.0.httml
	// @host localhost:8000
	// @BasePath /api/v1
	func main() {
		// Create context that listens for the interrupt signal from the OS
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop();

		// create fiber app
		app := routes.SetUpRouter()

		// initializing the server in goroutine so that
		// it won't block the graceful shutdown handling below
		go func() {
			if err != app.Listen(":" + "8000"); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		}()

		// Listen for the interrupt signal
		<-ctx.Done

		// Restore default behavior on the interrupt signal notify user of shutdown
		stop()

		// The context is used to inform the server it has 5 seconds to finish
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.Shutdown(); err != nil {
			panic(err)
		}

		// End of Server
	}
`
