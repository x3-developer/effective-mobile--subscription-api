package router

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"subscriptions/config"
	_ "subscriptions/docs"
	"subscriptions/internal/modules/subscription/v1/interfaces/rest"
	"subscriptions/internal/shared/application/container"
	"subscriptions/internal/shared/application/middleware"
)

func NewRouter(cfg *config.Config, container *container.Container) http.Handler {
	mux := http.NewServeMux()
	v1 := http.NewServeMux()

	rest.SubscriptionV1Routes(v1, container)

	apiHandler := middleware.ChainMiddleware(
		v1,
		middleware.CORSMiddleware(cfg.CORS),
		middleware.Recoverer,
		middleware.Logger,
	)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiHandler))

	if cfg.AppEnv != "production" {
		mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	}

	return mux
}
