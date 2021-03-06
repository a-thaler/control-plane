package gardener

type Config struct {
	Project        string `envconfig:"default=gardenerProject"`
	ShootDomain    string `envconfig:"optional"`
	KubeconfigPath string `envconfig:"default=./dev/kubeconfig.yaml"`
}
