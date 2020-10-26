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
  "pre_begin_time_str": "2020-08-08 00:00:00",
  "pre_end_time_str": "2020-08-11 00:00:00",
  "begin_time_str": "2020-08-11 00:00:00",
  "end_time_str": "2020-08-19 00:00:00",
  "share_combo_id": 6858121093648023822,
  "address_combo_id": 6858073458614731015,
  "show_more_talent_areas": true,
  "show_more_brand_areas": true,
  "more_talent_areas_link": "https://aweme.snssdk.com/magic/page/ejs/5f1eb28e1a2ce502e2d4db61?appType=douyin",
  "more_brand_areas_link": "https://aweme.snssdk.com/magic/page/ejs/5f296d17f6a99702e4054f5c?appType=douyin",
  "more_talent_areas_link_hotsoon": "https://hotsoon.snssdk.com/magic/page/ejs/5f2a9e1ea13a0702da221300?appType=hotsoon",
  "more_brand_areas_link_hotsoon": "https://hotsoon.snssdk.com/magic/page/ejs/5f2bfedc02c20302eaab5b26?appType=hotsoon",
  "show_task_entry": true,
  "show_address_task_entry": true,
  "show_share_task_entry": true,
  "buffer_time": 1,
  "filter_versions": [
    {
      "app_id": 1128,
      "client_version": 120900,
      "begin_time": 1600617600,
      "end_time": 1600790400,
      "device_platform": "iphone"
    },
    {
      "app_id": 1128,
      "client_version": 120900,
      "begin_time": 1600617600,
      "end_time": 1600790400,
      "device_platform": "ipad"
    }
  ],
  "tabs": [
    {
      "id": 1,
      "title": "超级直播"
    },
    {
      "id": 2,
      "title": "达人种草"
    },
    {
      "id": 3,
      "title": "品牌专区"
    },
    {
      "id": 4,
      "title": "C位打榜"
    }
  ],
  "max_coupon_switch": false,
  "user_info_switch": false,
  "user_address_switch": false,
  "max_coupon_default": {
    "6858121093648023822": {
      "max_coupon": {
        "credit": 818
      },
      "total_stock": 2000
    },
    "6858073458614731015": {
      "max_coupon": {
        "credit": 818
      },
      "total_stock": 2000
    }
  }
}`, GoStructType)
	fmt.Printf("ret:\n %s, %v\n", ret, err)
}
