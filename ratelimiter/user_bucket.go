package ratelimiter

type UserBucket struct {
	userMap map[int]*LeakyBucket
}

func CreateUserBucket(user int, cap int64) UserBucket {
	bucket := CreateBucket(cap)
	uMap := make(map[int]*LeakyBucket)
	uMap[user] = bucket
	return UserBucket{uMap}
}

func (u UserBucket) AllowAccess(user int) bool {
	return u.userMap[user].GrantAccess()
}
