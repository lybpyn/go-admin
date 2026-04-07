package apis

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/utils"
	"github.com/google/uuid"

	"go-admin/common/file_store"
	"go-admin/config"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

// 默认上传路径
const defaultPath = "static/uploadfile/"

type File struct {
	api.Api
}

// ensureDirExists 确保目录存在，如果不存在则创建，权限设置为 0777 避免权限问题
func ensureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 使用 0777 权限创建目录
		if err := os.MkdirAll(path, 0777); err != nil {
			return err
		}
		// 显式设置目录权限为 0777，避免 umask 影响
		return os.Chmod(path, 0777)
	}
	return nil
}

// getUploadPath 获取上传路径，优先使用配置，否则使用默认路径
func getUploadPath() string {
	if config.ExtConfig.Upload.Path != "" {
		path := config.ExtConfig.Upload.Path
		// 确保路径以 / 结尾
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		return path
	}
	return defaultPath
}

// getDateUploadPath 获取带日期的上传路径，返回本地存储路径和相对URL路径
// 本地路径：/home/ll/release/uploads/static/images/2025/11/10/
// URL路径：/static/images/2025/11/10/
func getDateUploadPath() (localPath string, urlPath string) {
	basePath := config.ExtConfig.Upload.Path
	if basePath == "" {
		basePath = "."
	}
	// 确保路径以 / 结尾
	if !strings.HasSuffix(basePath, "/") {
		basePath = basePath + "/"
	}

	// 添加static/images和日期目录
	now := time.Now()
	datePath := fmt.Sprintf("images/%d/%02d/%02d/", now.Year(), now.Month(), now.Day())

	localPath = basePath + datePath
	urlPath = "/" + datePath
	return
}

// getFileURL 获取完整的文件访问URL
func getFileURL(relativePath string) string {
	domain := config.ExtConfig.Upload.Domain
	if domain == "" {
		return relativePath
	}
	// 确保domain不以/结尾
	domain = strings.TrimSuffix(domain, "/")
	// 确保relativePath以/开头
	if !strings.HasPrefix(relativePath, "/") {
		relativePath = "/" + relativePath
	}
	return domain + relativePath
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadFile [post]
// @Security Bearer
func (e File) UploadFile(c *gin.Context) {
	e.MakeContext(c)
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("%s://%s/", "http", c.Request.Host)

	switch tag {
	case "1": // 单图
		e.handleSingleFile(c, urlPrefix)
	case "2": // 多图
		e.handleMultipleFiles(c, urlPrefix)
	case "3": // base64
		e.handleBase64File(c, urlPrefix)
	default:
		e.handleSingleFile(c, urlPrefix)
	}
}

func (e File) handleSingleFile(c *gin.Context, urlPrefix string) {
	fileResponse, done := e.singleFile(c, FileResponse{}, urlPrefix)
	if done {
		return
	}
	e.OK(fileResponse, "上传成功")
}

func (e File) handleMultipleFiles(c *gin.Context, urlPrefix string) {
	multipartFile := e.multipleFile(c, urlPrefix)
	e.OK(multipartFile, "上传成功")
}

func (e File) handleBase64File(c *gin.Context, urlPrefix string) {
	fileResponse := e.baseImg(c, FileResponse{}, urlPrefix)
	e.OK(fileResponse, "上传成功")
}

func (e File) baseImg(c *gin.Context, fileResponse FileResponse, urlPrefix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	decodedData, _ := base64.StdEncoding.DecodeString(file2list[1])
	fileName := uuid.New().String() + ".jpg"

	uploadPath := getUploadPath()
	if err := ensureDirExists(uploadPath); err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
		return fileResponse
	}

	base64File := uploadPath + fileName
	_ = ioutil.WriteFile(base64File, decodedData, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)

	fileResponse = e.buildFileResponse(base64File, urlPrefix, "", typeStr)
	source, _ := c.GetPostForm("source")

	if err := thirdUpload(source, fileName, base64File); err != nil {
		e.Error(200, errors.New(""), "上传第三方失败")
		return fileResponse
	}

	if source != "1" {
		fileResponse.Path = "/" + uploadPath + fileName
		fileResponse.FullPath = "/" + uploadPath + fileName
	}
	return fileResponse
}

func (e File) multipleFile(c *gin.Context, urlPrefix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	source, _ := c.GetPostForm("source")
	var multipartFile []FileResponse

	uploadPath := getUploadPath()
	for _, f := range files {
		fileName := uuid.New().String() + utils.GetExt(f.Filename)

		if err := ensureDirExists(uploadPath); err != nil {
			e.Error(500, errors.New(""), "初始化文件路径失败")
			continue
		}

		multipartFileName := uploadPath + fileName
		if err := c.SaveUploadedFile(f, multipartFileName); err != nil {
			continue
		}

		fileType, _ := utils.GetType(multipartFileName)
		if err := thirdUpload(source, fileName, multipartFileName); err != nil {
			e.Error(500, errors.New(""), "上传第三方失败")
			continue
		}

		fileResponse := e.buildFileResponse(multipartFileName, urlPrefix, f.Filename, fileType)
		if source != "1" {
			fileResponse.Path = "/" + uploadPath + fileName
			fileResponse.FullPath = "/" + uploadPath + fileName
		}
		multipartFile = append(multipartFile, fileResponse)
	}
	return multipartFile
}

func (e File) singleFile(c *gin.Context, fileResponse FileResponse, urlPrefix string) (FileResponse, bool) {
	files, err := c.FormFile("file")
	if err != nil {
		e.Error(200, errors.New(""), "图片不能为空")
		return FileResponse{}, true
	}

	// 直接使用配置路径，不创建日期子目录
	uploadPath := getUploadPath()
	fileName := uuid.New().String() + utils.GetExt(files.Filename)

	// 保存文件到本地
	localFile := uploadPath + fileName
	if err := c.SaveUploadedFile(files, localFile); err != nil {
		e.Error(500, errors.New(err.Error()), fmt.Sprint("文件保存失败-%s", err.Error()))
		return FileResponse{}, true
	}

	// 构建返回的URL
	fileType, _ := utils.GetType(localFile)
	relativePath := "static/" + fileName
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(localFile),
		Path:     relativePath,
		FullPath: getFileURL(relativePath),
		Name:     files.Filename,
		Type:     fileType,
	}
	return fileResponse, false
}

func (e File) buildFileResponse(filePath, urlPrefix, fileName, fileType string) FileResponse {
	return FileResponse{
		Size:     pkg.GetFileSize(filePath),
		Path:     filePath,
		FullPath: urlPrefix + filePath,
		Name:     fileName,
		Type:     fileType,
	}
}

func thirdUpload(source string, name string, path string) error {
	switch source {
	case "2":
		return ossUpload("img/"+name, path)
	case "3":
		return qiniuUpload("img/"+name, path)
	}
	return nil
}

func ossUpload(name string, path string) error {
	oss := file_store.ALiYunOSS{}
	return oss.UpLoad(name, path)
}

func qiniuUpload(name string, path string) error {
	oss := file_store.ALiYunOSS{}
	return oss.UpLoad(name, path)
}
