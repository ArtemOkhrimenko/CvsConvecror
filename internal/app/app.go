package app

type Repo interface {
}

type App struct {
	repo Repo
}

func New(repo Repo) *App {
	return &App{
		repo: repo,
	}
}
