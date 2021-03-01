package main

import "unsafe"

func main() {
	var b bool
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var s string
	var p unsafe.Pointer
	var up uintptr

	println("bool: ", unsafe.Sizeof(b))
	println("int: ", unsafe.Sizeof(i))
	println("int8: ",unsafe.Sizeof(i8))
	println("int16: ",unsafe.Sizeof(i16))
	println("int32: ",unsafe.Sizeof(i32))
	println("int64: ",unsafe.Sizeof(i64))
	println("uint: ",unsafe.Sizeof(u))
	println("uint8: ",unsafe.Sizeof(u8))
	println("uint16: ",unsafe.Sizeof(u16))
	println("uint32: ",unsafe.Sizeof(u32))
	println("uint64: ",unsafe.Sizeof(u64))
	println("float32: ",unsafe.Sizeof(f32))
	println("float64: ",unsafe.Sizeof(f64))
	println("complex64: ",unsafe.Sizeof(c64))
	println("complex128: ",unsafe.Sizeof(c128))
	println("string: ",unsafe.Sizeof(s))
	println("pointer: ",unsafe.Sizeof(p))
	println("uintptr: ",unsafe.Sizeof(up))
	println("------------------------------------------------")
	println("bool: ",unsafe.Alignof(b))
	println("int: ",unsafe.Alignof(i))
	println("int8: ",unsafe.Alignof(i8))
	println("int16: ",unsafe.Alignof(i16))
	println("int32: ",unsafe.Alignof(i32))
	println("int64: ",unsafe.Alignof(i64))
	println("uint: ",unsafe.Alignof(u))
	println("uint8: ",unsafe.Alignof(u8))
	println("uint16: ",unsafe.Alignof(u16))
	println("uint32: ",unsafe.Alignof(u32))
	println("uint64: ",unsafe.Alignof(u64))
	println("float32: ",unsafe.Alignof(f32))
	println("float64: ",unsafe.Alignof(f64))
	println("complex64: ",unsafe.Alignof(c64))
	println("complex128: ",unsafe.Alignof(c128))
	println("string: ",unsafe.Alignof(s))
	println("pointer: ",unsafe.Alignof(p))
	println("uintptr: ",unsafe.Alignof(up))
}
