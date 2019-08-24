package handler

import (
	"docter/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func UploadFile(ctx *gin.Context)  {

}*/

//将文件上传到服务器     并返回文件的名字和大小  用于入库
func UploadFile(ctx *gin.Context) ([]string,[]int64) {
	//获取解析后表单
	form,_:=ctx.MultipartForm()
	//这里是多文件上传 在之前单文件upload上传的基础上加 [] 变成upload[] 类似文件数组的意思
	files:=form.File["upload"]

	//创建文件名称的数组
	strNum:=len(files)
	fileNames:=make([]string,strNum)
	//创建文件大小的数组
	fileSizeNums:=make([]int64,strNum)

	//循环存文件到服务器本地
	for index,file:=range files{
		//dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		err:=ctx.SaveUploadedFile(file,"C:\\test\\"+file.Filename)
		fileNames[index]=file.Filename
		fileSizeNums[index]=file.Size
		if err!=nil {
			util.FailOnErr(err,"file upload err:")
			ctx.Writer.Write([]byte("上传失败！"))
			//终止操作
		}
	}
	ctx.String(http.StatusOK,fmt.Sprintf("%d 个文件被上传成功!", len(files)))
	return fileNames,fileSizeNums
}