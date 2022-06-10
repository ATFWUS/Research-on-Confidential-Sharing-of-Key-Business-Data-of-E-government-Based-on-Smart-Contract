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

type HashInfo struct {
	HashFile string `json:"hashfile"`
	SigHash string   `json:"sighash"`
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
	RECODE_USER_NOT_EXISTS = "101" //用户不存在
	RECODE_GET_FAIL = "105"  //获取内容失败
	RECODE_JSON_ERR = "106" //json传输失败
	RECODE_HASH_ERR ="1108" //文件完整性验证失败
)

//添加链码内容
//去验证文件哈希值
func (s *SmartContract) GetHashFile(ctx contractapi.TransactionContextInterface,
	HashFile,SigHash string) *BasicResponse {

	resp := BasicResponse{RECODE_OK,"OK",nil}

	//判断，根据云服务器对哈希值的签名拿到相关文件的哈希值
	HashData, err := ctx.GetStub().GetState(SigHash)
	if err != nil {
		resp.Code = RECODE_GET_FAIL
		resp.Msg = err.Error()
		return &resp
	}
	if HashData == nil{
		resp.Code = RECODE_HASH_ERR
		resp.Msg = fmt.Sprintf("HashFile(%s) is Not exit",HashFile)
		return &resp
	}

	//反序列化成定义的文件数据结构体
	var hashinfo HashInfo
	err = json.Unmarshal(HashData, &hashinfo)
	if err != nil {
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}
	if hashinfo.HashFile != HashFile {
		resp.Code = RECODE_HASH_ERR
		resp.Msg = fmt.Sprintf("(%s) de file hashfile (%s) not true and csp is false",SigHash,HashFile)
		return &resp
	}
	resp.Data = hashinfo.HashFile
	return &resp
}

//调用链码，存文件的哈希值
func (s *SmartContract) LocalHashFile(ctx contractapi.TransactionContextInterface,
	HashFile ,SigHash string) *BasicResponse {

	resp := BasicResponse{RECODE_OK,"OK",nil}

	hashinfo := HashInfo{HashFile,SigHash}

	jsonData,err := json.Marshal(hashinfo)
	if err != nil{
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}
	//SigHash作为key，在Fabric框架中都是Key-Value形式
	err = ctx.GetStub().PutState(SigHash,jsonData)
	if err != nil {
		resp.Code = RECODE_JSON_ERR
		resp.Msg = err.Error()
		return &resp
	}
	return &resp
}

func main() {
	HashChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating token-erc-20 chaincode: %v", err)
	}

	if err := HashChaincode.Start(); err != nil {
		log.Panicf("Error starting token-erc-20 chaincode: %v", err)
	}
}