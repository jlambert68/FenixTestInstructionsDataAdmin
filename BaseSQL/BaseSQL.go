package BaseSQL

const (
	// ImmatureTestInstructionInformation
	ImmatureTestInstructionInformationSQLDelete = "SET search_path TO \"FenixGuiBuilder\";" +
		"DELETE FROM \"ImmatureTestInstructionInformation\" " +
		"WHERE \"TestInstructionUuid\" IN "

	ImmatureTestInstructionInformationSQLInsert = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"ImmatureTestInstructionInformation\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"DropZoneUuid\", \"DropZoneName\", \"DropZoneDescription\", \"DropZoneMouseOver\"," +
		"\"DropZoneColor\", \"TestInstructionAttributeType\"," +
		"\"TestInstructionAttributeUuid\", \"TestInstructionAttributeName\"," +
		"\"AttributeValueAsString\", \"AttributeValueUuid\", \"FirstImmatureElementUuid\"," +
		"\"AttributeActionCommand\") "

	// TestInstructions
	TestInstructionsSQLDelete = "SET search_path TO \"FenixGuiBuilder\";" +
		"DELETE FROM \"TestInstructions\" " +
		"WHERE \"TestInstructionUuid\" IN "

	TestInstructionsSQLInsert = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"TestInstructions\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"TestInstructionTypeUuid\", \"TestInstructionTypeName\", \"TestInstructionDescription\"," +
		"\"TestInstructionMouseOverText\", \"Deprecated\", \"Enabled\", \"MajorVersionNumber\"," +
		"\"MinorVersionNumber\", \"UpdatedTimeStamp\") "

	// BasicTestInstructionInformation
	BasicTestInstructionInformationSQLDelete = "SET search_path TO \"FenixGuiBuilder\";" +
		"DELETE FROM \"BasicTestInstructionInformation\" " +
		"WHERE \"TestInstructionUuid\" IN "

	BasicTestInstructionInformationSQLInsert = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"BasicTestInstructionInformation\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"TestInstructionTypeUuid\", \"TestInstructionTypeName\", \"Deprecated\"," +
		"\"MajorVersionNumber\", \"MinorVersionNumber\", \"UpdatedTimeStamp\"," +
		"\"TestInstructionColor\", \"TCRuleDeletion\", \"TCRuleSwap\"," +
		"\"TestInstructionDescription\", \"TestInstructionMouseOverText\", \"Enabled\") "

	// TestInstructionAttributes
	TestInstructionAttributesSQLDelete = "SET search_path TO \"FenixGuiBuilder\";" +
		"DELETE FROM \"TestInstructionAttributes\" " +
		"WHERE \"TestInstructionUuid\" IN "

	TestInstructionAttributesSQLInsert = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"TestInstructionAttributes\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"TestInstructionAttributeUuid\", \"TestInstructionAttributeName\"," +
		"\"TestInstructionAttributeDescription\", \"TestInstructionAttributeMouseOver\"," +
		"\"TestInstructionAttributeTypeUuid\", \"TestInstructionAttributeTypeName\"," +
		"\"TestInstructionAttributeValueAsString\", \"TestInstructionAttributeValueUuid\"," +
		"\"TestInstructionAttributeVisible\", \"TestInstructionAttributeEnabled\"," +
		"\"TestInstructionAttributeMandatory\"," +
		"\"TestInstructionAttributeVisibleInTestCaseArea\"," +
		"\"TestInstructionAttributeIsDeprecated\", \"TestInstructionAttributeInputMask\"," +
		"\"TestInstructionAttributeType\") "

	// ImmatureElementModelMessage
	ImmatureElementModelMessageSQLDelete = "SET search_path TO \"FenixGuiBuilder\";" +
		"DELETE FROM \"ImmatureElementModelMessage\" " +
		"WHERE \"TopImmatureElementUuid\" IN "

	ImmatureElementModelMessageSQLInsert = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"ImmatureElementModelMessage\" " +
		"(\"DomainUuid\", \"DomainName\", \"ImmatureElementUuid\", \"ImmatureElementName\"," +
		"\"PreviousElementUuid\", \"NextElementUuid\", \"FirstChildElementUuid\"," +
		"\"ParentElementUuid\", \"TestCaseModelElementType\", \"OriginalElementUuid\"," +
		"\"TopImmatureElementUuid\", \"IsTopElement\") "
)
