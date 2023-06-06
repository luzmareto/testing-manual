package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
go test -v ./...   == untuk menjalankan test

go test -v == untuk melakukan semua test dalam satu modul/package

untuk test 1 func, terminal kita harus masuk ke dalam modul / package yang akan di test
go test -v -run nama func

SUBTEST :
Menjalankan lebih dari 1 test dari 1 func Testing

TEST MENGGUNAKAN TESTIFY :
assert.Equal == melakukan perbanding
assert.equal adalah fun Fail()
require.equal adalah fun FailNow()
- adalah expected atau output yang sudah kita tetapkan sebagai default
+ adalah output yang tidak sama dengan dengan expected atau gagal

TESTING DENGAN CARA MANUAL
pengecekan menggunkan fail() dan failnow()
t.Fail()  == jika ada kode yang gagal, kode selanjutnya tetap diproses walaupun outputnya tetap fail
t.FailNow() ==  jika ada kode yang gagal maka kode selanjutnya tidak di cek. namun test func selanjutnya akan ditest
arguments :
t.Error("pesan gagal") == akan memanggil func t.Fail()
t.Fatal("pesan gagal") == akan memanggil func t.FailNow()
*/

func TestSubTest(t *testing.T) { //go test -v -run=SubTest
	t.Run("Luz aja", func(t *testing.T) { //luz aja adalah nama SubTest
		result := HelloWorld("Luz")
		//t = testing, Hello Luz = expected, result = output, "" = message optional
		require.Equal(t, "Hello Luz", result, "Result must be Hello Luz")
	})
	t.Run("mareto doang", func(t *testing.T) { //mareto doang aja adalah nama SubTest
		result := HelloWorld("Mareto")
		//t = testing, Hello Luz = expected, result = output, "" = message optional
		require.Equal(t, "Hello Mareto", result, "Result must be Hello Mareto")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST") //mengeluarkan func yang sudah di test

	m.Run()

	//after
	fmt.Println("BEFORE UNIT TEST") //mengeluarkan hasil keseluruhan test fail

}

// TEST MENGGUNAKAN TESTIFY
func TestTestHelloWolrdRequire(t *testing.T) {
	result := HelloWorld("Luz")
	//t = testing, Hello Luz = expected, result = output, "" = message optional
	require.Equal(t, "Hello Luz", result, "Result must be Hello Luz")
	fmt.Println("TestHelloWorldAssert with Assert Done") //tidak akan melakukan println jika testnya gagal
}

// TEST MENGGUNAKAN TESTIFY
func TestTestHelloWolrdAssert(t *testing.T) {
	result := HelloWorld("Luz")
	//t = testing, Hello Luz = expected, result = output, "" = message optional
	assert.Equal(t, "Hello Luz", result, "Result must be Hello Luz")
	fmt.Println("TestHelloWorldAssert with Assert Done") //akan melakukan println walapun testnya fail atau ok
}

// TESTING DENGAN CARA MANUAL
func TestHelloWorld1(t *testing.T) { // test go test -v -run TestHelloWorld1
	result := HelloWorld("Luz")

	if result != "Hello Luz" {
		t.Error("Result must be Hello Luz") //output jika gagal
	}

	fmt.Println("TestHelloWorld1 done")
}

// TESTING DENGAN CARA MANUAL
func TestHelloWorld2(t *testing.T) { // test go test -v -run TestHelloWorld2
	result := HelloWorld("Mareto")

	if result != "Hello Mareto" {
		t.Fatal("Result must be Hello Mareto")
	}

	fmt.Println("TestHelloWorld2 done")
}
