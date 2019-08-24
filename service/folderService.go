package service

import (
	"docter/dao"
	"docter/handler"
	"docter/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建新的文件夹
func MkdirFolderService(ctx *gin.Context) {
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}

	folderRank := util.StringToInt(ctx.PostForm("folder_rank"))
	folderName := ctx.PostForm("folder_name")
	authority := util.StringToInt(ctx.PostForm("authority"))
	preFolderId := util.StringToInt(ctx.PostForm("pre_folder_id"))
	userId := util.StringToInt(ctx.PostForm("user_id"))
	dao.AddFolder(folderRank, folderName, authority, preFolderId, int64(userId))
}

//没测
//通过用户  id   包名  包等级   获得其下的包信息
func GetFolderInfosService(ctx *gin.Context) {
	//对用户进行鉴权
	//cus := &handler.CustomerClaim{}
	token := handler.GetToken(ctx.Request)
	cus,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}

	//token是自己的ID  是B
	userId1 := util.StringToInt(cus.Id)
	/*userName:=cus.UserName
	phone:=cus.Phone*/

	//直接拿的是那个文件主人的id   是A
	thisUserId := util.StringToInt(ctx.PostForm("user_id"))
	//文件所需权限   应该是直接获取----
	authority := util.StringToInt(ctx.PostForm("authority"))
	//通过ID查看自己的权限
	role := dao.QueryRole(thisUserId, userId1)

	fmt.Println("this id=", thisUserId, "user id=", userId1)
	fmt.Println("authority=",authority,"role=",role)
	//鉴权结果
	if thisUserId == userId1 {
		userId := util.StringToInt(ctx.PostForm("user_id"))
		folderName := ctx.PostForm("folder_name")
		folderRank := util.StringToInt(ctx.PostForm("folder_rank"))
		folderId := dao.GetFolderId(userId, folderName, folderRank)
		//获得包id后根据包id获得子节点
		tableInfoList := dao.GetFolderInfos(folderId)
		fileInfoList := dao.GetFileInfos(folderId)
		for i := 0; i < len(tableInfoList); i++ {
			ctx.JSON(http.StatusOK, gin.H{
				"folder_rank": tableInfoList[i].FolderRank,
				"folder_name": tableInfoList[i].FolderName,
				"authority":   tableInfoList[i].Authority,
				"id":          tableInfoList[i].Id,
			})
		}
		for i := 0; i < len(fileInfoList); i++ {
			ctx.JSON(http.StatusOK, gin.H{
				"file_name":      fileInfoList[i].FileName,
				"id":             fileInfoList[i].Id,
				"file_folder_id": fileInfoList[i].FileFolderId,
			})
		}
	} else if role >= authority {
		userId := util.StringToInt(ctx.PostForm("user_id"))
		folderName := ctx.PostForm("folder_name")
		folderRank := util.StringToInt(ctx.PostForm("folder_rank"))
		folderId := dao.GetFolderId(userId, folderName, folderRank)
		//获得包id后根据包id获得子节点
		tableInfoList := dao.GetFolderInfos(folderId)
		fileInfoList := dao.GetFileInfos(folderId)
		for i := 0; i < len(tableInfoList); i++ {
			ctx.JSON(http.StatusOK, gin.H{
				"folder_rank": tableInfoList[i].FolderRank,
				"folder_name": tableInfoList[i].FolderName,
				"authority":   tableInfoList[i].Authority,
				"id":          tableInfoList[i].Id,
			})
		}
		for i := 0; i < len(fileInfoList); i++ {
			ctx.JSON(http.StatusOK, gin.H{
				"file_name":      fileInfoList[i].FileName,
				"id":             fileInfoList[i].Id,
				"file_folder_id": fileInfoList[i].FileFolderId,
			})
		}
	} else {
		ctx.Writer.Write([]byte("您的权限不足。"))
	}

}

//通过文件id更改文件名
func ChangeFolderNameService(ctx *gin.Context) {
	//token验证
	token := handler.GetToken(ctx.Request)
	_,err:= handler.Decode(token)
	if err!=nil {
		ctx.Writer.Write([]byte("您的身份已过期或身份错误，请重新登录！"))
		return
	}
	folderId := util.StringToInt(ctx.PostForm("folder_id"))
	folderName := ctx.PostForm("folder_name")
	dao.ChangeFolderName(folderName, folderId)
}
