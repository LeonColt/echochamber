package echochamber

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Controller struct{}

func (*Controller) Register(parent *echo.Echo) {}

func (ptr *Controller) HandleError(ctx echo.Context, err error) error {
	if httpErr, ok := err.(HttpException); ok {
		if err := ctx.JSON(httpErr.GetStatus(), HTTPError{
			Code:    httpErr.GetStatus(),
			Message: err.Error(),
		}); err != nil {
			return err
		}
		return nil
	} else {
		return ptr.InternalServerError(ctx, err)
	}
}

func (*Controller) OkHTMLBlob(ctx echo.Context, html []byte) error {
	if err := ctx.HTMLBlob(http.StatusOK, html); err != nil {
		log.Warn(fmt.Sprintf("error serving HTML: %#v", err))
		return err
	}
	return nil
}

func (*Controller) OkHTML(ctx echo.Context, html string) error {
	if err := ctx.HTML(http.StatusOK, html); err != nil {
		log.Warn(fmt.Sprintf("error serving HTML: %#v", err))
		return err
	}
	return nil
}

func (*Controller) OkJSON(ctx echo.Context, data interface{}) error {
	if err := ctx.JSON(http.StatusOK, data); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %#v", err))
		return err
	}
	return nil
}

func (*Controller) Created(ctx echo.Context, data interface{}) error {
	if err := ctx.JSON(http.StatusCreated, data); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) NoContent(ctx echo.Context) error {
	if err := ctx.NoContent(http.StatusNoContent); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) BadRequestError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, HTTPError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) Unauthorized(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusBadRequest, HTTPError{
		Code:    http.StatusUnauthorized,
		Message: "please sign in first",
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) UnauthorizedError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, HTTPError{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) Forbidden(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusBadRequest, HTTPError{
		Code:    http.StatusForbidden,
		Message: "you are not allowed to do this",
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) ForbiddenError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, HTTPError{
		Code:    http.StatusForbidden,
		Message: err.Error(),
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) NotFoundError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusNotFound, HTTPError{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) ConflictError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusConflict, HTTPError{
		Code:    http.StatusConflict,
		Message: err.Error(),
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) InternalServerError(ctx echo.Context, err error) error {
	log.Warn(err)
	if err := ctx.JSON(http.StatusInternalServerError, HTTPError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}

func (*Controller) ServiceUnavailableError(ctx echo.Context, err error) error {
	log.Warn(err)
	if err := ctx.JSON(http.StatusServiceUnavailable, HTTPError{
		Code:    http.StatusServiceUnavailable,
		Message: "Service Unavailable",
	}); err != nil {
		log.Warn(fmt.Sprintf("error serving JSON: %v", err))
		return err
	}
	return nil
}
