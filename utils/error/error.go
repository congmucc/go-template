package error

import "fmt"

/**
 * @title: error_util
 * @description:
 * @author: congmu
 * @date:    2024/6/23 13:42
 * @version: 1.0
 */

// AppendError 合并错误，构建错误链
func AppendError(existErr, newErr error) error {
	if existErr == nil {
		return newErr
	}
	return fmt.Errorf("%v, %w", existErr, newErr)
}
