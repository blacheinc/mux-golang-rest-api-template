package typing

type Env struct {
	// Port for the server to listen on
	Port string `env:"PORT"`
	//AppDBUri is the monogodb connection string for App database
	MongoDBUri string `env:"MONGO_DB_URI"`
	//AppDBName is the database name for App
	MongoDBName string `env:"MONGO_DB_NAME"`
	//PostgreSQLUri is the PostgreSQL database uri string for App
	PostgreSQLUri string `env:"POSTGRE_SQL_URI"`
}
