package flightORM

import (
	"github.com/fengleng/flight/sqlparser/sqlparser"
	"strconv"
	"testing"
)

func TestCopy(t *testing.T) {

	var sli = make([]byte, 0, 8)
	sli = strconv.AppendInt(sli, 9, 10)

	v := sqlparser.Where{Type: sqlparser.WhereStr,
		Expr: &sqlparser.ComparisonExpr{
			Left:     &sqlparser.ColName{Name: sqlparser.NewColIdent("id")},
			Right:    sqlparser.NewIntVal(sli),
			Operator: sqlparser.EqualStr},
	}

	v2 := v
	v2.Type = sqlparser.HavingStr

	t.Log(sqlparser.String(&v))
	t.Log(sqlparser.String(&v2))
	//t.Log(sqlparser.String(v2))
	//var vv = &sqlparser.Where{}
	////err := Copy(vv, v)
	//
	//hh, err := json.Marshal(vv)
	//err = json.Unmarshal(hh, vv)
	//t.Log(err)
	//
	////clone
	//t.Log(vv)
}
