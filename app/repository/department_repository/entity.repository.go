package department_repository

import "time"

type Department struct {
	ID        int
	ManagerId int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}