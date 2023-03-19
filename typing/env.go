package typing

type Env struct {
	// Port for the server to listen on
	Port string `env:"PORT"`
	//AppDBUri is the monogodb connection string for App database
	MongoDBURI string `env:"MONGO_DB_URI"`
	//AppDBName is the database name for App
	MongoDBName string `env:"MONGO_DB_NAME"`
}
