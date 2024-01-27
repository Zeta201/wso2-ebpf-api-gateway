TOKEN=$(shell curl -s -X POST https://auth.funnel-labs.io/auth/realms/funnel/protocol/openid-connect/token -d "client_id=service&username=demo&password=abcd1234&grant_type=password" | jq -r '.access_token')

enable-jwt-ratelimit:
	kubectl apply -f ratelimit-jwt/echo.yaml

test-echo-service:
	@echo "Requesting auth token from funnel-labs..."
	curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0' --header 'Host: echo.gw.wso2.com' --header "Authorization: Bearer $(TOKEN)"	

test-details-service:
	@curl -sS -k --location 'https://details.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com' | jq
