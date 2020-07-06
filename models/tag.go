package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 添加创建时间
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// 添加更新时间
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// 获取文章tag
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 获取标签总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 编辑根据名字标签
func ExitTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//新增标签
func AddTag(name string, state int, CreateBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: CreateBy,
	})
	fmt.Println("新增tag成功")
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	fmt.Println("修改tag成功")
	return true
}
