package echochamber

import (
	"net/http"

	"github.com/LeonColt/ez"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type MixinController struct{}

func (ptr *MixinController) BindAndValidate(ctx echo.Context, input interface{}) error {
	if err := ctx.Bind(input); err != nil {
		return ez.New(ez.ErrorCodeInvalidArgument, err.Error())
	}
	if err := ctx.Validate(input); err != nil {
		return ez.New(ez.ErrorCodeInvalidArgument, err.Error())
	}
	return nil
}

func (ptr *MixinController) HandleError(ctx echo.Context, err error) error {
	if httpErr, ok := err.(*ez.Error); ok {
		if err := ctx.JSON(httpErr.GetHttpStatusCode(), HTTPError{
			Code:    httpErr.GetHttpStatusCode(),
			Message: err.Error(),
		}); err != nil {
			return err
		}
		return nil
	} else {
		return ptr.InternalServerError(ctx, err)
	}
}

func (*MixinController) OkHTMLBlob(ctx echo.Context, html []byte) error {
	if err := ctx.HTMLBlob(http.StatusOK, html); err != nil {
		log.Error().Stack().Err(err).Msg("error serving HTML Blob")
		return err
	}
	return nil
}

func (*MixinController) OkHTML(ctx echo.Context, html string) error {
	if err := ctx.HTML(http.StatusOK, html); err != nil {
		log.Error().Stack().Err(err).Msg("error serving HTML")
		return err
	}
	return nil
}

func (*MixinController) OkJSON(ctx echo.Context, data interface{}) error {
	if err := ctx.JSON(http.StatusOK, data); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) OkBlob(ctx echo.Context, contentType string, data []byte) error {
	if err := ctx.Blob(http.StatusOK, contentType, data); err != nil {
		log.Error().Stack().Err(err).Msg("error serving blob")
		return err
	}
	return nil
}

func (ptr *MixinController) OkTextPlain(ctx echo.Context, data string) error {
	return ptr.OkBlob(ctx, "text/plain", []byte(data))
}

func (*MixinController) Created(ctx echo.Context, data interface{}) error {
	if err := ctx.JSON(http.StatusCreated, data); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) NoContent(ctx echo.Context) error {
	if err := ctx.NoContent(http.StatusNoContent); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) BadRequestError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusBadRequest, HTTPError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) Unauthorized(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusUnauthorized, HTTPError{
		Code:    http.StatusUnauthorized,
		Message: "please sign in first",
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) UnauthorizedError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusUnauthorized, HTTPError{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) Forbidden(ctx echo.Context) error {
	if err := ctx.JSON(http.StatusForbidden, HTTPError{
		Code:    http.StatusForbidden,
		Message: "you are not allowed to do this",
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) ForbiddenError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusForbidden, HTTPError{
		Code:    http.StatusForbidden,
		Message: err.Error(),
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) NotFoundError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusNotFound, HTTPError{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) ConflictError(ctx echo.Context, err error) error {
	if err := ctx.JSON(http.StatusConflict, HTTPError{
		Code:    http.StatusConflict,
		Message: err.Error(),
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) InternalServerError(ctx echo.Context, err error) error {
	log.Error().Stack().Err(err).Msg("error occurred internal server error")
	if err := ctx.JSON(http.StatusInternalServerError, HTTPError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}

func (*MixinController) ServiceUnavailableError(ctx echo.Context, err error) error {
	log.Error().Stack().Err(err).Msg("error occured unavailable error")
	if err := ctx.JSON(http.StatusServiceUnavailable, HTTPError{
		Code:    http.StatusServiceUnavailable,
		Message: "Service Unavailable",
	}); err != nil {
		log.Error().Stack().Err(err).Msg("error serving JSON")
		return err
	}
	return nil
}
