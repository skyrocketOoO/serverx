package api

import (
	"github.com/gin-gonic/gin"
	controller "github.com/skyrocketOoO/serverx/internal/controller"
	"github.com/skyrocketOoO/serverx/internal/controller/middleware"
	"github.com/skyrocketOoO/serverx/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterAPIHandlers(r *gin.Engine, h *controller.Handler, cognito *service.Cognito) {
	r.Use(middleware.Cors())
	r.Use(middleware.ErrorHttp)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	vr := r.Group("/v1")
	{
		vr.GET("/ping", h.General.Ping)
		vr.GET("/healthy", h.General.Healthy)

		vr.POST("/login", h.Auth.Login)
		vr.POST("/set-new-password", h.Auth.SetNewPassword)

		vr.POST("/sign-up", h.Auth.SignUp)
		vr.POST("/confirm-sign-up", h.Auth.ConfirmSignUp)

		vr.POST("/forgot-password", h.Auth.ForgotPassword)
		vr.POST("/confirm-forgot-password", h.Auth.ConfirmForgotPassword)

		vr.POST("/resend-confirmation-code", h.Auth.ResendConfirmationCode)
		vr.POST("/refresh-token", h.Auth.RefreshToken)
	}

	pR := r.Group("/")
	pR.Use(middleware.CheckAuthorization(cognito))

	vpr := pR.Group("/v1")
	{
		vpr.POST("/change-password", h.Auth.ChangePassword)
		vpr.POST("/invite-user", h.Auth.InviteUser)
	}
}
