package routes

import (
	"fmt"
	"globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-openapi-bp"
	"net/http"
)

const (
	msgSpanish = "Hola %s"
	msgEnglish = "Hello %s"
)

type Regards map[string]string

// get /v1/regards/{person}
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
	person := params["person"]
	logger.Debug(SummaryKey, fmt.Sprintf("Name: %s", person))

	msg := Regards{
		"english" : fmt.Sprintf(msgEnglish, person),
	}
	if isPremium := req.Context().Value(clientIsPremiumCtx).(bool); isPremium {
		msg["spanish"] = fmt.Sprintf(msgSpanish, person)
	}

	_ = wRes.SendJSON(http.StatusOK, msg)
}
