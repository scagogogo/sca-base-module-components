package db

import (
	"github.com/scagogogo/sca-base-module-dao/mysql"
)

func init() {

	if mysql.Gorm == nil {
		return
	}

}
