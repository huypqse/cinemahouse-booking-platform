package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"auth-svc/internal/config"
	"auth-svc/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().SetHandlers(glog.HandlerJson)
			glog.SetHandlers(glog.HandlerJson)

			config.InitConfig(ctx)

			s := g.Server()
			s.Logger().SetHandlers(glog.HandlerJson)
			//set up for prometheus
			// metrics, err := metrics.NewMetrics("auth-svc")
			// if err != nil {
			// 	g.Log().Fatal(ctx, "init metrics error: %v", err)
			// }
			// metrics.SetSkipPath([]string{
			// 	"/backend/auth/v1/swagger/api.json",
			// })

			// s.Use(service.Middleware().CompressMiddleware)

			// s.Group("/metrics", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(
			// 		service.Middleware().IPWhitelistMiddleware,
			// 	)
			// 	group.GET("/", func(r *ghttp.Request) {
			// 		metrics.GfServeHandler(r)
			// 	})
			// })

			// s.Use(
			// 	metrics.GfMetricsHttpMiddleware(),
			// 	service.Middleware().MiddlewareLog,
			// 	service.Middleware().CORS,
			// 	service.Middleware().SetStatusRes,
			// 	ghttp.MiddlewareHandlerResponse,
			// 	middleware.HiddenInternalErrResponse,
			// )
			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.Writeln("ok")
			})
			s.BindHandler("/backend/auth/v1", func(r *ghttp.Request) {
				r.Response.Writeln("ok")
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
