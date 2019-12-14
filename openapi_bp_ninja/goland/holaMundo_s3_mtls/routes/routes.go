package routes

import (
	"fmt"
	"globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-openapi-bp"
	"net/http"
)

// get /v1/regards/{name}
func (s *HolaMundo) GetRegards(res http.ResponseWriter, req *http.Request) {
	logger, params, wRes, err := openapi.CommonAPISetup(nil, true, res, req)
	if err != nil { // Error processing already done. Or not.
		_ = wRes.SendJSON(http.StatusBadRequest, openapi.ErrorInfo{
			Code: http.StatusBadRequest,
			Description: fmt.Sprintf("There was an error: %+v", err),
		})
		return
	}

	// Define and extract the parameter names
	name := params["name"]
	logger.Debug(SummaryKey, fmt.Sprintf("name: %s", name))

	msg := fmt.Sprintf("Hello %s", name)

	_ = wRes.SendJSON(http.StatusOK, msg)
}
