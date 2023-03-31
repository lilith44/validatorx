package validatorx

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var (
	// 根据 2^(index-1) % 11 得出位置:权重 哈希表 index从右至左算
	weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码值换算表
	codes = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
)

func checkLastNum(idCard string) bool {
	sum := 0
	for i := 0; i < 17; i++ {
		n, _ := strconv.Atoi(idCard[i : i+1])
		sum += n * weight[i]
	}

	sum = sum % 11

	return codes[sum] == idCard[17:]
}

// IsIdCard 是否是合法的18位身份证号
func IsIdCard(fl validator.FieldLevel) bool {
	idCard := fl.Field().String()
	return regexp.MustCompile(IdCardWith18Len).MatchString(idCard) && checkLastNum(idCard)
}

// IsIdCardWith15Len 是否是合法的15位身份证号
func IsIdCardWith15Len(fl validator.FieldLevel) bool {
	return regexp.MustCompile(IdCardWith15Len).MatchString(fl.Field().String())
}
