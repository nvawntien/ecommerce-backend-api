package routers

import "go-ecommerce-backend-api/internal/routers/user"

type RouterGroup struct {
	UserRouter user.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
