package utils

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

// 计算请求签名
func Sign(params map[string]string, accessKey string) string {

	keys := lo.Keys(params)
	sort.Strings(keys)

	//拼接签名原始字符串
	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "sign" && value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, value))
		}
	}
	sb.WriteString(fmt.Sprintf("key=%s", accessKey))
	signStr := sb.String()

	fmt.Printf("[rawString]%s\n", signStr)

	//4. 计算md5签名
	signResult := GetMD5([]byte(signStr))
	return signResult
}

// 验证签名
func VerifySign(params map[string]string, signKey string) bool {
	// Check if sign exists in params
	signValue, exists := params["sign"]
	if !exists {
		return false
	}

	// Get the sign value and remove it from params
	key := fmt.Sprintf("%v", signValue)
	delete(params, "sign")

	// Generate current signature
	currentKey := Sign(params, signKey)

	// Compare the signatures
	return key == currentKey
}
