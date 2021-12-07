package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"

	"github.com/mahfuz110244/api-mc/internal/models"
	"github.com/mahfuz110244/api-mc/internal/status/mock"
	"github.com/mahfuz110244/api-mc/pkg/converter"
	"github.com/mahfuz110244/api-mc/pkg/logger"
	"github.com/mahfuz110244/api-mc/pkg/utils"
)

func TestStatusHandlers_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusUC := mock.NewMockUseCase(ctrl)
	statusHandlers := NewStatusHandlers(nil, mockStatusUC, apiLogger)

	handlerFunc := statusHandlers.Create()

	userID := uuid.New()

	status := &models.Status{
		Name:        "estimate",
		Description: "TestStatusHandlers_Create Description",
		// CreatedBy:  "eef6841bf0ee41669721ee8e1ca5a996",
		// UpdatedBy:  "eef6841b-f0ee-4166-9721-ee8e1ca5a996",
	}

	buf, err := converter.AnyToBytesBuffer(status)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/status/create", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	u := &models.User{
		UserID: userID,
	}
	ctxWithValue := context.WithValue(context.Background(), utils.UserCtxKey{}, u)
	req = req.WithContext(ctxWithValue)
	e := echo.New()
	ctx := e.NewContext(req, res)
	ctxWithReqID := utils.GetRequestCtx(ctx)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctxWithReqID, "statusHandlers.Create")
	defer span.Finish()

	mockStatus := &models.Status{
		Name:        "estimate",
		Description: "TestStatusHandlers_Create Description",
	}

	mockStatusUC.EXPECT().Create(ctxWithTrace, gomock.Any()).Return(mockStatus, nil)

	err = handlerFunc(ctx)
	require.NoError(t, err)
}

func TestStatusHandlers_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusUC := mock.NewMockUseCase(ctrl)
	statusHandlers := NewStatusHandlers(nil, mockStatusUC, apiLogger)

	handlerFunc := statusHandlers.Update()

	userID := uuid.New()

	status := &models.Status{
		Name:        "issued",
		Description: "TestStatusHandlers_Update Description",
	}

	buf, err := converter.AnyToBytesBuffer(status)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/status/f8a3cc26-fbe1-4713-98be-a2927201356e", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	u := &models.User{
		UserID: userID,
	}
	ctxWithValue := context.WithValue(context.Background(), utils.UserCtxKey{}, u)
	req = req.WithContext(ctxWithValue)
	e := echo.New()
	ctx := e.NewContext(req, res)
	ctx.SetParamNames("status_id")
	ctx.SetParamValues("f8a3cc26-fbe1-4713-98be-a2927201356e")
	ctxWithReqID := utils.GetRequestCtx(ctx)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctxWithReqID, "statusHandlers.Update")
	defer span.Finish()

	mockStatus := &models.Status{
		Name:        "issued",
		Description: "TestStatusHandlers_Update Description",
	}

	mockStatusUC.EXPECT().Update(ctxWithTrace, gomock.Any()).Return(mockStatus, nil)

	err = handlerFunc(ctx)
	require.NoError(t, err)
}

func TestStatusHandlers_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockStatusUC := mock.NewMockUseCase(ctrl)
	statusHandlers := NewStatusHandlers(nil, mockStatusUC, apiLogger)

	handlerFunc := statusHandlers.GetByID()

	userID := uuid.New()
	statusID := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/status/f8a3cc26-fbe1-4713-98be-a2927201356e", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	u := &models.User{
		UserID: userID,
	}
	ctxWithValue := context.WithValue(context.Background(), utils.UserCtxKey{}, u)
	req = req.WithContext(ctxWithValue)
	e := echo.New()
	ctx := e.NewContext(req, res)
	ctx.SetParamNames("status_id")
	ctx.SetParamValues(statusID.String())
	ctxWithReqID := utils.GetRequestCtx(ctx)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctxWithReqID, "statusHandlers.GetByID")
	defer span.Finish()

	mockStatus := &models.StatusBase{
		ID:          statusID,
		Name:        "issued",
		Description: "TestStatusHandlers_Update Description",
	}

	mockStatusUC.EXPECT().GetStatusByID(ctxWithTrace, statusID).Return(mockStatus, nil)

	err := handlerFunc(ctx)
	require.NoError(t, err)
}
