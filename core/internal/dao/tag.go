package dao

import (
	"errors"

	"github.com/krau/manyacg/core/internal/common/logger"
	entityModel "github.com/krau/manyacg/core/internal/model/entity"
	"gorm.io/gorm"
)

func CheckTagExist(tagName string) bool {
	var tag entityModel.Tag
	err := db.Where("name = ?", tagName).First(&tag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else if err != nil {
		logger.L.Errorf("Failed to get tag by name: %s", err)
		return false
	}
	return true
}

// 批量添加标签, 存在则忽略, 在内部处理错误且不返回
func AddTags(tags []*entityModel.Tag) {
	for _, tag := range tags {
		err := db.FirstOrCreate(tag, "name = ?", tag.Name).Error
		if err != nil {
			logger.L.Errorf("Failed to create tag: %s", err)
		}
	}
}

// 批量添加标签, 存在则忽略. 如有错误将返回
func addTags(tags []*entityModel.Tag, tx *gorm.DB) error {
	if tx == nil {
		tx = db
	}
	for _, tag := range tags {
		err := tx.FirstOrCreate(tag, "name = ?", tag.Name).Error
		if err != nil {
			logger.L.Errorf("Failed to create tag: %s", err)
			return err
		}
	}
	return nil
}
