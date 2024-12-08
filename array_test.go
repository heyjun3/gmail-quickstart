package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkArray(b *testing.B) {
	base := []string{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, fmt.Sprintf("No: %d", i))
	}
}

func BenchmarkArray_AllocateOnceLen(b *testing.B) {
	base := make([]string, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("No: %d", i)
	}
}

func BenchmarkArray_AllocateOnceCap(b *testing.B) {
	base := make([]string, 0, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, fmt.Sprintf("No: %d", i))
	}

}

func Ptr[T any](v T) *T {
	return &v
}

func BenchmarkArrayPointer_AllocateOnceLen(b *testing.B) {
	base := make([]*string, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base[i] = Ptr(fmt.Sprintf("No: %d", i))
	}
}

func BenchmarkArrayPointer_AllocateOnceCap(b *testing.B) {
	base := make([]*string, 0, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, Ptr(fmt.Sprintf("No: %d", i)))
	}
}

func BenchmarkArray_AllocateOnceCapAndItoa(b *testing.B) {
	base := make([]string, 0, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, "No: "+strconv.Itoa(i))
	}
}

func BenchmarkArray_AllocateOnceCapAndFormatInt(b *testing.B) {
	base := make([]string, 0, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, "No: "+strconv.FormatInt(int64(i), 10))
	}
}

type User struct {
	name    string
	age     int
	address string
}

func BenchmarkUserArray(b *testing.B) {
	base := []User{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, User{
			name:    "name " + strconv.Itoa(i),
			age:     i,
			address: "address",
		})
	}
}

func BenchmarkUserArray_AllocateOnceLen(b *testing.B) {
	base := make([]User, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base[i] = User{
			name:    "name " + strconv.Itoa(i),
			age:     i,
			address: "address",
		}
	}
}

func BenchmarkUserArray_AllocateOnceCap(b *testing.B) {
	base := make([]User, 0, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, User{
			name:    "name " + strconv.Itoa(i),
			age:     i,
			address: "address",
		})
	}
}

func BenchmarkUserArrayPointer_AllocateOnceLen(b *testing.B) {
	base := make([]*User, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base[i] = &User{
			name:    "name " + strconv.Itoa(i),
			age:     i,
			address: "address",
		}
	}
}

func BenchmarkUserArrayPointer_AllocateOnceCap(b *testing.B) {
	base := make([]*User, 0, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, &User{
			name:    "name " + strconv.Itoa(i),
			age:     i,
			address: "address",
		})
	}
}
