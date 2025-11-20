package social_network

type SocialNetwork struct {
	friendships map[int][]int
	usernames   map[int]string
}

func NewSocialNetwork() *SocialNetwork {
	return &SocialNetwork{
		friendships: make(map[int][]int),
		usernames:   make(map[int]string),
	}
}

func (sn *SocialNetwork) AddUser(id int, username string) {
	sn.usernames[id] = username
	if _, exists := sn.friendships[id]; !exists {
		sn.friendships[id] = []int{}
	}
}

func (sn *SocialNetwork) AddFriendship(user1, user2 int) {
	// TODO(human): Implement
}

func (sn *SocialNetwork) SuggestFriends(userID int) []int {
	// TODO(human): Implement (friends-of-friends not already friends)
	return nil
}

func (sn *SocialNetwork) FindInfluencers() []int {
	// TODO(human): Implement (users with highest degree)
	return nil
}

func (sn *SocialNetwork) DetectCommunities() [][]int {
	// TODO(human): Implement (connected components)
	return nil
}

func (sn *SocialNetwork) SixDegreesOfSeparation(user1, user2 int) int {
	// TODO(human): Implement (shortest path distance)
	return -1
}

func (sn *SocialNetwork) MutualFriends(user1, user2 int) []int {
	// TODO(human): Implement (intersection of friend lists)
	return nil
}

func (sn *SocialNetwork) ClosestCommonFriend(user1, user2 int) int {
	// TODO(human): Implement (mutual friend closest to user1)
	return -1
}
