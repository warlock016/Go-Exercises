package implicit_satisfaction

import "testing"

func TestDogSpeak(t *testing.T) {
	dog := Dog{}
	got := dog.Speak()
	want := "Woof!"

	if got != want {
		t.Errorf("Dog.Speak() = %q, want %q", got, want)
	}
}

func TestCatSpeak(t *testing.T) {
	cat := Cat{}
	got := cat.Speak()
	want := "Meow!"

	if got != want {
		t.Errorf("Cat.Speak() = %q, want %q", got, want)
	}
}

func TestRobotSpeak(t *testing.T) {
	robot := Robot{}
	got := robot.Speak()
	want := "Beep boop!"

	if got != want {
		t.Errorf("Robot.Speak() = %q, want %q", got, want)
	}
}

func TestMakeSpeak(t *testing.T) {
	tests := []struct {
		name    string
		speaker Speaker
		want    string
	}{
		{"Dog speaks", Dog{}, "Woof!"},
		{"Cat speaks", Cat{}, "Meow!"},
		{"Robot speaks", Robot{}, "Beep boop!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeSpeak(tt.speaker)
			if got != tt.want {
				t.Errorf("MakeSpeak(%T) = %q, want %q", tt.speaker, got, tt.want)
			}
		})
	}
}

func TestInterfaceVariable(t *testing.T) {
	// Test that we can assign different types to Speaker variable
	var s Speaker

	s = Dog{}
	if s.Speak() != "Woof!" {
		t.Errorf("Speaker variable with Dog: got %q, want %q", s.Speak(), "Woof!")
	}

	s = Cat{}
	if s.Speak() != "Meow!" {
		t.Errorf("Speaker variable with Cat: got %q, want %q", s.Speak(), "Meow!")
	}

	s = Robot{}
	if s.Speak() != "Beep boop!" {
		t.Errorf("Speaker variable with Robot: got %q, want %q", s.Speak(), "Beep boop!")
	}
}

func TestPolymorphicSlice(t *testing.T) {
	// All speakers in one slice!
	speakers := []Speaker{
		Dog{},
		Cat{},
		Robot{},
		Dog{},
		Cat{},
	}

	wants := []string{"Woof!", "Meow!", "Beep boop!", "Woof!", "Meow!"}

	for i, speaker := range speakers {
		got := speaker.Speak()
		if got != wants[i] {
			t.Errorf("speakers[%d].Speak() = %q, want %q", i, got, wants[i])
		}
	}
}

// This test demonstrates implicit satisfaction
func TestImplicitSatisfaction(t *testing.T) {
	// These assignments only work if the types satisfy Speaker
	// No explicit "implements" declaration needed!
	var _ Speaker = Dog{}
	var _ Speaker = Cat{}
	var _ Speaker = Robot{}

	t.Log("✓ All types implicitly satisfy Speaker interface")
}
