package server

import (
	"Ali-DDNS/app/domain_service/internal/domain_task/conf"
	"Ali-DDNS/app/domain_service/internal/domain_task/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// InitSession initial session configs
func InitSession(engine *gin.Engine) {
	rdStore, err := redis.NewStore(
		conf.Session().Size(),
		conf.Session().Network(),
		conf.Session().Address()+":"+conf.Session().Port(),
		conf.Session().Password(),
		conf.Session().Secret(),
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// set session expiration time as 30 minutes
	rdStore.Options(sessions.Options{MaxAge: 60 * 30})
	engine.Use(sessions.Sessions("hominsu-ddns", rdStore))
}

// NewGinServer new a gin engine
func NewGinServer(service *service.DomainTaskService) *gin.Engine {
	router := gin.Default()

	InitSession(router)

	router.Handle("GET", "/ip", func(context *gin.Context) {
		context.String(http.StatusOK, context.ClientIP())
	})

	router.Handle("GET", "/login", service.LoginGet)
	router.Handle("POST", "/login", service.LoginPost)

	router.Handle("POST", "/register", service.RegisterPost)

	// session middleware, ensure that handlers after are checked by the middleware to see if they are logged in
	router.Use(service.SessionMiddleWare)

	router.Handle("DELETE", "/register", service.RegisterDelete)

	router.Handle("GET", "/home", service.Home)
	router.Handle("GET", "/home/:user_name", service.Home)

	router.Handle("GET", "/DomainName/:user_name", service.DomainNameGet)
	router.Handle("POST", "/DomainName/:user_name", service.DomainNamePost)
	router.Handle("DELETE", "/DomainName/:user_name", service.DomainNameDel)

	router.Handle("GET", "/Device/:user_name", service.DeviceGet)
	router.Handle("POST", "/Device/:user_name", service.DevicePost)
	router.Handle("DELETE", "/Device/:user_name", service.DeviceDel)

	router.Handle("DELETE", "/logout", service.Logout)

	return router
}
