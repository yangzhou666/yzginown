package db

type Pagination struct {
	Page   int   `json:"page"`             // 页码
	Limit  int   `json:"limit"`            // 每页条数
	Total  int32 `json:"total"`            // 总数据条数
	Offset int   `json:"offset,omitempty"` // 自定义偏移值
}

func NewPagination() *Pagination {
	return &Pagination{
		Page:  1,
		Limit: 10,
	}
}

// GetOffset 获取分页偏移值
func (p *Pagination) GetOffset() int {
	if p.Offset > 0 {
		return p.Offset
	}
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

// TotalPage 获取总页数
func (p *Pagination) TotalPage() int32 {
	if p.Total == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := p.Total / int32(p.Limit)
	if p.Total%int32(p.Limit) > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}
