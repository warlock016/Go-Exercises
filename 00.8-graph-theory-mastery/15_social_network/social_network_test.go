package social_network

import "testing"

func TestSocialNetwork(t *testing.T) {
	sn := NewSocialNetwork()
	sn.AddUser(0, "Alice")
	sn.AddUser(1, "Bob")
	sn.AddUser(2, "Charlie")
	sn.AddUser(3, "Dave")

	sn.AddFriendship(0, 1)  // Alice-Bob
	sn.AddFriendship(1, 2)  // Bob-Charlie
	sn.AddFriendship(2, 3)  // Charlie-Dave

	// Test six degrees of separation
	degrees := sn.SixDegreesOfSeparation(0, 3)
	if degrees != 3 {
		t.Errorf("Expected 3 degrees of separation, got %d", degrees)
	}
}

func TestSuggestFriends(t *testing.T) {
	sn := NewSocialNetwork()
	sn.AddUser(0, "Alice")
	sn.AddUser(1, "Bob")
	sn.AddUser(2, "Charlie")

	sn.AddFriendship(0, 1)
	sn.AddFriendship(1, 2)

	suggestions := sn.SuggestFriends(0)  // Should suggest Charlie (friend of Bob)

	if len(suggestions) == 0 {
		t.Error("Expected friend suggestions")
	}
}
