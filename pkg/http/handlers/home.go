package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fathimtiaz/galaxy-merchant/config"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		fmt.Println("--a", err)
	}

    t.Execute(w, struct{Host string}{Host: config.CONF.Host})
}
