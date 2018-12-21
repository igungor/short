package short

import "testing"

func TestEncode(t *testing.T) {
	tt := []struct {
		input    int
		expected string
	}{
		{0, "mmmmm"},
		{1, "867nv"},
		{10, "csqsc"},
		{100, "6d7gw"},
		{1000, "ndrsn"},
		{10000, "mhz8p"},
		{100000, "mpxf9"},
		{1000000, "m2q43"},
		{10000000, "mjjcq"},
		{100000000, "63uqz6"},
		{1000000000, "njzupqe"},
	}

	for i, tc := range tt {
		if got, expected := Encode(tc.input), tc.expected; got != expected {
			t.Fatalf("test case %v. got: %v, expected: %v", i, got, expected)
		}
	}
}

func TestDecode(t *testing.T) {
	tt := []struct {
		input    string
		expected int
	}{
		{"mmmmm", 0},
		{"867nv", 1},
		{"csqsc", 10},
		{"6d7gw", 100},
		{"ndrsn", 1000},
		{"mhz8p", 10000},
		{"mpxf9", 100000},
		{"m2q43", 1000000},
		{"mjjcq", 10000000},
		{"63uqz6", 100000000},
		{"njzupqe", 1000000000},
	}

	for i, tc := range tt {
		if got, expected := Decode(tc.input), tc.expected; got != expected {
			t.Fatalf("test case %v. got: %v, expected: %v", i, got, expected)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	const count = 1e7

	result := make(map[int]string)
	for i := 0; i < count; i++ {
		result[i] = Encode(i)
	}

	for k, v := range result {
		if got, expected := Decode(v), k; got != expected {
			t.Fatalf("Decode failed for %q. got: %v, expected: %v", v, got, expected)
		}
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := Encode(i)
		_ = Decode(x)
	}
}
