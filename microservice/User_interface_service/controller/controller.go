package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"user_interface/model"
	userpb "user_interface/userpb"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

var jwtKey = []byte("secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi a")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var a model.Login
	json.NewDecoder(r.Body).Decode(&a)
	// conn, err := grpc.Dial("authentication:3001", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Hi a")
		log.Fatal(err)
	}
	defer conn.Close()
	client := userpb.NewAllTheServicesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	req := &userpb.LoginDataSend{UserId: a.Username, Pass: a.Password}
	res, err := client.Login(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	if !res.Status {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println(res)

	expirationTime := time.Now().Add(time.Minute * 10)
	claims := &model.Claims{
		Data: *res,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

func Add_people(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	/////////GRPC dial host 3002

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a11 model.User
	json.NewDecoder(r.Body).Decode(&a11)
	conn, err := grpc.Dial("user_management:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUsermanagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.Send_User_Data{
		Name:       a11.Name,
		Password:   a11.Password,
		Age:        a11.Age,
		Gmail:      a11.Email,
		Contact:    a11.Contact,
		Manager:    a11.Manager,
		Department: a11.Department,
	}
	res, err := client.Add_Users(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Response from server: %v\n", res)
	return

}

func Delete_People(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	/////////GRPC dial host 3002

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a11 model.User_delete
	json.NewDecoder(r.Body).Decode(&a11)
	conn, err := grpc.Dial("user_management:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUsermanagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.SendDeleteUserData{
		Name: a11.Name,
	}
	res, err := client.DeleteUser(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Response from server: %v\n", res)
	return

}

func Room_Addition(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "Room management" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u model.Room

	json.NewDecoder(r.Body).Decode(&u)

	conn, err := grpc.Dial("room_management:3003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewRoommanagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.SendRoomDetails{
		Roomnumber:      u.Room_number,
		Price:           u.Price,
		Features:        u.Features,
		Occupencystatus: u.OccupencyStatus,
	}
	res, err := client.AddRooms(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Response from server: %v\n", res)
	return

}

func Room_Deletion(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "Room management" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var a model.Room_Delete
	json.NewDecoder(r.Body).Decode(&a)

	conn, err := grpc.Dial("room_management:3003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewRoommanagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.SendDeleteRoomDetail{
		Roomnumber: a.Room_number,
	}
	res, err := client.DeleteRoom(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Response from server: %v\n", res)
	return

}

func Room_Feature_Updation(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "Room management" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a model.Room_feature_update
	json.NewDecoder(r.Body).Decode(&a)

	conn, err := grpc.Dial("room_management:3003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewRoommanagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.SendFeatureUpdate{
		Roomnumber: a.Room_numer,
		Features:   a.Features,
	}
	res, err := client.UpdateFeatures(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Response from server: %v\n", res)
	return

}

func Room_Price_Updation(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "Room management" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var a model.Room_price_update
	json.NewDecoder(r.Body).Decode(&a)

	conn, err := grpc.Dial("room_management:3003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewRoommanagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.SendPriceUpdate{
		Roomnumber: a.Room_number,
		Price:      a.Price,
	}
	res, err := client.UpdatePrice(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}

	fmt.Printf("Response from server: %v\n", res)
	return

}

func Add_Customer(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		log.Fatal(err)
	}

	tokenStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if claims.Data.Department != "Customer management" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var q model.Customer

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewDecoder(r.Body).Decode(&q)

	// conn, err := grpc.Dial("email:3005", grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:3005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := userpb.NewEmailVerificationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &userpb.Send_CustomerEmail{
		Email: q.Email,
	}
	res, err := client.Send_OTP(ctx, req)
	fmt.Println(res.Otp)
	if err != nil {
		log.Fatalf("Error calling the function: %v", err)
	}
	v := res.Otp

	var g model.Total

	g.Contact = q.Contact
	g.Email = q.Email
	g.Name = q.Name
	g.OTP = v
	//Make it as a cookie

	expirationTime := time.Now().Add(time.Minute * 10)
	new_claims := &model.Claims_total{
		Data: g,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, new_claims)
	tokenString1, err := token1.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "Data",
			Value:   tokenString1,
			Expires: expirationTime,
		})

	fmt.Printf("Response from server: %v\n", res)
	return
}
