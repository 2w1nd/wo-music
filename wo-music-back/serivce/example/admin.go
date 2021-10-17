package example

import (
	"github.com/wo-music-back/server/global"
	"github.com/wo-music-back/server/model/example"
)

type AdminService struct {
}

// CreateAdmin
// @Description: 创建用户
// @receiver: exa
// @param: e
// @return: err
func (exa *AdminService) CreateAdmin(e example.Admin) (err error) {
	err = global.DB.Create(&e).Error
	return err
}

// DeleteAdmin
// @Description: 删除用户
// @receiver: exa
// @param: e
// @return: err
func (exa *AdminService) DeleteAdmin(e example.Admin) (err error) {
	err = global.DB.Delete(&e).Error
	return err
}

// UpdateAdmin
// @Description: 更新用户
// @receiver: exa
// @param: e
// @return: err
func (exa *AdminService) UpdateAdmin(e *example.Admin) (err error) {
	err = global.DB.Save(e).Error
	return err
}
