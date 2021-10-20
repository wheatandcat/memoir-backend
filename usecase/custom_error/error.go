package custom_error

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func GetCustomStackTrace(err error) errors.StackTrace {
	var fs errors.StackTrace
	st, ok := err.(interface{ StackTrace() errors.StackTrace })
	if !ok {
		return fs
	}

	frames := st.StackTrace()

	for _, frame := range frames {
		pc := uintptr(frame)
		fun := runtime.FuncForPC(pc)
		f, _ := fun.FileLine(pc)

		// 不要なStackトレースが多すぎるのでフィルタリング
		if strings.Contains(f, "memoir-backend") || strings.Contains(f, "pkg/error") {
			if !strings.Contains(f, "usecase/custom_error") {
				fs = append(fs, frame)
			}
		}
	}

	if os.Getenv("APP_ENV") == "local" {
		// エラー出力
		fmt.Printf("■ error: %+v\n", err.Error())
		if len(fs) > 3 {
			fs = fs[:3]
		}

		fmt.Printf("■ stack trace: %+v\n", fs)
	}

	return fs
}

func CustomError(err error) error {
	e := errors.WithStack(err)
	GetCustomStackTrace(e)

	return e
}

func CustomErrorWrap(err error, message string) error {
	e := errors.Wrap(err, message)
	GetCustomStackTrace(e)

	return e
}
