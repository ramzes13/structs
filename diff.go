package structs

import "reflect"

var DiffDefaultTagName = "db"

func GenerateDiff(s1, s2 interface{}) map[string]interface{} {
	diff := make(map[string]interface{})
	s1Value := strctVal(s1)
	s2Value := strctVal(s2)

	fields := strctFields(s1Value.Type(), DiffDefaultTagName)

	for _, field := range fields {
		name := field.Name
		s1Val := s1Value.FieldByName(name)
		s2Val := s2Value.FieldByName(name)

		tagName, _ := parseTag(field.Tag.Get(DiffDefaultTagName))

		if tagName != "" {
			name = tagName
		}

		if !reflect.DeepEqual(s1Val.Interface(), s2Val.Interface()) {
			diff[name] = s1Val.Interface()
		}

	}
	return diff
}
