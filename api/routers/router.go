package routers

import (
	"websocket_p4/api/controller"
	"websocket_p4/common/configs"
	"websocket_p4/websocket"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	user *controller.ControllerUser,
	managerClient *websocket.ManagerClient,
	cf *configs.Configs,
) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())
	//middlewares.recovy
	engine.Use(gin.Recovery())

	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/add", user.AddAcount)
	r.GET("/find", user.GetUserByUserName)
	r.POST("/login", user.Login)
	r.GET("/all", user.GetAllUser)
	//socket
	r.GET("/ws", managerClient.ServerWs)

	return &ApiRouter{
		Engine: engine,
	}
}
