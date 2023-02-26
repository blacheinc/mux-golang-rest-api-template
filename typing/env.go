package typing

type Env struct {
	// Port for the server to listen on
	Port string `env:"PORT"`
}
