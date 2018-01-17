package cli

import (
	"testing"
)

var testpair = []struct {
	in  string
	out string
}{
	//{"qwe", ""},
	//{"sad123", "[]"},
	//{"4-2", "[]"},
	{"1-10", "[1 2 3 4 5 6 7 8 9 10]"},
	{"6-3,4-6", "[- 1 3 - 4 5]"},
	//{"фыв", ""},
}

func TestParcePortRange(t *testing.T) {
	//t.Error("test")
	//fmt.Print("test\n")
	for _, pair := range testpair {
		//fmt.Printf("%v\n", pair.in)
		s, err := ParcePortRange(pair.in)
		//fmt.Printf("%v\n", s)
		if err != nil {
			t.Error("\tFor String: ", pair.in,
				"\n\t\tGot string:", s,
				"\n\t\tError: ", err)
		}
	}

}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, pair := range testpair {
			//fmt.Printf("%v\n", pair.in)
			s, err := ParcePortRange(pair.in)
			//fmt.Printf("%v\n", s)
			if err != nil {
				b.Log("\tFor String: ", pair.in,
					"\n\t\tGot string:", s,
					"\n\t\tError: ", err)
			}
		}
	}

}
