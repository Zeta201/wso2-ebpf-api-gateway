# Setup Cilium 

## Cluster Setup

Create a new kind cluster without the default CNI.

```bash
make cluster-setup
```
## Install Cilium without KubeProxy

The below command will install Cilium as the CNI Plugin without KubeProxy and it will also install Hubble for observability.

```bash
make install-cilium
```
