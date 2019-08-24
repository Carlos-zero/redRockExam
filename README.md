#redRockExam

接口：        
"/upload"

参数          
file_table_num   文件所属文件夹的id                  
file_user_id    用户的id           
file_pre_folder_id   上一级文件夹id              


接口：   get     
"/download"        

参数            
file_name   文件名             
file_table_id   文件所在的文件夹id            
file_user_id    文件所属用户的id              

接口：

```
"/loginByName"	通过用户名登录
```

参数：

```
user_name：用户名
password：用户密码
```

接口：

```
"/loginByPhone"		通过手机号登录
```

参数：

```
phone：   手机号
password：	密码
```



接口：

```
"/register"		注册新账号
```

参数：

```
user_name：	用户名
password：   密码
phone：		电话号
```



接口：

```
"/mkdirFolder"	  增加一个空文件夹
```

参数：

```
folder_rank		文件夹所属层级    1为最高层级  一进去就看见    就像我的电脑  的CDEF盘
folder_name		文件名
authority		文件权限  数字1234代表   文件权限越高   别人访问就要有大于等于其对应的权限
pre_folder_id	这个文件的上一层文件的id
user_id			这个文件所属的用户的id
```



接口：

```
sonFolder		通过这个在客户端一层层显示每个文件夹的子文件
```

参数：

```
user_id			用户的id
authority		用户的权限    如果是自己就不需要这个   如果是好友或者陌生人访问则需要
folder_name		文件夹名
folder_rank		文件夹层级
```



接口：

```
"/changeFolderName"		更改文件夹的名称
```

参数：

```
folder_id		文件夹id
folder_name		文件夹的新名字
```

接口：

```
"/addFriend"	添加好友
```

参数：

```
id				自己的id
BName			对方id
```



接口：

```
deleteFriend		删除好友

```

参数：

```
id				自己的id
BName			对方id
```

接口

```
"/giveFriendRole"	给予好友访问权限
```

参数：

```
id		自己的id
BName	对方的id			写快了写错了
role	给予对方的权限等级
```

接口：

```
"/share"		单独分享某个文件
```

参数：

```
file_name		文件名
file_table_id   文件所在的文件夹的id
file_user_id	文件所属用户id
share_password	分享是主人设置密码
```



接口：get

```
"/share/:timemark"		分享文件产生的连接
```

参数：

```
share_password			该连接需要密码才可以下载  密码错误则不能下载
```

其余全是post




