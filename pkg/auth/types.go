package auth

import (
	"github.com/jenkins-x/jx/pkg/secreturl"
	"github.com/jenkins-x/jx/pkg/vault"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type AuthServer struct {
	URL   string      `json:"url"`
	Users []*UserAuth `json:"users"`
	Name  string      `json:"name"`
	Kind  string      `json:"kind"`

	CurrentUser string `json:"currentuser"`
}

type UserAuth struct {
	Username    string `json:"username"`
	ApiToken    string `json:"apitoken"`
	BearerToken string `json:"bearertoken"`
	Password    string `json:"password,omitempty"`
}

type AuthConfig struct {
	Servers []*AuthServer `json:"servers"`

	DefaultUsername  string `json:"defaultusername"`
	CurrentServer    string `json:"currentserver"`
	PipeLineUsername string `json:"pipelineusername"`
	PipeLineServer   string `json:"pipelineserver"`
}

// AuthConfigService implements the generic features of the ConfigService because we don't have superclasses
type AuthConfigService struct {
	config *AuthConfig
	saver  ConfigSaver
}

// FileAuthConfigSaver is a ConfigSaver that saves its config to the local filesystem
type FileAuthConfigSaver struct {
	FileName string
}

// VaultAuthConfigSaver is a ConfigSaver that saves configs to Vault
type VaultAuthConfigSaver struct {
	vaultClient vault.Client
	secretName  string
}

// MemoryAuthConfigSaver uses memory
type MemoryAuthConfigSaver struct {
	config AuthConfig
}

// ConfigMapVaultConfigSaver save the config in a config map and the secrets in vault
type ConfigMapVaultConfigSaver struct {
	secretName      string
	configMapClient v1.ConfigMapInterface
	secretURLClient secreturl.Client
}
