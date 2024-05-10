package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/savi2w/pupper/service"
	"github.com/savi2w/pupper/util/resutil"
	"github.com/savi2w/pupper/validation"
)

type Controller struct {
	resutil *resutil.ResUtil
	svc     *service.Service
}

func New(svc *service.Service, resutil *resutil.ResUtil) *Controller {
	return &Controller{
		resutil: resutil,
		svc:     svc,
	}
}

func (ctrl *Controller) HandleNewUser(ctx echo.Context) error {
	req, err := validation.VerifyNewUserRequest(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	if err := ctrl.svc.User.New(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(nil, nil, http.StatusCreated))
}
