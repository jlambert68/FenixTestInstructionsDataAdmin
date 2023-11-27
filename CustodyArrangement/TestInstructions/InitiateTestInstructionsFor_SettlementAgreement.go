package TestInstructions

func InitiateTestInstructionsFor_SettlementAgreement() AllTestInstructions_CA_TestCaseSetUpStruct {

	var allTestInstructions_CA_SettlementAgreement AllTestInstructions_CA_TestCaseSetUpStruct

	// Generate data for TestInstructionsMap within 'CustodyAccount'
	testInstruction_CA_SettlementAgreement_Edit := Initate_TestInstruction_CA_SettlementAgreement_Edit()
	testInstruction_CA_SettlementAgreement_AddSelectedSwift := Initate_TestInstruction_CA_SettlementAgreement_AddSelectedSwift()
	testInstruction_CA_SettlementAgreement_DeleteSelectedSwift := Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift()
	testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties := Initate_TestInstruction_CA_SettlementAgreement_AddSelectedInstructedParties()
	testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties := Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties()

	// Append 'SettlementAgreement_Edit' TestInstructionsMap into one structure
	allTestInstructions_CA_SettlementAgreement.TestInstruction = append(
		allTestInstructions_CA_SettlementAgreement.TestInstruction,
		testInstruction_CA_SettlementAgreement_Edit.TestInstruction)
	allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_Edit.BasicTestInstructionInformation)
	allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_Edit.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute = append(
		allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute,
		testInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute...)
	allTestInstructions_CA_SettlementAgreement.ImmatureElementModel = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureElementModel,
		testInstruction_CA_SettlementAgreement_Edit.ImmatureElementModel...)
	allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes,
		testInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes)

	// Append 'SettlementAgreement_AddSelectedSwift' TestInstructionsMap into one structure
	allTestInstructions_CA_SettlementAgreement.TestInstruction = append(
		allTestInstructions_CA_SettlementAgreement.TestInstruction,
		testInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstruction)
	allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_AddSelectedSwift.BasicTestInstructionInformation)
	allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_AddSelectedSwift.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute = append(
		allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute,
		testInstruction_CA_SettlementAgreement_AddSelectedSwift.TestInstructionAttribute...)
	allTestInstructions_CA_SettlementAgreement.ImmatureElementModel = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureElementModel,
		testInstruction_CA_SettlementAgreement_AddSelectedSwift.ImmatureElementModel...)
	allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes,
		testInstruction_CA_SettlementAgreement_AddSelectedSwift.FangEngineClassesMethodsAttributes)

	// Append 'SettlementAgreement_DeleteSelectedSwift' TestInstructionsMap into one structure
	allTestInstructions_CA_SettlementAgreement.TestInstruction = append(
		allTestInstructions_CA_SettlementAgreement.TestInstruction,
		testInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstruction)
	allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_DeleteSelectedSwift.BasicTestInstructionInformation)
	allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_DeleteSelectedSwift.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute = append(
		allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute,
		testInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstructionAttribute...)
	allTestInstructions_CA_SettlementAgreement.ImmatureElementModel = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureElementModel,
		testInstruction_CA_SettlementAgreement_DeleteSelectedSwift.ImmatureElementModel...)
	allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes,
		testInstruction_CA_SettlementAgreement_DeleteSelectedSwift.FangEngineClassesMethodsAttributes)

	// Append 'SettlementAgreement_AddSelectedInstructedParties' TestInstructionsMap into one structure
	allTestInstructions_CA_SettlementAgreement.TestInstruction = append(
		allTestInstructions_CA_SettlementAgreement.TestInstruction,
		testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.TestInstruction)
	allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.BasicTestInstructionInformation)
	allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute = append(
		allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute,
		testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.TestInstructionAttribute...)
	allTestInstructions_CA_SettlementAgreement.ImmatureElementModel = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureElementModel,
		testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.ImmatureElementModel...)
	allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes,
		testInstruction_CA_SettlementAgreement_AddSelectedInstructedParties.FangEngineClassesMethodsAttributes)

	// Append 'SettlementAgreement_DeleteSelectedInstructedParties' TestInstructionsMap into one structure
	allTestInstructions_CA_SettlementAgreement.TestInstruction = append(
		allTestInstructions_CA_SettlementAgreement.TestInstruction,
		testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.TestInstruction)
	allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.BasicTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.BasicTestInstructionInformation)
	allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureTestInstructionInformation,
		testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.ImmatureTestInstructionInformation...)
	allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute = append(
		allTestInstructions_CA_SettlementAgreement.TestInstructionAttribute,
		testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.TestInstructionAttribute...)
	allTestInstructions_CA_SettlementAgreement.ImmatureElementModel = append(
		allTestInstructions_CA_SettlementAgreement.ImmatureElementModel,
		testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.ImmatureElementModel...)
	allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes = append(
		allTestInstructions_CA_SettlementAgreement.FangEngineClassesMethodsAttributes,
		testInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.FangEngineClassesMethodsAttributes)

	return allTestInstructions_CA_SettlementAgreement

}
