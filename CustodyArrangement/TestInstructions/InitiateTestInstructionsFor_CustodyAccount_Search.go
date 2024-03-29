package TestInstructions

func InitiateTestInstructionsFor_CustodyAccount() AllTestInstructions_CA_TestCaseSetUpStruct {

	var allTestInstructions_CA_TestCaseSetUp AllTestInstructions_CA_TestCaseSetUpStruct

	// Generate data for TestInstructionsMap within 'CustodyAccount'
	testInstruction_CA_CustodyAccount_Search := Initate_TestInstruction_CA_CustodyAccount_Search()

	// Append 'TestCaseSetUp' TestInstructionsMap into one structure
	allTestInstructions_CA_TestCaseSetUp.TestInstruction = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstruction,
		testInstruction_CA_CustodyAccount_Search.TestInstruction)
	allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation,
		testInstruction_CA_CustodyAccount_Search.BasicTestInstructionInformation)
	allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation,
		testInstruction_CA_CustodyAccount_Search.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute,
		testInstruction_CA_CustodyAccount_Search.TestInstructionAttribute...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel,
		testInstruction_CA_CustodyAccount_Search.ImmatureElementModel...)
	allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes,
		testInstruction_CA_CustodyAccount_Search.FangEngineClassesMethodsAttributes)

	return allTestInstructions_CA_TestCaseSetUp

}
