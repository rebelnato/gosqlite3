package examples

import (
	"fmt"
	"syscall"

	"github.com/rebelnato/gosqlite3/crud"
	"golang.org/x/term"
)

/*
 Below function can be called from parent program to initiate interactive examples of CRUD operations
 How to call ?
 import the "github.com/rebelnato/examples" package
 Call the flow like examples.entryExamplesFlow()
 This function isn't gonna return any value , it will only initiate an interactive loop
 Loop progression depends on users inputs
*/

func entryExamplesFlow() {

	var userId, selectedOption string
	fmt.Print("Please provide your name : ")
	fmt.Scanln(&userId)
	fmt.Printf("Hello ! %q , would you like to initiate sqlite CRUD examples ? (Please respond with yes/no)\n> ", userId)
	fmt.Scanln(&selectedOption)
	if selectedOption == "no" {
		fmt.Println("Examples test program interupted")
		return
	} else {
		initiateExampleTest()
		return
	}
}

/*
 Below function will be called from within entryExamplesFlow if user wishes to proceed with CRUD examples
 How to call ? - Gets called automatically based on input provided in entryExamplesFlow function call

 This function isn't gonna return any value , it will only initiate an interactive loop
 Loop progression depends on users inputs
 As part of the loop user will receive option to perform read , insert , updaet and delete on available data in db
*/

func initiateExampleTest() {

	var exampleInputOption, inputUserName, inputPassword string
	fmt.Printf("Please select one of the CRUD operations (read,insert,update,delete). Enter exit if you want to exit the testing.\n> ")
	fmt.Scanln(&exampleInputOption)
	switch exampleInputOption {
	case "read":
		var readType string
		fmt.Println("Starting user data fetch example")
		fmt.Printf("What would you like to fetch from db ? (alluserslist/singleuserdata) \n> ")
		fmt.Scanln(&readType)
		if readType == "alluserslist" {
			fmt.Println("Starting process to fetch all available users list")
			users, err := crud.QueryUserList()
			if err != nil {
				fmt.Printf("Failed to fetch data from db due to error message :%q\n", err)
				initiateExampleTest()
				return
			} else {
				fmt.Printf("List of users from db is : %q \n\n", users)
				initiateExampleTest()
				return
			}
		} else if readType == "singleuserdata" {
			fmt.Println("Please provide username of the user to perform search")
			fmt.Print("User name : ")
			fmt.Scanln(&inputUserName)
			id, username, passwordFromDb, err := crud.QueryData(inputUserName)
			if err != nil {
				fmt.Printf("Failed to fetch data from db due to error message :%q\n", err)
				fmt.Println("Reinitiating the example test flow as the provided user is not found")
				initiateExampleTest()
			} else {
				fmt.Println("Fetched user data is as follows :")
				fmt.Printf("User name : %q\nIndex ID : %d\nUser password : %q", username, id, passwordFromDb)
			}
		} else {
			fmt.Println("Invalid input , please type exact command.")
			fmt.Println("Reinitiating the example test flow as the provided user is not found")
			initiateExampleTest()
			return
		}

		fmt.Println("") // Just adding an extra blank line for better clarity in terminal output
		initiateExampleTest()
		return
	case "insert":
		fmt.Println("Starting user data insertion example")
		fmt.Println("Please provide user name and password in same flow")
		fmt.Print("User name : ")
		fmt.Scanln(&inputUserName)
		fmt.Print("Password : ")
		pass, passReadError := term.ReadPassword(int(syscall.Stdin))
		if passReadError != nil {
			fmt.Printf("Password read failed with error %q", passReadError)
			return
		}
		inputPassword = string(pass)

		insertStatus := crud.InsertData(inputUserName, inputPassword)
		if insertStatus != nil {
			fmt.Printf("\nFailed to insert with error : %q", insertStatus)
		} else {
			fmt.Printf("\nSuccessfully inserted data for user : %q", inputUserName)
		}
		fmt.Println("") // Just adding an extra blank line for better clarity in terminal output
		initiateExampleTest()
		return
	case "update":
		var option, newUsername, newPassword string
		var updateStatus error
		fmt.Println("Starting user data update example")
		fmt.Println("Please provide username of the user to perform update")
		fmt.Print("User name : ")
		fmt.Scanln(&inputUserName)
		fmt.Printf("please provide the purpose of update task ? (username/password)\n> ")
		fmt.Scanln(&option)
		if option == "username" {
			fmt.Printf("Proceeding with username update .\nPlease provide new proposed username : ")
			fmt.Scanln(&newUsername)
			tmp := crud.UpdateUsername(inputUserName, newUsername)
			updateStatus = tmp
		} else if option == "password" {
			fmt.Printf("Proceeding with password update .\nPlease provide new proposed password for username %q : ", inputUserName)
			pass, passReadError := term.ReadPassword(int(syscall.Stdin))
			if passReadError != nil {
				fmt.Printf("Password read failed with error %q", passReadError)
				return
			}
			newPassword = string(pass)
			tmp := crud.UpdatePassword(inputUserName, newPassword)
			updateStatus = tmp
		}

		if updateStatus != nil {
			fmt.Printf("Failed to fetch data from db due to error message :%q\n", updateStatus)
			fmt.Println("Reinitiating the example test flow as the provided user is not found")
			initiateExampleTest()
			return
		} else {
			switch option {
			case "username":
				fmt.Printf("User name for old user name %q has successfully been updated to new user name %q .", inputUserName, newUsername)
			case "password":
				fmt.Printf("Password successfully updated for user %q.", inputUserName)
			default:
				fmt.Println("Invalid option entered, reinitiating the example test flow.")
				initiateExampleTest()
				return
			}
		}
		fmt.Println("") // Just adding an extra blank line for better clarity in terminal output
		initiateExampleTest()
		return

	case "delete":
		var confirmation string
		fmt.Println("Starting user data deletion example")
		fmt.Println("Please provide user name and password for the user whose data needs to be deleted")
		fmt.Print("User name : ")
		fmt.Scanln(&inputUserName)
		fmt.Print("Password : ")
		pass, passReadError := term.ReadPassword(int(syscall.Stdin))
		if passReadError != nil {
			fmt.Printf("Password read failed with error %q", passReadError)
			return
		}
		inputPassword = string(pass)
		fmt.Printf("Please confim whether you want to delete the data associated with user %q ? (yes/no)", inputUserName)
		fmt.Scanln(&confirmation)
		if confirmation == "yes" {
			deleteStatus := crud.DeleteUser(inputUserName, inputPassword)
			if deleteStatus != nil {
				fmt.Printf("\nFailed to delete user data with error : %q", deleteStatus)
			}
			fmt.Println("") // Just adding an extra blank line for better clarity in terminal output
			initiateExampleTest()
			return
		} else if confirmation == "no" {
			fmt.Printf("User denied deletion confimration, restarting the example test flow\n")
			initiateExampleTest()
			return
		} else {
			fmt.Printf("Invalid input provided by user, restarting the example test flow\n")
			initiateExampleTest()
			return
		}

	case "exit":
		fmt.Println("Exiting script based on user input")
		return
	default:
		fmt.Println("Provided input is invalid reinitiating the process.")
		initiateExampleTest()
		return
	}
}
