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
	"time"

	"github.com/spf13/viper"
)

type ViperX struct {
	// ins viper instance
	ins *viper.Viper
	// ConfigType config type
	ConfigType string
}

// New new viperx
func New(configType string) *ViperX {
	viperX := &ViperX{
		ins:        viper.New(),
		ConfigType: configType,
	}
	viperX.ins.SetConfigType(configType)
	return viperX
}

// ReadConfig read config from reader
func (vx *ViperX) ReadConfig(readers ...io.Reader) error {
	for _, reader := range readers {
		v := viper.New()
		v.SetConfigType(vx.ConfigType)
		err := v.ReadConfig(reader)
		if err != nil {
			return err
		}
		for k, v := range v.AllSettings() {
			vx.ins.SetDefault(k, v)
		}
	}
	return nil
}

// GetBool get value(bool) from config
func (vx *ViperX) GetBool(key string) bool {
	return vx.ins.GetBool(key)
}

// GetInt get value(int) from config
func (vx *ViperX) GetInt(key string) int {
	return vx.ins.GetInt(key)
}

// GetIntDefault get value(int) from config,
// if value is 0, return the default value
func (vx *ViperX) GetIntDefault(key string, defaultValue int) int {
	v := vx.ins.GetInt(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetInt32 get value(int32) from config
func (vx *ViperX) GetInt32(key string) int32 {
	return vx.ins.GetInt32(key)
}

// GetInt32Default get value(int32) from config,
// if value is 0, return the default value
func (vx *ViperX) GetInt32Default(key string, defaultValue int32) int32 {
	v := vx.ins.GetInt32(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetInt64 get value(int64) from config
func (vx *ViperX) GetInt64(key string) int64 {
	return vx.ins.GetInt64(key)
}

// GetInt64Default get value(int64) from config,
// if value is 0, return the default value
func (vx *ViperX) GetInt64Default(key string, defaultValue int64) int64 {
	v := vx.ins.GetInt64(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetUint get value(uint) from config
func (vx *ViperX) GetUint(key string) uint {
	return vx.ins.GetUint(key)
}

// GetUintDefault get value(uint) from config,
// if value is 0, return the default value
func (vx *ViperX) GetUintDefault(key string, defaultValue uint) uint {
	v := vx.ins.GetUint(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetUint32 get value(uint32) from config
func (vx *ViperX) GetUint32(key string) uint32 {
	return vx.ins.GetUint32(key)
}

// GetUint32Default get value(uint32) from config,
// if value is 0, return the default value
func (vx *ViperX) GetUint32Default(key string, defaultValue uint32) uint32 {
	v := vx.ins.GetUint32(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetUint64 get value(uint64) from config
func (vx *ViperX) GetUint64(key string) uint64 {
	return vx.ins.GetUint64(key)
}

// GetUint64Default get value(uint64) from config,
// if value is 0, return the default value
func (vx *ViperX) GetUint64Default(key string, defaultValue uint64) uint64 {
	v := vx.ins.GetUint64(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetFloat64 get value(float64) from config
func (vx *ViperX) GetFloat64(key string) float64 {
	return vx.ins.GetFloat64(key)
}

// GetFloat64Default get value(float64) from config,
// if value is 0, return the default value
func (vx *ViperX) GetFloat64Default(key string, defaultValue float64) float64 {
	v := vx.ins.GetFloat64(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetTime get value(time.Time) from config
func (vx *ViperX) GetTime(key string) time.Time {
	return vx.ins.GetTime(key)
}

// GetTimeDefault get value(time.Time) from config,
// if value is 0, return the default value
func (vx *ViperX) GetTimeDefault(key string, defaultValue time.Time) time.Time {
	v := vx.ins.GetTime(key)
	if v.IsZero() {
		return defaultValue
	}
	return v
}

// GetDuration get value(time.Duration) from config
func (vx *ViperX) GetDuration(key string) time.Duration {
	return vx.ins.GetDuration(key)
}

// GetDurationDefault get value(time.Duration) from config,
// if value is 0, return the default value
func (vx *ViperX) GetDurationDefault(key string, defaultValue time.Duration) time.Duration {
	v := vx.ins.GetDuration(key)
	if v == 0 {
		return defaultValue
	}
	return v
}

// GetIntSlice get value([]int) from config
func (vx *ViperX) GetIntSlice(key string) []int {
	return vx.ins.GetIntSlice(key)
}

// GetIntSliceDefault get value([]int) from config,
// it len(value) is 0, return the default value
func (vx *ViperX) GetIntSliceDefault(key string, defaultValue []int) []int {
	v := vx.ins.GetIntSlice(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringSlice get value([]string) from config
func (vx *ViperX) GetStringSlice(key string) []string {
	return vx.ins.GetStringSlice(key)
}

// GetStringSliceDefault get value([]string) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringSliceDefault(key string, defaultValue []string) []string {
	v := vx.ins.GetStringSlice(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringMap get value(map[string]interface{}) from config
func (vx *ViperX) GetStringMap(key string) map[string]interface{} {
	return vx.ins.GetStringMap(key)
}

// GetStringMapDefault get value(map[string]interface{}) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringMapDefault(key string, defaultValue map[string]interface{}) map[string]interface{} {
	v := vx.ins.GetStringMap(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringMapString get value(map[string]string) from config
func (vx *ViperX) GetStringMapString(key string) map[string]string {
	return vx.ins.GetStringMapString(key)
}

// GetStringMapStringDefault get value(map[string]string) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringMapStringDefault(key string, defaultValue map[string]string) map[string]string {
	v := vx.ins.GetStringMapString(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringMapStringSlice get value(map[string][]string) from config
func (vx *ViperX) GetStringMapStringSlice(key string) map[string][]string {
	return vx.ins.GetStringMapStringSlice(key)
}

// GetStringMapStringSliceDefault get value(map[string][]string) from config,
// if len(value) is 0, return the default value
func (vx *ViperX) GetStringMapStringSliceDefault(key string, defaultValue map[string][]string) map[string][]string {
	v := vx.ins.GetStringMapStringSlice(key)
	if len(v) == 0 {
		return defaultValue
	}
	return v
}

// GetStringFromENV get string from env, it will get the value from config by key,
// then use the value as key, get result from env,
// if len(value) is 0, then return the config's value, otherwise return the env's value
func (vx *ViperX) GetStringFromENV(key string) string {
	v := vx.ins.GetString(key)
	value := os.Getenv(v)
	if len(value) != 0 {
		return value
	}
	return v
}
