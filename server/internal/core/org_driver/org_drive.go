package driver

import (
	"context"
	"github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
)

type OrgDriver struct {
	*mysql.Driver
}

func init() {
	var (
		err         error
		driverObj   = &OrgDriver{}
		driverNames = g.SliceStr{"mysql", "mariadb", "tidb"}
	)
	for _, driverName := range driverNames {
		if err = gdb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *OrgDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &OrgDriver{
		&mysql.Driver{
			Core: core,
		},
	}, nil
}

func (d *OrgDriver) DoQuery(ctx context.Context, link gdb.Link, sql string, args ...interface{}) (result gdb.Result, err error) {
	hasTable, newSql := Stmt(ctx, sql)
	if hasTable {
		replace, err := gregex.Replace(`:v\d+`, []byte("?"), []byte(newSql))
		if err == nil {
			sql = string(replace)
		}
	}
	return d.Driver.DoQuery(ctx, link, sql, args...)
}
