package handlers

import (
	"errors"
	"fmt"
	"github.com/utkarsh-1905/Hello-World/pkg/render"
	"net/http"
)

//home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

//about page handler
func About(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the about page")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	res, err := divideValues(10.0, 0.0)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f is divided by %f is %f", 10.0, 0.0, res))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cant divide by 0")
		return 0, err
	}
	res := x / y
	return res, nil
}

func Hello(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}
