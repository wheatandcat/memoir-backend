package main

import (
	"log"
	"os"
)

var _ = func() interface{} {
	var envValues = map[string]string{"APP_ENV": "local"}

	for key, value := range envValues {
		if err := os.Setenv(key, value); err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}()
