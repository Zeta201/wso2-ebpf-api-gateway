ACCESS_TOKEN=$(shell curl -sS -k --location 'https://idp.am.wso2.com:9095/oauth2/token' --header 'Host: idp.am.wso2.com' --header 'Authorization: Basic NDVmMWM1YzgtYTkyZS0xMWVkLWFmYTEtMDI0MmFjMTIwMDAyOjRmYmQ2MmVjLWE5MmUtMTFlZC1hZmExLTAyNDJhYzEyMDAwMg==' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'grant_type=client_credentials' | jq -r '.access_token')

current_dir = $(shell pwd)

prepare:
	sudo apt-get install jq
	
apk-install:
	helm install apk ../apk-charts/apk-helm/ 

deploy-demo-apps:
	@echo "Deploying demo apps..."
	@kubectl apply -f bookinfo.yaml
	kubectl wait deployment -l wso2=demo --timeout=10m --for=condition=Available
	kubectl apply -f backend.yaml
	kubectl apply -f httproute.yaml
	kubectl apply -f api.yaml
	kubectl apply -f oauth2.yaml
	@echo "Deployment success!"
	@echo "Run \e[1mmake test-echo-service\e[0m to test the echo api"
	@echo "Run \e[1mmake test-details-service\e[0m to test the details api"
clean:
	kubectl delete -f bookinfo.yaml
	helm uninstall apk

test-wasm:
	curl -sS -I -k --location 'https://http-bin.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com'  --header "Authorization: bearer $(ACCESS_TOKEN)" 

test-echo-service:
	@curl -sS -I -k --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0/' --header 'Host: echo.gw.wso2.com' 

test-details-service:
	@curl -sS -k --location 'https://details.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com' | jq
