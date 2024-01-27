TOKEN=$(shell curl -s -X POST https://auth.funnel-labs.io/auth/realms/funnel/protocol/openid-connect/token -d "client_id=service&username=demo&password=abcd1234&grant_type=password" | jq -r '.access_token')
AGENT1 := $(shell kubectl get pods -n kube-system -l app.kubernetes.io/name=cilium-agent -o 'jsonpath={.items[0].metadata.name}')
AGENT2 := $(shell kubectl get pods -n kube-system -l app.kubernetes.io/name=cilium-agent -o 'jsonpath={.items[1].metadata.name}')
enable-jwt-ratelimit:
	kubectl apply -f ratelimit-jwt/echo.yaml
	kubectl apply -f ratelimit-jwt/ratelimit-config.yaml
	kubectl apply -f ratelimit-jwt/redis.yaml
	kubectl apply -f ratelimit-jwt/ratelimit-deployment.yaml

test-echo-without-token:
	@echo "Requesting auth token from funnel-labs..."
	curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0' --header 'Host: echo.gw.wso2.com' 	
# @curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v
	@curl -sS  -H "key1:value1" -H "Authorization: Bearer $(ACCESS_TOKEN)" http://$(GAT_ADD)/ -v

test-echo-with-token:
	@echo "Requesting auth token from funnel-labs..."
	curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0' --header 'Host: echo.gw.wso2.com' --header "Authorization: Bearer $(TOKEN)"	
test-ratelimit:
	curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0/' --header 'Host: echo.gw.wso2.com' -H "key1:value1" --header "Authorization: Bearer $(TOKEN)"
	curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0/' --header 'Host: echo.gw.wso2.com' -H "key1:value1" --header "Authorization: Bearer $(TOKEN)"
	curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0/' --header 'Host: echo.gw.wso2.com' -H "key1:value1" --header "Authorization: Bearer $(TOKEN)"

test-details-service:
	@curl -sS -k --location 'https://details.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com' | jq

test-wasm:
	@curl -sS -k --location 'https://details.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com' -I

copy-wasm:
	- kubectl cp ./wasm/optimized.wasm kube-system/$(AGENT1):/usr/bin
	- kubectl cp ./wasm/optimized.wasm kube-system/$(AGENT2):/usr/bin

clean:
	kubectl delete -f ratelimit-jwt/echo.yaml
	kubectl delete -f ratelimit-jwt/ratelimit-config.yaml
	kubectl delete -f ratelimit-jwt/redis.yaml
	kubectl delete -f ratelimit-jwt/ratelimit-deployment.yaml