package cmd

type Config struct {
	Name       string
	UserConfig *UserConfig
}

type UserConfig struct {
	UserListConfig *UserListConfig
}

type UserListConfig struct {
	Name string
}

func newConfig() *Config {
	return &Config{
		Name: "default_root_name",
		UserConfig: &UserConfig{
			UserListConfig: &UserListConfig{
				Name: "default_namespace_list_name",
			},
		},
	}
}
