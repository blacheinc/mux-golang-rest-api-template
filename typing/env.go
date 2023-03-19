package typing

type Env struct {
	// Port for the server to listen on
	Port string `env:"PORT"`
	//PostgreSQLURI is the PostgreSQL database uri string for App
	PostgreSQLURI string `env:"POSTGRE_SQL_URI"`
}
