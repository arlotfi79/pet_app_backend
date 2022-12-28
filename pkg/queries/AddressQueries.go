package queries

import (
	DataSignatures "github.com/mahmoudKheyrati/pet_app_backend/pkg/dataSignature"
	Database "github.com/mahmoudKheyrati/pet_app_backend/pkg/database"
	"log"
)

type AddressQuery struct {
	dbClient *Database.Postgresql
}

func NewAddressQuery(dbClient *Database.Postgresql) *AddressQuery {
	return &AddressQuery{dbClient: dbClient}
}

// -------------------------------- GET ----------------------------------

func (addressQuery *AddressQuery) GetAddressesByAccountID(id uint64) ([]DataSignatures.GetAddress, error) {
	db := addressQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT address_id, country, city, street, plaque
    								FROM address
    								WHERE account_id = $1`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	row, err := query.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var addresses []DataSignatures.GetAddress
	for row.Next() {
		var address DataSignatures.GetAddress

		err = row.Scan(&address.Id, &address.Country, &address.City, &address.Street, &address.Plaque)

		if err != nil {
			log.Fatal(err)
		}

		addresses = append(addresses, address)
	}

	return addresses, nil
}

// -------------------------------- POST ----------------------------------

func (addressQuery *AddressQuery) PutAddressUsingAccountID(accountID uint64, address *DataSignatures.PostAddress) error {
	db := addressQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO address (account_id, country, city, street, plaque)
									VALUES ($1, $2, $3, $4, $5)`)

	if err != nil {
		log.Fatal(err)
	}

	defer query.Close()

	_, err = query.Exec(accountID, address.Country, address.City, address.Street, address.Plaque)

	if err != nil {
		return err
	}

	return nil
}
