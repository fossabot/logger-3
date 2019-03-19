package file_provider

import (
	"fmt"
	"os"
	"time"
)

func TimeFileProvider(format string, timeFormat string) func(f *os.File) (*os.File, error) {
	return func(f *os.File) (*os.File, error) {
		if f != nil {
			err := f.Close()
			if err != nil {
				return nil, err
			}
		}
		return os.Create(fmt.Sprintf(format, time.Now().Format(timeFormat)))
	}
}