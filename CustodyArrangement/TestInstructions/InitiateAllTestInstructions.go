package TestInstructions

func InitiateAllTestInstructionsForCA() AllTestInstructions_CA_TestCaseSetUpStruct {

	var allTestInstructions_CA_TestCaseSetUp AllTestInstructions_CA_TestCaseSetUpStruct

	// Generate data for TestInstructions within 'Custody Arrangement'
	allTestInstructionsForGeneralSetupTearDown := InitiateTestInstructionsFor_GeneralSeupTearDown()
	allTestInstructionsForCustodyAccount := InitiateTestInstructionsFor_CustodyAccount()
	allTestInstructionsForSettlementAgreement := InitiateTestInstructionsFor_SettlementAgreement()

	// Append 'GeneralSetupTearDown' TestInstructions into one structure
	allTestInstructions_CA_TestCaseSetUp.TestInstruction = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstruction,
		allTestInstructionsForGeneralSetupTearDown.TestInstruction...)
	allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation,
		allTestInstructionsForGeneralSetupTearDown.BasicTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation,
		allTestInstructionsForGeneralSetupTearDown.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute,
		allTestInstructionsForGeneralSetupTearDown.TestInstructionAttribute...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel,
		allTestInstructionsForGeneralSetupTearDown.ImmatureElementModel...)
	allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes,
		allTestInstructionsForGeneralSetupTearDown.FangEngineClassesMethodsAttributes...)

	// Append 'CustodyAccount' TestInstructions into one structure
	allTestInstructions_CA_TestCaseSetUp.TestInstruction = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstruction,
		allTestInstructionsForCustodyAccount.TestInstruction...)
	allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation,
		allTestInstructionsForCustodyAccount.BasicTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation,
		allTestInstructionsForCustodyAccount.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute,
		allTestInstructionsForCustodyAccount.TestInstructionAttribute...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel,
		allTestInstructionsForCustodyAccount.ImmatureElementModel...)
	allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes,
		allTestInstructionsForCustodyAccount.FangEngineClassesMethodsAttributes...)

	// Append 'SettlementAgreement' TestInstructions into one structure
	allTestInstructions_CA_TestCaseSetUp.TestInstruction = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstruction,
		allTestInstructionsForSettlementAgreement.TestInstruction...)
	allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.BasicTestInstructionInformation,
		allTestInstructionsForSettlementAgreement.BasicTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureTestInstructionInformation,
		allTestInstructionsForSettlementAgreement.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute = append(
		allTestInstructions_CA_TestCaseSetUp.TestInstructionAttribute,
		allTestInstructionsForSettlementAgreement.TestInstructionAttribute...)
	allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel = append(
		allTestInstructions_CA_TestCaseSetUp.ImmatureElementModel,
		allTestInstructionsForSettlementAgreement.ImmatureElementModel...)
	allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes,
		allTestInstructionsForSettlementAgreement.FangEngineClassesMethodsAttributes...)

	return allTestInstructions_CA_TestCaseSetUp

}
