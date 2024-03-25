package md

import (
	"encoding/json"
	"fmt"
	"os"
)

var GlobalBlogData map[int]BlogData

type BlogData struct {
	ID          int
	Html        string
	ImagePath   string
	Title       string
	Description string
}

type PostData struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Postfile  string `json:"postfile"`
	Imagefile string `json:"imagefile"`
	Desc      string `json:"desc"`
}

type PostsDataArr struct {
	Posts []PostData `json:"posts"`
}

func ProcessBlogPosts() {

	GlobalBlogData = make(map[int]BlogData, 0)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("could not get pwd")
	}

	content, err := os.ReadFile(pwd + "/assets/blog/blogdata.json")
	if err != nil {
		panic(err)
	}

	var posts PostsDataArr
	err = json.Unmarshal(content, &posts)
	if err != nil {
		panic(err)
	}

	for _, post := range posts.Posts {
		blogHtml := ProcessFile("/assets/blog/posts/" + post.Postfile)
		GlobalBlogData[post.ID] = BlogData{
			ID:          post.ID,
			Html:        blogHtml,
			Title:       post.Title,
			Description: post.Desc,
			ImagePath:   "/assets/blog/img/" + post.Imagefile,
		}
	}
}
