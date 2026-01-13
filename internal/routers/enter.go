package routers

import (
	"go-ecommerce-backend-api/internal/routers/admin"
	"go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	UserRouter  user.UserRouterGroup
	AdminRouter admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
