package http

type App struct{
	Config *EnvConfig
}

func NewApp(cfg *EnvConfig) *App {
	return &App{Config: cfg}
}
