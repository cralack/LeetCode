package designtwitter

import "testing"

type Tweet struct {
	TweetId  int
	PostTime int
	Next     *Tweet
}
type User struct {
	UserId   int
	Followed map[int]*User
	Head     *Tweet
}

func (u *User) NewUser(uid int) {

}
func (u *User) Follow(uid int) {
	u.Followed[uid] = u
}

type Twitter struct {
	TweetList map[int]*Tweet
	UserList  map[int]*User
}

func Constructor() Twitter {
	return Twitter{}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {

}

func (t *Twitter) GetNewsFeed(userId int) []int {
	return []int{}
}

func (t *Twitter) Follow(followerId int, followeeId int) {

}

func (t *Twitter) Unfollow(followerId int, followeeId int) {

}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
func Test_design_twitter(t *testing.T) {

}
