package utils

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Total        int64       `json:"total"`
	PerPage      int         `json:"per_page"`
	CurrentPage  int         `json:"current_page"`
	LastPage     int         `json:"last_page"`
	NextPage     int         `json:"next_page"`
	PreviousPage int         `json:"previous_page"`
	Data         interface{} `json:"data"`
}

// Paginate 执行分页查询并序列化数据
func Paginate(db *gorm.DB, table interface{}, conditions map[string]interface{}, page, pageSize int, serializer func(interface{}) interface{}) (*Pagination, error) {
	var count int64
	query := db.Model(table).Where(conditions)
	err := query.Count(&count).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize
	data := make([]interface{}, 0)
	err = query.Limit(pageSize).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, err
	}

	serializedData := make([]interface{}, 0, len(data))
	for _, item := range data {
		serializedData = append(serializedData, serializer(item))
	}

	lastPage := int(count / int64(pageSize))
	if count%int64(pageSize) > 0 {
		lastPage++
	}

	pagination := &Pagination{
		Total:        count,
		PerPage:      pageSize,
		CurrentPage:  page,
		LastPage:     lastPage,
		NextPage:     0,
		PreviousPage: 0,
		Data:         serializedData,
	}

	if page > 1 {
		pagination.PreviousPage = page - 1
	}

	if page < lastPage {
		pagination.NextPage = page + 1
	}

	return pagination, nil
}
