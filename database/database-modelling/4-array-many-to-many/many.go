package main

import "fmt"

type PhoneRow struct {
	ID          int // primary key
	CountryCode int
	Number      int
}

type UserRow struct {
	ID   int // primary key
	Name string
	Age  int
}

type UserPhoneRow struct {
	UserID   int // primary key
	UserRow  UserRow
	PhoneID  int // primary key
	PhoneRow PhoneRow
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

type UserPhoneTable []UserPhoneRow

func (db *UserPhoneTable) Insert(userID int, phoneID int) {
	(*db) = append(*db, UserPhoneRow{
		UserID:  userID,
		PhoneID: phoneID,
	})
}

func (db *UserPhoneTable) Join(phoneDB *PhoneTable, userDB *UserTable) {
	for i := 0; i < len(*db); i++ {
		(*db)[i].PhoneRow = phoneDB.Where((*db)[i].PhoneID)
		(*db)[i].UserRow = userDB.Where((*db)[i].UserID)
	}
}

func main() {
	phoneDB := make(PhoneTable, 0)
	userDB := make(UserTable, 0)
	userPhoneDB := make(UserPhoneTable, 0)

	phoneDB.Insert(62, 1234567890)
	phoneDB.Insert(62, 2345678901)

	userDB.Insert("John", 20)
	userDB.Insert("Gina", 20)

	userPhoneDB.Insert(1, 1)
	userPhoneDB.Insert(1, 2)
	userPhoneDB.Insert(2, 1)
	userPhoneDB.Insert(2, 2)

	fmt.Println("Phones:")
	for _, row := range phoneDB {
		fmt.Println("| ID: ", row.ID, "| Country Code: ", row.CountryCode, "| Number: ", row.Number, "|")
	}
	fmt.Println("Users:")
	for _, row := range userDB {
		fmt.Println("| ID: ", row.ID, "| Name: ", row.Name, "| Age: ", row.Age, "|")
	}
	fmt.Println("UserPhone:")
	for _, row := range userPhoneDB {
		fmt.Println("| User ID: ", row.UserID, "| Phone ID: ", row.PhoneID, "|")
	}

	userPhoneDB.Join(&phoneDB, &userDB)
	fmt.Println("UserPhone:")
	fmt.Printf("%#v\n", userPhoneDB)
}
