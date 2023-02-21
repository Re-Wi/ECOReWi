package entity

import (
	"github.com/XM-GO/PandaKit/model"
	"time"
)

type EcoDynamic struct {
	Date   time.Time `json:"date" gorm:"column:日期"  form:"date"`
	Sign   string    `json:"sign" gorm:"type:varchar(1);comment:符号" form:"sign"`
	Amount int64     `json:"amount" gorm:"type:int(11);column:金额" form:"amount"`
	Item   string    `json:"item" gorm:"type:varchar(255);comment:事项" form:"item"`
	Remark string    `json:"remark" gorm:"type:varchar(255);comment:备注" form:"remark"`
	model.BaseAutoModel
}
