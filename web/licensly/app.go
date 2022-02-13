package licensly

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

type App struct{}

func (app *App) ErrorHandler(rw http.ResponseWriter, req *http.Request, err error) {
	errors.HandleMiddlewareError(rw, req, err, nil)
}

func NewApp() *App {
	return &App{}
}
