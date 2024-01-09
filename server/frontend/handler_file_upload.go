package frontend

import (
	"demo1/service/service_file"
	"demo1/utils"
	"demo1/utils/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)


func (s *server) InitRouterUpDownload(Router *gin.RouterGroup) {
	router := Router.Group("fileUpload").Use(s.OperationRecord())

	router.POST("upload", s.UploadFile)         // 上传文件
	router.POST("getFileList", s.GetFileList)   // 获取上传文件列表
	router.POST("deleteFile", s.DeleteFile)     // 删除指定文件
	router.POST("editFileName", s.EditFileName) // 编辑文件名或者备注
}

func (s *server) UploadFile(c *gin.Context) {
	var file service_file.ModUpload

	_, header, err := c.Request.FormFile("file")
	if err != nil {
		s.Logger(c).Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}

	err = c.Request.ParseMultipartForm(4 << 20) // 4Mb
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("文件太大,请小于4M:%v", err), c)
		return
	}

	headerFile := strings.Split(header.Filename, ".")
	if len(headerFile) < 1 {
		s.Logger(c).Error("上传文件没有文件类型!")
		response.FailWithMessage("上传文件没有文件类型", c)
		return
	}

	fileType := headerFile[len(headerFile)-1]
	uid, _ := uuid.NewV4()
	ustring := uid.String()

	saveFile := ustring + "." + fileType

	file = service_file.ModUpload{
		Url:  s.swaConfig.Local.StorePath + "/" + saveFile, // 返回路径
		Name: header.Filename,
		Tag:  headerFile[len(headerFile)-1],
		Key:  ustring,
	}
	if err := c.SaveUploadedFile(header, s.swaConfig.Local.StorePath+"/"+saveFile); err != nil {
		s.Logger(c).Error("保存上传文件失败!", zap.Error(err))
		response.FailWithMessage("保存上传文件失败", c)
		return
	}

	err = s.serviceFile.Get().Upload(c, file)
	if err != nil {
		s.Logger(c).Error("保存上传文件信息失败!", zap.Error(err))
		response.FailWithMessage("保存上传文件信息失败", c)
		return
	}

	var fldList []string

	response.OkWithDetailed(gin.H{"file": file, "flds": fldList}, "上传成功", c)
}

func (s *server) EditFileName(c *gin.Context) {
	var file service_file.ModUpload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceFile.Get().EditFileName(c, file)
	if err != nil {
		s.Logger(c).Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

func (s *server) DeleteFile(c *gin.Context) {
	var file service_file.ModUpload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := s.serviceFile.Get().DeleteFile(c, file); err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *server) GetFileList(c *gin.Context) {
	var pageInfo service_file.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := s.serviceFile.Get().GetFileRecordInfoList(c, pageInfo)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}


func (s *server) srvcUploadFile(c *gin.Context, file *multipart.FileHeader) (string, string, error) {
	ext := path.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	mkdirErr := os.MkdirAll(s.swaConfig.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		s.Logger(c).Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	p := s.swaConfig.Local.StorePath + "/" + filename
	filepath := s.swaConfig.Local.Path + "/" + filename

	f, openError := file.Open()
	if openError != nil {
		s.Logger(c).Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()

	out, createErr := os.Create(p)
	if createErr != nil {
		s.Logger(c).Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()

	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		s.Logger(c).Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}


func (s *server) srvcDeleteFile(key string) error {
	p := s.swaConfig.Local.StorePath + "/" + key
	if strings.Contains(p, s.swaConfig.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
