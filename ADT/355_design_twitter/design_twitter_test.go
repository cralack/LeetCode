package designtwitter

import (
	"container/heap"
	"testing"
)

type Tweet struct {
	TweetId  int
	PostTime int
	Next     *Tweet
}
type PQ []*Tweet

func (p PQ) Len() int           { return len(p) }
func (p PQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PQ) Less(i, j int) bool { return p[i].PostTime > p[j].PostTime }
func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.(*Tweet))
}
func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

type User struct {
	UserId   int
	Followed map[int]*User
	Posted   *Tweet
}

type Twitter struct {
	TimeStamp int
	UserList  map[int]*User
}

func Constructor() Twitter {
	return Twitter{
		TimeStamp: 0,
		UserList:  make(map[int]*User),
	}
}

func (t *Twitter) NewUser(uid int) *User {
	new := &User{
		UserId:   uid,
		Followed: make(map[int]*User),
		Posted:   nil,
	}
	new.Followed[uid] = new
	return new
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	//发推用户不存在就新建个用户
	if _, ok := t.UserList[userId]; !ok {
		t.UserList[userId] = t.NewUser(userId)
	}
	//loc user
	user := t.UserList[userId]
	//new tweet
	tmp_tweet := &Tweet{
		TweetId:  tweetId,
		PostTime: t.TimeStamp,
	}
	tmp_tweet.Next = user.Posted
	user.Posted = tmp_tweet
	t.TimeStamp++
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	//刷推用户不存在
	if _, ok := t.UserList[userId]; !ok {
		return res
	}
	//loc user
	user := t.UserList[userId]

	//检索当前用户新闻推送中最近 10 条推文的 ID
	pq := &PQ{}
	for _, fol := range user.Followed {
		if fol.Posted == nil {
			continue
		}
		heap.Push(pq, fol.Posted)
	}
	for pq.Len() > 0 {
		if len(res) >= 10 {
			break
		}
		tmp := heap.Pop(pq).(*Tweet)
		res = append(res, tmp.TweetId)
		if tmp.Next != nil {
			heap.Push(pq, tmp.Next)
		}
	}
	return res
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	//check user1
	if _, ok := t.UserList[followerId]; !ok {
		t.UserList[followerId] = t.NewUser(followerId)
	}
	follower := t.UserList[followerId]
	//check user2
	if _, ok := t.UserList[followeeId]; !ok {
		t.UserList[followeeId] = t.NewUser(followeeId)
	}
	followee := t.UserList[followeeId]

	follower.Followed[followeeId] = followee
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	user := t.UserList[followerId]
	delete(user.Followed, followeeId)

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
	// c1 := []string{"Twitter", "postTweet", "getNewsFeed",
	// 	"follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"}
	// c2 := [][]int{{-1}, {1, 5}, {1},
	// // 	{1, 2}, {2, 6}, {1}, {1, 2}, {1}}
	// c1 := []string{"Twitter", "postTweet", "getNewsFeed", "follow",
	// 	"getNewsFeed", "unfollow", "getNewsFeed"}
	// c2 := [][]int{{-1}, {1, 1}, {1}, {2, 1},
	// 	{2}, {2, 1}, {2}}
	// c1 := []string{"Twitter", "postTweet", "getNewsFeed", "follow", "getNewsFeed", "unfollow", "getNewsFeed"}
	// c2 := [][]int{{-1}, {1, 1}, {1}, {2, 1}, {2}, {2, 1}, {2}}
	// c1 := []string{"Twitter", "postTweet", "postTweet", "unfollow", "follow", "getNewsFeed"}
	// c2 := [][]int{{-1}, {1, 4}, {2, 5}, {1, 2}, {1, 2}, {1}}
	c1 := []string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed"}
	c2 := [][]int{{-1}, {1, 5}, {1, 3}, {1, 101}, {1, 13}, {1, 10}, {1, 2}, {1, 94}, {1, 505}, {1, 333}, {1, 22}, {1, 11}, {1}}
	var twi Twitter
	for idx, command := range c1 {
		switch command {
		case "Twitter":
			twi = Constructor()
		case "postTweet":
			twi.PostTweet(c2[idx][0], c2[idx][1])
		case "getNewsFeed":
			t.Log(twi.GetNewsFeed(c2[idx][0]))
		case "follow":
			twi.Follow(c2[idx][0], c2[idx][1])
		case "unfollow":
			twi.Unfollow(c2[idx][0], c2[idx][1])
		}
	}
}
