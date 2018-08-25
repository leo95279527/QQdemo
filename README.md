## 做一个简单的类似QQ聊天的工具，带注册、登录、拉取房间列表、选择房间、进入房间、群聊、一对一聊天；  
#### 版本更新日志  
- 2018.8.24：完成客户端与服务器端的连接，完成协议的制定；目前尝试用Go语言完成使用协议的逻辑和对MySQL的操作。  
- 2018.8.24.23点晚上：完成Go语言操作MySQL数据库的初始化连接、插入数据和查找数据的逻辑，并在编译时碰到两个bug:  
Server\DBServ\InsertDB.go:6:12: undefined: db  
Server\DBServ\InsertDB.go:10:3: undefined: checkErr  
明天继续coding。  
- 2018.8.25下午：完成了使用Go语言操作MySQL数据库的删除数据和更新的数据的逻辑，在编码过程中感到技术知识的不足，先列举如下  
1，函数定义、调用和传参；  
2，数据库：SQL语句上的条件约束和联合查询；  
3，Protobuf协议的运用；  
4，对指针的深入理解；  
5，计算机网络：Socket编程，TCP机制，分包粘包；  
6，操作系统：加锁的概念，为什么多线程会不安全；  
- 2018.8.25晚：添加了定义协议号的Server/userproto/userhandle.go文件；写了使用protobuf完成用户登录的逻辑；接下来应完成  
1，在Server/proto/message/Login.proto中的Login_ToC添加msg字段，并运用到登录逻辑中去，并加锁；   
2，定义存储客户端连接服务端用的哈希表；  
3，写好协议收发的逻辑；  