# viperx

[![license](https://img.shields.io/badge/https://img.shields.io/badge/license-apache-blue.svg)](https://github.com/vicanso/viperx/blob/master/LICENSE)
[![Build Status](https://github.com/vicanso/viperx/workflows/Test/badge.svg)](https://github.com/vicanso/viperx/actions)

增强了viper的配置获取方式，允许获取时使用默认值。允许指定多个配置文件源，如果后面的配置源存在相同的配置则覆盖已存在的配置，适用于将配置区分为默认配置与个性化配置的场景

```bash
	vx := NewViperX("yml")
	defaultValue := `
bValue: true
iValue: 1
i32Value: 2
i64Value: 3
uiValue: 4
ui32Value: 5
ui64Value: 6
f64Value: 7.1
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
envValue: __abc__
`

	_ = vx.ReadConfig(
		bytes.NewBufferString(defaultValue),
		bytes.NewBufferString(value),
	)
	fmt.Println(vx.GetInt("iValue"))
```