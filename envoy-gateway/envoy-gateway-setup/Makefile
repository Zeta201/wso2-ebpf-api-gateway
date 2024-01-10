ENVOY_SERVICE=$(shell kubectl get svc -n envoy-gateway-system --selector=gateway.envoyproxy.io/owning-gateway-namespace=default,gateway.envoyproxy.io/owning-gateway-name=eg -o jsonpath='{.items[0].metadata.name}')
GAT_ADD := $(shell kubectl get gtw -o 'jsonpath={.items[0].status.addresses[0].value}')

install-envoy:
	- kubectl apply -f https://github.com/envoyproxy/gateway/releases/download/v0.3.0/install.yaml
	- kubectl wait --timeout=5m -n envoy-gateway-system deployment/envoy-gateway --for=condition=Available

configure-envoy:
	- kubectl apply -f apps.yaml
	- kubectl -n envoy-gateway-system port-forward service/$(ENVOY_SERVICE) 8888:80 &

run-tests:
	@echo "Testing details service..."
	curl -sS -I --verbose --header "Host: www.details.com" http://$(GAT_ADD)/details/1 | jq
	@echo "Testing product service..."
	curl -sS -I --header "Host: www.products.com" http://$(GAT_ADD)/
	@echo "Testing echo service..."
	curl -sS -I --header "Host: www.echo.com" http://$(GAT_ADD)/ -v

clean:
	- kubectl delete -f apps.yaml
	- kubectl delete -f https://github.com/envoyproxy/gateway/releases/download/v0.3.0/install.yaml



