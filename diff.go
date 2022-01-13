package structs

import (
	"encoding/json"
	"reflect"
)

var DiffDefaultTagName = "db"

func GetFields(s1 interface{}) []string {
	response := []string{}
	s1Value := strctVal(s1)
	fields := strctFields(s1Value.Type(), DiffDefaultTagName)

	for _, field := range fields {
		name := field.Name
		tagName, _ := parseTag(field.Tag.Get(DiffDefaultTagName))

		if tagName != "" {
			name = tagName
		}

		response = append(response, name)
	}

	return response

}
func GenerateDiff(s1, s2 interface{}) map[string]interface{} {
	diff := make(map[string]interface{})
	s1Value := strctVal(s1)
	s2Value := strctVal(s2)

	fields := strctFields(s1Value.Type(), DiffDefaultTagName)

	for _, field := range fields {
		name := field.Name
		s1Val := s1Value.FieldByName(name)
		s2Val := s2Value.FieldByName(name)

		s1ValInterface := s1Val.Interface()
		s2ValInterface := s2Val.Interface()
		tagName, tagOpts := parseTag(field.Tag.Get(DiffDefaultTagName))

		if tagOpts.Has("omitcompare") {
			continue
		}

		if tagName != "" {
			name = tagName
		}

		if !DeepEqual(s1ValInterface, s2ValInterface) {
			diff[name] = s1Val.Interface()
		}

	}
	return diff
}

func DeepEqualJson(v1, v2 interface{}) bool {
	//the equality of two objects can easily be tested by testing the equality of their canonical forms
	var x1 interface{}
	bytesA, _ := json.Marshal(v1)
	_ = json.Unmarshal(bytesA, &x1)
	var x2 interface{}
	bytesB, _ := json.Marshal(v2)
	_ = json.Unmarshal(bytesB, &x2)

	return reflect.DeepEqual(x1, x2)
}

func DeepEqual(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}

	v := reflect.ValueOf(v1)
	typeName := v.Type().String()

	if typeName != "pgtype.JSONB" {
		return false
	}

	return DeepEqualJson(v1, v2)
}
