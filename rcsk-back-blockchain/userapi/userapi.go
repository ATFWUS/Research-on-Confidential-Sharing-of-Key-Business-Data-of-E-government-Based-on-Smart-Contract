package userapi

import (
	"blockchain/arraylist"
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

//连接数据库
//注册用户
func Reguser(username, userpwd string,secretkey1,secretkey2,publickey1x,publickey1y,publickey2x,publickey2y string) {
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}

	//插入数据
	ret, _ := db.Prepare("insert userkey SET username=?, userpwd=?,secretkey1=?,secretkey2=?,publickey1x=?,publickey1y=?,publickey2x=?,publickey2y=?")
	res,_ := ret.Exec(username,userpwd,secretkey1,secretkey2,publickey1x,publickey1y,publickey2x,publickey2y)
	//自增id
	id,_ := res.LastInsertId()
	fmt.Println(id)
}


//登录用户
func UserLogin(username string) string{
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}
	//查询用户
	var userpwd string
	ret,err := db.Query("SELECT userpwd From userkey where username = ?",username)
	defer ret.Close()
	for ret.Next(){
		ret.Scan(&userpwd)
	}
	if err == nil {
		fmt.Println("success")
		fmt.Println(userpwd)
		return userpwd
	}else {
		fmt.Println("error")
		return ""
	}
}




//取出发送者的所有密钥
func GetUserKey(username string) [6]string{
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}
	//获取用户的所有密钥
	var secretkey1 string
	var secretkey2 string
	var publickey1x string
	var publickey1y string
	var publickey2x string
	var publickey2y string
	var key [6]string
	ret,err := db.Query("SELECT secretkey1,secretkey2,publickey1x,publickey1y,publickey2x,publickey2y From userkey where username = ?",username)
	defer ret.Close()
	for ret.Next(){
		ret.Scan(&secretkey1,&secretkey2,&publickey1x,&publickey1y,&publickey2x,&publickey2y)
	}
	key[0] = secretkey1
	key[1] = secretkey2
	key[2] = publickey1x
	key[3] = publickey1y
	key[4] = publickey2x
	key[5] = publickey2y
	if err == nil {
		fmt.Println("success")
		return key
	}else {
		fmt.Println("error")
		key[0] = "1"
		return key
	}
}


func Storage(ct0,ct1,ct2,text,textname,sendname,receivenme string){
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}
	fmt.Println("密文text：%v",text)
	//插入数据
	ret, _ := db.Prepare("insert file SET ct0=?, ct1=?,ct2=?,text=?,textname=?,sendname=?,receivename=?")
	res,_ := ret.Exec(ct0,ct1,ct2,text,textname,sendname,receivenme)
	//自增id
	fmt.Println("",res)
}


//根据接收者的名字去取发送者的名字
func GetSearchName(receivename string) []string{
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}
	//查询发送者
	var name []string
	var sendname string
	ret,err := db.Query("SELECT sendname From file where receivename = ?",receivename)
	defer ret.Close()
	for ret.Next(){
		ret.Scan(&sendname)
		name=append(name,sendname)
	}
	if err == nil {
		fmt.Println("success")
		fmt.Println()
		return name
	}else {
		fmt.Println("error")
		return nil
	}
}


func Search(receivename string) *arraylist.ArrayList {
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}
	//获取对应的的所有索引
	var ct0 string
	var ct1 string
	var ct2 string
	var text string
	var textname string
	var arr [5]string
	list := arraylist.New()
	ret,err := db.Query("SELECT ct0,ct1,ct2,text,textname From file where receivename = ?",receivename)
	defer ret.Close()
	for ret.Next(){
		ret.Scan(&ct0,&ct1,&ct2,&text,&textname)
		arr[0] = ct0
		arr[1] = ct1
		arr[2] = ct2
		arr[3] = text
		arr[4] = textname
 		list.Add(arr)
	}

	if err == nil {
		fmt.Println("success")
		return list
	}else {
		fmt.Println("error")
			return nil
	}

}

//存储密钥的
func KeyStore(name,privkeyx,privkeyy,pubkeyx,pubkeyy,D string){
	var db *sql.DB
	var err error
	db,err = sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success1")
	}

	//插入数据
	ret, _ := db.Prepare("insert keyfile SET username=?, privkeyx=?,privkeyy=?,pubkeyx=?,pubkeyy=?,D=?")
	res,_ := ret.Exec(name,privkeyx,privkeyy,pubkeyx,pubkeyy,D)
	//自增id
	Kid,_ := res.LastInsertId()
	fmt.Println(Kid)
}

//取用户加密用的所有公私钥
func KeyGet(name string) [5]string{
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1)/user")
	defer db.Close()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("success")
	}
	//获取用户的所有密钥
	var privkeyx string
	var privkeyy string
	var pubkeyx string
	var pubkeyy string
	var D string

	var key [5]string
	ret,err := db.Query("SELECT privkeyx,privkeyy,pubkeyx,pubkeyy,D From keyfile where username = ?",name)
	defer ret.Close()
	for ret.Next(){
		ret.Scan(&privkeyx,&privkeyy,&pubkeyx,&pubkeyy,&D)
	}
	key[0] = privkeyx
	key[1] = privkeyy
	key[2] = pubkeyx
	key[3] = pubkeyy
	key[4] = D
	fmt.Println("这个key",key)

	if err == nil {
		fmt.Println("success")
		return key
	}else {
		fmt.Println("error")
		key[0] = "1"
		return key
	}

}
