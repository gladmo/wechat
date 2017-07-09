package models

type Img_joke struct {
	I_id        int64 `gorm:"primary_key"`
	C_id        int64
	Comment     string // image comment
	Img_list    string // 图片列表
	Source_list string // 原图片地址
	Img_count   int64  // 图片数量
	Create_at   int64
}

func (c *Img_joke) Save() (res bool) {
	db := Connect()
	defer db.Close()
	res = db.NewRecord(c)
	db.Create(c)
	return
}
