package model

type Player struct {
	Id        int    `gorm:"AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(256);not null;"`
	Password  string `gorm:"type:varchar(256);not null;"`
	History   string `gorm:"type:varchar(1024);"`
	Prepare   int    `gorm:"type:int(255);"`//是否准备好了，是 1 否 0 ,开局之后用prepare记录落子情况
	Identify string `gorm:"type:varchar(256);"`//进入身份，是 gobanger 否 audience
}

type Room struct {
	Id        int    `gorm:"type:varchar(256);not null;"`
	Password  string `gorm:"type:varchar(256);not null;"`
	Member1   string `gorm:"type:varchar(256);not null;"`
	Member2   string `gorm:"type:varchar(256);not null;"`
	Statu     int    `gorm:"type:int(255);not null;"`//是否满员，是 1 否 0
	Balance   string `gorm:"type:varchar(256);not null;"`//是否禁手，是 balance 否 unbalance
}

type History struct{
	Id        int    `gorm:"type:varchar(256);not null;"`
	Step      string `gorm:"type:varchar(4048);"`
}