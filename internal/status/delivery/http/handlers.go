package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/mahfuz110244/api-mc/config"
	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status"
	"github.com/mahfuz110244/api-mc/pkg/httpErrors"
	"github.com/mahfuz110244/api-mc/pkg/logger"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

// Status handlers
type statusHandlers struct {
	cfg      *config.Config
	statusUC status.UseCase
	logger   logger.Logger
}

// NewStatusHandlers Status handlers constructor
func NewStatusHandlers(cfg *config.Config, statusUC status.UseCase, logger logger.Logger) status.Handlers {
	return &statusHandlers{cfg: cfg, statusUC: statusUC, logger: logger}
}

// Create godoc
// @Summary Create status
// @Description Create status handler
// @Tags Status
// @Accept json
// @Produce json
// @Success 201 {object} models.Status
// @Router /status/create [post]
func (h statusHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "statusHandlers.Create")
		defer span.Finish()

		n := &models.Status{}
		if err := c.Bind(n); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdStatus, err := h.statusUC.Create(ctx, n)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdStatus)
	}
}

// Update godoc
// @Summary Update status
// @Description Update status handler
// @Tags Status
// @Accept json
// @Produce json
// @Param id path int true "status_id"
// @Success 200 {object} models.Status
// @Router /status/{id} [put]
func (h statusHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "statusHandlers.Update")
		defer span.Finish()

		statusUUID, err := uuid.Parse(c.Param("status_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		n := &models.Status{}
		if err = c.Bind(n); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		n.StatusID = statusUUID

		updatedStatus, err := h.statusUC.Update(ctx, n)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedStatus)
	}
}

// GetByID godoc
// @Summary Get by id status
// @Description Get by id status handler
// @Tags Status
// @Accept json
// @Produce json
// @Param id path int true "status_id"
// @Success 200 {object} models.Status
// @Router /status/{id} [get]
func (h statusHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "statusHandlers.GetByID")
		defer span.Finish()

		statusUUID, err := uuid.Parse(c.Param("status_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		statusByID, err := h.statusUC.GetStatusByID(ctx, statusUUID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, statusByID)
	}
}

// Delete godoc
// @Summary Delete status
// @Description Delete by id status handler
// @Tags Status
// @Accept json
// @Produce json
// @Param id path int true "status_id"
// @Success 200 {string} string	"ok"
// @Router /status/{id} [delete]
func (h statusHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "statusHandlers.Delete")
		defer span.Finish()

		statusUUID, err := uuid.Parse(c.Param("status_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.statusUC.Delete(ctx, statusUUID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetStatus godoc
// @Summary Get all status
// @Description Get all status with pagination
// @Tags Status
// @Accept json
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.StatusList
// @Router /status [get]
func (h statusHandlers) GetStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "statusHandlers.GetStatus")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		statusList, err := h.statusUC.GetStatus(ctx, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, statusList)
	}
}

// SearchByTitle godoc
// @Summary Search by title
// @Description Search status by title
// @Tags Status
// @Accept json
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.StatusList
// @Router /status/search [get]
func (h statusHandlers) SearchByTitle() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "statusHandlers.SearchByTitle")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		statusList, err := h.statusUC.SearchByTitle(ctx, c.QueryParam("title"), pq)

		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, statusList)
	}
}
