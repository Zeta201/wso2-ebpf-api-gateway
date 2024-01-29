package dp

import (
	"fmt"

	"github.com/cilium/cilium/pkg/envoy"
	envoy_config_listener "github.com/cilium/proxy/go/envoy/config/listener/v3"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type ListenerMutator func(*envoy_config_listener.Listener) *envoy_config_listener.Listener

const (
	httpConnectionManagerType = "envoy.filters.network.http_connection_manager"
	tcpProxyType              = "envoy.filters.network.tcp_proxy"
	tlsInspectorType          = "envoy.filters.listener.tls_inspector"
	proxyProtocolType         = "envoy.filters.listener.proxy_protocol"
	tlsTransportSocketType    = "envoy.transport_sockets.tls"

	rawBufferTransportProtocol = "raw_buffer"
)

// NewHTTPListenerWithDefaults same as NewListener but with default mutators applied.
func NewHTTPListenerWithDefaults(name string, authProvider *AuthProvider, mutatorFunc ...ListenerMutator) (ciliumv2.XDSResource, error) {

	return NewHTTPListener(name, mutatorFunc...)
}

// NewHTTPListener creates a new Envoy listener with the given name.
// The listener will have both secure and insecure filters.
// Secret Discovery Service (SDS) is used to fetch the TLS certificates.
func NewHTTPListener(name string, mutatorFunc ...ListenerMutator) (ciliumv2.XDSResource, error) {
	var filterChains []*envoy_config_listener.FilterChain

	insecureHttpConnectionManagerName := fmt.Sprintf("%s-insecure", name)
	insecureHttpConnectionManager, err := NewHTTPConnectionManager(
		insecureHttpConnectionManagerName,
		insecureHttpConnectionManagerName,
	)
	if err != nil {
		return ciliumv2.XDSResource{}, err
	}

	filterChains = append(filterChains, &envoy_config_listener.FilterChain{
		FilterChainMatch: &envoy_config_listener.FilterChainMatch{TransportProtocol: rawBufferTransportProtocol},
		Filters: []*envoy_config_listener.Filter{
			{
				Name: httpConnectionManagerType,
				ConfigType: &envoy_config_listener.Filter_TypedConfig{
					TypedConfig: insecureHttpConnectionManager.Any,
				},
			},
		},
	})

	listener := &envoy_config_listener.Listener{
		Name:         name,
		FilterChains: filterChains,
	}

	for _, fn := range mutatorFunc {
		listener = fn(listener)
	}

	listenerBytes, err := proto.Marshal(listener)
	if err != nil {
		return ciliumv2.XDSResource{}, err
	}
	return ciliumv2.XDSResource{
		Any: &anypb.Any{
			TypeUrl: envoy.ListenerTypeURL,
			Value:   listenerBytes,
		},
	}, nil
}
