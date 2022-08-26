package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	authhttp "github.com/skinnykaen/robbo_student_personal_account.git/package/auth/http"
	cohortshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/cohorts/http"
	coursepackethttp "github.com/skinnykaen/robbo_student_personal_account.git/package/coursePacket/http"
	courseshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/courses/http"
	projectpagehttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage/http"
	projectshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/projects/http"
	robbogrouphttp "github.com/skinnykaen/robbo_student_personal_account.git/package/robboGroup/http"
	robbounitshttp "github.com/skinnykaen/robbo_student_personal_account.git/package/robboUnits/http"
	usershtpp "github.com/skinnykaen/robbo_student_personal_account.git/package/users/http"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(lifecycle fx.Lifecycle,
	authhandler authhttp.Handler,
	projecthttp projectshttp.Handler,
	projectpagehttp projectpagehttp.Handler,
	coursehttp courseshttp.Handler,
	cohortshttp cohortshttp.Handler,
	usershttp usershtpp.Handler,
	robbounitshttp robbounitshttp.Handler,
	robbogrouphttp robbogrouphttp.Handler,
	coursepackethttp coursepackethttp.Handler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				router := gin.Default()
				router.Use(
					gin.Recovery(),
					gin.Logger(),
				)
				authhandler.InitAuthRoutes(router)
				projecthttp.InitProjectRoutes(router)
				projectpagehttp.InitProjectRoutes(router)
				coursehttp.InitCourseRoutes(router)
				cohortshttp.InitCohortRoutes(router)
				usershttp.InitUsersRoutes(router)
				robbounitshttp.InitRobboUnitsRoutes(router)
				robbogrouphttp.InitRobboGroupRoutes(router)
				coursepackethttp.InitCoursePacketRoutes(router)
				server := &http.Server{
					Addr: viper.GetString("server.address"),
					Handler: cors.New(
						// TODO make config
						cors.Options{
							AllowedOrigins:   []string{"http://0.0.0.0:3030", "http://0.0.0.0:8601", "http://localhost:3030"},
							AllowCredentials: true,
							AllowedMethods: []string{
								http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodOptions,
							},
							//AllowedHeaders: []string{"*"},
							AllowedHeaders: []string{
								"Origin", "X-Requested-With", "Content-Type", "Accept", "Set-Cookie", "Authorization",
							},
						},
					).Handler(router),
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}
				go func() {
					if err := server.ListenAndServe(); err != nil {
						log.Fatalf("Failed to listen and serve", err)
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}
