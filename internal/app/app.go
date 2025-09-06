package app

type App struct {
	Name string
}

func NewApp(name string) *App {
	return &App{Name: name}
}

func (a *App) Run() string {
	return "App " + a.Name + " is running!"
}
