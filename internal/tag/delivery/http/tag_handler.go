package http

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/util"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// TagHandler ...
type TagHandler struct {
	TagUseCase domain.TagUseCase
}

// NewTagHandler ...
func NewTagHandler(e *echo.Echo, us domain.TagUseCase) {
	handler := &TagHandler{
		TagUseCase: us,
	}
	// Get Operation
	e.GET("/tags", handler.GetAll)
	e.GET("/tags/:id", handler.GetByID)
	e.GET("/tags/:id/", handler.GetByID)

	// Create/Add Operation
	e.POST("/tags", handler.CreateTag)

	// Update Operation
	e.PUT("/tags/:id", handler.GetByID)
	e.PUT("/tags/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/tags/:id", handler.GetByID)
}

// GetAll ...
func (c *TagHandler) GetAll(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	start, limit := 0, 10
	var err error
	for k, v := range echoContext.QueryParams() {
		switch k {
		case "start":
			val := strings.TrimSpace(v[0])
			if start, err = strconv.Atoi(val); err != nil {
				return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
			}
		case "limit":
			val := strings.TrimSpace(v[0])
			if limit, err = strconv.Atoi(val); err != nil {
				return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
			}
		}
	}

	list, err := c.TagUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// GetByID ...
func (c *TagHandler) GetByID(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	list, err := c.TagUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// CreateTag ...
func (c *TagHandler) CreateTag(echoContext echo.Context) error {
	var tag domain.Tag
	err := echoContext.Bind(&tag)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&tag); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.TagUseCase.CreateTag(ctx, &tag)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, tag)

}

// UpdateTag ...
func (c *TagHandler) UpdateTag(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	var tag domain.Tag
	err = echoContext.Bind(&tag)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&tag); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.TagUseCase.UpdateTag(ctx, &tag, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, tag)

}

// DeleteTag godoc
// @Summary Delete existing tag
// @Description delete tag by given parameter id
// @Tags tags
// @Accept */*
// @Produce json
// @Param id path int true "Tag Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /courses/{id} [delete]
func (c *TagHandler) DeleteTag(echoContext echo.Context) error {
	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := echoContext.Request().Context()

	err = c.TagUseCase.DeleteTag(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusNoContent)
}
