package bottlerocket

import "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"

// Worker node configuration for bottlerocket is as same as for controlplane
// Only the cloudinit userdata is different, which cloudinit package handles
func NewNode(cloudinitInput string, sshAuthKeys []v1alpha3.User) ([]byte, error) {
	return NewInitControlPlane(cloudinitInput, sshAuthKeys)
}
