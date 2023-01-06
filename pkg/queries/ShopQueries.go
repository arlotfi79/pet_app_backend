package queries

import (
	DataSignatures "github.com/mahmoudKheyrati/pet_app_backend/pkg/dataSignature"
	Database "github.com/mahmoudKheyrati/pet_app_backend/pkg/database"
	"log"
)

type ShopQuery struct {
	dbClient *Database.Postgresql
}

func NewShopQuery(dbClient *Database.Postgresql) *ShopQuery {
	return &ShopQuery{dbClient: dbClient}
}

func (UserQuery *UserQuery) CreateShop(shop *DataSignatures.PostShop) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Shop (shop_owner_id, shop_name)
									VALUES ($1, $2)`)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	// acc.Password = hashedPass
	_, err = query.Exec(
		shop.ShopOwnerID,
		shop.ShopName,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
