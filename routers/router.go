package routers

import (
	"cloudreve/middleware"
	"cloudreve/pkg/conf"
	"cloudreve/routers/controllers"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// SetupTestEngine 设置测试用gin.Engine
func SetupTestEngine(engine *gin.Engine) {
	r = engine
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	if r == nil {
		r = gin.Default()
	}

	/*
		中间件
	*/
	r.Use(middleware.Session(conf.SystemConfig.SessionSecret))

	// 测试模式加加入Mock助手中间件
	if gin.Mode() == gin.TestMode {
		r.Use(middleware.MockHelper())
	}

	r.Use(middleware.CurrentUser())

	/*
		路由
	*/
	v3 := r.Group("/Api/V3")
	{
		// 测试用路由
		v3.GET("Ping", controllers.Ping)
		// 用户登录
		v3.POST("User/Session", controllers.UserLogin)
		// 验证码
		v3.GET("Captcha", controllers.Captcha)

		// 需要登录保护的
		auth := v3.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// 用户类
			user := auth.Group("User")
			{
				// 当前登录用户信息
				user.GET("Me", controllers.UserMe)
			}

		}

	}
	return r
}