package unsafe

import "testing"

func TestStringHeader(t *testing.T) {
	StringHeader()
}

func TestString2Bytes(t *testing.T) {
	String2Bytes()
}

func TestReadOnlyBytes(t *testing.T) {
	ReadOnlyBytes()
}

func TestStringIsReferenceType(t *testing.T) {
	StringIsReferenceType()
}

func TestBytes2String(t *testing.T) {
	Bytes2String()
}
