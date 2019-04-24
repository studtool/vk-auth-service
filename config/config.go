package config

import (
	"github.com/studtool/common/config"
	"github.com/studtool/common/consts"
)

var (
	ServerPort = config.NewStringDefault("STUDTOOL_VK_AUTH_SERVICE_PORT", "80")

	OAuth2ClientId     = config.NewStringDefault("STUDTOOL_VK_OAUTH2_CLIENT_ID", consts.EmptyString)
	OAuth2ClientSecret = config.NewStringDefault("STUDTOOL_VK_OAUTH2_CLIENT_SECRET", consts.EmptyString)

	ShouldAllowCORS   = config.NewFlagDefault("STUDTOOL_VK_AUTH_SERVICE_SHOULD_ALLOW_CORS", true)
	ShouldLogRequests = config.NewFlagDefault("STUDTOOL_VK_AUTH_SERVICE_SHOULD_LOG_REQUESTS", true)
)
