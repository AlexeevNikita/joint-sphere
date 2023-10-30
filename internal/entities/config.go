package entities

type Config struct {
	Database database
}

type database struct {
	User     string
	Password string
}
