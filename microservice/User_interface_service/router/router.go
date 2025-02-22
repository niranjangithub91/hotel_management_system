package router

import (
	"user_interface/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/add_people", controller.Add_people).Methods("POST")
	router.HandleFunc("/delete_staff", controller.Delete_People).Methods("DELETE")
	router.HandleFunc("/add_rooms", controller.Room_Addition).Methods("POST")
	router.HandleFunc("/delete_room", controller.Room_Deletion).Methods("DELETE")
	router.HandleFunc("/update_room_features", controller.Room_Feature_Updation).Methods("PUT")
	router.HandleFunc("/update_room_price", controller.Room_Price_Updation).Methods("PUT")
	router.HandleFunc("/add_customer", controller.Add_Customer).Methods("POST")
	return router
}
