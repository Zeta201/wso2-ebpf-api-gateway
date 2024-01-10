FORTIO_POD := $(shell kubectl get pods -l app=fortio -o 'jsonpath={.items[0].metadata.name}')

GAT_ADD := $(shell kubectl get gtw -o 'jsonpath={.items[0].status.addresses[0].value}')

CON := 2

install-circuit-breaker:
	kubectl apply -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/kubernetes/servicemesh/envoy/envoy-circuit-breaker.yaml
	
ew-test:
	@kubectl exec $(FORTIO_POD) -c fortio -- /usr/bin/fortio load -c $(CON) -qps 0 -n 20 http://echo-service:8080

ns-test:
	fortio load -H 'Host:www.echo.com' -c $(CON) -qps 0 -n 20 http://$(GAT_ADD)/

clean:
	kubectl delete -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/kubernetes/servicemesh/envoy/envoy-circuit-breaker.yaml

