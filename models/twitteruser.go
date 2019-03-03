package models

import (
	"github.com/dghubble/go-twitter/twitter"
)

type TwitterUser twitter.User

// Get user data from twitter user's data
func (t *TwitterUser) ToUser() User {
	return User{
		Name:            t.Name,
		Email:           t.Email,
		FavouritesCount: t.FavouritesCount,
		Following:       t.Following,
		FollowersCount:  t.FollowersCount,
		FriendsCount:    t.FriendsCount,
		StatusesCount:   t.StatusesCount,
	}
}

// Sync new twitter data
func (u *User) SyncTwitterData(t *TwitterUser) {
	u.Name = t.Name
	u.FavouritesCount = t.FavouritesCount
	u.Following = t.Following
	u.FollowersCount = t.FollowersCount
	u.StatusesCount = t.StatusesCount
}
