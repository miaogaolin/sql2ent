package converter

import (
	"fmt"
	"strings"

	"github.com/miaogaolin/ddlparser/parser"
)

var fieldMap = map[int]string{
	// For consistency, all integer types are converted to int64
	// number
	parser.Bool:      "Bool",
	parser.Boolean:   "Bool",
	parser.TinyInt:   "Int8",
	parser.SmallInt:  "Int16",
	parser.MediumInt: "Int32",
	parser.Int:       "Int32",
	parser.MiddleInt: "Int32",
	parser.Int1:      "Int8",
	parser.Int2:      "Int16",
	parser.Int3:      "Int32",
	parser.Int4:      "Int32",
	parser.Int8:      "int64",
	parser.Integer:   "Int32",
	parser.BigInt:    "Int64",
	parser.Float:     "Float",
	parser.Float4:    "Float",
	parser.Float8:    "Float",
	parser.Double:    "Float",
	parser.Decimal:   "Float",
	// date&time
	parser.Date:      "Time",
	parser.DateTime:  "Time",
	parser.Timestamp: "Time",
	parser.Time:      "Time",
	parser.Year:      "Time",
	// string
	parser.Char:       "String",
	parser.VarChar:    "String",
	parser.Binary:     "String",
	parser.VarBinary:  "String",
	parser.TinyText:   "String",
	parser.Text:       "Text",
	parser.MediumText: "Text",
	parser.LongText:   "Text",
	parser.Enum:       "Enum",
	parser.Set:        "String",
	parser.Json:       "String",
}

// ConvertDataType converts column type into ent function
func ConvertField(dataBaseType int) (string, error) {
	tp, ok := fieldMap[dataBaseType]
	if !ok {
		return "", fmt.Errorf("unsupported database type: %v", dataBaseType)
	}

	return tp, nil
}

func ConvertDefaultValue(dataType parser.DataType, val string, isHas bool) (imports []string, fields string) {
	defaultVal := ""
	isTime := false
	switch dataType.Type() {
	case parser.Bool,
		parser.Boolean,
		parser.TinyInt,
		parser.SmallInt,
		parser.MediumInt,
		parser.Int,
		parser.MiddleInt,
		parser.Int1,
		parser.Int2,
		parser.Int3,
		parser.Int4,
		parser.Int8,
		parser.Integer,
		parser.BigInt,
		parser.Float,
		parser.Float4,
		parser.Float8,
		parser.Double,
		parser.Decimal:
		defaultVal = val

	case parser.Date,
		parser.DateTime,
		parser.Time,
		parser.Year,
		parser.Timestamp:
		if isHas && strings.HasPrefix(val, "CURRENT_TIMESTAMP") {
			isTime = true
			defaultVal = "time.Now"
		}
		if strings.Contains(val, "ONUPDATECURRENT_TIMESTAMP") {
			isTime = true
			fields += ".UpdateDefault(time.Now)"
		}
	case parser.Enum:
		defaultVal = `"` + val + `"`
		fields += fmt.Sprintf(`.Values("%s")`, strings.Join(dataType.Value(), `","`))

	case parser.Char,
		parser.VarChar,
		parser.Binary,
		parser.VarBinary,
		parser.TinyText,
		parser.Text,
		parser.MediumText,
		parser.LongText,
		parser.Set,
		parser.Json:
		defaultVal = `"` + val + `"`
	}
	if isTime {
		imports = append(imports, "time")
	}
	if isHas && defaultVal != "" {
		fields = fmt.Sprintf(`.Default(%s)`, defaultVal) + fields
	}
	return
}
