package database

var connection string

func init() {
	connection = "PostgreSQL"
}

func GetConnection() string {
	return connection
}
