package person

import "log"

func (m Model) Delete() error {

	// Attempt to run the SQL query against the database
	_, err := db.Exec(`
		DELETE
        FROM person
        WHERE
        id = ?
    `, m.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
