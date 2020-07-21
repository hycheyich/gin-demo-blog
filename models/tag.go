package models

import "fmt"

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy   string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
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

func AddTag(name string, state int, CreateBy string) bool {
	db.Create(&Tag{
		Name:     name,
		State:    state,
		CreatedBy: CreateBy,
	})
	fmt.Println("新增tag成功")
	return true
}
//
//func (tag *Tag)()  {
//
//}