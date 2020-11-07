package goenv

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func String(key string) string {
	return os.Getenv(key)
}

func StringOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}

func MustString(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("no key " + key + " in environment variables")
	}

	return v
}

func Int(key string) int {
	v := os.Getenv(key)

	i, _ := strconv.Atoi(v)
	return i
}

func IntOrDefault(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}

	return i
}

func MustInt(key string) int {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("key %s is empty", key))
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer", key))
	}

	return i
}

func DurationOrDefault(key string, multiplier, def time.Duration) time.Duration {
	i := IntOrDefault(key, -1)
	if i == -1 {
		return def
	}

	return time.Duration(i) * multiplier
}

func IsTruthy(key string) bool {
	v := strings.ToLower(os.Getenv(key))

	if v == "" {
		return false
	}

	if v == "0" {
		return false
	}

	if v == "false" || v == "off" || v == "no" {
		return false
	}

	return true
}

func IsFalsy(key string) bool {
	return !IsTruthy(key)
}
