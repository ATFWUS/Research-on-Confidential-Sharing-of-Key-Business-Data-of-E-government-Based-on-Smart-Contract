package main

import (
	"crypto/elliptic"
	"crypto/sha256"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
	"math/big"
)

type SmartContract struct {
	contractapi.Contract
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
//hash到zq*上
func hashInZq(b []byte) *big.Int{
	N,_ := new(big.Int).SetString("115792089210356248762697446949407573529996955224135760342422259061068512044369", 10)
	h := sha256.Sum256(b)
	var arr  = make([]byte, sha256.Size)
	copy(arr[:],h[:])
	a := new(big.Int).SetBytes(arr)

	return new(big.Int).Mod(a,N)
}



func (s *SmartContract) Test(ctx contractapi.TransactionContextInterface,
	ct2, trapdoorx, trapdoory, ct0, ct1 string) *BasicResponse {

	ct00, _ := new(big.Int).SetString(ct0, 10)
	ct11, _ := new(big.Int).SetString(ct1, 10)
	trapdoorxx ,_:= new(big.Int).SetString(trapdoorx, 10)
	trapdooryy ,_:= new(big.Int).SetString(trapdoory, 10)
	//定义椭圆曲线
	ec := elliptic.P256()
	x, y := ec.Add(trapdoorxx, trapdooryy, ct00, ct11)

	a := x.Bytes()
	b := y.Bytes()
	a = append(a, b...)

	v := hashInZq(a)

	resp := BasicResponse{RECODE_OK, "OK", nil}
	if v.String() == ct2 {
		return &resp
	} else {
		resp.Msg = "error"
		return &resp
	}
}

func main() {
	TestChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating token-erc-20 chaincode: %v", err)
	}

	if err := TestChaincode.Start(); err != nil {
		log.Panicf("Error starting token-erc-20 chaincode: %v", err)
	}
}