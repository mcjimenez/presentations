package routes

import (
	"context"
	context2 "globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-openapi-bp/context"
	"net/http"
	"strings"
)

const cnPremium = "mybotservice.com"
const clientIsPremiumCtx = "isPremium"

func (s *HolaMundo) CheckPremium(res http.ResponseWriter, req *http.Request) (context.Context, error) {
	ctx := context2.New(req.Context())

	isPremium := strings.ToLower(ctx.ClientCert().Subject.CommonName) == cnPremium

	return ctx.SetValue(clientIsPremiumCtx, isPremium).Context(), nil
}