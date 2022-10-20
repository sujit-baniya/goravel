package models

import (
	"github.com/sujit-baniya/framework/database/orm"
)

type User struct {
	orm.Model
	Name   string
	Avatar string
	orm.SoftDeletes
}
