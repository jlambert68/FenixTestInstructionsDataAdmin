package shared_code

import (
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
)

// SendGrpcTestInstructionsAndTestInstructionContainersAndAllowedUsersMessageToWorker
func SendGrpcTestInstructionsAndTestInstructionContainersAndAllowedUsersMessageToWorker(
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage *fenixExecutionWorkerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage) (
	err error) {

	fmt.Println(supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage)

	return err
}
