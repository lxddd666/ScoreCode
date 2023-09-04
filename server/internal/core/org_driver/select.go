package driver

import (
	"context"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/library/contexts"
	"hotgo/internal/service"
)

var (
	orgIdKey = "org_id"
	tables   = "hg_admin_post"
)

func Stmt(ctx context.Context, sql string) (bool, string) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return false, sql
	}
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		return Select(ctx, stmt)
	}
	return false, sql
}

func Select(ctx context.Context, stmt *sqlparser.Select) (bool, string) {
	hasTable := false
	for _, expr := range stmt.From {
		switch tableExpr := expr.(type) {
		case *sqlparser.AliasedTableExpr:
			if !hasTable {
				hasTable = SetOrgId(ctx, stmt, tableExpr)
			} else {
				SetOrgId(ctx, stmt, tableExpr)
			}
		case *sqlparser.JoinTableExpr:
			if !hasTable {
				hasTable = JoinTableExpr(ctx, stmt, tableExpr)
			} else {
				JoinTableExpr(ctx, stmt, tableExpr)
			}

		}
	}
	return hasTable, sqlparser.String(stmt)
}

func JoinTableExpr(ctx context.Context, stmt *sqlparser.Select, expr *sqlparser.JoinTableExpr) bool {
	hasTable := false
	switch expr := expr.RightExpr.(type) {
	case *sqlparser.AliasedTableExpr:
		if !hasTable {
			hasTable = SetOrgId(ctx, stmt, expr)
		} else {
			SetOrgId(ctx, stmt, expr)
		}
	case *sqlparser.JoinTableExpr:
		if !hasTable {
			hasTable = JoinTableExpr(ctx, stmt, expr)
		} else {
			JoinTableExpr(ctx, stmt, expr)
		}
	}
	switch expr := expr.LeftExpr.(type) {
	case *sqlparser.AliasedTableExpr:
		if !hasTable {
			hasTable = SetOrgId(ctx, stmt, expr)
		} else {
			SetOrgId(ctx, stmt, expr)
		}
	case *sqlparser.JoinTableExpr:
		if !hasTable {
			hasTable = JoinTableExpr(ctx, stmt, expr)
		} else {
			JoinTableExpr(ctx, stmt, expr)
		}
	}

	return hasTable
}

func AliasedTableExpr(table *sqlparser.AliasedTableExpr) (tableName, alias string) {
	switch expr := table.Expr.(type) {
	case sqlparser.TableName:

		if table.As.String() != "" {
			alias = table.As.String()
		}
		if expr.Name.String() == tables {
			tableName = expr.Name.String()
			if alias != "" {
				return

			}
			return
		}
	}
	return
}

func SetOrgId(ctx context.Context, stmt *sqlparser.Select, table *sqlparser.AliasedTableExpr) bool {
	tableName, alias := AliasedTableExpr(table)
	user := contexts.GetUser(ctx)
	if tableName != "" && !service.AdminMember().VerifySuperId(ctx, user.Id) {
		stmt.AddWhere(buildOrgIdExpr(ctx, alias, user.OrgId))
		return true
	}
	return false
}

func buildOrgIdExpr(ctx context.Context, alias string, orgId int64) *sqlparser.ComparisonExpr {
	return &sqlparser.ComparisonExpr{
		Operator: "=",
		Left: &sqlparser.ColName{
			Metadata:  nil,
			Name:      sqlparser.NewColIdent(orgIdKey),
			Qualifier: sqlparser.TableName{Name: sqlparser.NewTableIdent(alias)},
		},
		Right:  sqlparser.NewStrVal([]byte(gconv.String(orgId))),
		Escape: nil,
	}
}
