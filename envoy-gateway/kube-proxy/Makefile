view-config:
	kubectl -n kube-system exec ds/cilium -- cilium status | grep KubeProxyReplacement

view-config-verbose:
	kubectl -n kube-system exec ds/cilium -- cilium status --verbose

view-ip-tables:
	sudo iptables-legacy-save | grep KUBE-SVC
