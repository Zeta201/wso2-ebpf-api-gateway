ACCESS_TOKEN=$(shell curl -sS -k --location 'https://idp.am.wso2.com:9095/oauth2/token' --header 'Host: idp.am.wso2.com' --header 'Authorization: Basic NDVmMWM1YzgtYTkyZS0xMWVkLWFmYTEtMDI0MmFjMTIwMDAyOjRmYmQ2MmVjLWE5MmUtMTFlZC1hZmExLTAyNDJhYzEyMDAwMg==' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'grant_type=client_credentials' | jq -r '.access_token')

current_dir = $(shell pwd)

deploy-apps:
	kubectl apply -f backend.yaml
	kubectl apply -f httproute.yaml
	kubectl apply -f api.yaml

port-forward:
	kubectl port-forward svc/apk-wso2-apk-gateway-service 9095:9095

gen-apk-conf:
	curl -k --location 'https://api.am.wso2.com:9095/api/configurator/1.0.0/apis/generate-configuration' --header 'Host: api.am.wso2.com' --form 'definition=@"/home/dinitha/wso2/envoy-gateway/demo/apk-conf/EmployeeServiceDefinition.json"' > EmployeeService.apk-conf


deploy-api:
	curl -k --location 'https://api.am.wso2.com:9095/api/deployer/1.0.0/apis/deploy' --header 'Host: api.am.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" --form 'apkConfiguration=@"/home/dinitha/wso2/envoy-gateway/demo/apk-conf/EmployeeService.apk-conf"' --form 'definitionFile=@"/home/dinitha/wso2/envoy-gateway/demo/apk-conf/EmployeeServiceDefinition.json"'

test-wasm:
	curl -sS -I -k --location 'https://http-bin.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com'  --header "Authorization: bearer $(ACCESS_TOKEN)" 

test-echo:
	curl -sS -k --location 'https://http-bin.gw.wso2.com:9095/http-bin-api/1.0.0/' --header 'Host: http-bin.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v

test-details:
	curl -sS -k --location 'https://http-bin.gw.wso2.com:9095/details-api/1.0.0/details/1' --header 'Host: details.gw.wso2.com'  --header "Authorization: bearer $(ACCESS_TOKEN)" -v | jq
	
test-api:
	curl -sS -k --location 'https://default.gw.wso2.com:9095/RW1wbG95ZWVTZXJ2aWNlQVBJMy4xNA/3.14/employee' --header 'Host: default.gw.wso2.com' --header "Authorization: bearer $(ACCESS_TOKEN)" -v
	
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

