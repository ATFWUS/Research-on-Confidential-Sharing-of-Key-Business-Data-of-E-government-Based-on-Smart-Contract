package main

import (
	"encoding/json"

	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"

)



type SmartContract struct {
	contractapi.Contract
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PbKey1 string	`json:"pbKey1"`
	PbKey2 string 	`json:"pbkey2"`
	PvKey1 string     `json:"pvkey1"`
	PvKey2 string `json:"pvkey2"`
}

//链码响应消息格式，后端前端交互采用相同消息格式
type BasicResponse struct {
	Code string `json:"code"` //告诉调用方有关链码执行、调用的问题
	Msg string 	`json:"msg"`	//消息的内容
	Data interface{} `json:"data"` //给调用方传递的消息

}

const (
	RECODE_OK = "0" //响应成功
	RECODE_UNKNOWN = "999" //未知错误
	RECODE_USER_ALREADY_EXISTS = "100" //用户已经存在
	RECODE_USER_NOT_EXISTS = "101" //用户不存在
	RECODE_PASSWORD_ERR = "102" //密码错误
	RECODE_AUTH_ERR = "103" //权限不足
	RECODE_USER_ALREDY_REGISTOR ="104" //用户已经注册
	RECODE_GET_FAIL = "105"  //获取内容失败
	RECODE_JSON_ERR = "106" //json传输失败
)

//添加链码内容

//仅限于组织1的人可以注册，权限控制

func (s *SmartContract) Registor(ctx contractapi.TransactionContextInterface,
	name, pass, PbKey1, PbKey2, PvKey1,PvKey2 string) *BasicResponse {

	resp := BasicResponse{RECODE_OK,"OK",nil}

	UserData, err := ctx.GetStub().GetState(name)
	if err != nil {
		resp.Code = RECODE_GET_FAIL
		resp.Msg = err.Error()
		return &resp
	}
	if UserData != nil{
		resp.Code = RECODE_USER_ALREDY_REGISTOR
		resp.Msg = fmt.Sprintf("user(%s) already registor",name)
		return &resp
	}

	//构造注册数据结构，保存
	userinfo := UserInfo{name,pass,PbKey1,PbKey2,PvKey1,PvKey2}
	jsonData,err := json.Marshal(userinfo)
	if err != nil{
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}
	//name作为key，在Fabric框架中都是Key-Value形式
	err = ctx.GetStub().PutState(name,jsonData)
	if err != nil {
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}
	return &resp
}


//登录
func (s *SmartContract) Login(ctx contractapi.TransactionContextInterface,
	name, pass string) *BasicResponse {

	resp := BasicResponse{RECODE_OK,"OK",nil}

	//判断用户能否登录，根据用户名拿到用户相关信息
	UserData, err := ctx.GetStub().GetState(name)
	if err != nil {
		resp.Code = RECODE_GET_FAIL
		resp.Msg = err.Error()
		return &resp
	}
	if UserData == nil{
		resp.Code = RECODE_USER_NOT_EXISTS
		resp.Msg = fmt.Sprintf("user(%s) already registor",name)
		return &resp
	}

	//反序列化成定义的用户数据结构体
	var userinfo UserInfo
	err = json.Unmarshal(UserData, &userinfo)
	if err != nil {
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}
	if userinfo.Password != pass {
		resp.Code = RECODE_PASSWORD_ERR
		resp.Msg = fmt.Sprintf("user(%s) already registor",name)
		return &resp
	}
	resp.Data = userinfo.Username
	return &resp
}

//取秘钥,这里连续调用2次合约，去分别取发送者和接收者的秘钥
func (s *SmartContract) GetKey(ctx contractapi.TransactionContextInterface,
	name string) *BasicResponse {

	resp := BasicResponse{RECODE_OK,"OK",nil}

	//根据用户名拿到用户相关信息
	UserData, err := ctx.GetStub().GetState(name)
	if err != nil {
		resp.Code = RECODE_GET_FAIL
		resp.Msg = err.Error()
		return &resp
	}
	if UserData == nil{
		resp.Code = RECODE_USER_NOT_EXISTS
		resp.Msg = fmt.Sprintf("user(%s) already registor",name)
		return &resp
	}

	//反序列化成定义的用户数据结构体
	var userinfo UserInfo
	err = json.Unmarshal(UserData, &userinfo)
	if err != nil {
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}

	resp.Data = userinfo
	return &resp
}

func main() {
	userChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating token-erc-20 chaincode: %v", err)
	}

	if err := userChaincode.Start(); err != nil {
		log.Panicf("Error starting token-erc-20 chaincode: %v", err)
	}
}