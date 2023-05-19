package huggingface

import (
	"code.gitea.io/gitea/modules/web"
	"context"
)

func Routes(_ context.Context, prefix string) *web.Route {
	m := web.NewRoute()
	m.Get(prefix+"/api/whoami-v2", WhoAMIHandler)
	return m
}
