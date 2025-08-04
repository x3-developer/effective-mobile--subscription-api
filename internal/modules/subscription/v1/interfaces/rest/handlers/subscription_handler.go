package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"subscriptions/internal/modules/subscription/v1/application/dto"
	"subscriptions/internal/modules/subscription/v1/application/mapper"
	"subscriptions/internal/modules/subscription/v1/application/usecase"
	"subscriptions/internal/shared/lib/req"
	"subscriptions/internal/shared/lib/res"
	"subscriptions/internal/shared/lib/val"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetTotalCost(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	createUC       usecase.CreateUseCase
	getAllUC       usecase.GetAllUseCase
	getByIdUC      usecase.GetByIdUseCase
	updateUC       usecase.UpdateUseCase
	deleteUC       usecase.DeleteUseCase
	getTotalCostUC usecase.GetTotalCostUseCase
}

func NewHandler(
	createUC usecase.CreateUseCase,
	getAllUC usecase.GetAllUseCase,
	getByIdUC usecase.GetByIdUseCase,
	updateUC usecase.UpdateUseCase,
	deleteUC usecase.DeleteUseCase,
	getTotalCostUC usecase.GetTotalCostUseCase,
) Handler {
	return &handler{
		createUC:       createUC,
		getAllUC:       getAllUC,
		getByIdUC:      getByIdUC,
		updateUC:       updateUC,
		deleteUC:       deleteUC,
		getTotalCostUC: getTotalCostUC,
	}
}

// Create creates a new subscription
//
//	@Summary		Create a new subscription
//	@Description	Create a new subscription
//	@Tags			Subscription
//	@Accept			json
//	@Produce		json
//	@Param			subscription	body		dto.CreateDTO						true	"Subscription to create"
//	@Success		201				{object}	docsResponse.SubscriptionCreate201	"Subscription created successfully"
//	@Failure		400				{object}	docsResponse.SubscriptionCreate400	"Bad Request or Validation Error"
//	@Failure		500				{object}	docsResponse.Response500			"Server Error"
//	@Router			/api/v1/subscription [post]
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	createDto, err := req.DecodeBody[dto.CreateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	errFields := val.ValidateDTO(createDto)
	if errFields != nil {
		msg := "validation errors occurred"
		res.SendValidationError(w, http.StatusBadRequest, msg, res.BadRequest, errFields)
		return
	}

	model := mapper.ToModelFromCreateDTO(&createDto)
	if model == nil {
		msg := "invalid subscription data"
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	createdModel, errFields, err := h.createUC.Execute(ctx, model)
	if err != nil {
		msg := fmt.Sprintf("failed to create subscription: %v", err)
		res.SendError(w, http.StatusBadRequest, msg, res.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		res.SendValidationError(w, http.StatusBadRequest, msg, res.BadRequest, errFields)
		return
	}

	responseDTO := mapper.ToResponseDTOFromModel(createdModel)

	res.SendSuccess(w, http.StatusCreated, responseDTO)
}

// GetAll retrieves all subscriptions
//
//	@Summary		Get all subscriptions
//	@Description	Retrieve all subscriptions
//	@Tags			Subscription
//	@Produce		json
//	@Success		200	{object}	docsResponse.SubscriptionList200	"List of subscriptions"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/subscription [get]
func (h *handler) GetAll(w http.ResponseWriter) {
	ctx := context.Background()
	models, err := h.getAllUC.Execute(ctx)
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve subscriptions: %v", err)
		res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
		return
	}

	responseDTOs := make([]*dto.ResponseDTO, len(models))
	for i, model := range models {
		responseDTOs[i] = mapper.ToResponseDTOFromModel(model)
	}

	res.SendSuccess(w, http.StatusOK, responseDTOs)
}

// GetById retrieves a subscription by its ID
//
//	@Summary		Get subscription by ID
//	@Description	Retrieve subscription by its ID
//	@Tags			Subscription
//	@Produce		json
//	@Param			id	path		int									true	"Subscription ID"
//	@Success		200	{object}	docsResponse.SubscriptionGetById200	"Subscription found"
//	@Failure		400	{object}	docsResponse.Response400			"Invalid ID"
//	@Failure		404	{object}	docsResponse.Response404			"Subscription not found"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/subscription/{id} [get]
func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing subscription id"
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid subscription id: %s", idStr)
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	model, err := h.getByIdUC.Execute(ctx, uint(id))
	if model == nil {
		msg := fmt.Sprintf("subscription with id %d not found", id)
		res.SendError(w, http.StatusNotFound, msg, res.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to retrieve subscription: %v", err)
		res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
		return
	}

	responseDTO := mapper.ToResponseDTOFromModel(model)

	res.SendSuccess(w, http.StatusOK, responseDTO)
}

// Update updates a subscription by its ID
//
//	@Summary		Update subscription
//	@Description	Update subscription by ID
//	@Tags			Subscription
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int									true	"Subscription ID"
//	@Param			subscription	body		dto.UpdateDTO						true	"Subscription update payload"
//	@Success		200				{object}	docsResponse.SubscriptionUpdate200	"Subscription updated"
//	@Failure		400				{object}	docsResponse.SubscriptionUpdate400	"Bad request or validation error"
//	@Failure		500				{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/subscription/{id} [patch]
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing subscription id"
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid subscription id: %s", idStr)
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	updateDto, err := req.DecodeBody[dto.UpdateDTO](r.Body)
	if err != nil {
		msg := fmt.Sprintf("invalid request body: %v", err)
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	errFields := val.ValidateDTO(updateDto)
	if errFields != nil {
		msg := "validation errors occurred"
		res.SendValidationError(w, http.StatusBadRequest, msg, res.BadRequest, errFields)
		return
	}

	model, err := h.getByIdUC.Execute(ctx, uint(id))
	if err != nil {
		msg := fmt.Sprintf("failed to update subscription: %v", err)
		res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
		return
	}
	if model == nil {
		msg := fmt.Sprintf("subscription with id %d not found", id)
		res.SendError(w, http.StatusNotFound, msg, res.NotFound)
		return
	}

	model = mapper.ToModelFromUpdateDTO(&updateDto, model)
	if model == nil {
		msg := "invalid subscription data"
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	updatedModel, errFields, err := h.updateUC.Execute(ctx, uint(id), model)
	if err != nil {
		msg := fmt.Sprintf("failed to update subscription: %v", err)
		res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
		return
	}
	if errFields != nil {
		msg := "validation errors occurred"
		res.SendValidationError(w, http.StatusBadRequest, msg, res.BadRequest, errFields)
		return
	}

	responseDTO := mapper.ToResponseDTOFromModel(updatedModel)

	res.SendSuccess(w, http.StatusOK, responseDTO)
}

// Delete deletes a subscription by its ID
//
//	@Summary		Delete subscription
//	@Description	Delete subscription by ID
//	@Tags			Subscription
//	@Produce		json
//	@Param			id	path		int									true	"Subscription ID"
//	@Success		200	{object}	docsResponse.SubscriptionDelete200	"Subscription deleted"
//	@Failure		400	{object}	docsResponse.Response400			"Invalid ID"
//	@Failure		404	{object}	docsResponse.Response404			"Subscription not found"
//	@Failure		500	{object}	docsResponse.Response500			"Server error"
//	@Router			/api/v1/subscription/{id} [delete]
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	if idStr == "" {
		msg := "missing subscription id"
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid subscription id: %s", idStr)
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	model, err := h.deleteUC.Execute(ctx, uint(id))
	if model == nil {
		msg := fmt.Sprintf("subscription with id %d not found", id)
		res.SendError(w, http.StatusNotFound, msg, res.NotFound)
		return
	}
	if err != nil {
		msg := fmt.Sprintf("failed to delete subscription: %v", err)
		res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
		return
	}

	responseDTO := mapper.ToResponseDTOFromModel(model)

	res.SendSuccess(w, http.StatusOK, responseDTO)
}

// GetTotalCost retrieves a total cost of subscriptions by a given date range
//
//	@Summary		Get total cost of subscriptions
//	@Description	Retrieve total cost of subscriptions by a given date range
//	@Tags			Subscription
//	@Produce		json
//	@Param			startDate			query		string										false	"Start date in format YYYY-MM-DD"
//	@Param			endDate				query		string										false	"End date in format YYYY-MM-DD"
//	@param			subscriptionName	query		string										false	"Filter by subscription name"
//	@Param			userId				query		int											false	"Filter by user ID"
//	@Success		200					{object}	docsResponse.SubscriptionGetTotalCost200	"Total cost of subscriptions"
//	@Failure		400					{object}	docsResponse.Response400					"Validation error or bad request"
//	@Failure		500					{object}	docsResponse.Response500					"Server error"
//	@Router			/api/v1/subscription/total-cost [get]
func (h *handler) GetTotalCost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	totalCostDTO, err := req.DecodeQuery[dto.TotalCostDTO](r.URL.Query())
	if err != nil {
		msg := fmt.Sprintf("invalid query parameters: %v", err)
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	errFields := val.ValidateDTO(totalCostDTO)
	if errFields != nil {
		msg := "validation errors occurred"
		res.SendValidationError(w, http.StatusBadRequest, msg, res.BadRequest, errFields)
		return
	}

	totalCostFilterVO := mapper.ToTotalCostFilterVOFromDTO(&totalCostDTO)
	if totalCostFilterVO == nil {
		msg := "invalid total cost filter"
		res.SendError(w, http.StatusBadRequest, msg, res.BadRequest)
		return
	}

	totalCost, err := h.getTotalCostUC.Execute(ctx, totalCostFilterVO)
	if err != nil {
		msg := fmt.Sprintf("failed to calculate total cost: %v", err)
		res.SendError(w, http.StatusInternalServerError, msg, res.ServerError)
		return
	}

	res.SendSuccess(w, http.StatusOK, totalCost)
}
