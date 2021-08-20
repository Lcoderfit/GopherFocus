// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// FileDao is the manager for logic model data accessing and custom defined data operations functions management.
type FileDao struct {
	gmvc.M             // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      fileColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB      // DB is the raw underlying database management object.
	Table  string      // Table is the underlying table name of the DAO.
}

// FileColumns defines and stores column names for table gf_file.
type fileColumns struct {
	Id        string // 自增ID
	Name      string // 文件名称
	Src       string // 本地文件存储路径
	Url       string // URL地址，可能为空
	UserId    string // 操作用户
	CreatedAt string //
}

// NewFileDao creates and returns a new DAO object for table data access.
func NewFileDao() *FileDao {
	columns := fileColumns{
		Id:        "id",
		Name:      "name",
		Src:       "src",
		Url:       "url",
		UserId:    "user_id",
		CreatedAt: "created_at",
	}
	return &FileDao{
		C:     columns,
		M:     g.DB("default").Model("gf_file").Safe(),
		DB:    g.DB("default"),
		Table: "gf_file",
	}
}