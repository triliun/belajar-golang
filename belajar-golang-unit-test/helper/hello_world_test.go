package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	Benchmarks := []struct {
		name    string
		request string
	}{{
		name:    "Varrel",
		request: "Varrel",
	},
		{
			name:    "Budi",
			request: "Budi",
		},
		{
			name:    "Joko",
			request: "Joko",
		},
	}

	for _, bench := range Benchmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Varrel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Varrel")
		}
	})
	b.Run("Al Jabaar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Al Jabaar")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Varrel")
	}
}

func BenchmarkHelloWorldAlJabaar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Al Jabaar")
	}
}

// Konsep Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct { // Slice Struct
		name     string
		request  string
		expected string
	}{
		{
			name:     "Varrel",
			request:  "Varrel",
			expected: "Hello Varrel",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Varrel", func(t *testing.T) {
		result := HelloWorld("Varrel")
		assert.Equal(t, "Hello Varrel", result, "Result must be 'Hello Varrel'")
	})
	t.Run("Al Jabaar", func(t *testing.T) {
		result := HelloWorld("Al Jabaar")
		assert.Equal(t, "Hello Al Jabaar", result, "Result must be 'Hello Al Jabaar'")

	})
}

// TestMain hanya bisa digunakan di satu package, TestMain bisa digunakan sebagai Before & After Test
func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	// After
	fmt.Println("AFTER UNIT TEST")
}
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" { // GOOS artinya GO Operating System
		t.Skip("Cannot run on Windows OS")
	}
	result := HelloWorld("Varrel")
	assert.Equal(t, "Hello Varrel", result, "Result must be 'Hello Varrel'")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Varrel")
	assert.Equal(t, "Hello Varrel", result, "Result must be 'Hello Varrel'") // Memanggil Fail()
	fmt.Println("TestHelloWorld with Assert Done")

}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Varrel")
	require.Equal(t, "Hello Varrel", result, "Result must be 'Hello Varrel'") // Memanggil FailNow()
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldVarrel(t *testing.T) { // Jangan sampai ada return value nya
	result := HelloWorld("Varrel")

	if result != "Hello Varrel" {
		// Error / Salah Respon nya
		// panic("Result is not Hello Varrel")
		// t.Fail() // Kode program dibawah nya tetap dieksekusi
		t.Error("Result must be 'Hello Varrel'") // sama seperti Fail() bedanya bisa di kasih argument
	}
	fmt.Println("TestHelloWorldVarrel Done")

}
func TestHelloWorldAljabaar(t *testing.T) { // Jangan sampai ada return value nya
	result := HelloWorld("Al Jabaar")

	if result != "Hello Al Jabaar" {
		// Error / Salah Respon nya
		// panic("Result is not Hello Al Jabaar")
		// t.FailNow() // Kode program dibawah nya tetap dieksekusi
		t.Fatal("Result must be 'Hello Al Jabaar'") // sama seperti FailNow() bedanya bisa di kasih argument
	}
	fmt.Println("TestHelloWorldAljabaar Done")
}
