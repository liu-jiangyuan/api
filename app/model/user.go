package model

import (
	"github.com/liu-jiangyuan/api/lib"
	"log"
)

type Nft struct {
	Id                        int `xorm:"pk autoincr"`
	ImgId                     int `xorm:"img_id"`
	Quality                   int
	Name                      string `xorm:"text notnull"`
	FullDescription           string `xorm:"full_description text notnull"`
	Description               string `xorm:"text notnull"`
}

func (u *Nft) TableName() string  {
	return "nft_names"
}



func (u *Nft) Before(closures func(interface{})) {
	log.Println("before")
}

func (u *Nft) After(closures func(interface{})) {
	log.Println("after")
}

func (u *Nft) Get() (Nft,error) {
	var nft Nft
	b, err := lib.Engine.Id(u.Id).Get(&nft)
	log.Println(b,err)
	return nft,err
}

func (u *Nft) Find() ([]Nft,error) {
	nft := make([]Nft,0)
	err := lib.Engine.Where("id < ?",9).Find(&nft)
	return nft,err
}