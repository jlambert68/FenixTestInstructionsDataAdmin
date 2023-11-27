package GenerateSqlData

import (
	"fmt"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/BaseSQL"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

// Initate_TestInstruction_CA_CustodyAccount_Search
// Function that creates all data for the TestInstruction
func GenerateSqlInsert_For_TestInstructions_CA(testInstructionSetUp *TestInstructions.TestInstruction_CA_TestCaseSetUpStruct) {

	testInstructionsSQL := BaseSQL.TestInstructionsSQLInsert
	basicTestInstructionInformationSQL := BaseSQL.BasicTestInstructionInformationSQLInsert
	immatureTestInstructionInformationSQL := BaseSQL.ImmatureTestInstructionInformationSQLInsert
	testInstructionAttributesSQL := BaseSQL.TestInstructionAttributesSQLInsert
	immatureElementModelMessageSQL := BaseSQL.ImmatureElementModelMessageSQLInsert

	var dataRowToBeInsertedMultiType []interface{}
	var dataRowsToBeInsertedMultiType [][]interface{}

	// 'TestInstructionsSQLInsert'
	// Create data to be inserted in the DB-table 'TestInstructionsMap'
	dataRowsToBeInsertedMultiType = nil
	dataRowToBeInsertedMultiType = nil

	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.DomainUUID)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.DomainName)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.TestInstructionUUID)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.TestInstructionName)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.TestInstructionTypeUUID)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.TestInstructionTypeName)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.TestInstructionDescription)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.TestInstructionMouseOverText)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.Deprecated)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.Enabled)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.MajorVersionNumber)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.MinorVersionNumber)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.TestInstruction.UpdatedTimeStamp)

	dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)

	if dataRowsToBeInsertedMultiType != nil {
		testInstructionsSQL = testInstructionsSQL + shared_code.GenerateSQLInsertValues(dataRowsToBeInsertedMultiType)
		testInstructionsSQL = testInstructionsSQL + ";"
	} else {
		testInstructionsSQL = ""
	}

	// 'BasicTestInstructionInformationSQLInsert'
	// Create data to be inserted in the DB-table 'BasicTestInstructionInformation'
	dataRowsToBeInsertedMultiType = nil
	dataRowToBeInsertedMultiType = nil

	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.DomainUUID)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.DomainName)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionUUID)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionName)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionTypeUUID)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionTypeName)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.Deprecated)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.MajorVersionNumber)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.MinorVersionNumber)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.UpdatedTimeStamp)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionColor)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TCRuleDeletion)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TCRuleSwap)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionDescription)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.TestInstructionMouseOverText)
	dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, testInstructionSetUp.BasicTestInstructionInformation.Enabled)

	dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)

	if dataRowsToBeInsertedMultiType != nil {
		basicTestInstructionInformationSQL = basicTestInstructionInformationSQL + shared_code.GenerateSQLInsertValues(dataRowsToBeInsertedMultiType)
		basicTestInstructionInformationSQL = basicTestInstructionInformationSQL + ";"
	} else {
		basicTestInstructionInformationSQL = ""
	}

	// 'ImmatureTestInstructionInformationSQLInsert'
	// Create data to be inserted in the DB-table 'ImmatureTestInstructionInformation'
	dataRowsToBeInsertedMultiType = nil

	// Loop all ImmatureTestInstructionInformation-messages and add to data-structure used for converting into SQL-Values-data
	for _, tempImmatureTestInstructionInformation := range testInstructionSetUp.ImmatureTestInstructionInformation {
		dataRowToBeInsertedMultiType = nil

		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DomainUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DomainName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.TestInstructionUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.TestInstructionName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DropZoneUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DropZoneName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DropZoneDescription)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DropZoneMouseOver)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.DropZoneColor)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.TestInstructionAttributeType)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.TestInstructionAttributeUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.TestInstructionAttributeName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.AttributeValueAsString)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.AttributeValueUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.FirstImmatureElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureTestInstructionInformation.AttributeActionCommand)

		dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)
	}

	if dataRowsToBeInsertedMultiType != nil {
		immatureTestInstructionInformationSQL = immatureTestInstructionInformationSQL + shared_code.GenerateSQLInsertValues(dataRowsToBeInsertedMultiType)
		immatureTestInstructionInformationSQL = immatureTestInstructionInformationSQL + ";"
	} else {
		immatureTestInstructionInformationSQL = ""
	}

	// 'TestInstructionAttributesSQLInsert'
	// Create data to be inserted in the DB-table 'TestInstructionAttributes'
	dataRowsToBeInsertedMultiType = nil

	// Loop all TestInstructionAttribute-messages and add to data-structure used for converting into SQL-Values-data
	for _, tempTestInstructionAttribute := range testInstructionSetUp.TestInstructionAttribute {
		dataRowToBeInsertedMultiType = nil

		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.DomainUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.DomainName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeDescription)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeMouseOver)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeTypeUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeTypeName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeValueAsString)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeValueUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeVisible)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeEnabled)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeMandatory)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeVisibleInTestCaseArea)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeIsDeprecated)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeInputMask)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempTestInstructionAttribute.TestInstructionAttributeType)

		dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)
	}

	if dataRowsToBeInsertedMultiType != nil {
		testInstructionAttributesSQL = testInstructionAttributesSQL + shared_code.GenerateSQLInsertValues(dataRowsToBeInsertedMultiType)
		testInstructionAttributesSQL = testInstructionAttributesSQL + ";"
	} else {
		testInstructionAttributesSQL = ""
	}

	// 'ImmatureElementModelMessageSQLInsert'
	// Create data to be inserted in the DB-table 'ImmatureElementModelMessage'
	dataRowsToBeInsertedMultiType = nil

	// Loop all ImmatureElementModels and add to data-structure used for converting into SQL-Values-data
	for _, tempImmatureElementModel := range testInstructionSetUp.ImmatureElementModel {
		dataRowToBeInsertedMultiType = nil

		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.DomainUUID)

		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.DomainName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.ImmatureElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.ImmatureElementName)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.PreviousElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.NextElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.FirstChildElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.ParentElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.TestCaseModelElementType)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.OriginalElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.TopImmatureElementUUID)
		dataRowToBeInsertedMultiType = append(dataRowToBeInsertedMultiType, tempImmatureElementModel.IsTopElement)

		dataRowsToBeInsertedMultiType = append(dataRowsToBeInsertedMultiType, dataRowToBeInsertedMultiType)
	}

	if dataRowsToBeInsertedMultiType != nil {
		immatureElementModelMessageSQL = immatureElementModelMessageSQL + shared_code.GenerateSQLInsertValues(dataRowsToBeInsertedMultiType)
		immatureElementModelMessageSQL = immatureElementModelMessageSQL + ";"
	} else {
		immatureElementModelMessageSQL = ""
	}

	// Output all SQL-data
	fmt.Println(testInstructionsSQL)
	fmt.Println(basicTestInstructionInformationSQL)
	fmt.Println(immatureTestInstructionInformationSQL)
	fmt.Println(testInstructionAttributesSQL)
	fmt.Println(immatureElementModelMessageSQL)

}
