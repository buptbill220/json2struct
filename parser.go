package json2struct

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"strconv"
	"reflect"
	"regexp"
	"unsafe"
)

var _ = fmt.Printf
type JsonType uint
type GeneratorType uint

const (
	InterfaceType JsonType = iota
	BoolType
	IntType
	Int64Type
	FloatType
	StringType
	SliceType
	MapType
	StructType
	KeywordStringType
	
	GoStructType GeneratorType = iota
	PbMessageType
	IntMax = int64((1 << 31) - 1)
	IntMin = int64(-(1 << 31))
)

func (t JsonType) string() string {
	switch t {
	case InterfaceType:
		return "interface{}"
	case BoolType:
		return "bool"
	case IntType:
		return "int"
	case Int64Type:
		return "int64"
	case FloatType:
		return "float64"
	case StringType, KeywordStringType:
		return "string"
	case SliceType:
		return "[]"
	case MapType:
		return "map"
	case StructType:
		return "*struct"
	}
	panic("unknown type: ")
}

var (
	commonInitialisms = []string{
		"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP",
		"HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA",
		"SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID",
		"URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS",
	}
)

func str2Bytes(s string) []byte {
	x := (*reflect.StringHeader)(unsafe.Pointer(&s))
	h := reflect.SliceHeader{x.Data, x.Len, x.Len}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func mergeValueType (typeMap map[JsonType]bool) JsonType {
	if len(typeMap) == 2 {
		// all type are nil/struct/slice/map
		if typeMap[InterfaceType] && (typeMap[SliceType] || typeMap[MapType] || typeMap[StructType]) {
			delete(typeMap, InterfaceType)
			goto end
		}
		// all type are string
		if typeMap[StringType] && typeMap[KeywordStringType] {
			return StringType
		}
		if typeMap[MapType] && typeMap[StructType] {
			return MapType
		}
		if typeMap[StringType] || typeMap[KeywordStringType] {
			return InterfaceType
		}
	}
	if len(typeMap) >= 2 {
		if typeMap[SliceType] || typeMap[MapType] || typeMap[BoolType] {
			return InterfaceType
		}
		// all type are float64/int64/int
		if typeMap[FloatType] {
			return FloatType
		}
		// all type are int64/int
		return Int64Type
	}
	end:
	for jt, _ := range typeMap {
		return jt
	}
	return InterfaceType
}

func mergeMapKeyType (typeMap map[JsonType]bool) JsonType {
	if len(typeMap) <= 1 {
		for jt, _ := range typeMap {
			return jt
		}
		return StringType
	}
	if typeMap[StringType] || typeMap[KeywordStringType] || typeMap[BoolType] {
		return StringType
	}
	if typeMap[FloatType] {
		return FloatType
	}
	// all type are int64/int
	return Int64Type
}

func parseSliceType(s []interface{}) JsonType {
	if len(s) == 0 {
		return InterfaceType
	}
	typeMap := make(map[JsonType]bool, 2)
	for _, v := range s {
		typeMap[parseValueType(v)] = true
	}
	return mergeValueType(typeMap)
}

// key type of map must be simple type string/int/float/bool, interface is impossible
func parseMapType(m map[string]interface{}) (isMap, isObject bool, jt JsonType){
	if len(m) == 0 {
		return false, true, KeywordStringType
	}
	typeMap := make(map[JsonType]bool, 2)
	for k, _ := range m {
		t, _ :=  toSimpleValue(k)
		typeMap[t] = true
	}
	jt = mergeMapKeyType(typeMap)
	if jt == KeywordStringType {
		return false, true, jt
	}
	if jt == InterfaceType {
		jt = StringType
	}
	return true, false, jt
}

func toProperCase(name string) string {
	buf := bytes.NewBuffer(make([]byte, 0, (len(name) << 1) + (len(name) >> 1) >> 1))
	slices := strings.Split(name, "_")
	for _, s := range slices {
		if len(s) == 0 {
			continue
		}
		buf.WriteString(strings.ToUpper(s[0:1]) + s[1:])
	}
	return buf.String()
}

func toValue(data string) (JsonType, interface{}) {
	var interV interface{}
	e := json.Unmarshal(str2Bytes(data), &interV)
	if e == nil {
		if interV == nil {
			return InterfaceType, nil
		}
		t := reflect.TypeOf(interV).Kind()
		switch t {
		case reflect.Slice, reflect.Array: // []interface{}
			return SliceType, interV
		case reflect.Map: // map[string]interface{}
			isM, _, _ := parseMapType(interV.(map[string]interface{}))
			if isM {
				return MapType, interV
			}
			return StructType, interV
		}
	}
	return toSimpleValue(data)
}

func toSimpleValue(data string) (JsonType, interface{}) {
	if data == "" {
		return StringType, data
	}
	v, e := strconv.ParseFloat(data, 64)
	if e == nil {
		return checkFloatType(v)
	}
	if data == "true" || data == "false" {
		b, _ := strconv.ParseBool(data)
		return BoolType, b
	}
	if isValidKeyword(data) {
		return KeywordStringType, data
	}
	return StringType, data
}

func checkFloatType(v float64) (JsonType, interface{}) {
	ifv := int64(v)
	if math.Abs(float64(v) - float64(ifv)) > 1e-15 {
		return FloatType, v
	}
	if (ifv >= IntMin && ifv <= IntMax) {
		return IntType, int(ifv)
	}
	return Int64Type, ifv
}

func parseValueType(v interface{}) JsonType {
	if v == nil {
		return InterfaceType
	}
	if _, ok := v.([]interface{}); ok {
		return SliceType
	}
	if _, ok := v.(map[string]interface{}); ok {
		isMap, _, _ := parseMapType(v.(map[string]interface{}))
		if isMap {
			return MapType
		}
		return StructType
	}
	switch v.(type) {
	case bool:
		return BoolType
	case int, uint, int32, uint32, int64, uint64:
		return Int64Type
	case float64:
		jt, _ := checkFloatType(v.(float64))
		return jt
	case string:
		if isValidKeyword(v.(string)) {
			return KeywordStringType
		} else {
			return StringType
		}
	}
	return InterfaceType
}

var KeywordReg, _ = regexp.Compile(`^[_a-zA-Z]\w*$`)
func isValidKeyword(str string) bool {
	return KeywordReg.MatchString(str)
}

func isSimpleType(t JsonType) bool {
	switch t {
	case BoolType, IntType, Int64Type, FloatType, StringType, KeywordStringType, InterfaceType:
		return true
	}
	return false
}