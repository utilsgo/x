package ptr

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPtr(t *testing.T) {
	b := Ptr(true)
	NewWithT(t).Expect(b).To(Equal(Bool(true)))
}

func TestBool(t *testing.T) {
	NewWithT(t).Expect(bool(true)).To(Equal(*Bool(true)))
	NewWithT(t).Expect(bool(false)).To(Equal(*Bool(false)))
}

func TestInt(t *testing.T) {
	NewWithT(t).Expect(int(1)).To(Equal(*Int(1)))
}

func TestInt8(t *testing.T) {
	NewWithT(t).Expect(int8(1)).To(Equal(*Int8(1)))
}

func TestInt16(t *testing.T) {
	NewWithT(t).Expect(int16(1)).To(Equal(*Int16(1)))
}

func TestInt32(t *testing.T) {
	NewWithT(t).Expect(int32(1)).To(Equal(*Int32(1)))
}

func TestInt64(t *testing.T) {
	NewWithT(t).Expect(int64(1)).To(Equal(*Int64(1)))
}

func TestUint(t *testing.T) {
	NewWithT(t).Expect(uint(1)).To(Equal(*Uint(1)))
}

func TestUint8(t *testing.T) {
	NewWithT(t).Expect(uint8(1)).To(Equal(*Uint8(1)))
}

func TestUint16(t *testing.T) {
	NewWithT(t).Expect(uint16(1)).To(Equal(*Uint16(1)))
}

func TestUint32(t *testing.T) {
	NewWithT(t).Expect(uint32(1)).To(Equal(*Uint32(1)))
}

func TestUint64(t *testing.T) {
	NewWithT(t).Expect(uint64(1)).To(Equal(*Uint64(1)))
}

func TestUintptr(t *testing.T) {
	NewWithT(t).Expect(uintptr(1)).To(Equal(*Uintptr(1)))
}

func TestFloat32(t *testing.T) {
	NewWithT(t).Expect(float32(1)).To(Equal(*Float32(1)))
}

func TestFloat64(t *testing.T) {
	NewWithT(t).Expect(float64(1)).To(Equal(*Float64(1)))
}

func TestComplex64(t *testing.T) {
	NewWithT(t).Expect(complex64(1)).To(Equal(*Complex64(1)))
}

func TestComplex128(t *testing.T) {
	NewWithT(t).Expect(complex128(1)).To(Equal(*Complex128(1)))
}

func TestString(t *testing.T) {
	NewWithT(t).Expect(string("string")).To(Equal(*String("string")))
}

func TestByte(t *testing.T) {
	NewWithT(t).Expect(byte([]uint8{
		98,
		121,
		116,
		101,
		115,
	}[0])).To(Equal(*Byte([]uint8{
		98,
		121,
		116,
		101,
		115,
	}[0])))
}

func TestRune(t *testing.T) {
	NewWithT(t).Expect(rune('r')).To(Equal(*Rune('r')))
}
