package store

import "Angular/api-rest/models"

func (store *DbStore) GetCardz() ([]models.Cardz, error) {
	var c []models.Cardz
	var l []models.CardzLink
	var b []models.Back

	if err := store.db.Select(&c, "SELECT * FROM cardz"); err != nil {
		return nil, err
	}
	if err := store.db.Select(&l, "SELECT l.* FROM cardz c	INNER JOIN link l ON l.idCardz = c.id "); err != nil {
		return nil, err
	}
	if err := store.db.Select(&b, "SELECT b.* FROM cardz c INNER JOIN backgroud b ON b.id = c.backgroudId"); err != nil {
		return nil, err
	}
	for i := range c {
		c[i].Links = l
		c[i].Background = b[i]
	}
	return c, nil
}

func (store *DbStore) GetCardzById(id int64) (models.Cardz, error) {
	var c models.Cardz
	var l []models.CardzLink
	var b models.Back
	err := store.db.Get(&c, "SELECT * FROM cardz WHERE id = ?", id)
	if err != nil && err.Error() == noRowsSql {
		return c, nil
	}
	if err != nil {
		return c, err
	}

	err = store.db.Select(&l, `SELECT l.* FROM cardz c	
	INNER JOIN link l ON l.idCardz = c.id
	WHERE c.id = ? `, id)
	if err != nil && err.Error() == noRowsSql {
		return c, nil
	}
	if err != nil {
		return c, err
	}

	err = store.db.Get(&b, `SELECT b.* FROM cardz c 
	INNER JOIN backgroud b ON b.id = c.backgroudId
	WHERE c.id = ?`, id)
	if err != nil && err.Error() == noRowsSql {
		return c, nil
	}
	if err != nil {
		return c, err
	}

	c.Links = l
	c.Background = b

	return c, nil

}

func (store *DbStore) DeleteCardz(id int64) error {
	if _, err := store.db.Exec(`DELETE c,l,b FROM cardz c
	INNER JOIN link l ON l.idCardz = c.id
	INNER JOIN backgroud b ON c.backgroudId = b.id
	WHERE c.id = ? `, id); err != nil {
		return err
	}
	return nil
}

func (store *DbStore) CreateCardz(c models.Cardz) error {
	idC, idB := 0, 0

	if _, err := store.db.NamedExec(`INSERT INTO backgroud (image,animation)
	 VALUES (:image,:animation)`, &c.Background); err != nil {
		return err
	}

	if err := store.db.Get(&idB, "SELECT max(id) as id FROM backgroud"); err != nil {
		return err
	}
	c.BacgroudId = idB
	if _, err := store.db.NamedExec(`INSERT INTO cardz (name,biography,avatar,animation,backgroudId) 
	VALUES (:name,:biography,:avatar,:animation,:backgroudId)`, &c); err != nil {
		return err
	}

	if err := store.db.Get(&idC, `SELECT max(id) as id FROM cardz`); err != nil {
		return err
	}

	for i := range c.Links {
		c.Links[i].CardzId = idC
		if _, err := store.db.NamedExec(`INSERT INTO link (url,text,icon,idCardz)
		VALUES (:url,:text,:icon,:idCardz)`, &c.Links[i]); err != nil {
			return err
		}
	}

	return nil
}
