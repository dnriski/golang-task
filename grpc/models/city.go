package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"unsia/pb/cities"
)

type City struct {
	Pb  cities.City
	Log *log.Logger
}

func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `SELECT id, name FROM cities where id = $1`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name)
	if err != nil {
		u.Log.Println("Error on Query : ", err)
		return err
	}
	return nil
}

func (u *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id`
	stat, err := db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	err = stat.QueryRowContext(ctx, in.Name).Scan(&u.Pb.Id)
	if err != nil {
		return err
	}

	u.Pb.Name = in.Name
	// fmt.Println(rs)

	return nil
}

func (u *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `DELETE FROM cities WHERE id = $1;`
	stat, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	rs, err := stat.ExecContext(ctx, in.Id)
	if err != nil {
		return err
	}

	affected, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("data not found")
	}
	return nil
}

// TRY Update Data
func (u *City) Update(ctx context.Context, db *sql.DB, in *cities.City) error {
	query := `UPDATE cities SET name = $2 WHERE id = $1;`

	_, err := db.ExecContext(ctx, query, in.Id, in.Name)
	if err != nil {
		return err
	}

	return nil
}
