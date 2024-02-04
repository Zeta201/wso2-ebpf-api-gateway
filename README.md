.. raw:: html

   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://cilium.io/static/cilium-dark-1-8c7358e82e52a0ccab8d1d056de00c50.svg" width="200" alt="Cilium Logo">
      <img src="https://cilium.io/static/cilium-dark-1-8c7358e82e52a0ccab8d1d056de00c50.svg" width="200" alt="Cilium Logo">
   </picture>
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://github.com/envoyproxy/artwork/blob/main/PNG/Envoy_Logo_Final_CMYK.png" width="200" alt="Cilium Logo">
      <img src="https://github.com/envoyproxy/artwork/blob/main/PNG/Envoy_Logo_Final_CMYK.png" width="200" alt="Cilium Logo">
   </picture>
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://kubernetes.io/images/favicon.png" width="100" alt="K8 Logo">
      <img src="https://kubernetes.io/images/favicon.png" width="200" alt="K8 Logo">
   </picture>
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="100" height="80" alt="Golang Logo">
      <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="140" alt="Golang Logo">
   </picture>
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/530px-Tux.svg.png" width="100" height="100" alt="Linux Logo">
      <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/35/Tux.svg/530px-Tux.svg.png" width="140" alt="Linux Logo">
   </picture>
   
WSO2-Implementing an eBPF based API Gateway
===========================================
### Table of Contents
- **[Project Introduction](#project-introduction)**<br>
- **[Technologies Used](#technologies-used)**<br>
- **[System Architecture](#system-architecture)**<br>
- **[Developer Guide](#developer-guide)**<br>
   - [Environment Setup](#environment-setup)<br>
   - [Deploy WSO2APK](#deploy-wso2apk)<br>
   - [Configuring Cilium Envoy](#configuring-cilium-envoy)<br>
   - [Debugging](#debugging)<br>
- **[References](#references)**<br> 
   

## Project Introduction

Enterprise grade web applications are mostly architectured in the form of micro-services where
multiple services with well-defined and distinct functionalities communicate with each other to
perform a certain job instead of following a monolithic approach. Within a microservices
architecture, each microservice is a single service built to accommodate an application feature
and handle discrete tasks. Each microservice communicates with other services through simple
interfaces to solve business problems. In a competitive business environment developers should
be able to deliver changes rapidly, frequently and reliably in order for the business to thrive in
today’s volatile, uncertain, complex and ambiguous world.
Most micro-services architecture will at some point require exposing some services to the
outside of the cluster and securely routing traffic into it. An API Gateway accepts API requests
from a client, processes them based on defined policies, directs them to the appropriate services,
and combines the responses for a simplified user experience. Typically, it handles a request by
invoking multiple microservices and aggregating the results. It can also translate between
protocols in legacy deployments.
An API Gateway integrated to a production level microservices environment should be capable
of implementing, Routing Policies such as HTTP Routing, Rate limiting, Request/Response
Manipulation, Circuit Breakers, Canary Testing, A/B Scenarios, Health Checks, Load Balancing.
It should also be able to implement API Security Policies such as Traffic Authentication, Authorization, Access Control and TLS Termination. The Gateway should also be able to provide observability and visibility by real-time traffic, logging, and tracing.

## Technologies Used

### Extended Berkeley Packet Filter (eBPF)

eBPF is a revolutionary kernel technology that allows the developers to write custom code that can be loaded into the kernel dynamically, changing the way the kernel behaves. This enables a new generation of highly performant networking, observability, and security tools.
These are just a few of the many things eBPF enables us to do.

- Performance tracing of pretty much any aspect of a system
- High-performance networking, with built-in visibility
- Detecting and (optionally) preventing malicious activity
- Network security features that involve dropping certain incoming packets and allowing
  others. These features include firewalling, DDoS protection, and mitigating packet-ofdeath vulnerabilities.
- L3/L4 Load Balancing and Packet Parsing using XDP. (eBPF programs attached to the
  XDP hook on a network interface drop certain packets. The program is executed on the
  network card itself whenever a new packet arrives).
- Network Layer Security

eBPF is continuously evolving and being integrated into various enterprise-grade production
systems. Here are some examples.

- [Cilium project](https://github.com/cilium/cilium): The first networking project to use eBPF to replace the entire datapath in
  container environments.
- [Facebook](https://facebook.com) (now Meta) developed [Katran](https://github.com/facebookincubator/katran) which is an open-source, Layer 4 load balancer,
  for Facebook’s need for a highly scalable and fast solution. Every single packet to
  Facebook.com since 2017 has passed through eBPF.
- Android uses eBPF for traffic management and traffic control. Every network packet that
  received by our mobile phones has gone through an eBPF program.

### Cilium Project

Cilium is a CNI Plugin for Kubernetes which uses eBPF as a platform for Kubernetes
networking, TCP/IP layer load balancing, TCP/IP layer policy enforcement, deep observability
and highly performant sidecarless service meshes.
It is open-source, CNCF software for transparently securing the network connectivity between
application services deployed using Linux container management platforms like Docker and
Kubernetes.
At the foundation of Cilium is the Linux kernel technology, eBPF, as discussed previously
enabling the dynamic insertion of powerful security visibility and control logic within Linux
itself. Because eBPF runs inside the Linux kernel, Cilium security policies can be applied and
updated without any changes to the application code or container configuration

### Envoy Proxy

[Envoy](https://github.com/envoyproxy/envoy) is a high-performance C++ distributed L4/L7 proxy designed for single services and
applications, as well as a communication bus and “universal data plane” designed for large
microservice “service mesh” architectures. Built on the learnings of solutions such as NGINX,
HAProxy, hardware load balancers, and cloud load balancers, Envoy runs alongside every
application and abstracts the network by providing common features in a platform-agnostic
manner.

### Golang

The [Go](https://go.dev/doc/) is an open-source, expressive, concise, fast, statically typed, compiled programming
language developed by Google. Its concurrency mechanisms make it easy to write programs that
get the most out of multicore and networked machines, while its novel type system enables
flexible and modular program construction. Go compiles quickly to machine code yet has the
convenience of garbage collection and the power of run-time reflection.

### WSO2APK

[APK](https://github.com/wso2/apk) is WSO2's cloud native API management platform. APK is designed to help users to build,
deploy, and manage APIs in a cloud environment. This platform is built on top of a microservices
architecture and uses containerization technologies to ensure scalability and flexibility. With
features like automatic failover and load balancing, WSO2APK platform is designed to be highly
available and able to handle large numbers of API requests without performance degradation.
They have also added support for continuous delivery and deployment, so users can quickly and
easily push updates to their API services. WSO2APK is powered by envoy proxy for performing
complex Layer 7 operations such as routing, load balancing, rate limiting, and traffic
authentication to the APIs. It follows Kubernetes Gateway API Specification for the underlying
implementation

## System Architecture

   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53454670476_3143f79c52_b.jpg" width="100%" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53454670476_3143f79c52_b.jpg" width="100%" alt="Cilium Logo">
   </picture>

## System Functionality

In my proposed system, I have used `Cilium` as the CNI Layer for Kubernetes Networking. This
system enables `observability`, `visibility`, `policy enforcement` in `L3/L4 layers` in TCP/IP Stack
using eBPF. This system is also able to perform load balancing in L3/L4 layers. So, all the L3/L4
layer operations will be handled by using kernel code (eBPF).
When it comes to Layer 7 characteristics such as L7 network policies, load balancing, and
observability this system will make use of the Node Level running Envoy Proxy to perform these
L7 processing. When we apply any L7 policy for any K8 Service, the system will start an envoy
proxy if it’s not already running, and it will configure a listener which listens on a port. By
configuring this envoy proxy, the system is able to perform all the API Gateway functionalities
such as HTTP Routing, Load Balancing, TLS Termination, Rate Limiting, Circuit Breaking, etc.
for the traffic. The envoy proxy is configured by using CRDs that Cilium provides.
When it comes to L7-Aware Traffic Management, there are two kinds of network traffic in a
Kubernetes environment. They are Ingress Traffic (North-South) which comes from outside the
cluster and Service Mesh Traffic (East-West) which is the traffic between two internal K8
Services. This system is capable of manipulating both the N-S and E-W Traffic by enforcing L7-
Aware Policies.
The users have the ability to run their own custom build of Cilium Proxy to extend the Envoy
Proxy leveraging more extensions such as Lua, Compressor, gzip etc.
This system also supports Envoy Customization to write and run any Business Logic the user
wants in an enterprise environment. This is achieved through Web Assembly Extension provided
by Envoy Proxy. The users can write any Custom Business Logic they prefer and mount the
compiled WASM Binary into Cilium Agent

In summary the above system supports following features.

- HTTP Routing
- TLS Termination
- HTTP Request/Response Manipulation
- Observability – Cilium Hubble Tool
- Traffic Authentication - JWT/OAuth2
- API Aware Rate Limiting
- L7 Load Balancing
- WASM Support for custom business logic
- Circuit Breaker
- TCP/IP Layer Policy Enforcement and Load Balancing

## eBPF Action in Layer 7

The users can easily and efficiently enforce L3/L4 layer policies for their services. This is
highly performant as L3/L4 layer operations are processed by eBPF code and hence there is no
any involvement of complex Application Layer processing. This system leverages Cilium’s
highly performant Sidecarless Service Mesh where users can achieve Unix Domain Socket
Speed between services.

When it comes to L7 traffic processing eBPF enables Accelerated Redirection for the Envoy
Proxy. If the user is running Istio service mesh, all the connections go through the network stack
down to the Ethernet Level and Loopback and back up to the Envoy Proxy and back out to the
network interface. This is using TCP which has been written for lossy environments. Below diagrams
highlights this fact.
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53454815143_08bd23886d.jpg" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53454815143_08bd23886d.jpg" width="30%" alt="Cilium Logo">
   </picture>
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53453765597_afed2c37d4.jpg" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53453765597_afed2c37d4.jpg" width="30%" alt="Cilium Logo">
   </picture>
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53454815153_60932f74db.jpg" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53454815153_60932f74db.jpg" width="30%" alt="Cilium Logo">
   </picture>

Cilium uses socket level BPF and it will detect that your application is talking to the sidecar locally so it will simply copy the data from one socket to the other transparently acclerating envoy.

As depicted in the above performance graph, Y-axis is the number of requests per second the the X- axis is the number of persistent connections. Typically if you're running service mesh or you're running with with a sidecar the app does not open a new connection for every request. Most networking libraries will maintain connection pools and will reuse the same TCP connection multiple times for new requests. With Cilium. The above graph indicates 3 cases. Blue bar indicates the IP tables redirect. This is what you get when you deploy Istio Service Mesh. Orange is the case where you point the the app to the sidecar so you don't need the IP tables redirect rule so the difference between blue and orange is already that the cost of a single IP tables redirect. Then Yellow is Cilium accelerating this transparently and it's somewhere between 3-4X faster.

Users can also use Hubble under Cilium CNI Layer to gain a deep observability and visibility for
logging, tracing, and monitoring connections under this setup.

## Developer Guide
This guide will help you to understand on deploying the eBPF based API Gateway.

### Environment Setup

#### Clone the Repository

First clone the repository as given below.
```bash
git clone https://github.com/Zeta201/wso2-ebpf-api-gateway.git
```
Assuming your are under the **root directory**, change your working directory to cilium-setup inside the cloned repository directory.
```bash
cd cilium-setup
```

#### Setup a Kind Cluster

First we need to create a new kind cluster to install Cilium. The below command will create a new kind cluster with one worker node disabling the default CNI and kube-proxy.

```bash
make cluster-setup
```

#### Install Cilium

The below command will install Cilium as the CNI Plugin for Kubernetes and it will also install Hubble for metrics and logging.
**IMPORTANT**
Under this guide we will be deploying **Cilium Node Level Envoy Proxy** as an embedded process inside the Cilium Agent. However it is possible to deploy the Cilium Envoy Proxy as a standalone daemonset.
If you want to deploy it as a standalone daemonset set add the below configuration to the [cilium-config.yaml](https://github.com/Zeta201/wso2-ebpf-api-gateway/blob/main/cilium-setup%20/cilium-config.yaml) file under this directory.
```bash
envoy:
  enabled: true
```

```bash
make install-cilium
```
#### Install MetalLB

To access the service that will be exposed via the Gateway API, we need to allocate an external IP address. When a Gateway is created, an associated Kubernetes Services of the type LoadBalancer is created. When using a managed Kubernetes Service like EKS, AKS or GKE, the LoadBalancer is assigned an IP (or DNS name) automatically. For private cloud or for home labs, we need another tool – such as MetalLB below – to allocate an IP Address and to provide L2 connectivity.

```bash
make install-metallb
```
#### Configure MetalLB

Run the below command to apply L2 announcement policy for MetalLB.

```bash
make configure-metallb
```
#### Install Demo Applications

Install the applcations for testing the API gateway.

```bash
make install-demo-apps
```

#### Clean

Run below command to uninstall cilium, metallb and demo apps.

```bash
make clean
```

### Deploy WSO2APK

Change your working directory to apk-setup directory under the root directory.
Assuming your are under the **root directory** execute the following command.
```bash
cd apk-setup
```
#### Install required packages

This will install jq package for Ubuntu environment which will help us by formatting the response objects and making it easier for selecting specific attributes of resources.

```bash
make prepare
```
#### Install APK

Install APK from the customized helm chart.

```bash
make apk-install
```
#### Deploy the Demo Apps
Install and deploy the API, HTTPRoute, and Backend CRs for the demo application set.

```bash
make deploy-demo-apps
```
#### Test the setup

Test echo service

```bash
make test-echo-service
```
Test details service

```bash
make test-details-service
```
#### Clean
Uninstall APK and undeploy all APIs.

```bash
make clean
```
### Configuring Cilium Envoy

This guide explains how to configure node level cilium envoy for enforcing various policies on APIs.

Change your working directory to cilium-policies directory under the root directory.
Assuming your are under the **root directory** execute the following command.
```bash
cd cilium-policies
```

#### Enable JWT Authentication and Rate Limit Policies for an API

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
#### Configuring Cilium Envoy with Web Assembly

The below guide explains how to customize the Cilium Envoy possibly for running custom domain business logic using Web Assembly.

Run the below command to mount the web assembly compiled binary into Cilium Agents.
```bash
make copy-wasm
```
Run the following command to enable wasm filter using a **CEC** 
```bash
make enable-wasm
```
Run the following command to send a request to the details-api to test the Wasm filter. You should see some custom HTTP headers have been added to the response object as given below
```bash
make test-wasm
```
**Sample Output**

```bash 

```
### Debugging 

**CiliumEnvoyConfig** has minimal feedback to the user about the correctness of the configuration. So in the event a CEC does produce an undesirable outcome, troubleshooting will require inspecting the Envoy config and logs, rather than being able to look at the **CiliumEnvoyConfig** in question.  Check Cilium Agent logs if you are creating Envoy resources explicitly to make sure there is no conflict.

```bash
kubectl logs -n kube-system ds/cilium | grep -E "level=(error|warning)"
```    
## References
