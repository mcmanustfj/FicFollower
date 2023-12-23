package main

import (
	"time"

	"github.com/google/uuid"
)

type Model interface {
	id() 
}

type User struct {
	username string
}

type Story struct {
	id uuid.UUID
	title string
	site string
	link string
	updated string
	updatedlink string
	author string
	authorlink string
	checked time.Time
}

type UserFollows struct {
	username string
	storyId uuid.UUID
	readDate time.Time
	readLink string
}

type Model = (User | Story | UserFollows)
