package model

type Page struct {
	Books       []*Book
	PageNo      int64
	PageSize    int64
	TotalPageNo int64
	TotalRecord int64
	MinPrice    float64
	MaxPrice    float64
	IsLogin     bool
	Username    string
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// GetPrev 获取上一页
func (p *Page) GetPrev() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

// IsHasNext 判断是否有上一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetNext 获取下一页
func (p *Page) GetNext() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}

// IsByPrice 是否依据价格查询
func (p *Page) IsByPrice() bool {
	if p.MaxPrice == 0 && p.MinPrice == 0 {
		return false
	} else {
		return true
	}
}
