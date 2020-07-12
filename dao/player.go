package dao

import inits "test/model"

//注册的时候通过找用户名判断用户是否存在
func LoginSearch(user *inits.Player){
	inits.DB.Model(&inits.Player{}).Where("username = ?",user.Username).First(&user)
}

//注册的时候新建用户
func LoginInsert(user *inits.Player) error{
	err :=  inits.DB.Model(&inits.Player{}).Create(&user).Error
	return err
}

//登录的时候查找用户名和密码判断用户是否正确
func RegisterSearch(user *inits.Player){
	inits.DB.Model(&inits.Player{}).Where(&inits.Player{Username: user.Username, Password: user.Password}).First(&user)
}