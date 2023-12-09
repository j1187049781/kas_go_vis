package models


// comment "addr_tag"
type AddrTag struct {
     Model
     Address string  `gorm:"column:address;type:VARCHAR(255);not null;comment:address" json:"address"`
     Name string  `gorm:"column:name;type:VARCHAR(255);not null;comment:name" json:"name"`
     Link string  `gorm:"column:link;type:VARCHAR(255);null;comment:link" json:"link"`
}

func (AddrTag) TableName() string {
     return "addr_tag"
}