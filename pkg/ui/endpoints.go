package ui

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

func BodySwapHandler(w http.ResponseWriter, r *http.Request) {
	getSelection := func(sel string) int {
		switch sel {
		case "cv":
			return CvPage
		case "home":
			return LandingPage
		case "blog":
			return BlogPage
		default:
			fmt.Println("BodySwapHandler hit without valid Hx-Page-Selection")
		}
		return -1
	}

	selection := getSelection(r.Header.Get("Hx-Page-Selection"))

	buf := bytes.Buffer{}
	GetBody(selection, -1).Render(context.Background(), &buf)
	w.Write(buf.Bytes())
}

func BlogBodyHandler(w http.ResponseWriter, r *http.Request) {

	blogId := -1
	pageSelect := BlogPostView

	if r.Header.Get("Hx-Page-Selection") != "blog-post" {
		fmt.Println("Blog endpoint hit without correct page selection")
	}
	blogSelectStr := r.Header.Get("Hx-Blog-Selection")
	if blogSelectStr == "" {
		fmt.Println("Blog endpoint hit without a blog selection")
	} else {
		tid, err := strconv.ParseInt(blogSelectStr, 10, 8)
		if err != nil {
			fmt.Println("Error parsing blogId")
			tid = -1
		}
		blogId = int(tid)
	}
	//todo get blog ID if exists

	buf := bytes.Buffer{}
	GetBody(pageSelect, blogId).Render(context.Background(), &buf)
	w.Write(buf.Bytes())
}
