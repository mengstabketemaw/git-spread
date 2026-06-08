package models

import "time"

type Commit struct {
	Hash string
	Message string
	Date time.Time
}
