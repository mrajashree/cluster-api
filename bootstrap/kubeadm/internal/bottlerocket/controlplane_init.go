// TODO: make bottlerocket(init) more agnostic. In addition to other changes to make things
// less hacky, also move calling cloudinit from controller and passing it to
// bottlerocket bootstrap, to all control to bottlerocket bootstrap itself.
// That way, bottlerocket bootstrap will call cloudinit to generate that userdata
// which is much more cleaner.
package bottlerocket

import (
	"encoding/base64"
	"fmt"
	"strings"

	capbk "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
)

type ControlPlaneInput struct {
	BootstrapSettings     string
	KubernetesSettings    string
	HostContainerSettings string
}

func NewInitControlPlane(cloudinitInput string, sshAuthUsers []capbk.User) ([]byte, error) {
	// Parse ssh auth keys
	keys := ""
	for _, user := range sshAuthUsers {
		for _, key := range user.SSHAuthorizedKeys {
			keys += "\"" + key + "\","
		}
	}
	authInitData := fmt.Sprintf("{\"ssh\":{\"authorized-keys\":[%s]}}\n", strings.TrimRight(keys, ","))
	b64AuthInitString := base64.StdEncoding.EncodeToString([]byte(authInitData))

	cpInput := new(ControlPlaneInput)
	cpInput.BootstrapSettings = fmt.Sprintf(`[settings.host-containers.kubeadm-the-hard-way]
enabled = true
superpowered = true
source = "public.ecr.aws/k1e6s8o8/kubeadm-the-hard-way:0.0.1"
user-data = "%s"
`, cloudinitInput)

	cpInput.KubernetesSettings = `[settings.kubernetes]
cluster-domain = "cluster.local"
standalone-mode = true
authentication-mode = "tls"
server-tls-bootstrap = false`

	// TODO: replace user data??
	cpInput.HostContainerSettings = fmt.Sprintf(`[settings.host-containers.admin]
enabled = true
user-data = "%s"`, b64AuthInitString)

	userData := fmt.Sprintf("%s%s\n%s", cpInput.BootstrapSettings, cpInput.KubernetesSettings, cpInput.HostContainerSettings)
	return []byte(userData), nil
}
