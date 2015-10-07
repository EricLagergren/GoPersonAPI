package person

import "log"

func (m *Model) Create() error {

	result, err := db.Exec(`
		INSERT INTO person (
			firstname,
			lastname,
			emailaddress,
			phone,
			weight
		)
		VALUES (
			?, ?, ?, ?, ?
		)`,
		m.FirstName, m.LastName, m.EmailAddress, m.Phone, m.Weight)

	if err != nil {
		log.Print(err)
		return err
	}

	// Attempt to grab the ID of the newly created object
	newID, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return err
	}
	// No error when grabbing the newly created ID
	m.ID = newID
	return nil
}
