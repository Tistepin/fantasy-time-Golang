package utils

/**
* User:徐国纪
* Create_Time:下午 01:34
**/

import (
	jsoniter "github.com/json-iterator/go"
)

var jsons = jsoniter.ConfigCompatibleWithStandardLibrary

func ParamInt8Slice(str string) []int8 {
	res := make([]int8, 0)
	if IsEmptyString(str) {
		return res
	}

	if err := Deserialize(str, &res); err != nil {
		return res
	}

	return res
}

func IsEmptyString(str string) bool {
	return str == ""
}

func ParamInt64Slice(str string) []int64 {
	res := make([]int64, 0)
	if IsEmptyString(str) {
		return res
	}

	if err := Deserialize(str, &res); err != nil {
		return res
	}

	return res
}

func ParamStringSlice(str string) []string {
	res := make([]string, 0)
	if IsEmptyString(str) {
		return res
	}

	if err := Deserialize(str, &res); err != nil {
		return res
	}

	return res
}

func Deserialize(s string, obj interface{}) error {
	return jsons.UnmarshalFromString(s, obj)
}

func ParamUInt64Slice(str string) []uint64 {
	res := make([]uint64, 0)
	if IsEmptyString(str) {
		return res
	}

	if err := Deserialize(str, &res); err != nil {
		return res
	}

	return res
}
