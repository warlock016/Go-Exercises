package custom_types

// TODO(human): Define Temperature as a custom type based on float64

type Temperature float64

// ToFahrenheit converts Celsius to Fahrenheit
func (t Temperature) ToFahrenheit() float64 {
	// TODO(human): Apply conversion formula
	f := (t * 9 / 5) + 32
	return float64(f)
}

// ToCelsius returns the temperature in Celsius (identity function)
func (t Temperature) ToCelsius() float64 {
	// TODO(human): Return as float64
	// c := (t - 32) * 5 / 9
	return float64(t)
}

// Add adds two temperatures
func (t Temperature) Add(other Temperature) Temperature {
	// TODO(human): Add temperatures
	return t + other
}

// IsFreezing returns true if temperature is at or below 0°C
func (t Temperature) IsFreezing() bool {
	// TODO(human): Check if at or below freezing
	return t <= 0
}

// TODO(human): Define Distance as a custom type based on int
type Distance int

// ToKilometers converts meters to kilometers
func (d Distance) ToKilometers() float64 {
	// TODO(human): Apply conversion
	return float64(d) / 1000
}

// ToMeters returns the distance in meters (identity function)
func (d Distance) ToMeters() int {
	// TODO(human): Return as int
	return int(d)
}

// Add adds two distances
func (d Distance) Add(other Distance) Distance {
	// TODO(human): Add distances
	return d + other
}

// IsMarathon returns true if distance is at least 42195 meters (marathon distance)
func (d Distance) IsMarathon() bool {
	// TODO(human): Check if at least marathon distance
	return d >= 42195
}

// TODO(human): Define StringSet as a custom type for a set of unique strings
type StringSet map[string]struct{}

// NewStringSet creates a new empty StringSet
func NewStringSet() StringSet {
	// TODO(human): Initialize empty set
	return make(StringSet)
}

// Add adds a string to the set
func (s StringSet) Add(item string) {
	// TODO(human): Add item to set
	s[item] = struct{}{}
}

// Remove removes a string from the set
func (s StringSet) Remove(item string) {
	// TODO(human): Remove item from set
	delete(s, item)
}

// Contains returns true if the string is in the set
func (s StringSet) Contains(item string) bool {
	// TODO(human): Check if item exists in set
	_, ok := s[item]
	return ok
}

// Size returns the number of elements in the set
func (s StringSet) Size() int {
	// TODO(human): Return set size
	return len(s)
}

// ToSlice returns all elements as a slice
func (s StringSet) ToSlice() []string {
	// TODO(human): Convert set to slice
	result := []string{}

	for k := range s {
		result = append(result, k)
	}
	return result
}

// Union returns a new set containing all elements from both sets
func (s StringSet) Union(other StringSet) StringSet {
	// TODO(human): Combine both sets
	result := StringSet{}

	for k := range s {
		result.Add(k)
	}

	for k := range other {
		result.Add(k)
	}

	return result
}

// Intersection returns a new set containing only elements in both sets
func (s StringSet) Intersection(other StringSet) StringSet {
	// TODO(human): Find common elements

	result := StringSet{}

	for k := range s {
		if other.Contains(k) {
			result.Add(k)
		}
	}

	return result
}
