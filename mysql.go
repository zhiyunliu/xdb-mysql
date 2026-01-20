package xdbmysql

import (
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhiyunliu/glue/config"
	contribxdb "github.com/zhiyunliu/glue/contrib/xdb"
	"github.com/zhiyunliu/glue/contrib/xdb/expression"
	"github.com/zhiyunliu/glue/contrib/xdb/tpl"
	"github.com/zhiyunliu/glue/xdb"
	"github.com/zhiyunliu/golibs/xtypes"
)

const Proto = "mysql"

type mysqlResolver struct {
}

func (s *mysqlResolver) Name() string {
	return Proto
}

func (s *mysqlResolver) Resolve(connName string, setting config.Config, opts ...xdb.Option) (interface{}, error) {
	cfg := contribxdb.NewConfig(connName)
	err := setting.ScanTo(cfg.Cfg)
	if err != nil {
		return nil, fmt.Errorf("读取DB配置:%w", err)
	}
	return contribxdb.NewDB(Proto, cfg, opts...)
}
func init() {
	tplMatcher := xdb.NewTemplateMatcher(expression.DefaultExpressionMatchers...)
	tplstmtProcessor := xdb.NewStmtDbTypeProcessor()

	xdb.Register(&mysqlResolver{})
	_ = xdb.RegistTemplate(tpl.NewFixed(Proto, "?", tplMatcher, tplstmtProcessor))
	registerProtopDbtype()
}

func registerProtopDbtype() {

	_ = xdb.RegisterProtoDbType(Proto, "BIT", reflect.TypeOf((*[]byte)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "BOOLEAN", reflect.TypeOf((*bool)(nil)).Elem())

	_ = xdb.RegisterProtoDbType(Proto, "TINYINT", reflect.TypeOf((*int)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "SMALLINT", reflect.TypeOf((*int)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "INTEGER", reflect.TypeOf((*int)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "INT", reflect.TypeOf((*int)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "BIGINT", reflect.TypeOf((*int64)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "FLOAT", reflect.TypeOf((*float32)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "DOUBLE", reflect.TypeOf((*float64)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "DECIMAL", reflect.TypeOf((*xtypes.Decimal)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "MONEY", reflect.TypeOf((*xtypes.Decimal)(nil)).Elem())

	_ = xdb.RegisterProtoDbType(Proto, "CHAR", reflect.TypeOf((*string)(nil)).Elem())
	//_ = xdb.RegisterProtoDbType(Proto, "NCHAR", reflect.TypeOf((*string)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "VARCHAR", reflect.TypeOf((*string)(nil)).Elem())
	//_ = xdb.RegisterProtoDbType(Proto, "NVARCHAR", reflect.TypeOf((*string)(nil)).Elem())

	_ = xdb.RegisterProtoDbType(Proto, "TEXT", reflect.TypeOf((*string)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "MEDIUMTEXT", reflect.TypeOf((*string)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "LONGTEXT", reflect.TypeOf((*string)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "ENUM", reflect.TypeOf((*string)(nil)).Elem())

	_ = xdb.RegisterProtoDbType(Proto, "DATE", reflect.TypeOf((*time.Time)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "TIME", reflect.TypeOf((*string)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "DATETIME", reflect.TypeOf((*time.Time)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "TIMESTAMP", reflect.TypeOf((*time.Time)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "YEAR", reflect.TypeOf((*int)(nil)).Elem())

	_ = xdb.RegisterProtoDbType(Proto, "TINYBLOB", reflect.TypeOf((*[]byte)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "BLOB", reflect.TypeOf((*[]byte)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "MEDIUMBLOB", reflect.TypeOf((*[]byte)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "LONGBLOB", reflect.TypeOf((*[]byte)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "BINARY", reflect.TypeOf((*[]byte)(nil)).Elem())
	_ = xdb.RegisterProtoDbType(Proto, "VARBINARY", reflect.TypeOf((*[]byte)(nil)).Elem())

}
