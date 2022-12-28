package dataSignature

type PostAddress struct {
	Country string `json:"country" binding:"required"`
	City    string `json:"city" binding:"required"`
	Street  string `json:"street" binding:"required"`
	Plaque  string `json:"plaque" binding:"required"`
}

type GetAddress struct {
	Id uint64 `json:"id" binding:"required"`
	PostAddress
}
