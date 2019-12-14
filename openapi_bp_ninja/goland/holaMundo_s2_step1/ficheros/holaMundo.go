package routes

import "globaldevtools.bbva.com/bbva_global_secaas_swat_components/go-utils/logger"

const (
	SummaryKey = logger.SummaryKey
)

type HolaMundo struct {
	config *Config
}

func New (config Config) *HolaMundo {
	return &HolaMundo{
		config: &config,
	}
}