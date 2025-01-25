package image_api

import (
	"fast_gin/global"
	"fast_gin/utils/Random"
	"fast_gin/utils/find"
	"fast_gin/utils/md5"
	"fast_gin/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func (ImageApi) UploadView(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		res.FailWithMsg("请选择文件！", c)
		return
	}
	// 判断文件大小
	if file.Size > global.Config.Upload.MaxSize*1024*1024 {
		res.FailWithMsg("文件大小不能超过"+strconv.FormatInt(global.Config.Upload.MaxSize, 10)+"MB", c)
		return
	}
	// 白名单
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !find.InList(global.Config.Upload.AllowExt, ext) {
		res.FailWithMsg("文件格式不正确", c)
		return
	}

	// 处理文件名重复
	fp := path.Join("uploads", global.Config.Upload.Dir, file.Filename)

	_, err1 := os.Stat(fp)
	if !os.IsExist(err1) {
		// 文件存在
		// 算上传的图片与本身的图片相同，如果一样，则不保存，返回之前的地址
		uploadFile, _ := file.Open()
		oldFile, _ := os.Open(fp)
		uoloadFileHash := md5.MD5WithFile(uploadFile)
		oldFileHash := md5.MD5WithFile(oldFile)
		if uoloadFileHash == oldFileHash {
			res.Ok("/"+fp, "上传成功！", c)
			return
		}
		// 上传的图片名称一样，内容不同，则生成新的文件名
		fileNameNotExt := strings.TrimSuffix(file.Filename, ext)
		newFileName := fmt.Sprintf("%s_%s%s", fileNameNotExt, Random.RanDomString(3), ext)
		fp = path.Join("uploads", global.Config.Upload.Dir, newFileName)
	}
	c.SaveUploadedFile(file, fp)
	res.Ok("/"+fp, "上传成功！", c)

}
