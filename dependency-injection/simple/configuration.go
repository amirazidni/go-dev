package simple

// struct field provider
type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

func NewApplication() *Application {
	return &Application{
		Configuration: &Configuration{ // buat ini jadi provider
			Name: "struct field provider",
		},
	}
}
