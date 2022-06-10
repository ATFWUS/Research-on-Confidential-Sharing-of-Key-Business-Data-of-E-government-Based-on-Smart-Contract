package routeas

import (
	"blockchain/arraylist"
	"blockchain/sdkapi"
	"blockchain/suanfa"
	_ "blockchain/suanfa"
	"crypto/sha1"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	ginsession "github.com/go-session/gin-session"
	"io"

	"blockchain/userapi"
	"blockchain/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/cloudflare/cfssl/scan/crypto/md5"
	"github.com/gin-gonic/gin"
	_ "github.com/go-session/gin-session"
	"math/big"
	"net/http"
)

//通用响应消息格式
type ResponseMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 注册请求消息格式
type ReqRegsiter struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`

}

// 注册请求消息格式
type ReqLogin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// 通用响应消息处理函数
func responseData(c *gin.Context, resp *ResponseMsg) {
	resp.Msg = utils.GetCodeMsg(resp.Code)
	c.JSON(http.StatusOK, resp)
}

//随机生成在Zq*中的数
func ECrandom(ec elliptic.Curve) *big.Int{

	P := ec.Params().N
	key := [256]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}
	var arr = make([]byte, 256)
	copy(arr[:],key[:])
	arrint := new(big.Int).SetBytes(arr)
	modresult := new(big.Int).Mod(arrint , P)

	return modresult
}


//对字符串进行MD5哈希
func ahash(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行SHA1哈希
func bhash(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//ECCEncrypt 椭圆曲线加密
func ECCEncrypt(plain string, pubKey *ecies.PublicKey) ([]byte, error) {
	src := []byte(plain)
	return ecies.Encrypt(rand.Reader, pubKey, src, nil, nil)
}

//ECCDecrypt 椭圆曲线解密
func ECCDecrypt(cipher []byte, prvKey *ecies.PrivateKey) (string, error) {

	if src, err := prvKey.Decrypt(cipher, nil, nil); err != nil {
		return "", err
	} else {
		fmt.Println("src",src)
		return string(src), nil
	}
}

//注册服务： POST /register {"username":"yekai", "password":"123"}
func Register(c *gin.Context) {
	fmt.Println("begin Register.....")
	//0. 通用响应消息处理
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}
	//延迟执行,响应消息
	defer responseData(c, &resp)
	//1. 解析请求消息
	var reqUser ReqRegsiter
	err := c.ShouldBindJSON(&reqUser)
	if err != nil {
		// 需要有错误处理
		resp.Code = utils.RECODE_BIND_ERR
		return
	}

	fmt.Printf("Register:%+v\n", reqUser)

	//定义椭圆曲线
	ec := elliptic.P256()
	//生成用户的公私钥
	//生成私钥对
	sk1 := ECrandom(ec)
	sk2 := ECrandom(ec)

	//利用私钥对 生成对应的公钥,基点:gx,gy就是生成元p的坐标
	gx := ec.Params().Gx
	gy := ec.Params().Gy

	//第一个人的第一个公钥G*ks1
	skp1x, skp1y := ec.ScalarMult(gx, gy, sk1.Bytes())
	//组装起第一个公钥
	var skp1 ecdsa.PublicKey

	skp1.X = skp1x
	skp1.Y = skp1y

	//第一个人的第二个公钥
	skp11x, skp11y := ec.ScalarMult(gx, gy, sk2.Bytes())
	//组装成它的公钥
	var skp2 ecdsa.PublicKey
	skp2.X = skp11x
	skp2.Y = skp11y

	//将big int数据转为string
	secretkey1 := sk1.String()
	secretkey2 := sk2.String()
	publickey1x := skp1x.String()
	publickey1y := skp1y.String()
	publickey2x := skp11x.String()
	publickey2y := skp11y.String()

	fmt.Println(reqUser.UserName)
	//注册的用户 经过CA认证
	jsonData, err :=sdkapi.Register(reqUser.UserName,reqUser.PassWord)
	if err != nil {
		resp.Code = utils.RECODE_CC_CALL_ERR
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = utils.RECODE_JSON_ERR
		return
	}

	//其公私钥存在用户数据库中
	userapi.Reguser(reqUser.UserName, reqUser.PassWord, secretkey1, secretkey2, publickey1x, publickey1y, publickey2x, publickey2y)

	//存储用户的椭圆曲线加密的公私钥于数据库中
	prvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}
	//转为ecies密钥对

	var eciesPrivateKey *ecies.PrivateKey
	var eciesPublicKey *ecies.PublicKey
	eciesPrivateKey = ecies.ImportECDSA(prvKey)
	eciesPublicKey = &eciesPrivateKey.PublicKey

	//将这对密钥存储在数据库中
	userapi.KeyStore(reqUser.UserName,eciesPrivateKey.X.String(),eciesPrivateKey.Y.String(),
		eciesPublicKey.X.String(),eciesPublicKey.Y.String(),eciesPrivateKey.D.String())
	return

}

//登录服务： POST /login {"username":"wxc", "password":"123"}
func Login(c *gin.Context) {

	var reqUser ReqLogin
	//0. 通用响应消息处理
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}
	//延迟执行
	defer responseData(c, &resp)
	//解析请求消息
	err := c.ShouldBindJSON(&reqUser)
	if err != nil {
		// 需要有错误处理
		resp.Code = utils.RECODE_BIND_ERR
		return
	}
	fmt.Printf("Login:%+v\n", reqUser)

	//
	jsonData, err := sdkapi.Login(reqUser.UserName, reqUser.PassWord)
	//)
	if err != nil {
		resp.Code = utils.RECODE_CC_CALL_ERR
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = utils.RECODE_JSON_ERR
		return
	}
	//3. session处理
	store := ginsession.FromContext(c)
	store.Set("username", reqUser.UserName)
	err = store.Save()
	if err != nil {
		resp.Code = utils.RECODE_SESSION_ERR
		return
	}

}

type SendData struct {
	Sendname    string `json:"sendname"`
	Receivename string `json:"receivename"`
	Keyword  string `json:"keyword"`
	Content  string `json:"content"`
	Filename string `json:"filename"`
}


//定义关于哈希值的结构体
type hashfile struct {
	SigHash string `json:"sighash"`
	Hash    string `json:"hash"`
}

//发送索引、以及文件
func Send(c *gin.Context) {

	fmt.Println("begin Send.....")

	//定义椭圆曲线
	ec := elliptic.P256()

	//定义发送者的第一个私钥
	var sx *big.Int

	//定义接收者的第一个公钥
	var rxp ecdsa.PublicKey

	////定义接收者的第二个公钥
	var ryp ecdsa.PublicKey

	////定义关键词
	//var keyword string

	//0. 通用响应消息处理
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}

	//延迟执行,响应消息
	defer responseData(c, &resp)

	var senddata SendData

	err := c.ShouldBindJSON(&senddata)
	if err != nil {
		// 需要有错误处理
		resp.Code = utils.RECODE_BIND_ERR
		return
	}
	fmt.Printf("send:%+v\n", senddata)
	//根据发送者的名字取出其第一个私钥
	fmt.Println(senddata.Sendname)
	sendkey := userapi.GetUserKey(senddata.Sendname)
	sx,_ = new(big.Int).SetString(sendkey[0],10)

	//根据接收者的名字取出其所有的公钥对
	receivekey := userapi.GetUserKey(senddata.Receivename)
	rxpx := receivekey[2]
	rxpy := receivekey[3]
	rypx := receivekey[4]
	rypy := receivekey[5]

	//组装这个接收者的两个公钥
	rxp.X , _ = new(big.Int).SetString(rxpx,10)
	rxp.Y, _ =  new(big.Int).SetString(rxpy,10)
	ryp.X,_ = new(big.Int).SetString(rypx,10)
	ryp.Y,_ = new(big.Int).SetString(rypy,10)

	//1. 运算索引
	var ct [3] *big.Int
	ct = suanfa.Encryption(ec,sx,rxp,ryp,senddata.Keyword)

	//ct全部转string
	ct0 := ct[0].String()
	ct1 := ct[1].String()
	ct2 := ct[2].String()

	//取出接收者的公钥
	ReceiveEncKey := userapi.KeyGet(senddata.Receivename)
	var pubKey ecies.PublicKey
	pubKey.X,_ = new(big.Int).SetString(ReceiveEncKey[2],10)
	pubKey.Y,_ = new(big.Int).SetString(ReceiveEncKey[3],10)
	pubKey.Curve = elliptic.P256()

	//对Content加密
	cipher, err := ECCEncrypt(senddata.Content, &pubKey)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("密文：%v", cipher)

	//密文转为16进制字符串
	userapi.Storage(ct0,ct1,ct2,fmt.Sprintf("%x", cipher),senddata.Filename,senddata.Sendname,senddata.Receivename)

	var reqhash hashfile

	//对文档进行md5哈希运算
	fmt.Println("senddata.content",senddata.Content)
	htext := ahash(senddata.Content)
	reqhash.Hash = htext
	//对这个哈希值再一次哈希
	htextTohash := bhash(htext)
	reqhash.SigHash = htextTohash
	fmt.Println("regsighash",reqhash.SigHash)

	//哈希值存入区块链
	_, _ = sdkapi.LocalHash(reqhash.Hash,reqhash.SigHash )

}


type IndexInfo struct {
	Ct0 string `json:"ct0"`
	Ct1 string `json:"ct1"`
	Ct2      string `json:"ct2"`
	Id          string `json:"id"`
	Sendname    string `json:"sendname"`
	Receivename string `json:"receivename"`

}


type SearchInFile struct {
	Receivename string `json:"receivename"`
	Keyword string `json:"keyword"`
}

type Result struct {
	Content string `json:"content"`
	Filename string `json:"filename"`
	Hash string `json:"hash"`
	Sendname string `json:"sendname"`
}


//发送陷门，寻找文件
func Receive(c *gin.Context)  {
	fmt.Println("begin Search.....")

	//定义椭圆曲线
	ec := elliptic.P256()

	//定义接收者的第一个私钥
	var rx *big.Int

	//定义发送者的第一个公钥
	var sxp ecdsa.PublicKey

	////定义关键词
	var keyword string

	//定义计数器
	var count = 0

	//0. 通用响应消息处理
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}

	//延迟执行,响应消息
	defer responseData(c, &resp)

	var results Result
	var searchdata SearchInFile

	err := c.ShouldBindJSON(&searchdata)
	if err != nil {
		// 需要有错误处理
		resp.Code = utils.RECODE_BIND_ERR
		return
	}

	//根据接收者的名字取出其第一个私钥
	receivekey := userapi.GetUserKey(searchdata.Receivename)
	rx, _ = new(big.Int).SetString(receivekey[0], 10)

	//根据接收者的名字去遍历发送者的名字
	sendname := userapi.GetSearchName(searchdata.Receivename)

	//定义存储切片
	var arrData  = make([]Result,0,10)
	//开始大遍历:
	for i := 0; i < len(sendname); i++ {
		//根据发送者的名字取出其所有的公钥对

		sendkey := userapi.GetUserKey(sendname[i])

		sxpx := sendkey[2]
		sxpy := sendkey[3]

		//组装这个发送者的两个公钥
		sxp.X, _ = new(big.Int).SetString(sxpx, 10)
		sxp.Y, _ = new(big.Int).SetString(sxpy, 10)

		//计算陷门
		keyword = searchdata.Keyword
		trapdoorx, trapdoory := suanfa.Trapdoor(ec, keyword, rx, sxp)

		//测试方程，寻找对应文件
		var list *arraylist.ArrayList
		list = userapi.Search(searchdata.Receivename)
		fmt.Printf("list:",list)

		//这里进行测试，通过下面方法取出list中想要的值
		for i := 0; i < list.Size(); i++ {
			//var indexinfo IndexInfo
			count++

			var arr [5]string
			arr = (list.Get(i)).([5]string)

			ct0 := arr[0]
			ct1 := arr[1]
			ct2 := arr[2]

			//智能合约匹配
			jsonData, err := sdkapi.TestAA(ct2, trapdoorx.String(), trapdoory.String(), ct0, ct1 )

			if err != nil {
				resp.Code = utils.RECODE_CC_CALL_ERR
				return
			}
			err = json.Unmarshal(jsonData, &resp)
			if err != nil {
				resp.Code = utils.RECODE_JSON_ERR
				return
			}

			//测试成功
			if resp.Msg == "OK" {
				results.Filename = arr[4]
				//打印下内容
				fmt.Println("arr[3]",arr[3])
				//对content解密
				var prvKey ecies.PrivateKey
				//取其私钥
				var keyreceive [5]string
				keyreceive = userapi.KeyGet(searchdata.Receivename)
				//fmt.Println("这个数组",keyreceive)
				prvKey.X,_ = new(big.Int).SetString(keyreceive[0],10)
				prvKey.Y,_ = new(big.Int).SetString(keyreceive[1],10)
				//fmt.Println("D",keyreceive[4])
				prvKey.D ,_= new(big.Int).SetString(keyreceive[4],10)

				prvKey.Curve = elliptic.P256()
				//content内容以数组呈现
				contentofarr,_:=hex.DecodeString(arr[3])

				plain, err := ECCDecrypt(contentofarr, &prvKey)
				if err != nil {
					fmt.Println(err)
				}

				results.Content = plain
				results.Hash = ahash(results.Content)
				results.Sendname = sendname[i]

				arrData = append(arrData,results)
				resp.Data = arrData

				fmt.Println("zhaodao")
				fmt.Println(resp.Data)

			}

		}
	}
	return
}

//发送陷门，寻找文件
func Receivee(c *gin.Context)  {
	fmt.Println("begin Search2.....")

	//定义椭圆曲线
	ec := elliptic.P256()

	//定义接收者的第一个私钥
	var rx *big.Int

	//定义发送者的第一个公钥
	var sxp ecdsa.PublicKey

	////定义关键词
	var keyword string

	//定义计数器
	var count = 0

	//0. 通用响应消息处理
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}

	//延迟执行,响应消息
	defer responseData(c, &resp)

	var results Result
	var searchdata SearchInFile

	err := c.ShouldBindJSON(&searchdata)
	if err != nil {
		// 需要有错误处理
		resp.Code = utils.RECODE_BIND_ERR
		return
	}
	fmt.Printf("receive:%+v\n", searchdata)

	//根据接收者的名字取出其第一个私钥
	receivekey := userapi.GetUserKey(searchdata.Receivename)
	rx, _ = new(big.Int).SetString(receivekey[0], 10)

	//根据接收者的名字去遍历发送者的名字
	sendname := userapi.GetSearchName(searchdata.Receivename)

	//开始大遍历:
	for i := 0; i < len(sendname); i++ {
		//根据发送者的名字取出其所有的公钥对
		fmt.Println(sendname[i])
		sendkey := userapi.GetUserKey(sendname[i])

		sxpx := sendkey[2]
		sxpy := sendkey[3]

		//组装这个接收者的两个公钥
		sxp.X, _ = new(big.Int).SetString(sxpx, 10)
		sxp.Y, _ = new(big.Int).SetString(sxpy, 10)

		//1. 运算陷门
		//计算陷门
		keyword = searchdata.Keyword
		trapdoorx, trapdoory := suanfa.Trapdoor(ec, keyword, rx, sxp)

		//测试方程，寻找对应文件

		var list *arraylist.ArrayList

		list = userapi.Search(searchdata.Receivename)

		//这里进行测试，通过下面方法取出list中想要的值
		for i := 0; i < list.Size(); i++ {

			count++

			var arr [5]string
			arr = (list.Get(i)).([5]string)

			ct0 := arr[0]
			ct1 := arr[1]
			ct2 := arr[2]

			jsonData, err := sdkapi.TestAA(ct2, trapdoorx.String(), trapdoory.String(), ct0, ct1 )

			if err != nil {
				resp.Code = utils.RECODE_CC_CALL_ERR
				return
			}
			err = json.Unmarshal(jsonData, &resp)
			if err != nil {
				resp.Code = utils.RECODE_JSON_ERR
				return
			}

			//测试成功
			if resp.Msg == "OK" {
				results.Filename = arr[4]
				results.Content = arr[3]
				results.Hash = ahash(results.Content)
				results.Sendname = sendname[i]

				var arrData  = make([]Result,0,1)
				arrData = append(arrData,results)
				resp.Data = arrData
				fmt.Println("text", results.Content)
				fmt.Println("textname", arr[4])
				fmt.Println("zhaodao")
				return
			}
			if count >= list.Size() {
				fmt.Println("mei zhao dao")
				return
			}

		}
	}
}


type WanShu struct{
	Content string `json:"contenta"`
}

//验证文件的完整性的哈希值
func GetHashFile(c *gin.Context) {
	//0. 通用响应消息处理
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}
	//延迟执行
	defer responseData(c, &resp)

	//1. 解析请求消息
	var reqhash hashfile

	//2. 与区块链交互
	//对文档进行sha256哈希运算

	var data WanShu

	err := c.ShouldBindJSON(&data)
	fmt.Println("data",data)
	if err != nil {
		// 需要有错误处理
		resp.Code = utils.RECODE_BIND_ERR
		return
	}
	fmt.Printf("aa:",data.Content)
	htext := ahash(data.Content)
	reqhash.Hash = htext
	fmt.Println("htext:",htext)

	//对这个哈希值再一次哈希
	htextTohash := bhash(htext)
	reqhash.SigHash = htextTohash
	fmt.Println("regsighash",reqhash.SigHash)

	jsonData, err := sdkapi.GetHash(reqhash.Hash, reqhash.SigHash)
	if err != nil {
		resp.Code = utils.RECODE_CC_CALL_ERR
		return
	}
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		resp.Code = utils.RECODE_JSON_ERR
		return
	}

}

// 退出
func Logout(c *gin.Context) {
	// 响应消息
	resp := ResponseMsg{
		Code: utils.RECODE_OK,
		Data: nil,
	}
	//延迟执行
	defer responseData(c, &resp)
	// 删除session
	err := ginsession.Destroy(c)
	if err != nil {
		fmt.Println("failed to logout destory session", err)
		resp.Code = utils.RECODE_SESSION_ERR
		return
	}
	return
}










