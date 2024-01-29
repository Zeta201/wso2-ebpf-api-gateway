# Setup a Kind Cluster and Install Cilium 


## Cluster Setup

Create a new kind cluster without the default CNI.

```bash
make cluster-setup
```
## Install Cilium

The below command will install Cilium as the CNI Plugin without KubeProxy and it will also install Hubble for observability.

```bash
make install-cilium
```
## Install MetalLB

Install MetalLb to assign an external IP to the load balancer service.

```bash
make install-metallb
```
## Configure MetalLB

Allocate IP Pool and L2 Advertisement 

```bash
make configure-metallb
```
## Install Demo Applications

Install the applcations for testing the API gateway.

```bash
make install-demo-apps
```
## Clean

Run below command to uninstall cilium, metallb and demo apps

```bash
make clean
```
