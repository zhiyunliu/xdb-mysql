package demos

import "github.com/zhiyunliu/glue/server/api"

func HandleServices(apiSrv *api.Server) {
	apiSrv.Handle("/tests/query", Query)
}
