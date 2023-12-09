package models

// comment "addr"
type AddrBalance struct {
	Model
	Time    string  `gorm:"column:time;type:DateTime;not null;comment:time" json:"time"`
	Address string  `gorm:"column:address;type:VARCHAR(255);not null;comment:address" json:"address"`
	Balance float64 `gorm:"column:balance;type:DECIMAL(12,9);default:0.0;comment:balance" json:"balance"`

	Tags []AddrTag `gorm:"many2many:addr_tag;foreignKey:Address;joinForeignKey:Address;References:Address;JoinReferences:Address"`
}

func (AddrBalance) TableName() string {
	return "addr_balance"
}