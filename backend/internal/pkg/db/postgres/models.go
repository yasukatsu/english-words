package postgres

import (
	"time"
)

// Words ...
type Words struct {
	Id       int       `xorm:"not null pk autoincr INTEGER"`
	Spel     string    `xorm:"VARCHAR(255)"`
	Define   string    `xorm:"TEXT"`
	Pos      string    `xorm:"TEXT"`
	Meaning  string    `xorm:"TEXT"`
	UpdateAt time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	CreateAt time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}
