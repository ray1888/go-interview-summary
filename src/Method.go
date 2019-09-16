package interview

import "fmt"

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	// Body exactly the same as the Append function defined above.
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

func (p *ByteSlice) Append2(data []byte) {
	slice := *p
	// Body as above, without the return.
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Body as above, without the return.
	l := len(slice)
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
	*p = slice
	return len(data), nil
}

func MethodRecevier() {
	var b ByteSlice
	b = b.Append([]byte{1, 2, 3})
	fmt.Printf("byteSlice 1 is %v", b)
	b.Write([]byte{7, 8, 9})
	fmt.Printf("byteSlice 2 is %v", b)
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	b.Append2([]byte{4, 5, 6})
	fmt.Printf("byteSlice 3 is %v", b)
}
