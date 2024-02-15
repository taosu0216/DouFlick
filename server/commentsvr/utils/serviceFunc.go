package utils

import "strconv"

func CommentAdd(userId, videoId int64, commentText string, commentId int64) (*Comment, error) {
	comment, err := DbCommentAdd(userId, videoId, commentId, commentText)
	if err != nil {
		return nil, err
	}

	if err = CacheSetComment(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func CommentDelete(videoId, commentId int64) error {
	err := DbCommentDelete(commentId)
	if err != nil {
		return err
	}

	if err = CacheDeleteComment([]string{strconv.FormatInt(commentId, 10)}, videoId); err != nil {
		return err
	}
	return nil
}

func GetComments(videoId int64) ([]*Comment, error) {
	comments, err := CacheGetComments(videoId)
	if err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		comments, err = DbGetComments(videoId)
		if err != nil {
			return nil, err
		}
	}
	return comments, nil
}
