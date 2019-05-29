
package bean

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// 合同全集详情
// 本条记录主键key由成员ID和合同ID联合组成，具备唯一性
type Work struct {
	Timestamp        int64  `json:"timestamp"`        // 本条记录创建时间戳
	Uid              string `json:"uid"`              // 用户唯一ID（32位MD5值）
	Workexperience   string `json:"workexperience"`   // 用户工作经历
	ApplyDate        string `json:"applyDate"`        // 申请日期
	WorkStartDate    string `json:"workStartDate"`    // 工作开始日期
	WorkEndDate      string `json:"workEndDate"`      // 工作终止日期
}

// 贷款操作
// args：UID、工作经历、申请日期、工作开始日期、工作终止日期、简历ID
// name：成员名称
func Work(stub shim.ChaincodeStubInterface, args []string, name string) error {
	if len(args) != 6 {
		return fmt.Errorf("Parameter count error while Work, count must 5")
	}
	if len(args[0]) != 32 {
		return fmt.Errorf("Parameter uid length error while Work, 32 is right")
	}
	if len(args[2]) != 14 {
		return fmt.Errorf("Parameter ApplyDate length error while Work, 14 is right")
	}
	if len(args[3]) != 14 {
		return fmt.Errorf("Parameter WorkStartDate length error while Work, 14 is right")
	}
	if len(args[4]) != 14 {
		return fmt.Errorf("Parameter WorkEndDate length error while Work, 14 is right")
	}
	var work Work
	work.Uid = args[0]
	work.Workexperience = args[1]
	work.ApplyDate = args[2]
	work.WorkStartDate = args[3]
	work.WorkEndDate = args[4]
	work.Timestamp = time.Now().Unix()

	workJsonBytes, err := json.Marshal(&work) // Json序列化
	if err != nil {
		return fmt.Errorf("Json serialize Work fail while Work, work id = " + args[5])
	}
	// 生成联合主键
	key, err := stub.CreateCompositeKey("Work", []string{name, args[5]})
	if err != nil {
		return fmt.Errorf("Failed to CreateCompositeKey while Work")
	}
	// 保存工作经历信息
	err = stub.PutState(key, workJsonBytes)
	if err != nil {
		return fmt.Errorf("Failed to PutState while Work, work id = " + args[5])
	}
	return nil
}