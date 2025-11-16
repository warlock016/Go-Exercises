package custom_types

import (
	"math"
	"sort"
	"testing"
)

// Temperature tests

func TestTemperatureToFahrenheit(t *testing.T) {
	tests := []struct {
		name    string
		celsius Temperature
		want    float64
	}{
		{"freezing point", Temperature(0), 32.0},
		{"boiling point", Temperature(100), 212.0},
		{"room temperature", Temperature(20), 68.0},
		{"negative temperature", Temperature(-40), -40.0},
		{"body temperature", Temperature(37), 98.6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.celsius.ToFahrenheit()
			if math.Abs(got-tt.want) > 0.1 {
				t.Errorf("Temperature(%v).ToFahrenheit() = %v, want %v", tt.celsius, got, tt.want)
			}
		})
	}
}

func TestTemperatureToCelsius(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want float64
	}{
		{"zero", Temperature(0), 0.0},
		{"positive", Temperature(25.5), 25.5},
		{"negative", Temperature(-10.2), -10.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp.ToCelsius()
			if got != tt.want {
				t.Errorf("Temperature(%v).ToCelsius() = %v, want %v", tt.temp, got, tt.want)
			}
		})
	}
}

func TestTemperatureAdd(t *testing.T) {
	tests := []struct {
		name  string
		temp1 Temperature
		temp2 Temperature
		want  Temperature
	}{
		{"positive + positive", Temperature(20), Temperature(5), Temperature(25)},
		{"positive + negative", Temperature(20), Temperature(-5), Temperature(15)},
		{"zero + positive", Temperature(0), Temperature(10), Temperature(10)},
		{"negative + negative", Temperature(-10), Temperature(-5), Temperature(-15)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp1.Add(tt.temp2)
			if got != tt.want {
				t.Errorf("Temperature(%v).Add(%v) = %v, want %v", tt.temp1, tt.temp2, got, tt.want)
			}
		})
	}
}

func TestTemperatureIsFreezing(t *testing.T) {
	tests := []struct {
		name string
		temp Temperature
		want bool
	}{
		{"below freezing", Temperature(-5), true},
		{"at freezing", Temperature(0), true},
		{"above freezing", Temperature(5), false},
		{"room temperature", Temperature(20), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.temp.IsFreezing()
			if got != tt.want {
				t.Errorf("Temperature(%v).IsFreezing() = %v, want %v", tt.temp, got, tt.want)
			}
		})
	}
}

// Distance tests

func TestDistanceToKilometers(t *testing.T) {
	tests := []struct {
		name   string
		meters Distance
		want   float64
	}{
		{"1 km", Distance(1000), 1.0},
		{"5 km", Distance(5000), 5.0},
		{"500 meters", Distance(500), 0.5},
		{"marathon", Distance(42195), 42.195},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.meters.ToKilometers()
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("Distance(%d).ToKilometers() = %v, want %v", tt.meters, got, tt.want)
			}
		})
	}
}

func TestDistanceToMeters(t *testing.T) {
	tests := []struct {
		name string
		dist Distance
		want int
	}{
		{"1000 meters", Distance(1000), 1000},
		{"zero", Distance(0), 0},
		{"5000 meters", Distance(5000), 5000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dist.ToMeters()
			if got != tt.want {
				t.Errorf("Distance(%d).ToMeters() = %d, want %d", tt.dist, got, tt.want)
			}
		})
	}
}

func TestDistanceAdd(t *testing.T) {
	tests := []struct {
		name  string
		dist1 Distance
		dist2 Distance
		want  Distance
	}{
		{"1km + 2km", Distance(1000), Distance(2000), Distance(3000)},
		{"zero + 5km", Distance(0), Distance(5000), Distance(5000)},
		{"100m + 200m", Distance(100), Distance(200), Distance(300)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dist1.Add(tt.dist2)
			if got != tt.want {
				t.Errorf("Distance(%d).Add(%d) = %d, want %d", tt.dist1, tt.dist2, got, tt.want)
			}
		})
	}
}

func TestDistanceIsMarathon(t *testing.T) {
	tests := []struct {
		name string
		dist Distance
		want bool
	}{
		{"exactly marathon", Distance(42195), true},
		{"more than marathon", Distance(50000), true},
		{"less than marathon", Distance(40000), false},
		{"10k", Distance(10000), false},
		{"zero", Distance(0), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dist.IsMarathon()
			if got != tt.want {
				t.Errorf("Distance(%d).IsMarathon() = %v, want %v", tt.dist, got, tt.want)
			}
		})
	}
}

// StringSet tests

func TestNewStringSet(t *testing.T) {
	set := NewStringSet()

	if set == nil {
		t.Fatal("NewStringSet() returned nil")
	}

	if len(set) != 0 {
		t.Errorf("NewStringSet() size = %d, want 0", len(set))
	}
}

func TestStringSetAdd(t *testing.T) {
	set := NewStringSet()

	set.Add("apple")
	if !set.Contains("apple") {
		t.Error("StringSet.Add(\"apple\") - element not found")
	}

	if set.Size() != 1 {
		t.Errorf("StringSet.Add() size = %d, want 1", set.Size())
	}

	// Add duplicate
	set.Add("apple")
	if set.Size() != 1 {
		t.Errorf("StringSet.Add() duplicate - size = %d, want 1", set.Size())
	}

	set.Add("banana")
	if set.Size() != 2 {
		t.Errorf("StringSet.Add() multiple - size = %d, want 2", set.Size())
	}
}

func TestStringSetRemove(t *testing.T) {
	set := NewStringSet()
	set.Add("apple")
	set.Add("banana")

	set.Remove("apple")
	if set.Contains("apple") {
		t.Error("StringSet.Remove(\"apple\") - element still present")
	}

	if set.Size() != 1 {
		t.Errorf("StringSet.Remove() size = %d, want 1", set.Size())
	}

	// Remove non-existent
	set.Remove("cherry")
	if set.Size() != 1 {
		t.Errorf("StringSet.Remove() non-existent - size = %d, want 1", set.Size())
	}
}

func TestStringSetContains(t *testing.T) {
	set := NewStringSet()
	set.Add("apple")

	if !set.Contains("apple") {
		t.Error("StringSet.Contains(\"apple\") = false, want true")
	}

	if set.Contains("banana") {
		t.Error("StringSet.Contains(\"banana\") = true, want false")
	}
}

func TestStringSetSize(t *testing.T) {
	set := NewStringSet()

	if set.Size() != 0 {
		t.Errorf("Empty set size = %d, want 0", set.Size())
	}

	set.Add("a")
	set.Add("b")
	set.Add("c")

	if set.Size() != 3 {
		t.Errorf("Set with 3 elements size = %d, want 3", set.Size())
	}

	set.Remove("b")

	if set.Size() != 2 {
		t.Errorf("Set after removal size = %d, want 2", set.Size())
	}
}

func TestStringSetToSlice(t *testing.T) {
	set := NewStringSet()
	set.Add("apple")
	set.Add("banana")
	set.Add("cherry")

	slice := set.ToSlice()

	if len(slice) != 3 {
		t.Fatalf("StringSet.ToSlice() length = %d, want 3", len(slice))
	}

	// Sort for deterministic comparison
	sort.Strings(slice)
	expected := []string{"apple", "banana", "cherry"}

	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("StringSet.ToSlice()[%d] = %q, want %q", i, slice[i], val)
		}
	}
}

func TestStringSetToSliceEmpty(t *testing.T) {
	set := NewStringSet()
	slice := set.ToSlice()

	if slice == nil {
		t.Error("StringSet.ToSlice() empty set returned nil, want empty slice")
	}

	if len(slice) != 0 {
		t.Errorf("StringSet.ToSlice() empty set length = %d, want 0", len(slice))
	}
}

func TestStringSetUnion(t *testing.T) {
	set1 := NewStringSet()
	set1.Add("a")
	set1.Add("b")

	set2 := NewStringSet()
	set2.Add("b")
	set2.Add("c")

	union := set1.Union(set2)

	if union.Size() != 3 {
		t.Errorf("Union size = %d, want 3", union.Size())
	}

	expected := []string{"a", "b", "c"}
	for _, item := range expected {
		if !union.Contains(item) {
			t.Errorf("Union missing %q", item)
		}
	}

	// Original sets unchanged
	if set1.Size() != 2 {
		t.Error("Union modified original set1")
	}
	if set2.Size() != 2 {
		t.Error("Union modified original set2")
	}
}

func TestStringSetIntersection(t *testing.T) {
	set1 := NewStringSet()
	set1.Add("a")
	set1.Add("b")
	set1.Add("c")

	set2 := NewStringSet()
	set2.Add("b")
	set2.Add("c")
	set2.Add("d")

	inter := set1.Intersection(set2)

	if inter.Size() != 2 {
		t.Errorf("Intersection size = %d, want 2", inter.Size())
	}

	if !inter.Contains("b") || !inter.Contains("c") {
		t.Error("Intersection missing expected elements")
	}

	if inter.Contains("a") || inter.Contains("d") {
		t.Error("Intersection contains unexpected elements")
	}
}

func TestStringSetIntersectionDisjoint(t *testing.T) {
	set1 := NewStringSet()
	set1.Add("a")
	set1.Add("b")

	set2 := NewStringSet()
	set2.Add("c")
	set2.Add("d")

	inter := set1.Intersection(set2)

	if inter.Size() != 0 {
		t.Errorf("Intersection of disjoint sets size = %d, want 0", inter.Size())
	}
}

func TestStringSetIntersectionEmpty(t *testing.T) {
	set1 := NewStringSet()
	set2 := NewStringSet()

	inter := set1.Intersection(set2)

	if inter.Size() != 0 {
		t.Errorf("Intersection of empty sets size = %d, want 0", inter.Size())
	}
}
