package models


// comment "addr_tag"
type AddrTag struct {
     Name string  `gorm:"column:name;type:VARCHAR(255);null;comment:name;primary_key" json:"name"`
     Link string  `gorm:"column:link;type:VARCHAR(255);null;comment:link" json:"link"`
}

func (AddrTag) TableName() string {
     return "addr_tag"
}