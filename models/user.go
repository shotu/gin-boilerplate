package models

import (

)

//User ...
type User struct {
	Email     string 
	Password  string 
	Name      string 
}

//UserModel ...
type UserModel struct{}

// //Signin ...
// func (m UserModel) Signin(form forms.SigninForm) (user User, err error) {

// 	err = db.GetDB().SelectOne(&user, "SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

// 	if err != nil {
// 		return user, err
// 	}

// 	bytePassword := []byte(form.Password)
// 	byteHashedPassword := []byte(user.Password)

// 	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

// 	if err != nil {
// 		return user, errors.New("Invalid password")
// 	}

// 	return user, nil
// }

// //Signup ...
// func (m UserModel) Signup(c) (user User, err error) {
// 	getDb := db.GetDB()

// 	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

// 	if err != nil {
// 		return user, err
// 	}

// 	if checkUser > 0 {
// 		return user, errors.New("User exists")
// 	}

// 	bytePassword := []byte(form.Password)
// 	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	if err != nil {
// 		panic(err)
// 	}

// 	res, err := getDb.Exec("INSERT INTO public.user(email, password, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())

// 	if res != nil && err == nil {
// 		err = getDb.SelectOne(&user, "SELECT id, email, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)
// 		if err == nil {
// 			return user, nil
// 		}
// 	}

// 	return user, errors.New("Not registered")
// }

// //One ...
// func (m UserModel) One(userID int64) (user User, err error) {
// 	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1", userID)
// 	return user, err
// }
