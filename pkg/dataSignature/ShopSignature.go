package dataSignature

type PostShop struct {
	ShopOwnerID uint64 `json:"shopOwnerID" binding:"required"`
	ShopName    string `json:"shopName" binding:"required"`
}
