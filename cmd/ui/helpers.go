package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


func (app *application) render(w http.ResponseWriter, r *http.Request, name string) {

	ts, ok := app.template.pages[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", name))
		return
	}

	// Execute the template set, passing in any dynamic data.
	err := ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}
