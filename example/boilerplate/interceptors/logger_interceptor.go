package interceptors

import (
	"fmt"
	"time"

	"github.com/l1ttps/routix/interceptor"
)

func LoggerInterceptor(c *interceptor.InterceptorContext) func() {
	fmt.Println("Before")
	timeNow := time.Now()
	return func() {
		fmt.Println("After...", time.Since(timeNow))
	}
}
