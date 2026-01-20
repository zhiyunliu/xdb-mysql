package demos

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/zhiyunliu/glue"
	"github.com/zhiyunliu/glue/context"
	"github.com/zhiyunliu/golibs/xtypes"
	"github.com/zhiyunliu/golibs/xtypes/datetime"
)

type DataItem struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	Birthday1 *time.Time         `json:"birthday1"`
	Birthday2 *time.Time         `json:"birthday2"`
	Birthday3 *datetime.DateTime `json:"birthday3"`
	Value     float64            `json:"value"`
	Age       *int               `json:"age"`
	Addr      string             `json:"addr"`
	Money     xtypes.Decimal     `json:"money"`

	Bit        []byte         `json:"bit"`
	Boolean    bool           `json:"boolean"`
	Tinyint    int            `json:"tinyint"`
	Smallint   int            `json:"smallint"`
	Int        int            `json:"int"`
	Bigint     int64          `json:"bigint"`
	Float      float32        `json:"float"`
	Double     float64        `json:"double"`
	Decimal    xtypes.Decimal `json:"decimal"`
	Char       string         `json:"char"`
	Nchar      string         `json:"nchar"`
	Varchar    string         `json:"varchar"`
	Nvarchar   string         `json:"nvarchar"`
	Text       string         `json:"text"`
	Mediumtext string         `json:"mediumtext"`
	Longtext   string         `json:"longtext"`
	Enum       string         `json:"enum"`
	Time       string         `json:"time"`
	Year       *int32         `json:"year"`
	Tinyblob   []byte         `json:"tinyblob"`
	Blob       []byte         `json:"blob"`
	Mediumblob []byte         `json:"mediumblob"`
	Longblob   []byte         `json:"longblob"`
	Binary     Binary         `json:"binary"`
	Varbinary  []byte         `json:"varbinary"`
}

type Binary struct {
	Val []byte
}

func (b *Binary) Scan(src any) error {
	if src == nil {
		return nil
	}
	*b = Binary{
		Val: src.([]byte),
	}
	return nil
}

func (b Binary) MarshalJSON() ([]byte, error) {
	if b.Val == nil {
		return []byte("null"), nil
	}
	return json.Marshal(string(b.Val))
}

func (b Binary) Value() (driver.Value, error) {
	return b.Val, nil
}

func Query(ctx context.Context) any {
	id := ctx.Request().Query().Get("id")
	result := []DataItem{}

	dbObj := glue.DB("demo")
	err := dbObj.QueryAs(ctx.Context(), `select  * from test_mysql where id=@{id}`, map[string]any{"id": id}, &result)
	if err != nil {
		return err
	}

	return result
}
