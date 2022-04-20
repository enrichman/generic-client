package cmd

type Config struct {
	Name string

	NamespaceConfig NamespaceConfig
}

type NamespaceConfig struct {
	NamespaceListConfig NamespaceListConfig
}

type NamespaceListConfig struct {
	Name string
}

func NewConfig() Config {
	return Config{
		Name: "default_root_name",
		NamespaceConfig: NamespaceConfig{
			NamespaceListConfig: NamespaceListConfig{
				Name: "default_namespace_list_name",
			},
		},
	}
}
