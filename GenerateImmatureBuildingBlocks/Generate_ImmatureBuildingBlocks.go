package GenerateImmatureBuildingBlocks

import fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"

// GenerateImmatureBuildingBlocks
// Generates the full set of available building blocks for - Sub Custody -
func GenerateImmatureBuildingBlocks() (
	responseMessage *fenixTestCaseBuilderServerGrpcApi.
		AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage) {

	// Generate ImmatureTestInstructionsMessages
	var immatureTestInstructions []*fenixTestCaseBuilderServerGrpcApi.ImmatureTestInstructionMessage
	immatureTestInstructions = generateImmatureTestInstructionsMessagesSubCustody()

	// Generate ImmatureTestInstructionContainers
	var immatureTestInstructionContainers []*fenixTestCaseBuilderServerGrpcApi.ImmatureTestInstructionContainerMessage
	immatureTestInstructionContainers = generateImmatureTestInstructionContainersMessagesSubCustody()

	// Generate the full gRPC response for available TestInstructions and TestInstructionsContainers
	responseMessage = &fenixTestCaseBuilderServerGrpcApi.
		AvailableTestInstructionsAndPreCreatedTestInstructionContainersResponseMessage{
		ImmatureTestInstructions:          immatureTestInstructions,
		ImmatureTestInstructionContainers: immatureTestInstructionContainers,
		AckNackResponse: &fenixTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:                      false,
			Comments:                     "",
			ErrorCodes:                   nil,
			ProtoFileVersionUsedByClient: 0,
		},
	}

	return responseMessage

}
