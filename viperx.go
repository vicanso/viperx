// Copyright 2020 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package viperx

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type ViperX struct {
	*viper.Viper
	// ConfigType config type
	ConfigType string
}

func toENVKey(key string) string {
	arr := strings.Split(key, ".")
	for index, k := range arr {
		arr[index] = strings.ToUpper(k)
	}
	return strings.Join(arr, "_")
}

// New new viperx
func New(configType string) *ViperX {
	viperX := &ViperX{
		ConfigType: configType,
	}
	viperX.Viper = viper.New()
	viperX.SetConfigType(viperX.ConfigType)
	return viperX
}

// ReadConfig read config from reader
func (vx *ViperX) ReadConfig(readers ...io.Reader) error {
	size := len(readers)
	if size == 0 {
		return nil
	}
	for _, reader := range readers[0 : size-1] {
		v := viper.New()
		v.SetConfigType(vx.ConfigType)
		err := v.ReadConfig(reader)
		if err != nil {
			return err
		}
		for k, v := range v.AllSettings() {
			vx.SetDefault(k, v)
		}
	}
	return vx.Viper.ReadConfig(readers[size-1])
}

// GetIntDefault get value(int) from config,
// if value is 0, return the default value
func (vx *ViperX) GetIntDefault(key string, defaultValue int) int {
	v := vx.GetInt(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetInt32Default get value(int32) from config,
// if value is 0, return the default value
func (vx *ViperX) GetInt32Default(key string, defaultValue int32) int32 {
	v := vx.GetInt32(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetInt64Default get value(int64) from config,
// if value is 0, return the default value
func (vx *ViperX) GetInt64Default(key string, defaultValue int64) int64 {
	v := vx.GetInt64(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetUintDefault get value(uint) from config,
// if value is 0, return the default value
func (vx *ViperX) GetUintDefault(key string, defaultValue uint) uint {
	v := vx.GetUint(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetUint32Default get value(uint32) from config,
// if value is 0, return the default value
func (vx *ViperX) GetUint32Default(key string, defaultValue uint32) uint32 {
	v := vx.GetUint32(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetUint64Default get value(uint64) from config,
// if value is 0, return the default value
func (vx *ViperX) GetUint64Default(key string, defaultValue uint64) uint64 {
	v := vx.GetUint64(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetFloat64Default get value(float64) from config,
// if value is 0, return the default value
func (vx *ViperX) GetFloat64Default(key string, defaultValue float64) float64 {
	v := vx.GetFloat64(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetStringDefault get value(string) from config,
// if value is 0, return the default value
func (vx *ViperX) GetStringDefault(key, defaultValue string) string {
	v := vx.GetString(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetTimeDefault get value(time.Time) from config,
// if value is 0, return the default value
func (vx *ViperX) GetTimeDefault(key string, defaultValue time.Time) time.Time {
	v := vx.GetTime(key)
	if v.IsZero() {
		return defaultValue
	}
	return v
}

// GetDurationDefault get value(time.Duration) from config,
// if value is 0, return the default value
func (vx *ViperX) GetDurationDefault(key string, defaultValue time.Duration) time.Duration {
	v := vx.GetDuration(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetIntSliceDefault get value([]int) from config,
// it len(value) is 0, return the default value
func (vx *ViperX) GetIntSliceDefault(key string, defaultValue []int) []int {
	v := vx.GetIntSlice(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringSliceDefault get value([]string) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringSliceDefault(key string, defaultValue []string) []string {
	v := vx.GetStringSlice(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringMapDefault get value(map[string]interface{}) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringMapDefault(key string, defaultValue map[string]interface{}) map[string]interface{} {
	v := vx.GetStringMap(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringMapStringDefault get value(map[string]string) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringMapStringDefault(key string, defaultValue map[string]string) map[string]string {
	v := vx.GetStringMapString(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringMapStringSliceDefault get value(map[string][]string) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringMapStringSliceDefault(key string, defaultValue map[string][]string) map[string][]string {
	v := vx.GetStringMapStringSlice(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringFromENV returns the value from env first,
// then get from config if not exits.
// The key of env will be different of config, for example, the key `redis.uri` will be `REDIS_URI`.
func (vx *ViperX) GetStringFromENV(key string) string {
	value := os.Getenv(toENVKey(key))
	if len(value) != 0 {
		return value
	}
	return vx.GetString(key)
}

// GetStringFromENVDefault returns the value of `GetStringFromENV` if exits, otherwise returns the default value.
func (vx *ViperX) GetStringFromENVDefault(key, defaultValue string) string {
	value := vx.GetStringFromENV(key)
	if value != "" {
		return value
	}
	return defaultValue
}

// GetDurationFromENV returns the duration of key,
// it's the same as `GetStringFromENV` but returns duration.
func (vx *ViperX) GetDurationFromENV(key string) time.Duration {
	value := os.Getenv(toENVKey(key))
	if len(value) != 0 {
		v, _ := time.ParseDuration(value)
		return v
	}
	return vx.GetDuration(key)
}

// GetDurationFromENVDefault returns the value of `GetDurationFromENV` if exits, otherwise returns the default value
func (vx *ViperX) GetDurationFromENVDefault(key string, defaultValue time.Duration) time.Duration {
	value := vx.GetDurationFromENV(key)
	if value != 0 {
		return value
	}
	return defaultValue
}

// GetIntFromENV returns the int of key,
// it's the same as `GetStringFromENV` but returns int.
func (vx *ViperX) GetIntFromENV(key string) int {
	value := os.Getenv(toENVKey(key))
	if len(value) != 0 {
		v, _ := strconv.Atoi(value)
		return v
	}
	return vx.GetInt(key)
}

// GetIntFromENVDefault returns the value of `GetIntFromENV` if exists, otherwist returns the default value
func (vx *ViperX) GetIntFromENVDefault(key string, defaultValue int) int {
	value := vx.GetIntFromENV(key)
	if value != 0 {
		return value
	}
	return defaultValue
}

// GetBoolFromENV return the bool of key,
// get the value from env if exists, otherwist get the value from config
func (vx *ViperX) GetBoolFromENV(key string) bool {
	value := strings.ToLower(os.Getenv(toENVKey(key)))
	if len(value) != 0 {
		b, _ := strconv.ParseBool(value)
		return b
	}
	return vx.GetBool(key)
}

// GetStringSliceFromENV return string slice of key,
// get the value from env and split to slice by ',' if exists, otherwist get string slice from config
func (vx *ViperX) GetStringSliceFromENV(key string) []string {
	value := os.Getenv(toENVKey(key))
	if len(value) != 0 {
		return strings.Split(value, ",")
	}
	return vx.GetStringSlice(key)
}

// GetStringSliceFromENVDefault return string slice of key,
// if it is empty, return the default value
func (vx *ViperX) GetStringSliceFromENVDefault(key string, defaultValue []string) []string {
	values := vx.GetStringSliceFromENV(key)
	if len(values) != 0 {
		return values
	}
	return defaultValue
}
