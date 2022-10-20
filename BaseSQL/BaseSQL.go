package BaseSQL

const (
	ImmatureTestInstructionInformationSQL = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"ImmatureTestInstructionInformation\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"DropZoneUuid\", \"DropZoneName\", \"DropZoneDescription\", \"DropZoneMouseOver\"," +
		"\"DropZoneColor\", \"TestInstructionAttributeType\"," +
		"\"TestInstructionAttributeUuid\", \"TestInstructionAttributeName\"," +
		"\"AttributeValueAsString\", \"AttributeValueUuid\", \"FirstImmatureElementUuid\"," +
		"\"AttributeActionCommand\") "

	TestInstructionsSQL = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"TestInstructions\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"TestInstructionTypeUuid\", \"TestInstructionTypeName\", \"TestInstructionDescription\"," +
		"\"TestInstructionMouseOverText\", \"Deprecated\", \"Enabled\", \"MajorVersionNumber\"," +
		"\"MinorVersionNumber\", \"UpdatedTimeStamp\") "

	BasicTestInstructionInformationSQL = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"BasicTestInstructionInformation\" " +
		"(\"DomainUuid\", \"DomainName\", \"TestInstructionUuid\", \"TestInstructionName\"," +
		"\"TestInstructionTypeUuid\", \"TestInstructionTypeName\", \"Deprecated\"," +
		"\"MajorVersionNumber\", \"MinorVersionNumber\", \"UpdatedTimeStamp\"," +
		"\"TestInstructionColor\", \"TCRuleDeletion\", \"TCRuleSwap\"," +
		"\"TestInstructionDescription\", \"TestInstructionMouseOverText\", \"Enabled\") "

	TestInstructionAttributesSQL = "SET search_path TO \"FenixGuiBuilder\";" +
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

	ImmatureElementModelMessageSQL = "SET search_path TO \"FenixGuiBuilder\";" +
		"INSERT INTO \"ImmatureElementModelMessage\" " +
		"(\"DomainUuid\", \"DomainName\", \"ImmatureElementUuid\", \"ImmatureElementName\"," +
		"\"PreviousElementUuid\", \"NextElementUuid\", \"FirstChildElementUuid\"," +
		"\"ParentElementUuid\", \"TestCaseModelElementType\", \"OriginalElementUuid\"," +
		"\"TopImmatureElementUuid\", \"IsTopElement\") "
)
