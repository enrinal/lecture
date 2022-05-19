package main

import (
	"fmt"
)

type DriverRow struct {
	ID         int // primary key
	Name       string
	Age        int
	Region     string
	CustomerID int // foreign key
	Customer   CustomerRow
}

type CustomerRow struct {
	ID      int // primary key
	Name    string
	Address string
}

type DriverTable []DriverRow

func (d *DriverTable) Insert(name string, age int, region string, customerID int) {
	*d = append(*d, DriverRow{
		ID:         len(*d) + 1,
		Name:       name,
		Age:        age,
		Region:     region,
		CustomerID: customerID,
	})
}

func (d *DriverTable) Join(c *CustomerTable) {
	for i := 0; i < len(*d); i++ {
		(*d)[i].Customer = (*c).Where((*d)[i].CustomerID)
	}
}

func (d *DriverTable) Where(id int) DriverRow {
	for _, row := range *d {
		if row.ID == id {
			return row
		}
	}
	return DriverRow{}
}

type CustomerTable []CustomerRow

func (c *CustomerTable) Insert(name string, address string) {
	*c = append(*c, CustomerRow{
		ID:      len(*c) + 1,
		Name:    name,
		Address: address,
	})
}

func (c *CustomerTable) Where(id int) CustomerRow {
	for _, row := range *c {
		if row.ID == id {
			return row
		}
	}
	return CustomerRow{}
}

func main() {
	driverDB := make(DriverTable, 0)
	customerDB := make(CustomerTable, 0)

	driverDB.Insert("John", 21, "Jakarta", 1)
	driverDB.Insert("Jane", 22, "Bandung", 2)
	driverDB.Insert("Gina", 24, "Jakarta", 3)

	customerDB.Insert("Budi", "Jakarta")
	customerDB.Insert("Siti", "Bandung")
	customerDB.Insert("Sri", "Jakarta")

	// get data
	fmt.Println(driverDB.Where(1))
	fmt.Println(customerDB.Where(1))
	fmt.Println("===========================================")

	// join data
	driverDB.Join(&customerDB)
	fmt.Printf("%+v\n", driverDB.Where(1))
	fmt.Printf("%+v\n", driverDB.Where(2))
}
