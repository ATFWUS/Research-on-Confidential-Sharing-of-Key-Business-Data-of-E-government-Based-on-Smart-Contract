package sdkapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var  hashContract *gateway.Contract
var testContract *gateway.Contract
var indexContract *gateway.Contract
var searchindexContract *gateway.Contract

var reglogContract *gateway.Contract

//测试2合约
var testAAContract *gateway.Contract


func init() {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "false")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("wxcchannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}


	contract := network.GetContract("blocks_a")
	hashContract = contract


	contract = network.GetContract("testa")
	testContract = contract

	contract = network.GetContract("index")
	indexContract = contract

	contract = network.GetContract("searchindex")
	searchindexContract = contract

	contract = network.GetContract("regloguser")
	reglogContract = contract

	contract = network.GetContract("testAA")
	testAAContract = contract

}

func TestAA (ct2, trapdoorx, trapdoory, ct0, ct1 string)([]byte, error){
	fmt.Println("正在测试合约")
	return testAAContract.EvaluateTransaction("Test",ct2, trapdoorx, trapdoory, ct0, ct1)

}



func LocalHash(HashFile,SigHash string) ([]byte, error){
	fmt.Println("local",HashFile)
	return hashContract.SubmitTransaction("LocalHashFile",HashFile,SigHash)
}


func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}

//CA认证
func Register(username,password string)([]byte, error) {

	return reglogContract.SubmitTransaction("Registor",username,password)
}

//CA认证登录
func Login(username, password string) ([]byte, error) {
	return reglogContract.EvaluateTransaction("Login", username, password)
}


////调用验证文件完整性的合约
func GetHash(HashFile,SigHash string)([]byte, error){
	return hashContract.EvaluateTransaction("GetHashFile",HashFile,SigHash)
}



//智能合约测试
func Test(a,ct1 string)([]byte,error){

	return testContract.EvaluateTransaction("Test",a,ct1)
}










