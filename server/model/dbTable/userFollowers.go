package dbTable

type UserFollowers struct {
	UserId     uint `gorm:"column:user_id;uniqueIndex:following,sort:desc"`
	FollowerId uint `gorm:"column:user_id;uniqueIndex:following,sort:desc"`
}
