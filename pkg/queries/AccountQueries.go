package queries

import (
	DataSignatures "github.com/mahmoudKheyrati/pet_app_backend/pkg/dataSignature"
	Database "github.com/mahmoudKheyrati/pet_app_backend/pkg/database"
	pwd "github.com/mahmoudKheyrati/pet_app_backend/pkg/encryption"
	"log"
)

type UserQuery struct {
	dbClient *Database.Postgresql
}

func NewUserQuery(dbClient *Database.Postgresql) *UserQuery {
	return &UserQuery{dbClient: dbClient}
}

// CREATE

func (UserQuery *UserQuery) CreateNormalUser(account *DataSignatures.PostNormalAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Account (name, last_name, user_name, phone_number, password, email, gender, join_date)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	hashedPass, _ := pwd.Encrypt(account.Password)
	log.Print(hashedPass)
	// acc.Password = hashedPass
	_, err = query.Exec(
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		hashedPass,
		account.Email,
		account.Gender,
		account.JoinDate,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (UserQuery *UserQuery) CreateVetUser(account *DataSignatures.PostVetAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Vet (name, last_name, user_name, phone_number, password, email, gender, 
                 join_date, expertize_level, expertize_area, working_hours_start_date, working_hours_end_date,
                 working_hours_start_hour, working_hours_end_hour)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	hashedPass, _ := pwd.Encrypt(account.Password)
	log.Print(hashedPass)
	// acc.Password = hashedPass
	_, err = query.Exec(
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		hashedPass,
		account.Email,
		account.Gender,
		account.JoinDate,
		account.ExpertiseLevel,
		account.ExpertiseArea,
		account.WorkingHoursStartDate,
		account.WorkingHoursEndDate,
		account.WorkingHoursStartHour,
		account.WorkingHoursEndHour,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (UserQuery *UserQuery) CreatePetLoverUser(account *DataSignatures.PostPetLoverAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`INSERT INTO Vet (name, last_name, user_name, phone_number, password, email, gender, 
                 join_date, free_hours_start_date, free_hours_end_date, free_hours_start_hour, free_hours_end_hour)
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	hashedPass, _ := pwd.Encrypt(account.Password)
	log.Print(hashedPass)
	// acc.Password = hashedPass
	_, err = query.Exec(
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		hashedPass,
		account.Email,
		account.Gender,
		account.JoinDate,
		account.FreeTimeStartDate,
		account.FreeTimeEndDate,
		account.FreeTimeStartHour,
		account.FreeTimeEndHour,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

// UPDATE

func (UserQuery *UserQuery) UpdateNormalUser(account *DataSignatures.UpdateNormalAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE Account
									SET name = $1,
									    last_name = $2,
									    user_name = $3, 
									    password = $4, 
									    phone_number = $5, 
									    email = $6, 
									    gender = $7
									WHERE user_name = $3`)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	hashedPass, _ := pwd.Encrypt(account.Password)
	log.Print(hashedPass)
	// acc.Password = hashedPass
	_, err = query.Exec(
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		hashedPass,
		account.Email,
		account.Gender,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (UserQuery *UserQuery) UpdateVetUser(account *DataSignatures.UpdateVetAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE Vet
									SET name = $1,
									    last_name = $2,
									    user_name = $3, 
									    password = $4, 
									    phone_number = $5, 
									    email = $6, 
									    gender = $7,
									    expertize_level = $8,
									    expertize_area = $9,
									    working_hours_start_date = $10,
									    working_hours_end_date = $11,
									    working_hours_start_hour = $12,
									    working_hours_end_hour = $13
									    WHERE user_name = $3`)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	hashedPass, _ := pwd.Encrypt(account.Password)
	log.Print(hashedPass)
	// acc.Password = hashedPass
	_, err = query.Exec(
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		hashedPass,
		account.Email,
		account.Gender,
		account.JoinDate,
		account.ExpertiseLevel,
		account.ExpertiseArea,
		account.WorkingHoursStartDate,
		account.WorkingHoursEndDate,
		account.WorkingHoursStartHour,
		account.WorkingHoursEndHour,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (UserQuery *UserQuery) UpdatePetLoverUser(account *DataSignatures.UpdatePetLoverAccount) error {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`UPDATE Pet_Lover
									SET name = $1,
									    last_name = $2,
									    user_name = $3, 
									    password = $4, 
									    phone_number = $5, 
									    email = $6, 
									    gender = $7,
									    free_hours_start_date = $8,
									    free_hours_end_date = $9,
									    free_hours_start_hour = $10,
									    free_hours_end_hour = $11
									    WHERE user_name = $3`)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	if err != nil {
		log.Fatalln(err)
		return err
	}

	defer query.Close()

	hashedPass, _ := pwd.Encrypt(account.Password)
	log.Print(hashedPass)
	// acc.Password = hashedPass
	_, err = query.Exec(
		account.Name,
		account.LastName,
		account.UserName,
		account.PhoneNumber,
		hashedPass,
		account.Email,
		account.Gender,
		account.JoinDate,
		account.FreeTimeStartDate,
		account.FreeTimeEndDate,
		account.FreeTimeStartHour,
		account.FreeTimeEndHour,
	)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

// GET

func (UserQuery *UserQuery) GetUserByUname(username string) ([]DataSignatures.GetNormalAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE user_name=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(username)
	if err != nil {
		log.Fatalln(err)
		return nil, err

	}
	var accounts []DataSignatures.GetNormalAccount
	for rows.Next() {
		var account DataSignatures.GetNormalAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.Password, &account.PhoneNumber, &account.Email, &account.Gender, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}

func (UserQuery *UserQuery) GetUserByEmail(email string) ([]DataSignatures.GetNormalAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE email=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(email)
	if err != nil {
		log.Fatalln(err)
		return nil, err

	}
	var accounts []DataSignatures.GetNormalAccount
	for rows.Next() {
		var account DataSignatures.GetNormalAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.Password, &account.PhoneNumber, &account.Email, &account.Gender, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}

func (UserQuery *UserQuery) GetUserByPhoneNumber(phoneNumber string) ([]DataSignatures.GetNormalAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE phone_number=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(phoneNumber)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var accounts []DataSignatures.GetNormalAccount
	for rows.Next() {
		var account DataSignatures.GetNormalAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.Password, &account.PhoneNumber, &account.Email, &account.Gender, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}

func (UserQuery *UserQuery) GetUserById(id uint64) ([]DataSignatures.GetNormalAccount, error) {
	db := UserQuery.dbClient.GetDB()

	query, err := db.Prepare(`SELECT * FROM Account WHERE account_id=$1`)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer query.Close()
	rows, err := query.Query(id)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var accounts []DataSignatures.GetNormalAccount
	for rows.Next() {
		var account DataSignatures.GetNormalAccount

		err = rows.Scan(&account.Id, &account.Name, &account.LastName, &account.UserName, &account.Password, &account.PhoneNumber, &account.Email, &account.Gender, &account.JoinDate)

		if err != nil {
			log.Fatalln(err)
		}

		accounts = append(accounts, account)
	}
	return accounts, err
}
