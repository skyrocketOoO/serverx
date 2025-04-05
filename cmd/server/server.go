package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/skyrocketOoO/erx/erx"
	"github.com/skyrocketOoO/serverx/api"
	"github.com/skyrocketOoO/serverx/internal/boot"
	"github.com/skyrocketOoO/serverx/internal/controller"
	authcontroller "github.com/skyrocketOoO/serverx/internal/controller/auth"
	generalcontroller "github.com/skyrocketOoO/serverx/internal/controller/general"
	"github.com/skyrocketOoO/serverx/internal/controller/middleware"
	"github.com/skyrocketOoO/serverx/internal/domain"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
	"github.com/skyrocketOoO/serverx/internal/service"
	authucase "github.com/skyrocketOoO/serverx/internal/usecase/auth"
	generalucase "github.com/skyrocketOoO/serverx/internal/usecase/general"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "The main service command",
	Long:  ``,
	Run:   RunServer,
}

func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.InitAll(); err != nil {
		log.Fatal().Msgf("Initialization failed: %v", err)
	}

	handlers, err := newHandlers()
	if err != nil {
		log.Error().Err(err).Msg("Error creating handlers")
		return
	}

	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.ErrorHttp)
	api.RegisterAPIHandlers(router, handlers)
	port, _ := cmd.Flags().GetString("port")
	server := newHTTPServer(router, port)

	go func() {
		log.Info().Msgf("Starting server on port %s...", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Server failed")
			return
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server forced to shutdown: %v", err)
	} else {
		log.Info().Msg("Server shut down gracefully.")
	}
}

func init() {
	Cmd.Flags().StringP("port", "p", "8080", "port")
	Cmd.Flags().
		StringVarP(&domain.Database, `database`, "d", "postgres", `"postgres", "mysql"`)
	Cmd.Flags().
		StringVarP(&domain.Env, `env`, "e", "dev", `"local", "dev", "prod"`)

	Cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		validDatabases := map[string]bool{"postgres": true, "mysql": true}
		if !validDatabases[domain.Database] {
			return erx.Errorf("invalid database value: %s. Must be one of: postgres, mysql",
				domain.Database)
		}

		validEnvs := map[string]bool{"dev": true, "prod": true}
		if !validEnvs[domain.Env] {
			return erx.Errorf(
				"invalid environment value: %s. Must be one of: dev, prod",
				domain.Env,
			)
		}

		port, _ := cmd.Flags().GetString("port")
		if _, err := strconv.Atoi(port); err != nil || port == "" {
			return erx.Errorf("invalid port value: %s. Must be a valid number", port)
		}

		return nil
	}
}

func newHandlers() (*controller.Handler, error) {
	cognitoCli, err := service.NewCognito(context.TODO())
	if err != nil {
		return nil, er.W(err)
	}

	authUsecase := authucase.New(cognitoCli)
	generalUsecase := generalucase.New()

	authHandler := authcontroller.NewHandler(authUsecase)
	generalHandler := generalcontroller.NewHandler(generalUsecase)

	handlers := controller.NewHandler(authHandler, generalHandler)

	return handlers, nil
}

func newHTTPServer(router *gin.Engine, port string) *http.Server {
	return &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}
}
