package Config

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port int
}

type DatabaseConfiguration struct {
	Cluster      string
	Port         int
	DatabaseName string
	UserName     string
	Password     string
}
