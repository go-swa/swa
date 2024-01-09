package service_file

import (
	"github.com/ServiceWeaver/weaver"
	"net/textproto"
)

type ModUpload struct {
	weaver.AutoMarshal
	ModBase
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

func (ModUpload) TableName() string {
	return "swa_upload"
}

type FileHeader struct {
	weaver.AutoMarshal
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	content   []byte
	tmpfile   string
	tmpoff    int64
	tmpshared bool
}



