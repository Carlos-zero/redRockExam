package dao

import (
	"docter/util"
)

//tableRank   代表该table所在的层级   authority   权限
func AddFolder(folderRank int, folderName string, authority int, preFolderId int, userId int64) {
	db := util.Conn()

	_, err := db.Exec("insert into folder(folder_rank,folder_name,authority,pre_folder_id,user_id)value (?,?,?,?,?)", folderRank, folderName, authority, preFolderId, userId)
	util.FailOnErr(err, "AddTable err")
}

//通过用户id和包名和包的等级 获取包的id   然后通过包的id获得其下属包
func GetFolderId(userId int, folderName string, folderRank int) int {
	db := util.Conn()

	var folderId int
	err := db.QueryRow("select id from folder where user_id=? and folder_name=? and folder_rank=?", userId, folderName, folderRank).Scan(&folderId)
	util.FailOnErr(err, "获取文件id失败！")
	return folderId
}

//子文件信息结构体
type TableInfo struct{
	FolderRank int
	FolderName string
	Authority int
	Id int
}

//子文件信息
func GetFolderInfos(folderId int) []TableInfo{
	db := util.Conn()

	rows, err := db.Query("select folder_rank,folder_name,authority,id from folder where pre_folder_id=?", folderId)
	util.FailOnErr(err, "获取包信息失败！")

	/*var folderRank int
	var folderName string
	var authority int
	var id int


	var folderRanks []int
	var folderNames []string
	var authoritys []int
	var ids []int*/
	//tableInfoList ：=make([]*TableInfo)
	var tableInfoList []TableInfo
	for rows.Next() {
		var tableInfo TableInfo
		/*_=rows.Scan(folderRank)
		folderRanks= append(folderRanks, folderRank)
		_=rows.Scan(folderName)
		folderNames=append(folderNames,folderName)
		_=rows.Scan(authority)
		authoritys= append(authoritys, authority)
		_=rows.Scan(id)
		ids=append(ids,id)*/
		err:=rows.Scan(&tableInfo.FolderRank,&tableInfo.FolderName,&tableInfo.Authority,&tableInfo.Id)
		util.FailOnErr(err,"写入失败")
		tableInfoList= append(tableInfoList, tableInfo)
	}
	return tableInfoList
	//fmt.Println(tableInfoList[0].FolderRank,tableInfoList[0].Authority,tableInfoList[0].FolderName)
	//fmt.Println(len(tableInfoList),"????")
}
/*func GetFolderInfos(folderId int) {
	db := util.Conn()

	var tableInfoList []*TableInfo

	//rows, err := db.Query("select folder_rank,folder_name,authority,id from folder where pre_folder_id=?", folderId)
	err:=db.Select()
	util.FailOnErr(err, "获取包信息失败！")

	var folderRank int
	var folderName string
	var authority int
	var id int


	var folderRanks []int
	var folderNames []string
	var authoritys []int
	var ids []int
	for rows.Next() {
		_=rows.Scan(folderRank)
		folderRanks= append(folderRanks, folderRank)
		_=rows.Scan(folderName)
		folderNames=append(folderNames,folderName)
		_=rows.Scan(authority)
		authoritys= append(authoritys, authority)
		_=rows.Scan(id)
		ids=append(ids,id)
	}
	fmt.Println(folderRank,folderName,authority,id)
}
*/

//给文件夹改名
func ChangeFolderName(folderName string,folderId int)  {
	db := util.Conn()

	_, err := db.Exec("update folder set folder_name=? where id=?", folderName, folderId)
	util.FailOnErr(err, "改名失败！---ChangeFolderName")
}