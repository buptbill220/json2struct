package json2struct

import (
	"fmt"
	"testing"
)

func TestJson2Struct_int(t *testing.T) {
	ret, err := Json2Struct("auto_name", "123", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_int64(t *testing.T) {
	ret, err := Json2Struct("auto_name", "123123123123", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_float(t *testing.T) {
	ret, err := Json2Struct("auto_name", "123123.121212", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_string(t *testing.T) {
	ret, err := Json2Struct("auto_name", "wiyd", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_stringv2(t *testing.T) {
	ret, err := Json2Struct("auto_name", "wiyd", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_bool(t *testing.T) {
	ret, err := Json2Struct("auto_name", "true", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}


func TestJson2Struct_slice_int(t *testing.T) {
	ret, err := Json2Struct("auto_name", "[1,2,3,42]", GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_slice_string(t *testing.T) {
	ret, err := Json2Struct("auto_name", `["12","2","3","42"]`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_slice_bool(t *testing.T) {
	ret, err := Json2Struct("auto_name", `[true,false,false,false]`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_map(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"name": 23, "32": 23232323232323}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_struct(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"name": "fangming", "gender": true, "info": {}, "extra": [1,2]}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_slice_struct(t *testing.T) {
	ret, err := Json2Struct("auto_name", `[{"name": "fangming", "gender": true, "info": {}, "extra": [1,2]}]`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_slice_map(t *testing.T) {
	ret, err := Json2Struct("auto_name", `[{"name": 23, "32": 23}, {"name1": 2233, "dddd32": 23}]`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_map_slice(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"23":[1,2,3], "1": [3,4,1]}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_map_struct(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"23":{"int": 23, "string": "2323"}, "1": {"int": 23, "string": "2323", "float": 23.2}}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_struct_slice(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"23":{"int_slice": [23,12], "string": "2323"}, "1": {"int": 23, "string": "2323", "float": 23.2}}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_struct_map(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"23":{"a": "12", "string": "2323"}, "1": {"b": "23", "string": "2323", "float": "23.2"}}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_map_interface(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{"23": "23", "21": true}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_slice_interface(t *testing.T) {
	ret, err := Json2Struct("auto_name", `[1,2,3,"232343"]`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_b(t *testing.T) {
	ret, err := Json2Struct("auto_name", `[1,2,3,44444444444]`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}

func TestJson2Struct_a(t *testing.T) {
	ret, err := Json2Struct("auto_name", `{
  "timer": "2020-08-08 00:00:00",
  "timex": "2020-08-11 00:00:00",
  "mmx": "2020-08-11 00:00:00",
  "mtime": "2020-08-19 00:00:00",
  "idx": 1258121093648023822,
  "nxczx": 6858073458614731015,
  "dcsd": true,
  "xxx": true,
  "ddsdfsfdsf": "https://xx.cc",
  "limk": "https://xx.cc",
  "dse": "https://xx.cc",
  "sdf": "https://xx.cc",
  "show": true,
  "xshow": true,
  "shore": true,
  "buffer_time": 1,
  "filter_versions": [
    {
      "xeid": 1128,
      "dx": 120900,
      "eta": 100617600,
      "et": 10790400,
      "diisdf": "iphone"
    },
    {
      "xeid": 1128,
      "dx": 120900,
      "eta": 1600617600,
      "et": 1600790400,
      "xxdf": "ipad"
    }
  ],
  "tabs": [
    {
      "id": 1,
      "title": "test"
    },
    {
      "id": 2,
      "title": "test1"
    },
    {
      "id": 3,
      "title": "test2"
    },
    {
      "id": 4,
      "title": "test3"
    }
  ],
  "lsdkjfn": false,
  "ieir": false,
  "xkjkmlkj": false,
  "dfsdfsf": {
    "458122393648023822": {
      "max_coupon": {
        "credit": 818
      },
      "total_stock": 2000
    },
    "458122393sd8023822": {
      "max_coupon": {
        "credit": 818
      },
      "total_stock": 2000
    }
  }
}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}
