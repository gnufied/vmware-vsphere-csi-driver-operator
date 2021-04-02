module github.com/openshift/vmware-vsphere-csi-driver-operator

go 1.16

require (
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/openshift/api v0.0.0-20210315130343-ff0ef568c9ca
	github.com/openshift/build-machinery-go v0.0.0-20200917070002-f171684f77ab
	github.com/openshift/client-go v0.0.0-20201214125552-e615e336eb49
	github.com/openshift/library-go v0.0.0-20210113192829-cfbb3f4c80c2
	github.com/openshift/vsphere-problem-detector v0.0.0-20210316174602-c32cbf1855ba
	github.com/prometheus/client_golang v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/vmware/govmomi v0.23.1
	gopkg.in/gcfg.v1 v1.2.3
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	k8s.io/component-base v0.20.4
	k8s.io/klog/v2 v2.4.0
	k8s.io/legacy-cloud-providers v0.20.4
)
