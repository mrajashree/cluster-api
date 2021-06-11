package bottlerocket

import capbk "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"

// Worker node configuration for bottlerocket is as same as for controlplane
// Only the cloudinit userdata is different, which cloudinit package handles
func NewNode(cloudinitInput string, sshAuthKeys []capbk.User) ([]byte, error) {
	return NewInitControlPlane(cloudinitInput, sshAuthKeys)
}
