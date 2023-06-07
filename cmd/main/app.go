package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"proj/internal/config"
	"proj/internal/user"
	"proj/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("---> creating router...")
	router := httprouter.New()

	cfg := config.GetConfig() // cfg === config | "===" means the same | 

	logger.Info("---> registering user handler...")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("---> starting application...")

	var listener net.Listener
	var ListenErr error

	// sock === socket 
	if cfg.Listen.Type == "sock" { 	
		logger.Info("---> detecting application path...")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("---> creating socket...")
		socketPath := path.Join(appDir, "app.sock")

		logger.Debugf("CREATED SUCCESFULLY | socket path: %s", socketPath)

		logger.Info("---> listening unix socket...")
		listener, ListenErr = net.Listen("unix", socketPath)
		logger.Infof("application is listening unix sockets... %s", socketPath)
		
	} else {
		logger.Info("---> listening tcp...")
		listener, ListenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("application is running... on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if ListenErr != nil {
		logger.Fatal(ListenErr)
	}

	server := &http.Server {
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
