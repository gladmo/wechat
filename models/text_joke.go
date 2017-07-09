package models

type Text_joke struct {
	T_id      int64 `gorm:"primary_key"`
	C_id      int64
	Source    string `gorm:"size:50"`   // 来源
	Content   string `gorm:"type:text"` // 内容
	Create_at int64
}

func (c *Text_joke) Save() (res bool) {
	db := Connect()
	defer db.Close()
	res = db.NewRecord(c)
	db.Create(c)
	return
}
