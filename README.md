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

Project Introduction
--------------------
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

Technologies Used
-----------------

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

System Architecture
-------------------
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53454670476_3143f79c52_b.jpg" width="100%" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53454670476_3143f79c52_b.jpg" width="100%" alt="Cilium Logo">
   </picture>

System Functionality
--------------------
In my proposed system, I have used ``Cilium`` as the CNI Layer for Kubernetes Networking. This
system enables ``observability``, ``visibility``, ``policy enforcement`` in ``L3/L4 layers`` in TCP/IP Stack
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

eBPF Action in Layer 7
----------------------
The users can easily and efficiently enforce L3/L4 layer policies for their services. This is
highly performant as L3/L4 layer operations are processed by eBPF code and hence there is no
any involvement of complex Application Layer processing. This system leverages Cilium’s
highly performant Sidecarless Service Mesh where users can achieve Unix Domain Socket
Speed between services.

When it comes to L7 traffic processing eBPF enables Accelerated Redirection for the Envoy
Proxy. If the user is running Istio service mesh, all the connections go through the network stack
down to the Ethernet Level and Loopback and back up to the Envoy Proxy and back out to the
network interface. This is using TCP which was written for lossy environments. Below diagram
highlights this fact.
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53454815143_08bd23886d.jpg" width="300" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53454815143_08bd23886d.jpg" width="100%" alt="Cilium Logo">
   </picture>
But in this system, I leverage Cilium capabilities which will detect that the service is
communicating with its local sidecar, and it will simply copy the data from one socket to the
other achieving Unix Domain Socket Speed on TCP sockets as given below.
   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53453765597_afed2c37d4.jpg" width="300" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53453765597_afed2c37d4.jpg" width="100%" alt="Cilium Logo">
   </picture>
Users can achieve a performance improvement of 3-4 times faster and persistent HTTP
connections as given below.

   <picture>
      <source media="(prefers-color-scheme: light)" srcset="https://live.staticflickr.com/65535/53454815153_60932f74db.jpg" width=100%" alt="Cilium Logo">
      <img src="https://live.staticflickr.com/65535/53454815153_60932f74db.jpg" width="100%" alt="Cilium Logo">
   </picture>





Functionality Overview
======================






