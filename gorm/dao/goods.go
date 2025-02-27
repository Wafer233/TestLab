package dao

// Good -> Model
// 默认gorm对struct字段名使用Snake Case命名风格转换成mysql表字段名(需要转换成小写字母)。
// https://gorm.io/zh_CN/docs/models.html
type Good struct {
	Id         int     //表字段名为：id
	Name       string  //表字段名为：name
	Price      float64 //表字段名为：price
	TypeId     int     //表字段名为：type_id
	CreateTime int64   `gorm:"column:createtime"` //表字段名为：createtime
}
