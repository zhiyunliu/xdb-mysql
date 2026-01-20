package main

import (
	sctx "context"

	"github.com/zhiyunliu/glue"
	"github.com/zhiyunliu/glue/log"
	"github.com/zhiyunliu/xdb-mysql/examples/demos"

	_ "github.com/zhiyunliu/glue/contrib/metrics/prometheus"
	"github.com/zhiyunliu/glue/server/api"
	_ "github.com/zhiyunliu/xdb-mysql"
)

func main() {
	apiSrv := api.New("apiserver", api.Log(log.WithRequest(), log.WithResponse()))

	demos.HandleServices(apiSrv)

	app := glue.NewApp(glue.Server(apiSrv),
		glue.StartedHook(func(ctx sctx.Context) error {

			return nil
		}),
		glue.StartingHook(func(ctx sctx.Context) error {
			return nil
		}))
	app.Start()
}
