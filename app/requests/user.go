package requests

import "server/global"

type AdminList struct {
	global.List
}

type AdminUpdate struct {
	Id        uint   `json:"id"`
	Name      string `json:"name" binding:"required" msg:"用户名不能为空"`
	PassWord  string `json:"password,omitempty"`
	CPassWord string `json:"password2,omitempty"`
	Email     string `json:"email" binding:"required,email" msg:"邮件格式错误"`
	Phone     string `json:"phone" binding:"required" valid:"matches(^1[0-9]{10}$)" msg:"手机号格式错误"`
	Gender    int    `json:"gender" binding:"required,max=2" min_msg:"长度最大于1" msg:"性别不能为空"`
	RealName  string `json:"real_name" binding:"required,min=2" min_msg:"长度最小大于2" msg:"真实姓名不能为空"`
	Avatar    string `json:"avatar,omitempty"`
	Roles     []uint `json:"group"`
}

//	type UserAdd struct {
//		Name      string `json:"name" binding:"required" msg:"用户名不能为空"`
//		PassWord  string `json:"password" binding:"required,min=3,eqfield=CPassWord" min_msg:"长度最小大于3" eqfield_msg:"两次输入密码不一致" msg:"密码不能为空"`
//		CPassWord string `json:"password2" binding:"required,min=3" min_msg:"长度最小大于3" msg:"密码不能为空"`
//		RealName  string `json:"real_name" binding:"required,min=2" min_msg:"长度最小大于2" msg:"真实姓名不能为空"`
//		Avatar    string `json:"avatar" binding:"required,min=3" min_msg:"长度最小大于3" msg:"通向不能为空"`
//		Roles     []uint `json:"group"`
//	}
type UserAdd struct {
	Name     string `json:"name" binding:"required" msg:"用户名不能为空"`
	RealName string `json:"real_name" binding:"required,min=2" min_msg:"长度最小大于2" msg:"真实姓名不能为空"`
	Email    string `json:"email" binding:"required,email" msg:"邮件格式错误"`
	Phone    string `json:"phone" binding:"required" valid:"matches(^1[0-9]{10}$)" msg:"手机号格式错误"`
	PassWord string `json:"password" binding:"required,min=3" min_msg:"长度最小大于3" msg:"密码不能为空"`
	Gender   int    `json:"gender" binding:"required,max=2" min_msg:"长度最大于1" msg:"性别不能为空"`
	Roles    []uint `json:"group"`
}
