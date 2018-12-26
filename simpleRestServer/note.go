package main

import "time"

type Note struct {
	Title, Description string
	CreatedOn time.Time
}