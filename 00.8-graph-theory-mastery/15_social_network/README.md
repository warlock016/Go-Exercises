# Exercise 15: Social Network Analysis (Capstone)

**Tier:** 4 (Mastery)
**Time:** 120 min
**Goal:** Apply multiple graph algorithms to a realistic problem

## Functions

```go
type SocialNetwork struct {
    friendships map[int][]int
    usernames   map[int]string
}

func (sn *SocialNetwork) SuggestFriends(userID int) []int  // Friends-of-friends
func (sn *SocialNetwork) FindInfluencers() []int  // High degree centrality
func (sn *SocialNetwork) DetectCommunities() [][]int  // Connected components
func (sn *SocialNetwork) SixDegreesOfSeparation(user1, user2 int) int  // Shortest path
func (sn *SocialNetwork) MutualFriends(user1, user2 int) []int
func (sn *SocialNetwork) ClosestCommonFriend(user1, user2 int) int
```

## Challenge

Build a social network analyzer combining:
- BFS (shortest path, degrees of separation)
- DFS (community detection)
- Graph properties (influencer detection)
- Set operations (mutual friends)

**Real-world application:** LinkedIn connections, Twitter followers, Facebook friends.

This is your chance to synthesize everything you've learned!
