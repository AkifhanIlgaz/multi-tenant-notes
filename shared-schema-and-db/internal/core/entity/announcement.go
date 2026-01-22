package entity

import "time"

type Announcement struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time

	UserId   int
	TenantId int
}
