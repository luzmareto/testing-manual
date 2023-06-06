package tabletest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTablleHelloWorld(t *testing.T) {
	//membuat slice table Test
	tests := []struct {
		name     string //nama test
		request  string //output yang kita masukan
		expected string //expected dari request dan func yang di test
	}{
		//membuat banyak test
		{
			name:     "liza",
			request:  "Luz",
			expected: "Hello Luz",
		},
		{
			name:     "Mareto",
			request:  "Mareto",
			expected: "Hello Mareto",
		},
		{
			name:     "Luz Mareto",
			request:  "Luz Mareto",
			expected: "Hello Luz Mareto",
		},
	}

	//melakukn testing pada slice tests
	for _, test := range tests {
		//melakukan testing pada file name atau nama test
		t.Run(test.name, func(t *testing.T) {
			//membuat var penampung untuk func dari file yang akan di test lalu memasukan parameternya dari test.request
			result := HelloWorld2(test.request)
			//melakukan perbandingan. jika seluruh expected valuenya sama seperti value maka hasilnya ok/pass
			require.Equal(t, test.expected, result)
		})
	}
}
