# Setup WSO2 APK


## Install required packages

This will install jq package for ubuntu environment.

```bash
make prepare
```
## Install APK

Install APK from the customized helm chart.

```bash
make apk-install
```
## Deploy the Demo Apps
Install and deploy the API, HTTPRoute, and Backend CRs for the demo application set.

```bash
make deploy-demo-apps
```
## Test the setup

Test echo service

```bash
make test-echo-service
```
Test details service

```bash
make test-details-service
```
## Clean
Uninstall APK and undeploy all APIs.

```bash
make clean
```
