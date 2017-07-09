package models

type Crawl struct {
	C_id      int64  `gorm:"primary_key"`
	Unique    string `gorm:"size:50;unique_index"` // 唯一标识，用于查重
	Source    string `gorm:"size:50"`              // 来源
	Url       string `gorm:"size:255"`             // 爬取地址
	Is_del    int64  // 是否删除
	Create_at int64
}

func (c *Crawl) Save() (res bool) {
	db := Connect()
	defer db.Close()
	res = db.NewRecord(c)
	db.Create(c)
	return
}

func FindByUnique(unique string, source string) (ok bool) {
	db := Connect()
	defer db.Close()

	var crawl Crawl
	db.Select("c_id").Where("`unique` = ? AND source= ? AND is_del = 0", unique, source).First(&crawl)
	if crawl.C_id != 0 {
		return false
	}
	return true

}
