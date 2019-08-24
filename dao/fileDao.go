package dao

import (
	"docter/util"
	"fmt"
	"time"
)

//fileDao   用来向数据库中储存文件信息

type FileData struct {
	FileAddr     string
	FileSize     string
	FileName     string
	FileTableNum string
	FileUserNum  string
}

//对上传的文件进行信息入库
func FileUploadFished(fileAddr string, fileSize int64, fileName string, fileTableId int, fileUserId int, uploadTime time.Time, filePreFolderId int) {
	db := util.Conn()
	//defer db.Close()

	_, err := db.Exec("insert into file_data(file_addr,file_size,file_name,file_folder_id,file_user_id,upload_time,file_pre_folder_id)value (?,?,?,?,?,?,?)", fileAddr, fileSize, fileName, fileTableId, fileUserId, uploadTime, filePreFolderId)

	util.FailOnErr(err, "insert database failure!")
}

//根据用户信息对从数据库中提取文件路径
func GetFileAddr(fileName string, fileTableId int, fileUserId int) string {
	db := util.Conn()

	var fileAddr string

	//rows,err:=db.Query("select file_addr from file_data where file_name=? and file_table_id=? and file_user_id=?",fileName,fileTableId,fileUserId)
	err := db.QueryRow("select file_addr from file_data where file_name=? and file_folder_id=? and file_user_id=?", fileName, fileTableId, fileUserId).Scan(&fileAddr)
	util.FailOnErr(err, "查询数据失败！GetFileAddr")
	/*if rows.NextResultSet() {
		fmt.Println("查询数据为0")
	}else {
		fmt.Println("???????")
	}*/
	/*for rows.Next() {
		err = rows.Scan(&fileAddr)
		util.FailOnErr(err,"写入数据失败！")
	}*/
	return fileAddr
}

//子文件信息结构体
//因为文件里面不能再有文件 所以文件不需要等级   也不需要权限
type FileInfo struct {
	FileName     string
	Id           int
	FileFolderId int
}

//子文件信息
func GetFileInfos(folderId int) []FileInfo {
	db := util.Conn()

	fmt.Println("/////",folderId)
	rows, err := db.Query("select file_name,id,file_folder_id from file_data where file_pre_folder_id=?", folderId)
	util.FailOnErr(err, "获取包信息失败！")

	var fileInfoList []FileInfo
	for rows.Next() {
		var fileInfo FileInfo

		err := rows.Scan(&fileInfo.FileName, &fileInfo.Id, &fileInfo.FileFolderId)
		util.FailOnErr(err, "写入失败")
		fileInfoList = append(fileInfoList, fileInfo)
	}
	return fileInfoList
	//fmt.Println(tableInfoList[0].FolderRank,tableInfoList[0].Authority,tableInfoList[0].FolderName)
	//fmt.Println(len(tableInfoList),"????")
}
