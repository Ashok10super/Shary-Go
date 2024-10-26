package model

//creating a post model

type Post struct{
	id  int  `json:"id"`
	userid  int  `json:"userid"`
	title   string `json:"title`
	body    string `json:"body`
}

