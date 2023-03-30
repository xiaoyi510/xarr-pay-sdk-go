package xarr_pay_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cast"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func GenerateSign(params interface{}, secretKey string) string {
	// 将参数转换为url.Values
	values := url.Values{}
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("generateSign: params must be a struct or a pointer to struct")
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		filed := t.Field(i).Tag.Get("json")
		value := v.Field(i)
		if !value.IsZero() {
			switch value.Kind() {
			case reflect.String:
				values.Set(filed, value.String())
				break
			case reflect.Int32, reflect.Int, reflect.Int64:
				values.Set(filed, cast.ToString(value.Int()))
				break
			case reflect.Bool:
				values.Set(filed, cast.ToString(value.Bool()))
				break
			case reflect.Slice:
				values.Set(filed, cast.ToString(value.Slice(0, value.Len())))
				break
			}
		}
	}

	// 按照key进行排序
	var keys []string
	for k := range values {
		// 签名不进入计算
		if k == "sign" || k == "sign_type" || k == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 按照key=value的格式拼接参数
	var queryStr string
	for _, k := range keys {
		v := values[k][0]
		if (v) == "" {
			continue
		}
		queryStr += k + "=" + (v) + "&"
	}
	queryStr = strings.TrimRight(queryStr, "&")

	if secretKey != "" {
		// 添加密钥
		queryStr += secretKey
	}

	// 计算md5值
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(queryStr))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
