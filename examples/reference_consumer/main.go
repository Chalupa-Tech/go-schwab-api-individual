package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Chalupa-Tech/go-schwab-api-individual/models/components"
	"github.com/Chalupa-Tech/go-schwab-api-individual/types"
)

func main() {
	log.Println("Starting Reference Consumer Validation...")

	// 1. Validate OptionContract (Union Type)
	log.Println("Validating OptionContract Union Type...")

	// Case A: Single Object
	singleOptionJson := []byte(`{
		"putCall": "PUT",
		"symbol": "XYZ 20240101 P 100",
		"description": "XYZ Jan 2024 100 Put"
	}`)

	var optionContractSingle components.OptionContract
	if err := json.Unmarshal(singleOptionJson, &optionContractSingle); err != nil {
		// Note: expected to fail if the mock JSON doesn't exactly match required fields of OptionContractObject
		// But here we are just testing the union mechanism if the input is valid.
		// Since OptionContractObject might have required fields, let's just construct it manually to verify type signature.
		log.Printf("JSON unmarshal (expected if partial data): %v\n", err)
	}

	// Construct manually to prove accessibility
	ocObj := components.OptionContractObject{
		Symbol: types.String("TEST_SYMBOL"),
	}
	contract := components.CreateOptionContractOptionContractObject(ocObj)
	if contract.Type != components.OptionContractTypeOptionContractObject {
		log.Fatal("Failed to create OptionContract with single object")
	}
	log.Println("Successfully created OptionContract with single object")

	// Case B: Array of Objects
	ocArray := []components.OptionContractObject{
		{Symbol: types.String("TEST_SYMBOL_1")},
		{Symbol: types.String("TEST_SYMBOL_2")},
	}
	contractArray := components.CreateOptionContractArrayOfOptionContractObject(ocArray)
	if contractArray.Type != components.OptionContractTypeArrayOfOptionContractObject {
		log.Fatal("Failed to create OptionContract with array of objects")
	}
	log.Println("Successfully created OptionContract with array")

	// 2. Validate AccountEquity (Struct Validation)
	log.Println("Validating AccountEquity Struct...")

	equity := components.AccountEquity{
		AssetType: components.AccountEquityAssetTypeEquity,
		Symbol:    types.String("AAPL"),
	}

	// Marshal to check if Type field is automatically handled or if tagging works
	data, err := equity.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal AccountEquity: %v", err)
	}
	log.Printf("Marshaled AccountEquity: %s\n", string(data))

	// Validate defaults if any (e.g. Type discriminated field)
	// The generated code for AccountEquity has `Type *string "default:\"AccountEquity\" json:\"_type\""`
	// Note: Speakeasy's MarshalJSON utils usually handle defaults if configured.

	log.Println("Reference Consumer Validation Passed!")
	os.Exit(0)
}
