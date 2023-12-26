package dal_test

import (
	"kas_go_vis/db/dal"
	"testing"
	"gorm.io/gorm"
)
type User struct {
    gorm.Model
    Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
    Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
    gorm.Model
    Name      string
    UserRefer uint `gorm:"index:,unique"`
}
func TestDa1(t *testing.T) {
	db := dal.ConnectDB(dal.MySQLDSN)
	err := db.AutoMigrate(&User{},&Profile{})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	
}