package routes

import (
	"globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-openapi-bp"
)

type ServiceConfig struct {
	ValidOrigins []string `json:"validOrigins"`
}

type Config struct {
	OpenAPIOptions *openapi.Options
	ServiceConfig *ServiceConfig `json:"serviceConfig"`
}