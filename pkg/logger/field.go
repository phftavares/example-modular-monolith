package logger

import (
	"time"

	"go.uber.org/zap"
)

// String creates a zap.Field for a string value.
func String(key, value string) zap.Field {
	return zap.String(key, value)
}

// Bool creates a zap.Field for a bool value.
func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

// Int creates a zap.Field for an int value.
func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// Int8 creates a zap.Field for an int8 value.
func Int8(key string, value int8) zap.Field {
	return zap.Int8(key, value)
}

// Int16 creates a zap.Field for an int16 value.
func Int16(key string, value int16) zap.Field {
	return zap.Int16(key, value)
}

// Int32 creates a zap.Field for an int32 value.
func Int32(key string, value int32) zap.Field {
	return zap.Int32(key, value)
}

// Int64 creates a zap.Field for an int64 value.
func Int64(key string, value int64) zap.Field {
	return zap.Int64(key, value)
}

// Uint creates a zap.Field for a uint value.
func Uint(key string, value uint) zap.Field {
	return zap.Uint(key, value)
}

// Uint8 creates a zap.Field for a uint8 value.
func Uint8(key string, value uint8) zap.Field {
	return zap.Uint8(key, value)
}

// Uint16 creates a zap.Field for a uint16 value.
func Uint16(key string, value uint16) zap.Field {
	return zap.Uint16(key, value)
}

// Uint32 creates a zap.Field for a uint32 value.
func Uint32(key string, value uint32) zap.Field {
	return zap.Uint32(key, value)
}

// Uint64 creates a zap.Field for a uint64 value.
func Uint64(key string, value uint64) zap.Field {
	return zap.Uint64(key, value)
}

// Float32 creates a zap.Field for a float32 value.
func Float32(key string, value float32) zap.Field {
	return zap.Float32(key, value)
}

// Float64 creates a zap.Field for a float64 value.
func Float64(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}

// Duration creates a zap.Field for a time.Duration value.
func Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

// Error creates a zap.Field for an error value.
func Err(err error) zap.Field {
	return zap.Error(err)
}

// Any creates a zap.Field for a value of any type.
func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
