// Code generated by ydbgen; DO NOT EDIT.

package tests

import (
	"strconv"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/table"
)

var (
	_ = strconv.Itoa
	_ = ydb.StringValue
	_ = table.NewQueryParameters
)

func (c *ConvAssert) Scan(res *table.Result) (err error) {
	res.SeekItem("int8_int16")
	c.Int8Int16 = int16(res.OInt8())

	res.SeekItem("int32_int64")
	c.Int32Int64 = int64(res.OInt32())

	res.SeekItem("int16_int8")
	c.Int16Int8 = ydbConvI16ToI8(res.OInt16())

	res.SeekItem("uint64_int8")
	c.Uint64Int8 = ydbConvU64ToI8(res.OUint64())

	res.SeekItem("uint32_uint")
	c.Uint32Uint = uint(res.OUint32())

	res.SeekItem("int32_int")
	c.Int32Int = int(res.OInt32())

	res.SeekItem("int32_to_byte")
	c.Int32ToByte = ydbConvI32ToB(res.OInt32())

	return res.Err()
}

func ydbConvI16ToI8(x int16) int8 { 
	const (
		bits = 8
		mask = (1 << (bits - 1)) - 1
	)
	var abs uint64
	{
		v := int64(x)
		m := v >> 63
		abs = uint64(v ^ m - m)
	}
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatInt(int64(x), 10) +
				" (type int16) overflows int8",
		)
	}
	return int8(x)
}

func ydbConvI32ToB(x int32) byte { 
	if x < 0 {
		panic("ydbgen: convassert: conversion of negative int32 to byte")
	}
	const (
		bits = 8
		mask = (1 << (bits)) - 1
	)
	var abs uint64
	{
		v := int64(x)
		m := v >> 63
		abs = uint64(v ^ m - m)
	}
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatInt(int64(x), 10) +
				" (type int32) overflows byte",
		)
	}
	return byte(x)
}

func ydbConvU64ToI8(x uint64) int8 { 
	const (
		bits = 8
		mask = (1 << (bits - 1)) - 1
	)
	abs := uint64(x)
	if abs&mask != abs {
		panic(
			"ydbgen: convassert: " + strconv.FormatUint(uint64(x), 10) +
				" (type uint64) overflows int8",
		)
	}
	return int8(x)
}

