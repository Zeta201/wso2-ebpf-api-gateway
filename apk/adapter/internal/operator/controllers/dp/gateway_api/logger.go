package gateway_api

import (
	"github.com/cilium/proxy/pkg/logging"
	"github.com/cilium/proxy/pkg/logging/logfields"
)

const (
	Subsys = "gateway-controller"

	gatewayClass = "gatewayClass"
	gateway      = "gateway"
	httpRoute    = "httpRoute"
	grpcRoute    = "grpcRoute"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, Subsys)
