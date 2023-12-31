/*
WorldCat Search API v. 2

Searching of WorldCat

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding - struct for FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding
type FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding struct {
	ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner  *[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner
	ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 *[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1
}

// []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAsFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding is a convenience function that returns []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner wrapped in FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding
func ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInnerAsFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding(v *[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding {
	return FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding{
		ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner: v,
	}
}

// []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1AsFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding is a convenience function that returns []FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 wrapped in FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding
func ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1AsFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding(v *[]FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding {
	return FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding{
		ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner
	err = newStrictDecoder(data).Decode(&dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner)
	if err == nil {
		jsonArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner, _ := json.Marshal(dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner)
		if string(jsonArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner) == "{}" { // empty struct
			dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner = nil
	}

	// try to unmarshal data into ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1
	err = newStrictDecoder(data).Decode(&dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1)
	if err == nil {
		jsonArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1, _ := json.Marshal(dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1)
		if string(jsonArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1) == "{}" { // empty struct
			dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner = nil
		dst.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) MarshalJSON() ([]byte, error) {
	if src.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner != nil {
		return json.Marshal(&src.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner)
	}

	if src.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 != nil {
		return json.Marshal(&src.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner != nil {
		return obj.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner
	}

	if obj.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1 != nil {
		return obj.ArrayOfFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHoldingOneOfInner1
	}

	// all schemas are nil
	return nil
}

type NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding struct {
	value *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding
	isSet bool
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) Get() *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding {
	return v.value
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) Set(val *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) {
	v.value = val
	v.isSet = true
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) IsSet() bool {
	return v.isSet
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding(val *FindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding {
	return &NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding{value: val, isSet: true}
}

func (v NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindBibRetainedHoldings200ResponseBriefRecordsInnerInstitutionHolding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
