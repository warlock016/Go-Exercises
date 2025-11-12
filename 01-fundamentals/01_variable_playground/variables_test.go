package variables

import "testing"

func TestDeclareInteger(t *testing.T) {
	got := DeclareInteger()
	want := 42
	if got != want {
		t.Errorf("DeclareInteger() = %v, want %v", got, want)
	}
}

func TestDeclareFloat(t *testing.T) {
	got := DeclareFloat()
	want := 3.14
	if got != want {
		t.Errorf("DeclareFloat() = %v, want %v", got, want)
	}
}

func TestDeclareString(t *testing.T) {
	got := DeclareString()
	want := "Hello, Go!"
	if got != want {
		t.Errorf("DeclareString() = %q, want %q", got, want)
	}
}

func TestDeclareBoolean(t *testing.T) {
	got := DeclareBoolean()
	want := true
	if got != want {
		t.Errorf("DeclareBoolean() = %v, want %v", got, want)
	}
}

func TestDeclareMultiple(t *testing.T) {
	gotName, gotAge, gotHeight := DeclareMultiple()
	wantName, wantAge, wantHeight := "Alice", 25, 5.6

	if gotName != wantName {
		t.Errorf("DeclareMultiple() name = %q, want %q", gotName, wantName)
	}
	if gotAge != wantAge {
		t.Errorf("DeclareMultiple() age = %v, want %v", gotAge, wantAge)
	}
	if gotHeight != wantHeight {
		t.Errorf("DeclareMultiple() height = %v, want %v", gotHeight, wantHeight)
	}
}
