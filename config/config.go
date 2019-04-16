package config

import (
	"github.com/studtool/common/config"
)

var (
	ServerPort = config.NewStringDefault("STUDTOOL_VK_AUTH_SERVICE_PORT", "80")

	ShouldAllowCORS   = config.NewFlagDefault("STUDTOOL_VK_AUTH_SERVICE_SHOULD_ALLOW_CORS", true)
	ShouldLogRequests = config.NewFlagDefault("STUDTOOL_VK_AUTH_SERVICE_SHOULD_LOG_REQUEST", true)
)
