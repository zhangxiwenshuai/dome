package models
/**
* 构建用户结构体，用于表示数据结构的定义
 */
type User struct {
	User string 'json:"name"'
	Birthday string 'json:"birthday"'
	Address string 'json:"address"'
	Nick string 'json:"nick"'
}