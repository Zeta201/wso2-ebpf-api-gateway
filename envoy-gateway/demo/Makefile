ACCESS_TOKEN=$(shell curl -s -X POST https://auth.funnel-labs.io/auth/realms/funnel/protocol/openid-connect/token -d "client_id=service&username=demo&password=abcd1234&grant_type=password" | jq -r '.access_token')

RATINGS_POD := $(shell kubectl get pod -l app=ratings -o 'jsonpath={.items[0].metadata.name}')
AGENT1 := $(shell kubectl get pods -n kube-system -l app.kubernetes.io/name=cilium-agent -o 'jsonpath={.items[0].metadata.name}')
AGENT2 := $(shell kubectl get pods -n kube-system -l app.kubernetes.io/name=cilium-agent -o 'jsonpath={.items[1].metadata.name}')
AGENT3 := $(shell kubectl get pods -n kube-system -l app.kubernetes.io/name=cilium-agent -o 'jsonpath={.items[2].metadata.name}')

GAT_ADD := $(shell kubectl get gtw -o 'jsonpath={.items[0].status.addresses[0].value}')

prep:
	- kubectl apply -f ratelimit-config.yaml
	- kubectl apply -f redis.yaml
	- kubectl apply -f ratelimit-deployment.yaml
	- kubectl apply -f envoy-echo.yaml

internal-test-without-token:
	kubectl exec $(RATINGS_POD) -- curl -sS -I http://echo-service:8080/ -v
    
internal-test-with-token:
	kubectl exec $(RATINGS_POD) -- curl -sS -I -H "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/ -v

internal-test-wrong-token:
	kubectl exec $(RATINGS_POD) -- curl -sS -I -H "Authorization: Bearer $(ACCESS_TOKEN)ab" http://echo-service:8080/ -v
	
internal-test-invalid-token:
	kubectl exec $(RATINGS_POD) -- curl -sS -H "Authorization: Bearer abcd" http://echo-service:8080/ -v
		
external-test-with-token:
	curl -sS -o /dev/null -H "Host: www.echo.com" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	
external-test-without-token:
	curl -sS -o /dev/null -H "Host: www.echo.com"  http://$(GAT_ADD)/ -v
test-wasm:
	kubectl exec $(RATINGS_POD) -- curl -sS -I http://details:9080/details/1 
internal-test-ratelimit:
	- @kubectl exec $(RATINGS_POD) -- curl -sS -I --header "key1:value1" --header "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/ 
	- @kubectl exec $(RATINGS_POD) -- curl -sS -I --header "key1:value1" --header "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/
	- @kubectl exec $(RATINGS_POD) -- curl -sS -I --header "key1:value1" --header "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/
	- @kubectl exec $(RATINGS_POD) -- curl -sS -I --header "key1:value1" --header "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/
	- @kubectl exec $(RATINGS_POD) -- curl -sS -I --header "key1:value1" --header "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/
	- @kubectl exec $(RATINGS_POD) -- curl -sS -I --header "key1:value1" --header "Authorization: Bearer $(ACCESS_TOKEN)" http://echo-service:8080/

external-test-ratelimit:
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	
restart:
	- kubectl rollout restart ds/cilium -n kube-system
	- kubectl rollout status ds/cilium -n kube-system --timeout 5m

copy-bin:
	- kubectl cp ./cilium-envoy kube-system/$(AGENT1):/usr/bin
	- kubectl cp ./cilium-envoy kube-system/$(AGENT2):/usr/bin
	- kubectl cp ./cilium-envoy kube-system/$(AGENT3):/usr/bin

copy-wasm:
	- kubectl cp ./optimized.wasm kube-system/$(AGENT1):/usr/bin
	- kubectl cp ./optimized.wasm kube-system/$(AGENT2):/usr/bin
	- kubectl cp ./optimized.wasm kube-system/$(AGENT3):/usr/bin
	
clean-cilium:
	- helm uninstall cilium -n kube-system
login-agent:
	- kubectl exec -it $(AGENT1) -n kube-system -- bash	
clean:
	- kubectl delete -f ratelimit-config.yaml
	- kubectl delete -f redis.yaml
	- kubectl delete -f ratelimit-deployment.yaml
	- kubectl delete -f envoy-echo.yaml
