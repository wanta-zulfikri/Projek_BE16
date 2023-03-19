package user

import (
	"database/sql"
	"errors"
	"fmt"
)



type UserModel struct {
	conn *sql.DB
}

func (R *UserModel) Create(newuser *User) error {
	res, err := R.conn.Exec("INSERT INTO user(Nama, No_HP, Password) VALUES (?, ?, ?)", newuser.Nama, newuser.No_HP, newuser.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if aff <= 0 {
		return errors.New("register anda gagal")
	}
	return nil
}

func (R *UserModel) SetSQLConnection(db *sql.DB) {
	R.conn = db
}

func (R *UserModel) GetByNoHP(noHP string) (*User, error) {
	user := &User{}
	row := R.conn.QueryRow("SELECT * FROM user WHERE No_HP = ?", noHP)
	err := row.Scan(&user.Nama, &user.No_HP, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (R *UserModel) Deactivate(noHP string) error {
	res, err := R.conn.Exec("DELETE FROM user WHERE No_HP = ?", noHP)
	if err != nil {
		fmt.Println(err)
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if aff <= 0 {
		return errors.New("deactivate user failed")
	}
	return nil
}

func (R *UserModel) Update(user *User) error {
	res, err := R.conn.Exec("UPDATE user SET Nama = ?, Password = ? WHERE No_HP = ?", user.Nama, user.Password, user.No_HP)
	if err != nil {
		fmt.Println(err)
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if aff <= 0 {
		return errors.New("update user failed")
	}
	return nil
}

func (R *UserModel) GetAllUsers() ([]User, error) {
	allusers := []User{}
	rows, err := R.conn.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return allusers, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = User{}
		if err := rows.Scan(&user.Nama, &user.No_HP, &user.Password); err != nil {
			fmt.Println(err)
			return nil, err
		}
		allusers = append(allusers, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return allusers, nil
}

