package example

type User struct {
	Username string `json:"userName" form:"userName" gorm:"comment:用户名"`
	UserPassword string `json:"userPassword" form:"userPassword" gorm:"comment:用户密码"`
	Type int `json:"type" form:"type" gorm:"comment:类型"`
	Phone int `json:"phone" form:"phone" gorm:"comment:手机号"`
	Salt string `json:"salt" form:"Salt" gorm:"comment:密钥"`
}
