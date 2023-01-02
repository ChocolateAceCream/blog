package request

type SendEmail struct {
	Email string `json:"email" binding:"required"`
}
