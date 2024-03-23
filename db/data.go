package db

type blog struct{
	ID int	`json:"id"`
	Title string	`json:"title"`
	Content string	`json:"content"`
}

var Blogs = [] blog{
	{ID: 1, Title: "First Blog", Content: "This is the first blog"},
	{ID: 2, Title: "Second Blog", Content: "This is the second blog"},
	{ID: 3, Title: "Third Blog", Content: "This is the third blog"},
}