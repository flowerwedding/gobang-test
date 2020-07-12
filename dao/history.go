package dao

import (
	"strings"
	inits "test/model"
)

func UnbalanceSelect(h *inits.History,id int){
	inits.DB.Model(&inits.History{}).Where("id = ?",id).First(&h)
}

func History(id int,s string) error {
	var h inits.History
	inits.DB.Model(&inits.History{}).Where("id = ?",id).First(&h)
	err := inits.DB.Model(inits.History{}).Where("id = ?",id).Update("step",h.Step + " # " + s).Error

	if err != nil {
		return err
	}
	return nil
}

func SelectHistory(id int) []string {
	var h inits.History
	inits.DB.Model(&inits.History{}).Where("id = ?",id).First(&h)
	return strings.Split(h.Step, " # ")
}