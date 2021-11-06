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
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTestViper() *ViperX {
	vx := New("yml")
	defaultValue := `
bValue: true
iValue: 1
i32Value: 2
i64Value: 3
uiValue: 4
ui32Value: 5
ui64Value: 6
f64Value: 7.1
sValue: abcd
smValue:
  c: true
`
	value := `
tValue: 2020-09-06T01:51:16.749Z
dValue: 1ms
isValue:
- 1
- 2
ssValue:
- a
- b
smValue:
  a: 1
  b: def
smsValue:
  a: b
  b: c
smssValue:
  a:
  - d
  - e
  - f
envValue:
  uri: __abc__
  timeout: 3s
`

	err := vx.ReadConfig(
		bytes.NewBufferString(defaultValue),
		bytes.NewBufferString(value),
	)
	if err != nil {
		panic(err)
	}
	return vx
}

const (
	unknown = "unknown"
)

func TestGetBool(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.False(vx.GetBool(unknown))
	assert.True(vx.GetBool("bValue"))
	assert.True(vx.GetBool("smValue.c"))
}

func TestGetInt(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetInt(unknown))
	assert.Equal(1, vx.GetInt("iValue"))
	assert.Equal(2, vx.GetIntDefault(unknown, 2))
}

func TestGetInt32(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetInt32(unknown))
	assert.Equal(int32(2), vx.GetInt32("i32Value"))
	assert.Equal(int32(2), vx.GetInt32Default("i32Value", 3))
	assert.Equal(int32(3), vx.GetInt32Default(unknown, 3))
}

func TestGetInt64(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetInt64(unknown))
	assert.Equal(int64(3), vx.GetInt64("i64Value"))
	assert.Equal(int64(3), vx.GetInt64Default("i64Value", 4))
	assert.Equal(int64(4), vx.GetInt64Default(unknown, 4))
}

func TestGetUint(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetUint(unknown))
	assert.Equal(uint(4), vx.GetUint("uiValue"))
	assert.Equal(uint(4), vx.GetUintDefault("uiValue", 5))
	assert.Equal(uint(5), vx.GetUintDefault(unknown, 5))
}

func TestGetUint32(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetUint32(unknown))
	assert.Equal(uint32(5), vx.GetUint32("ui32Value"))
	assert.Equal(uint32(5), vx.GetUint32Default("ui32Value", 6))
	assert.Equal(uint32(6), vx.GetUint32Default(unknown, 6))
}

func TestGetUint64(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetUint64(unknown))
	assert.Equal(uint64(6), vx.GetUint64("ui64Value"))
	assert.Equal(uint64(6), vx.GetUint64Default("ui64Value", 7))
	assert.Equal(uint64(7), vx.GetUint64Default(unknown, 7))
}

func TestGetFloat64(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetFloat64(unknown))
	assert.Equal(7.1, vx.GetFloat64("f64Value"))
	assert.Equal(7.1, vx.GetFloat64Default("f64Value", 7.2))
	assert.Equal(7.2, vx.GetFloat64Default(unknown, 7.2))
}

func TestGetString(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetString(unknown))
	assert.Equal("abcd", vx.GetString("sValue"))
	assert.Equal("abcd", vx.GetStringDefault("sValue", "def"))
	assert.Equal("def", vx.GetStringDefault(unknown, "def"))
}

func TestGetTime(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetTime(unknown))
	assert.Equal("2020-09-06 01:51:16.749 +0000 UTC", vx.GetTime("tValue").UTC().String())
	now := time.Now()
	assert.Equal("2020-09-06 01:51:16.749 +0000 UTC", vx.GetTimeDefault("tValue", now).String())
	assert.Equal(now.String(), vx.GetTimeDefault(unknown, now).String())
}

func TestGetDuration(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetDuration(unknown))
	assert.Equal(time.Millisecond, vx.GetDuration("dValue"))
	assert.Equal(time.Millisecond, vx.GetDurationDefault("dValue", time.Microsecond))
	assert.Equal(time.Microsecond, vx.GetDurationDefault(unknown, time.Microsecond))
}

func TestGetIntSlice(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetIntSlice(unknown))
	assert.Equal([]int{1, 2}, vx.GetIntSlice("isValue"))
	assert.Equal([]int{1, 2}, vx.GetIntSliceDefault("isValue", []int{2}))
	assert.Equal([]int{2}, vx.GetIntSliceDefault(unknown, []int{2}))
}

func TestGetStringSlice(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetStringSlice(unknown))
	assert.Equal([]string{"a", "b"}, vx.GetStringSlice("ssValue"))
	assert.Equal([]string{"a", "b"}, vx.GetStringSliceDefault("ssValue", []string{"b"}))
	assert.Equal([]string{"b"}, vx.GetStringSliceDefault(unknown, []string{"b"}))
}

func TestGetStringMap(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetStringMap(unknown))
	assert.Equal(map[string]interface{}{
		"a": 1,
		"b": "def",
	}, vx.GetStringMap("smValue"))
	assert.Equal(map[string]interface{}{
		"a": 1,
		"b": "def",
	}, vx.GetStringMapDefault("smValue", map[string]interface{}{
		"a": 2,
	}))

	assert.Equal(map[string]interface{}{
		"a": 2,
	}, vx.GetStringMapDefault(unknown, map[string]interface{}{
		"a": 2,
	}))
}

func TestGetStringMapString(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetStringMapString(unknown))
	assert.Equal(map[string]string{
		"a": "b",
		"b": "c",
	}, vx.GetStringMapString("smsValue"))
	assert.Equal(map[string]string{
		"a": "b",
		"b": "c",
	}, vx.GetStringMapStringDefault("smsValue", map[string]string{
		"a": "c",
	}))
	assert.Equal(map[string]string{
		"a": "c",
	}, vx.GetStringMapStringDefault(unknown, map[string]string{
		"a": "c",
	}))
}

func TestGetStringMapStringSlice(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetStringMapStringSlice(unknown))

	assert.Equal(map[string][]string{
		"a": {
			"d",
			"e",
			"f",
		},
	}, vx.GetStringMapStringSlice("smssValue"))

	assert.Equal(map[string][]string{
		"a": {
			"d",
			"e",
			"f",
		},
	}, vx.GetStringMapStringSliceDefault("smssValue", map[string][]string{
		"d": {
			"a",
		},
	}))

	assert.Equal(map[string][]string{
		"d": {
			"a",
		},
	}, vx.GetStringMapStringSliceDefault(unknown, map[string][]string{
		"d": {
			"a",
		},
	}))
}

func TestGetStringFromENV(t *testing.T) {
	assert := assert.New(t)
	vx := newTestViper()
	assert.Empty(vx.GetStringFromENV(unknown))
	value := "__abc__"
	assert.Equal(value, vx.GetStringFromENV("envValue.uri"))
	// 不存在的取默认值
	assert.Equal("def", vx.GetStringFromENVDefault("envValue", "def"))

	// 设置env配置
	envValue := "__d__"
	err := os.Setenv(toENVKey("envValue.uri"), envValue)
	assert.Nil(err)
	assert.Equal(envValue, vx.GetStringFromENV("envValue.uri"))

	timeout := 3 * time.Second
	assert.Equal(timeout, vx.GetDurationFromENV("envValue.timeout"))
	// 不存在的取默认值
	assert.Equal(5*time.Second, vx.GetDurationFromENVDefault("envValue.timeout1", 5*time.Second))

	// 设置env配置
	err = os.Setenv(toENVKey("envValue.timeout"), "4s")
	assert.Nil(err)
	assert.Equal(4*time.Second, vx.GetDurationFromENV("envValue.timeout"))
}
