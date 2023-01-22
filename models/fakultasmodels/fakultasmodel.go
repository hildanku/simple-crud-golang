package fakultasmodels

import (
	"golang-crud/config"
	"golang-crud/entities"
)

func GetAll() []entities.Fakultas_ent {
	rows, err := config.DB.Query(`SELECT * FROM fakultas`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// buat variable untuk tampung data dari entites fakultas
	var fakultasRows []entities.Fakultas_ent

	for rows.Next() {
		// var fakultas mengambil dari nama variable entities dan entities pun sama
		var fakultas entities.Fakultas_ent
		if err := rows.Scan(&fakultas.Id, &fakultas.Nama_fakultas, &fakultas.Created_at, &fakultas.Updated_at); err != nil {
			panic(err)
		}
		fakultasRows = append(fakultasRows, fakultas)
	}
	return fakultasRows
}

func Add(inputfakultas entities.Fakultas_ent) bool {
	result, err := config.DB.Exec(`
	INSERT into fakultas (nama_fakultas, created_at, updated_at)
	VALUE ( ?, ?, ?)`,
		inputfakultas.Nama_fakultas, inputfakultas.Created_at, inputfakultas.Updated_at,
	)
	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

// KARENA NELENPARNYA SINGLE DATA KITA WAJIB NENANBAH KEY fakultasRows
func Detail(id int) entities.Fakultas_ent {

	row := config.DB.QueryRow(`select id, nama_fakultas from fakultas where id = ?`, id)

	var fakultasRows entities.Fakultas_ent
	if err := row.Scan(&fakultasRows.Id, &fakultasRows.Nama_fakultas); err != nil {
		panic(err)
	}
	return fakultasRows
}

func Edit(id int, fakultas entities.Fakultas_ent) bool {
	query, err := config.DB.Exec(`UPDATE fakultas SET nama_fakultas = ?, updated_at = ? WHERE id = ?`, fakultas.Nama_fakultas, fakultas.Updated_at, id)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`delete from fakultas where id =?`, id)
	return err
}
