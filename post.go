package main

type post struct {
	slug string
}

func newPost(slug string) *post {
	thePost := post{slug: slug}
	return &thePost
}
