// Map常用操作函数
package hmap

import (
	"errors"
	"fmt"
)

// Map转Struct
// 输入参数为Map类型值，输出为结构体
func MapToStruct(input interface{}, output interface{}) {
	err := Decode(input, output)
	panic(errors.New(fmt.Sprintf("MapToStruct error %s", err.Error())))
}
