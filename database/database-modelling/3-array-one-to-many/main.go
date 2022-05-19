package main

import "fmt"

type PhoneRow struct {
	ID          int // primary key
	CountryCode int
	Number      int
}

type UserRow struct {
	ID       int // primary key
	Name     string
	Age      int
	PhoneIDs []int // foreign key
	PhoneRow []PhoneRow
}

type UserTable []UserRow

func (db *UserTable) InsertUser(name string, age int, phoneIDs []int) {
	(*db) = append(*db, UserRow{
		ID:       len(*db) + 1,
		Name:     name,
		Age:      age,
		PhoneIDs: phoneIDs,
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
		for j := 0; j < len((*db)[i].PhoneIDs); j++ {
			(*db)[i].PhoneRow = append((*db)[i].PhoneRow, (*phoneDB).Where((*db)[i].PhoneIDs[j]))
		}
	}
}

type PhoneTable []PhoneRow

func (db *PhoneTable) Insert(countryCode int, number int) {
	(*db) = append(*db, PhoneRow{
		ID:          len(*db) + 1,
		CountryCode: countryCode,
		Number:      number,
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

func main() {
	phoneDB := make(PhoneTable, 0)
	userDB := make(UserTable, 0)

	phoneDB.Insert(62, 1234567890)
	phoneDB.Insert(62, 2345678901)
	phoneDB.Insert(62, 3243434343)

	userDB.InsertUser("John", 20, []int{1, 2})
	userDB.InsertUser("Gina", 20, []int{3})

	userDB.Join(&phoneDB)

	fmt.Printf("%+v\n", userDB.Where(1))
	fmt.Printf("%+v\n", userDB.Where(2))

}
