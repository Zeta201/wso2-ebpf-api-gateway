cluster-setup:
	- @sudo sysctl fs.inotify.max_user_watches=524288
	- @sudo sysctl fs.inotify.max_user_instances=512
	- kind create cluster --config cluster-config.yaml
install-cilium:
	- helm upgrade --install --namespace kube-system --repo https://helm.cilium.io cilium cilium --values cilium-config.yaml

install-metallb:
	- helm repo add metallb https://metallb.github.io/metallb
	- kubectl create ns metallb-system
	- helm install metallb metallb/metallb --version 0.13.11 --namespace=metallb-system

configure-metallb:
	- kubectl apply -f service-ip.yaml
	- kubectl apply -f metallb_values.yaml

install-demo-apps:
	- kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.11/samples/bookinfo/platform/kube/bookinfo.yaml
	- kubectl apply -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/kubernetes/servicemesh/envoy/test-application-proxy-circuit-breaker.yaml

clean:
	- helm uninstall cilium -n kube-system
	- helm uninstall metallb -n metallb-system
	- kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.11/samples/bookinfo/platform/kube/bookinfo.yaml
	- kubectl delete -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/kubernetes/servicemesh/envoy/test-application-proxy-circuit-breaker.yaml
	- kind delete cluster











	




