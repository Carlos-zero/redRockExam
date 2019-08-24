package controller

import (

	"docter/service"
	"github.com/gin-gonic/gin"
)

func Gogogo()  {
	router:=gin.Default()
	router.POST("/upload", func(ctx *gin.Context) {

		service.UploadService(ctx)
	})

	router.GET("/download", func(ctx *gin.Context) {
		service.DownloadService(ctx)
	})

	//--ok
	router.POST("/loginByName", func(ctx *gin.Context) {
		service.LoginServiceByName(ctx)
	})
	//--ok
	router.POST("/loginByPhone", func(ctx *gin.Context) {
		service.LoginServiceByPhone(ctx)
	})

	//---ok
	router.POST("/register", func(ctx *gin.Context) {
		service.Register(ctx)
	})

	//创建文件夹的接口---ok
	router.POST("/mkdirFolder", func(ctx *gin.Context) {
		service.MkdirFolderService(ctx)
	})

	//展示子文件夹的接口--ok
	router.POST("/sonFolder", func(ctx *gin.Context) {
		service.GetFolderInfosService(ctx)

	})

	router.POST("/changeFolderName", func(ctx *gin.Context) {
		service.ChangeFolderNameService(ctx)
	})

	//添加好友
	router.POST("/addFriend", func(ctx *gin.Context) {
		service.AddFriendService(ctx)
	})
	//删除好友
	router.POST("/deleteFriend", func(ctx *gin.Context) {
		service.DeleteFriendService(ctx)
	})
	//给予好友权限
	router.POST("/giveFriendRole", func(ctx *gin.Context) {
		service.GiveRoleService(ctx)
	})

	//生成分享连接
	router.POST("/share", func(ctx *gin.Context) {
		service.ShareFileService(ctx)
	})
	//验证
	router.GET("/share/:timemark", func(ctx *gin.Context) {
		service.AuthShareService(ctx)
	})


	router.Run(":8080")


}
