package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"   //官方文件
	"github.com/hyperledger/fabric/protos/peer"  //官方文件
	"github.com/hyperledger/fabric/aberic/chaincode/go/finance/bean" //有
	"github.com/hyperledger/fabric/aberic/chaincode/go/finance/utils" //有
)

type Experience struct {
}

func (t *Experience) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 0 {
		return shim.Error("Parameter error while Init")
	}
	return shim.Success(nil)
}

func (t *Experience) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case "work": // 记录工作
		return work(stub, args)
	default:
		return shim.Error("Unknown func type while Invoke, please check")
	}
}

// 记录贷款数据
func work(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	name, err := utils.GetCreatorName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = bean.Work(stub, args, name)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("记录工作经历成功"))
}

func main() {
	if err := shim.Start(new(Finance)); err != nil {
		fmt.Printf("Chaincode startup error: %s", err)
	}
}