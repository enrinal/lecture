package main

import "fmt"

type PhoneRow struct {
	ID          int // primary key
	CountryCode int
	Number      int
	UserID      int // foreign key
}

type UserRow struct {
	ID        int // primary key
	Name      string
	Age       int
	PhoneRows []PhoneRow
}

type PhoneTable []PhoneRow

func (db *PhoneTable) Insert(countryCode int, number int, userID int) {
	(*db) = append(*db, PhoneRow{
		ID:          len(*db) + 1,
		CountryCode: countryCode,
		Number:      number,
		UserID:      userID,
	})
}

func (db *PhoneTable) Where(phoneID int) PhoneRow {
	var result PhoneRow
	for _, row := range *db {
		if row.ID == phoneID {
			return row
		}
	}
	return result
}

type UserTable []UserRow

func (db *UserTable) Insert(name string, age int) {
	(*db) = append(*db, UserRow{
		ID:   len(*db) + 1,
		Name: name,
		Age:  age,
	})
}

func (db *UserTable) Where(id int) UserRow {
	for _, row := range *db {
		if row.ID == id {
			return row
		}
	}
	return UserRow{}
}

func (db *UserTable) Join(phoneDB *PhoneTable) {
	for i := 0; i < len(*db); i++ {
		for j := 0; j < len(*phoneDB); j++ {
			if (*phoneDB)[j].UserID == (*db)[i].ID {
				(*db)[i].PhoneRows = append((*db)[i].PhoneRows, (*phoneDB)[j])
			}
		}
	}
}

func main() {
	phoneDB := make(PhoneTable, 0)
	userDB := make(UserTable, 0)

	phoneDB.Insert(62, 1234567890, 1)
	phoneDB.Insert(62, 2234567890, 1)
	phoneDB.Insert(62, 3234567890, 2)
	phoneDB.Insert(62, 4234567890, 2)
	phoneDB.Insert(62, 5234567890, 1)

	userDB.Insert("John", 30)
	userDB.Insert("Jane", 25)

	userDB.Join(&phoneDB)

	fmt.Printf("%+v\n", userDB.Where(1))
	fmt.Printf("%+v\n", userDB.Where(2))
}
