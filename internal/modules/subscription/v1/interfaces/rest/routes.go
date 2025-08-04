package rest

import (
	"net/http"
	"subscriptions/internal/modules/subscription/v1/application/usecase"
	infrastructure "subscriptions/internal/modules/subscription/v1/infrastructure/persistence"
	"subscriptions/internal/modules/subscription/v1/interfaces/rest/handlers"
	"subscriptions/internal/shared/application/container"
	"subscriptions/internal/shared/lib/res"
)

func SubscriptionV1Routes(mux *http.ServeMux, container *container.Container) {
	repo := infrastructure.NewRepository(container.PostgresDB)
	createUC := usecase.NewCreateUseCase(repo)
	getAllUC := usecase.NewGetAllUseCase(repo)
	getByIdUC := usecase.NewGetByIdUseCase(repo)
	updateUC := usecase.NewUpdateUseCase(repo)
	deleteUC := usecase.NewDeleteUseCase(repo)
	getTotalCostUC := usecase.NewGetTotalCostUseCase(repo)
	h := handlers.NewHandler(createUC, getAllUC, getByIdUC, updateUC, deleteUC, getTotalCostUC)

	mux.Handle("/subscription",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				h.Create(w, r)
			case http.MethodGet:
				h.GetAll(w)
			default:
				msg := "Method not allowed. Allowed methods: POST, GET"
				res.SendError(w, http.StatusMethodNotAllowed, msg, res.MethodNotAllowed)
			}
		}),
	)

	mux.HandleFunc("/subscription/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetById(w, r)
		case http.MethodPatch:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		default:
			msg := "Method not allowed. Allowed methods: GET, PATCH, DELETE"
			res.SendError(w, http.StatusMethodNotAllowed, msg, res.MethodNotAllowed)
		}
	})

	mux.HandleFunc("/subscription/total-cost", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			msg := "Method not allowed. Allowed method: GET"
			res.SendError(w, http.StatusMethodNotAllowed, msg, res.MethodNotAllowed)
			return
		}
		h.GetTotalCost(w, r)
	})
}
