package utils

func GetFavoriteList(userId int64) ([]int64, error) {
	list, err := DBGetFavoriteList(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func IsFavoriteDict(userId, videoId int64) (bool, error) {
	isFav, err := DBIsFavoriteDict(userId, videoId)
	if err != nil {
		return false, err
	}
	return isFav, nil
}

func FavoriteAdd(userId, videoId int64) error {
	err := DBFavoriteAdd(userId, videoId)
	if err != nil {
		return err
	}
	return nil
}

func FavoriteDelete(userId, videoId int64) error {
	err := DBFavoriteDelete(userId, videoId)
	if err != nil {
		return err
	}
	return nil
}
