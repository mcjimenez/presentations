package routes

import (
	"context"
	context2 "globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-openapi-bp/context"
	"net/http"
	"strings"
)

const cnBoss = "mybotservice.com"
const clientIsTheBossCtx = "isTheBoss"

func (s *HolaMundo) CheckBossCert(res http.ResponseWriter, req *http.Request) (context.Context, error) {
	ctx := context2.New(req.Context())
	isABoss := false
	if cn := ctx.ClientCert().Subject.CommonName; strings.ToLower(cn) == cnBoss {
		isABoss = true
	}
	return ctx.SetValue(clientIsTheBossCtx, isABoss).Context(), nil
}