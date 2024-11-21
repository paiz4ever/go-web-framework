package gromore

type CommonBody struct {
	CsjUserID string `json:"csj_user_id" binding:"required"`
}
