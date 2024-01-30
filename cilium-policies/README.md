# Configuring Cilium Envoy

This guide explains how to configure node level cilium envoy for enforcing various policies on APIs.

## Enable JWT Authentication and Rate Limit Policies for an API

Run the below command to enforce a policy to authenticate the traffic reaching `echo-api` using JWT Authentication.

```bash
make enable-jwt-ratelimit
```
Run the below command to send traffic to echo-api without a Bearer token. The request should fail and you should see **401** returned as the response.
```bash
make test-echo-without-token
```
**Sample output**
```bash
Requesting auth token from funnel-labs...
curl -sS -k -I --location 'https://echo.gw.wso2.com:9095/echo-api/1.0.0' --header 'Host: echo.gw.wso2.com' 	
HTTP/2 401 
www-authenticate: Bearer realm="https://echo-service.default/"
content-length: 14
content-type: text/plain
date: Tue, 30 Jan 2024 04:56:36 GMT
server: envoy

* Closing connection -1
curl: (3) URL using bad/illegal format or missing URL
make: *** [Makefile:17: test-echo-without-token] Error 3
```

Run the below command to send traffic to echo-api with a Bearer token. You should see **200** returned as the response.
```bash
make test-echo-with-token
```
**Sample output**
```bash

HTTP/2 200 
x-powered-by: Express
vary: Origin, Accept-Encoding
access-control-allow-credentials: true
accept-ranges: bytes
cache-control: public, max-age=0
last-modified: Wed, 21 Sep 2022 10:25:56 GMT
etag: W/"809-1835f952f20"
content-type: text/html; charset=UTF-8
content-length: 2057
date: Tue, 30 Jan 2024 05:00:26 GMT
x-envoy-upstream-service-time: 1
server: envoy
```

Run the below command to test the rate limit policy enforced on the **echo-api**. You should see only the first request have succeeded and the rest of the requests have been rate-limited.
```bash
make test-ratelimit 
```
**Sample output**
```bash
HTTP/2 200 
x-powered-by: Express
vary: Origin, Accept-Encoding
access-control-allow-credentials: true
accept-ranges: bytes
cache-control: public, max-age=0
last-modified: Wed, 21 Sep 2022 10:25:56 GMT
etag: W/"809-1835f952f20"
content-type: text/html; charset=UTF-8
content-length: 2057
date: Tue, 30 Jan 2024 04:28:37 GMT
x-envoy-upstream-service-time: 1
x-ratelimit-limit: 1, 1;w=60
x-ratelimit-remaining: 0
x-ratelimit-reset: 23
server: envoy

HTTP/2 429 
x-envoy-ratelimited: true
x-ratelimit-limit: 1, 1;w=60
x-ratelimit-remaining: 0
x-ratelimit-reset: 23
date: Tue, 30 Jan 2024 04:28:37 GMT
server: envoy
vary: Accept-Encoding

HTTP/2 429 
x-envoy-ratelimited: true
x-ratelimit-limit: 1, 1;w=60
x-ratelimit-remaining: 0
x-ratelimit-reset: 23
date: Tue, 30 Jan 2024 04:28:37 GMT
server: envoy
vary: Accept-Encoding

```
