# TypeConverter
---

![go](https://img.shields.io/github/go-mod/go-version/kisschou/TypeConverter?color=green&style=flat-square) ![commit](https://img.shields.io/github/last-commit/kisschou/TypeConverter?color=green&style=flat-square) ![license](https://img.shields.io/github/license/kisschou/TypeConverter?color=green&style=flat-square)

## 关于

这不过就是一个类型转换器. 因为太多的时候无法分出到底interface{}过来的数据的类型是哪个, 所以就想着弄了个这玩意.

<br />

## 用法

首先你得依靠自己的力量在你的项目上安装好`Golang环境`, 并且确认你即将导入包的`Golang项目`是可用的, 然后:

1. 在需要使用的脚本中引入包:

```go
import tc "ateshit.icu/kisschou/TypeConverter"
```

<br />

2. 在需要使用的地方使用:

```go
var t interface{} // 假定一个不知道类型的数据

tcImpl := tc.New(t) // 将未知类型的数据初始化脚本

str := tcImpl.String // 获取string类型的数据
i64 := tcImpl.Int64 // 获取int64的数据
// 可选的类型:
// .String => string
// .Byte => byte
// .Bytes => []byte
// .Bool => bool
// .Int => int
// .Int8 => int8
// .Int16 => int16
// .Int32 => int32
// .Int64 => int64
// .Uint => uint
// .Uint8 => uint8
// .Uint16 => uint16
// .Uint32 => uint32
// .Uint64 => uint64
// .Float32 => float32
// .Float64 => float64
// .Complex64 => complex64
// .Complex128 => complex128
// InterfaceSlice => []interface{}
// StringMapInterface => map[string]interface{}

tcImpl.Equal("hello") // 判断是否相等
tcImpl.Compare("EGT", 13) // 判断是否大于等于13
```

<br />

## 通用函数

总结一些我觉得可用的上的函数:

### 1. (*TypeConverter) Equal(compareData interface{}) bool

比对新传入的对象与当前的对象是否相等.

Example:
```go
import tc "ateshit.icu/kisschou/TypeConverter"

// 待比较对象
var d1, d2 interface{}

// 比较
if tc.New(d1).Equal(d2) {
    // true...
} else {
    // false...
}
```

<br />

### 2. (*TypeConverter) Compare(operator string, compareData interface{}) bool

传入新对象和运算符以与当前对象进行比较.

Example:
```go
import tc "ateshit.icu/kisschou/TypeConverter"

// 待比较对象
var d1, d2 interface{}

// 比较
if tc.New(d1).Compare("GT", d2) {
    // true...
} else {
    // false...
}
```

> 可选运算符: `EQ`(=), `NEQ`(!=), `EGT`(>=), `GT`(>), `LT`(<), `ELT`(<=)

<br />

### 3. SliceEqual(d1, d2 []interface{}) bool

传入两个切片, 比对他们是否相等.

Example:
```go
import tc "ateshit.icu/kisschou/TypeConverter"

// 待比较对象
var d1, d2 []interface{}

// 比较
if tc.SliceEqual(d1, d2) {
    // true...
} else {
    // false...
}
```

<br />

### 4. MapEqual(d1, d2 map[string]interface{}) bool

传入两个map, 比对他们是否相等.

Example:
```go
import tc "ateshit.icu/kisschou/TypeConverter"

// 待比较对象
var d1, d2 map[string]interface{}

// 比较
if tc.MapEqual(d1, d2) {
    // true...
} else {
    // false...
}
```

<br />

## Licence

Copyright (c) 2022-present Kisschou.
