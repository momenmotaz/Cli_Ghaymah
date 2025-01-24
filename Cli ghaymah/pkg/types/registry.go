package types

// DockerRegistry represents Docker registry credentials
type DockerRegistry struct {
    RegistryURL string `yaml:"registryUrl" json:"registryUrl"`
    Username    string `yaml:"username" json:"username"`
    Password    string `yaml:"password" json:"password,omitempty"`
}

// Validate checks if the registry configuration is valid
func (r *DockerRegistry) Validate() bool {
    return r.RegistryURL != "" && r.Username != "" && r.Password != ""
}
