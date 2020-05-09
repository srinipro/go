package config

// Config to hold the Properties of initializing the Application.
type Config struct {
	Port       int64  `json:"port"`
	Server     string `json:"server"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Drivername string `json:"drivername"`
}
