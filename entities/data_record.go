package entities

import "time"

type DataRecord struct {
	UUID             int
	Author           string
	Message          string
	Likes            int
	AuthorUpdatedAt  *time.Time
	MessageUpdatedAt *time.Time
	LikesUpdatedAt   *time.Time
	CreatedAt        time.Time
	UpdatedAt        bool
	DeletedAt        bool
}
