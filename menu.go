package main

import (
	"encoding/json"
	"fmt"
	// "github.com/gorilla/mux"
	// "log"
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	// "strings"
	"database/sql"
	"sort"
)

type Passenger struct {
	PassengerId  int    `json:"passengerId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	EmailAddr    string `json:"emailAddr"`
	Password     string `json:"password"`
}

type Driver struct {
	DriverId      int    `json:"driverId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	MobileNumber  string `json:"mobileNumber"`
	EmailAddr     string `json:"emailAddr"`
	Password      string `json:"password"`
	LicenseNumber string `json:"licenseNumber"`
	IdNumber      string `json:"idNumber"`
	DriverStatus  string `json:"driverStatus"`
}

type Trip struct {
	TripId            int       `json:"tripId"`
	PickUpPostalCode  string    `json:"pickUpPostalCode"`
	DropOffPostalCode string    `json:"dropOffPostalCode"`
	PassengerId       int       `json:"passengerId"`
	StartTime         sql.NullTime `json:"startTime"`
	EndTime           sql.NullTime `json:"endTime"`
	TripStatus        string    `json:"tripStatus"`
	RequestTime       sql.NullTime `json:"requestTime"`
	DriverId	int	`json:"driverId"`
}

type TripDriver struct {
	TripId           int    `json:"tripId"`
	DriverId         int    `json:"driverId"`
	TripDriverStatus string `json:"tripDriverStatus"`
}

func mainMenuOptions() {
	fmt.Println()
	fmt.Println("==========Main Menu==========")
	fmt.Println("1. Login as Passenger")
	fmt.Println("2. Login as Driver")
	fmt.Println("3. Create Passenger Account")
	fmt.Println("4. Create Driver Account")
	fmt.Println("0. Quit")
	fmt.Println("=============================")
	fmt.Print("Please enter an option: ")
}

func main() {
	flag := true
	for flag {
		mainMenuOptions()
		mainMenuInput := 10000
		fmt.Scan(&mainMenuInput)
		switch mainMenuInput {
		case 0:
			flag = false
		case 1:
			//Passenger auth info input
			var emailAddr string
			var password string
			fmt.Print("Please enter an email address: ")
			fmt.Scan(&emailAddr)
			fmt.Print("Please enter a password: ")
			fmt.Scan(&password)
			//Passenger auth endpoint
			client := &http.Client{}
			var responseObject Passenger
			url := "http://localhost:5000/auth/passenger?emailAddr=" + emailAddr + "&&password=" + password
			loginFlag := false
			if req, err := http.NewRequest("GET", url, nil); err == nil {
				if res, err := client.Do(req); err == nil {
					if body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
						if res.StatusCode == 200 {
							json.Unmarshal(body, &responseObject)
							loginFlag = true
						}
					}
				}
			}
			if loginFlag {
				//Passenger Menu
				fmt.Println("Welcome " + responseObject.FirstName)
				passengerMenu(responseObject)
			} else {
				fmt.Println("User not found")
			}

		case 2:
			//Driver auth info input
			var emailAddr string
			var password string
			fmt.Print("Please enter an email address: ")
			fmt.Scan(&emailAddr)
			fmt.Print("Please enter a password: ")
			fmt.Scan(&password)
			//Driver login endpoint
			client := &http.Client{}
			var responseObject Driver
			url := "http://localhost:5000/auth/driver?emailAddr=" + emailAddr + "&&password=" + password
			loginFlag := false
			if req, err := http.NewRequest("GET", url, nil); err == nil {
				if res, err := client.Do(req); err == nil {
					if body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
						if res.StatusCode == 200 {
							json.Unmarshal(body, &responseObject)
							loginFlag = true
						}
					}
				}
			}
			if loginFlag {
				//Driver Menu
				fmt.Println("Welcome " + responseObject.FirstName)
				driverMenu(responseObject)
			} else {
				fmt.Println("User not found")
			}

		case 3:
			//Passenger acccount info input
			var p Passenger
			fmt.Print("Please enter your first name: ")
			fmt.Scan(&p.FirstName)
			fmt.Print("Please enter your last name: ")
			fmt.Scan(&p.LastName)
			fmt.Print("Please enter a mobile number: ")
			fmt.Scan(&p.MobileNumber)
			fmt.Print("Please enter an email address: ")
			fmt.Scan(&p.EmailAddr)
			fmt.Print("Please enter a password: ")
			fmt.Scan(&p.Password)
			//Passenger creation endpoint
			client := &http.Client{}
			url := "http://localhost:5000/passenger"
			postBody, _ := json.Marshal(p)
			resBody := bytes.NewBuffer(postBody)
			if req, err := http.NewRequest("POST", url, resBody); err == nil {
				if res, err2 := client.Do(req); err2 == nil {
					if res.StatusCode == 202 {
						fmt.Println("Successfully created account")
					} else if res.StatusCode == 400 {
						fmt.Println("Error - Bad Request")
					}
				}
			}
		case 4:
			//Driver acccount info input
			var d Driver
			fmt.Print("Please enter your first name: ")
			fmt.Scan(&d.FirstName)
			fmt.Print("Please enter your last name: ")
			fmt.Scan(&d.LastName)
			fmt.Print("Please enter a mobile number: ")
			fmt.Scan(&d.MobileNumber)
			fmt.Print("Please enter an email address: ")
			fmt.Scan(&d.EmailAddr)
			fmt.Print("Please enter a password: ")
			fmt.Scan(&d.Password)
			fmt.Print("Please enter a id number: ")
			fmt.Scan(&d.IdNumber)
			fmt.Print("Please enter a license number: ")
			fmt.Scan(&d.LicenseNumber)
			//Driver creation endpoint
			client := &http.Client{}
			url := "http://localhost:5000/driver"
			postBody, _ := json.Marshal(d)
			resBody := bytes.NewBuffer(postBody)
			if req, err := http.NewRequest("POST", url, resBody); err == nil {
				if res, err2 := client.Do(req); err2 == nil {
					if res.StatusCode == 202 {
						fmt.Println("Successfully created account")
					} else if res.StatusCode == 400 {
						fmt.Println("Error - Bad Request")
					}
				}
			}
		}
	}
}

func passengerMenuOptions() {
	fmt.Println()
	fmt.Println("========Passenger Menu========")
	fmt.Println("1. View Trips")
	fmt.Println("2. Request Trip")
	fmt.Println("3. Update account info")
	fmt.Println("0. Log Out")
	fmt.Println("==============================")
	fmt.Print("Please enter an option: ")
}

func passengerMenu(p Passenger) {
	flag := true
	for flag {
		passengerMenuOptions()
		passengerMenuInput := 10000
		fmt.Scan(&passengerMenuInput)
		switch passengerMenuInput {
		case 0:
			fmt.Println("Loggin out")
			flag = false
		case 1:
			passengerTripMenu(p.PassengerId)
			//Trip endpoint
			type TripDetails struct {
				PassengerId int `json:"passengerId"`
			}
			var td TripDetails
			td.PassengerId = p.PassengerId
			type TripsResponse struct {
				Trips map[string]Trip `json:"Trips"`
			}
			client := &http.Client{}
			url := "http://localhost:5001/trip"
			getBody, _ := json.Marshal(td)
			resBody := bytes.NewBuffer(getBody)
			if req, err := http.NewRequest("GET", url, resBody); err == nil {
				if res, err := client.Do(req); err == nil {
					if body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
						var responseObject TripsResponse
						json.Unmarshal(body, &responseObject)
						tripArray := []Trip{}
						for _, element := range responseObject.Trips {
							tripArray = append(tripArray,element)
						}
						//Sort trip array reverse chronologically
						sort.Slice(tripArray, func(i, j int) bool {
							return tripArray[i].RequestTime.Time.After(tripArray[j].RequestTime.Time)
						})
						for _, element := range tripArray {
							tripArray = append(tripArray,element)
							fmt.Println()
							fmt.Println("========Trip "+ strconv.Itoa(element.TripId) + "========")
							fmt.Println("Request Time: " + element.RequestTime.Time.Format(time.RFC822))
							fmt.Println("Pickup Postal Code: " + element.PickUpPostalCode)
							fmt.Println("Drop off Postal Code: " + element.DropOffPostalCode)
							if element.StartTime.Time.IsZero(){
								fmt.Println("Start Time: Pending" )
							}else{
								fmt.Println("Start Time: " + element.StartTime.Time.Format(time.RFC822))
							}
							if element.EndTime.Time.IsZero(){
								fmt.Println("End Time: Pending" )
							}else{
								fmt.Println("End Time: " + element.EndTime.Time.Format(time.RFC822))
							}
							fmt.Println("Trip Status: " + element.TripStatus)
							fmt.Println("==============================")
							fmt.Println()
						}
					}
				}
			}
			//Result (1:25pm Pending Driver)
		case 2:
			//Trip request info input
			type TripRequest struct {
				PickUpPostalCode  string `json:"pickUpPostalCode"`
				DropOffPostalCode string `json:"dropOffPostalCode"`
				PassengerId       int    `json:"passengerId"`
				RequestTime       string `json:"requestTime"`
			}
			var tr TripRequest
			fmt.Print("Please enter a pickup postal code: ")
			fmt.Scan(&tr.PickUpPostalCode)
			fmt.Print("Please enter a drop off postal code: ")
			fmt.Scan(&tr.DropOffPostalCode)
			tr.PassengerId = p.PassengerId
			tr.RequestTime = time.Now().Format(time.RFC3339)
			//Trip request endpoint
			client := &http.Client{}
			url := "http://localhost:5001/trip"
			postBody, _ := json.Marshal(tr)
			resBody := bytes.NewBuffer(postBody)
			if req, err := http.NewRequest("POST", url, resBody); err == nil {
				if res, err2 := client.Do(req); err2 == nil {
					if res.StatusCode == 202 {
						fmt.Println("Successfully requested for trip")
					} else if res.StatusCode == 400 {
						fmt.Println("No driver avaialable")
					} else if res.StatusCode == 409 {
						fmt.Println("Existing conflicting ride")
					}
				}
			}
			//Result
		case 3:
			//Passenger account info current
			fmt.Print("First Name: "+p.FirstName)
			fmt.Print("Last Name: "+p.LastName)
			fmt.Print("Mobile Number: "+p.MobileNumber)
			fmt.Print("Email Addr: "+p.EmailAddr)
			fmt.Print("Password : "+p.Password)
			
			//Passenger acccount info input
			var tempP Passenger = p
			fmt.Print("Please enter your first name: ")
			fmt.Scan(&tempP.FirstName)
			fmt.Print("Please enter your last name: ")
			fmt.Scan(&tempP.LastName)
			fmt.Print("Please enter a mobile number: ")
			fmt.Scan(&tempP.MobileNumber)
			fmt.Print("Please enter an email address: ")
			fmt.Scan(&tempP.EmailAddr)
			fmt.Print("Please enter a password: ")
			fmt.Scan(&tempP.Password)
			//Passenger update endpoint
			client := &http.Client{}
			url := "http://localhost:5000/passenger"
			patchBody, _ := json.Marshal(tempP)
			resBody := bytes.NewBuffer(patchBody)
			if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
				if res, err2 := client.Do(req); err2 == nil {
					if res.StatusCode == 202 {
						fmt.Println("Successfully updated account")
						p = tempP
					} else if res.StatusCode == 400 {
						fmt.Println("Error - Bad Request")
					}
				}
			}
		}
	}
}

func passengerTripMenu(passengerId int) {
	flag := true
	for flag {
		//Trip endpoint
		type TripDetails struct {
			PassengerId int `json:"passengerId"`
		}
		var td TripDetails
		td.PassengerId = passengerId
		type TripsResponse struct {
			Trips map[string]Trip `json:"Trips"`
		}
		client := &http.Client{}
		url := "http://localhost:5001/trip"
		getBody, _ := json.Marshal(td)
		resBody := bytes.NewBuffer(getBody)
		rideMap := make(map[int]Trip)
		fmt.Println()
		fmt.Println("=========Passenger Trips=========")
		if req, err := http.NewRequest("GET", url, resBody); err == nil {
			if res, err := client.Do(req); err == nil {
				if body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
					var responseObject TripsResponse
					json.Unmarshal(body, &responseObject)
					tripArray := []Trip{}
					for _, element := range responseObject.Trips {
						tripArray = append(tripArray,element)
					}
					//Sort trip array reverse chronologically
					sort.Slice(tripArray, func(i, j int) bool {
						return tripArray[i].RequestTime.Time.After(tripArray[j].RequestTime.Time)
					})
					var i=1
					for _, element := range tripArray {
						fmt.Println(strconv.Itoa(i)+". Trip " + "(" + element.RequestTime.Time.Format(time.RFC822) + "): "+ element.TripStatus)
						rideMap[i]=element
						i++
					}
				}
			}
		}
		// if req, err := http.NewRequest("GET", url, resBody); err == nil {
		// 	if res, err := client.Do(req); err == nil {
		// 		if body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
		// 			var responseObject TripsResponse
		// 			json.Unmarshal(body, &responseObject)
		// 			tripArray := []Trip{}
		// 			for _, element := range responseObject.Trips {
		// 				tripArray = append(tripArray,element)
		// 			}
		// 			//Sort trip array reverse chronologically
		// 			sort.Slice(tripArray, func(i, j int) bool {
		// 				return tripArray[i].RequestTime.Time.After(tripArray[j].RequestTime.Time)
		// 			})
		// 			for _, element := range tripArray {
		// 				tripArray = append(tripArray,element)
		// 				fmt.Println()
		// 				fmt.Println("========Trip "+ strconv.Itoa(element.TripId) + "========")
		// 				fmt.Println("Request Time: " + element.RequestTime.Time.Format(time.RFC822))
		// 				fmt.Println("Pickup Postal Code: " + element.PickUpPostalCode)
		// 				fmt.Println("Drop off Postal Code: " + element.DropOffPostalCode)
		// 				fmt.Println("Start Time: " + element.StartTime.Time.Format(time.RFC822))
		// 				fmt.Println("End Time: " + element.EndTime.Time.Format(time.RFC822))
		// 				fmt.Println("Trip Status: " + element.TripStatus)
		// 				fmt.Println("==============================")
		// 				fmt.Println()
		// 			}
		// 		}
		// 	}
		// }
		fmt.Println("0. Quit")
		fmt.Println("=============================")
		fmt.Print("Please enter an option: ")
		passengerTripMenuInput := 10000
		fmt.Scan(&passengerTripMenuInput)
		switch passengerTripMenuInput {
		case 0:
			fmt.Println("Exiting")
			flag=false
		default:
			val, ok := rideMap[passengerTripMenuInput]
			// If the key exists
			if ok {
				passengerTripItemMenu(val)
			}
		}
	}
}

func passengerTripItemMenu(trip Trip) {
	type TripUpdate struct {
		TripId            int       `json:"tripId"`
		StartTime         string `json:"startTime"`
		EndTime           string `json:"endTime"`
		TripStatus        string    `json:"tripStatus"`
		DriverId string `json:"driverId"`
	}

	flag := true
	for flag {
		fmt.Println()
		fmt.Println("=========Trip "+ strconv.Itoa(trip.TripId) +"=========")
		fmt.Println("Request Time: " + trip.RequestTime.Time.Format(time.RFC822))
		fmt.Println("Pickup Postal Code: " + trip.PickUpPostalCode)
		fmt.Println("Drop off Postal Code: " + trip.DropOffPostalCode)
		if trip.StartTime.Time.IsZero(){
			fmt.Println("Start Time: Pending" )
		}else{
			fmt.Println("Start Time: " + trip.StartTime.Time.Format(time.RFC822))
		}
		if trip.EndTime.Time.IsZero(){
			fmt.Println("End Time: Pending" )
		}else{
			fmt.Println("End Time: " + trip.EndTime.Time.Format(time.RFC822))
		}
		fmt.Println("Trip Status: " + trip.TripStatus)
		fmt.Println()
		fmt.Println("0. Quit")
		fmt.Println("=============================")
		fmt.Print("Please enter an option: ")
		driverTripItemMenuInput := 10000
		fmt.Scan(&driverTripItemMenuInput)
		switch driverTripItemMenuInput {
		case 0:
			fmt.Println("Exiting")
			flag=false
		}
	}
}

func driverMenuOptions() {
	fmt.Println()
	fmt.Println("=========Driver Menu=========")
	fmt.Println("1. View Trips")
	fmt.Println("2. Update account info")
	fmt.Println("0. Quit")
	fmt.Println("=============================")
	fmt.Print("Please enter an option: ")
}
func driverMenu(d Driver) {
	flag := true
	for flag {
		driverMenuOptions()
		driverMenuInput := 10000
		fmt.Scan(&driverMenuInput)
		switch driverMenuInput {
		case 0:
			fmt.Println("Loggin out")
			//Log out endpoint
			client := &http.Client{}
			url := "http://localhost:5000/driver/logout"
			patchBody, _ := json.Marshal(d)
			resBody := bytes.NewBuffer(patchBody)
			if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
				if res, err2 := client.Do(req); err2 == nil {
					if res.StatusCode == 202 {
						fmt.Println("Successfully logged out")
					} else if res.StatusCode == 400 {
						fmt.Println("Error - Bad Request")
					}
				}
			}
			flag = false
		case 1:
			driverTripMenu(d.DriverId)
		case 2:
			//Driver account info current
			fmt.Print("First Name: "+d.FirstName)
			fmt.Print("Last Name: "+d.LastName)
			fmt.Print("Mobile Number: "+d.MobileNumber)
			fmt.Print("Email Addr: "+d.EmailAddr)
			fmt.Print("Password : "+d.Password)
			fmt.Print("License Number : "+d.LicenseNumber)
			fmt.Print("Id Number : "+d.IdNumber)
			//Driver acccount info input
			var tempD Driver = d
			fmt.Print("Please enter your first name: ")
			fmt.Scan(&tempD.FirstName)
			fmt.Print("Please enter your last name: ")
			fmt.Scan(&tempD.LastName)
			fmt.Print("Please enter a mobile number: ")
			fmt.Scan(&tempD.MobileNumber)
			fmt.Print("Please enter an email address: ")
			fmt.Scan(&tempD.EmailAddr)
			fmt.Print("Please enter a password: ")
			fmt.Scan(&tempD.Password)
			fmt.Print("Please enter a license number: ")
			fmt.Scan(&tempD.LicenseNumber)
			//Driver update endpoint
			client := &http.Client{}
			url := "http://localhost:5000/driver"
			patchBody, _ := json.Marshal(tempD)
			resBody := bytes.NewBuffer(patchBody)
			if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
				if res, err2 := client.Do(req); err2 == nil {
					if res.StatusCode == 202 {
						fmt.Println("Successfully updated account")
						d = tempD
					} else if res.StatusCode == 400 {
						fmt.Println("Error - Bad Request")
					}
				}
			}
		}
	}
}

func driverTripMenu(driverId int) {
	flag := true
	for flag {
		type TripDetails struct {
			DriverId int `json:"driverId"`
		}
		var td TripDetails
		td.DriverId = driverId
		type TripsResponse struct {
			Trips map[string]Trip `json:"Trips"`
		}
		client := &http.Client{}
		url := "http://localhost:5001/driver/trip"
		getBody, _ := json.Marshal(td)
		resBody := bytes.NewBuffer(getBody)
		rideMap := make(map[int]Trip)
		fmt.Println()
		fmt.Println("=========Driver Trips=========")
		if req, err := http.NewRequest("GET", url, resBody); err == nil {
			if res, err := client.Do(req); err == nil {
				if body, err2 := ioutil.ReadAll(res.Body); err2 == nil {
					var responseObject TripsResponse
					json.Unmarshal(body, &responseObject)
					tripArray := []Trip{}
					for _, element := range responseObject.Trips {
						tripArray = append(tripArray,element)
					}
					//Sort trip array reverse chronologically
					sort.Slice(tripArray, func(i, j int) bool {
						return tripArray[i].RequestTime.Time.After(tripArray[j].RequestTime.Time)
					})
					var i=1
					for _, element := range tripArray {
						fmt.Println(strconv.Itoa(i)+". Trip " + "(" + element.RequestTime.Time.Format(time.RFC822) + "): "+ element.TripStatus)
						rideMap[i]=element
						i++
					}
				}
			}
		}
		fmt.Println("0. Quit")
		fmt.Println("=============================")
		fmt.Print("Please enter an option: ")
		driverTripMenuInput := 10000
		fmt.Scan(&driverTripMenuInput)
		switch driverTripMenuInput {
		case 0:
			fmt.Println("Exiting")
			flag=false
		default:
			val, ok := rideMap[driverTripMenuInput]
			// If the key exists
			if ok {
				driverTripItemMenu(val)
			}
		}
	}
}

func driverTripItemMenu(trip Trip) {
	type TripUpdate struct {
		TripId            int       `json:"tripId"`
		StartTime         string `json:"startTime"`
		EndTime           string `json:"endTime"`
		TripStatus        string    `json:"tripStatus"`
		DriverId string `json:"driverId"`
	}

	flag := true
	for flag {
		fmt.Println()
		fmt.Println("=========Trip "+ strconv.Itoa(trip.TripId) +"=========")
		fmt.Println("Request Time: " + trip.RequestTime.Time.Format(time.RFC822))
		fmt.Println("Pickup Postal Code: " + trip.PickUpPostalCode)
		fmt.Println("Drop off Postal Code: " + trip.DropOffPostalCode)
		if trip.StartTime.Time.IsZero(){
			fmt.Println("Start Time: Pending" )
		}else{
			fmt.Println("Start Time: " + trip.StartTime.Time.Format(time.RFC822))
		}
		if trip.EndTime.Time.IsZero(){
			fmt.Println("End Time: Pending" )
		}else{
			fmt.Println("End Time: " + trip.EndTime.Time.Format(time.RFC822))
		}
		fmt.Println("Trip Status: " + trip.TripStatus)
		fmt.Println()
		fmt.Println("1. Start Trip")
		fmt.Println("2. End Trip")
		fmt.Println("0. Quit")
		fmt.Println("=============================")
		fmt.Print("Please enter an option: ")
		driverTripItemMenuInput := 10000
		fmt.Scan(&driverTripItemMenuInput)
		switch driverTripItemMenuInput {
		case 0:
			fmt.Println("Exiting")
			flag=false
		case 1:
			if trip.StartTime.Time.IsZero(){
				//Driver trip update endpoint
				var tu TripUpdate
				tu.TripId=trip.TripId
				tu.StartTime=time.Now().Format(time.RFC3339)
				tu.EndTime=trip.EndTime.Time.Format(time.RFC3339)
				tu.TripStatus="Ongoing"
				tu.DriverId=strconv.Itoa(trip.DriverId)
				client := &http.Client{}
				url := "http://localhost:5001/driver/trip"
				patchBody, _ := json.Marshal(tu)
				resBody := bytes.NewBuffer(patchBody)
				if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
					if res, err2 := client.Do(req); err2 == nil {
						if res.StatusCode == 202 {
							fmt.Println("Started Trip")
							flag=false
						} else if res.StatusCode == 400 {
							fmt.Println("Error - Bad Request")
						}
					}
				}
			}else{
				fmt.Println("Error - Trip already started")
			}
		case 2:

			if trip.EndTime.Time.IsZero(){
				if trip.StartTime.Time.IsZero(){
					//Driver trip update endpoint ->Reject trip
					var tu TripUpdate
					tu.TripId=trip.TripId
					tu.StartTime=time.Now().Format(time.RFC3339)
					tu.EndTime=time.Now().Format(time.RFC3339)
					tu.TripStatus="Rejected"
					tu.DriverId=strconv.Itoa(trip.DriverId)
					client := &http.Client{}
					url := "http://localhost:5001/driver/trip"
					patchBody, _ := json.Marshal(tu)
					resBody := bytes.NewBuffer(patchBody)
					if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
						if res, err2 := client.Do(req); err2 == nil {
							if res.StatusCode == 202 {
								fmt.Println("Rejected Trip")
								flag=false
							} else if res.StatusCode == 400 {
								fmt.Println("Error - Bad Request")
							}
						}
					}
				}else{
					//Driver trip update endpoint ->End ongoing trip
					var tu TripUpdate
					tu.TripId=trip.TripId
					tu.StartTime=trip.StartTime.Time.Format(time.RFC3339)
					tu.EndTime=time.Now().Format(time.RFC3339)
					tu.TripStatus="Ended"
					tu.DriverId=strconv.Itoa(trip.DriverId)
					client := &http.Client{}
					url := "http://localhost:5001/driver/trip"
					patchBody, _ := json.Marshal(tu)
					resBody := bytes.NewBuffer(patchBody)
					if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
						if res, err2 := client.Do(req); err2 == nil {
							if res.StatusCode == 202 {
								fmt.Println("Ended Trip")
								flag=false
							} else if res.StatusCode == 400 {
								fmt.Println("Error - Bad Request")
							}
						}
					}
				}
			}else{
				fmt.Println("Error - Trip already ended")
			}
			
		}
	}
}