package main

import (
	"fmt"
	"github.com/a-h/templ"
	"main/pkg/md"
	"main/pkg/ui"
	"net/http"
)

func main() {

	md.ProcessBlogPosts()
	mainComponent := ui.MainPage(ui.LandingPage)
	http.Handle("/", templ.Handler(mainComponent))

	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("./lib"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/body-swap", ui.BodySwapHandler)
	http.HandleFunc("/blog-swap", ui.BlogBodyHandler)

	fmt.Println("Server Starting.")

	http.ListenAndServe(":3000", nil)

}
