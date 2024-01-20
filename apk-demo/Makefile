ACCESS_TOKEN=$(shell curl -sS -k --location 'https://idp.am.wso2.com:9095/oauth2/token' --header 'Host: idp.am.wso2.com' --header 'Authorization: Basic NDVmMWM1YzgtYTkyZS0xMWVkLWFmYTEtMDI0MmFjMTIwMDAyOjRmYmQ2MmVjLWE5MmUtMTFlZC1hZmExLTAyNDJhYzEyMDAwMg==' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'grant_type=client_credentials' | jq -r '.access_token')

current_dir = $(shell pwd)

apk-install:
	helm install apk ../apk-charts/apk-helm/ 

deploy-demo-apps:
	@echo "Deploying demo apps..."
	@kubectl apply -f bookinfo.yaml
	kubectl wait deployment -l wso2=demo --timeout=5m --for=condition=Available
	kubectl apply -f backend.yaml
	kubectl apply -f httproute.yaml
	kubectl apply -f api.yaml
	@echo "Deployment success!"
	@echo "Run \e[1mmake test-echo-service\e[0m to test the api"

clean-demo-apps:
	kubectl delete -f bookinfo.yaml

test-wasm:
	curl -sS -I -k --location 'https://http-bin.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com'  --header "Authorization: bearer $(ACCESS_TOKEN)" 

test-echo-service:
	@curl -sS -I -k --location 'https://http-bin.gw.wso2.com:9095/http-bin-api/1.0.0/' --header 'Host: http-bin.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" 

test-details-service:
	@curl -sS -k --location 'https://http-bin.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com'  --header "Authorization: bearer $(ACCESS_TOKEN)" | jq
		
test-ratelimit:
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	- curl -I -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v

