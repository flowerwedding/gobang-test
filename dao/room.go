package dao

import inits "test/model"

//通过查找房间号判断房间是否已经存在
func GetroomSearch(room *inits.Room){
	inits.DB.Model(&inits.Room{}).Where(&inits.Room{Id: room.Id, Password: room.Password}).First(&room)
}

func GetroomInsert(room *inits.Room)error{
	err := inits.DB.Model(&inits.Room{}).Create(&room).Error
	return err
}

func OutroomDelete(r *inits.Room,perid string){
	inits.DB.Model(&inits.Room{}).Where("id = ?",perid).First(&r)
	inits.DB.Model(&inits.Room{}).Where("id = ?",perid).Delete(&r)
}

func RandomSelect(number string,r *inits.Room){
	inits.DB.Model(&inits.Room{}).Where("id = ?",number).First(&r)
}

func RandomSelect2(balance string,r *inits.Room){
	inits.DB.Model(&inits.Room{}).Where("balance = ? and statu = 0 ",balance).Take(&r)
}

func RandomUpdate(r *inits.Room,premember2 interface{})(error,error){
	err1 := inits.DB.Model(inits.Room{}).Where("id = ?",r.Id).Update("member2",premember2).Error
	err2 := inits.DB.Model(inits.Room{}).Where("id = ?",r.Id).Update("statu",1).Error
	return err1,err2
}

func RandomUpdate2(premember2 interface{},identify string)error{
	err3 := inits.DB.Model(inits.Player{}).Where("username = ?",premember2).Update("identify",identify).Error
	return err3
}