package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/liu-jiangyuan/api/app/model"
	"math/big"
)

type Index struct {
	Ctx iris.Context
	Num *big.Int
}

func (i *Index) Get () iris.Map {
	n := model.Nft{
		Id:              3,
	}
	nft , err := n.Find()
	return iris.Map{
		"nft": nft,
		"err": err,
	}
}
func(i *Index) GetWelcome() interface{} {
	return iris.Map{
		"Num": i.Num,
	}
}
