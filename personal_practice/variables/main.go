package main

import "fmt"

// !IMR:  First type: declare variable outside of the function: declaration +  type + initiaization in same line
// !IMR:  Here we provide explicit type:
var name string = "Mostafizur Rahaman"
var age int = 20
var height float64 = 5.6
var isStudent bool = false

// !IMR: Second type: declare and initialize in same line.
// !IMR: from first initialized value type infered
var firstName = "Mostafiz"
var myAge = 20
var myHeight = 3.5
var amIJobHolder = true

// !IMR: Third type: Declare only with type definication
// !If  we don't assign any value, the default value will be "FALSY" Value
// !Falsy values are ("", 0, false, nil)
var lastName string
var lastAttemps int
var retryCount float64
var isRetryable bool

func main() {

	fmt.Println("-------------First One---------------")
	fmt.Println("Name", name)
	fmt.Println("Age", age)
	fmt.Println("Height", height)
	fmt.Println("Is Student", isStudent)
	fmt.Println()
	fmt.Println("------MODIFICATION  Variables (First) ------------")
	fmt.Println()

	// !IMR: Here we are going to modify  first block  of variables:
	name = "Fahim Hossain"
	age = 24
	height = 5.7
	isStudent = true
	fmt.Println("Name", name)
	fmt.Println("Age", age)
	fmt.Println("Height", height)
	fmt.Println("Is Student", isStudent)
	fmt.Println()
	fmt.Println("------MODIFICATION------------")
	fmt.Println()

	fmt.Println("---------------------------")

	// !IMR: Second type of variables:
	fmt.Println("-------------Second of Variables---------------")
	fmt.Println()
	fmt.Println("First Name", firstName)
	fmt.Println("MyAge", myAge)
	fmt.Println("myHeight", myHeight)
	fmt.Println("amIJobHolder? ", amIJobHolder)

	fmt.Println("-------------- we can  do modifiication (Reassingnment)")
	firstName = "Ratul Hossain"
	myAge = 30
	myHeight = 6.2
	amIJobHolder = false

	fmt.Println("FirstName -> ", firstName, "myAge ->", myAge, "myHeight ->", myHeight, "amIJobHolder ->", amIJobHolder)

	fmt.Println()
	fmt.Println("-------------Second of Variables End---------------")

	// !IMR: Third type of variable:
	fmt.Println("-------------Third block of variabls---------------")
	fmt.Println("-")
	fmt.Println("lastName", lastName)
	fmt.Println("lastAttemps", lastAttemps)
	fmt.Println("retryCount", retryCount)
	fmt.Println("isRetryable", isRetryable)
	fmt.Println("-")

	// ! now we are going to assign value into unassign third type of variables
	lastName = "Rony Broker"
	lastAttemps = 3
	retryCount = 5
	isRetryable = true

	fmt.Println("After assgining value -:", "lastName :", lastName, "lastAttemps : ", lastAttemps, "retryCount :", retryCount, "isRetryable :", isRetryable)

	fmt.Println("-------------Third block of variabls---------------")

	// ! IMR: Forth type variable: (Shorthand method)
	// should assign value same line
	// cannot declare outside of the main function
	// cannot access outside of the function
	// reassignable
	schoolName := "Bamni Adarsha High School"
	schoolEIIN := 107323
	schoolScore := 5.0
	isWellKnownSchool := true

	fmt.Println("-------------Forth block of variabls---------------")
	fmt.Println("-")
	fmt.Println("schoolName", schoolName)
	fmt.Println("schoolEIIN", schoolEIIN)
	fmt.Println("schoolScore", schoolScore)
	fmt.Println("isWellKnownSchool", isWellKnownSchool)
	fmt.Println("-")

	// override the value of the shorthand variables:
	schoolName = "BASH"
	schoolEIIN = 107223
	schoolScore = 5.2
	isWellKnownSchool = false

	fmt.Println("After override the value : _", "schoolName", schoolName, "schoolEIIN", schoolEIIN, "schoolScore", schoolScore, "isWellKnownSchool", isWellKnownSchool)

	fmt.Println("-------------Forth block of variabls---------------")

	// ! Multiple variable declaration:
	// ! For same type we can declare multiple with type :
	var i, f int = 10, 20
	fmt.Println(i, f)

	// ! If types are different don't need to sepcify  the type, without explicit type declaration we have to assign variable
	//! here we can use var or := both

	var a, b, c = 20, "Mostafiz", true
	fmt.Println(a, b, c)

	n1, n2, n3 := 1, "Rahul", false
	fmt.Println(n1, n2, n3)

	// ! Declare block of (Group of variable):

	var (
		friendName    string  = "Fahim"
		friendAge     int     = 20
		friendHeight  float32 = 3.4
		isBestFriend  bool    = false
		friend2Name           = "Ami"
		friend2Age            = 20
		friend2Height         = 5.6
		isRealFriend          = true
		isUndefined   bool
	)

	fmt.Println("friendName", friendName)
	fmt.Println("friendAge", friendAge)
	fmt.Println("friendHeight", friendHeight)
	fmt.Println("isBestFriend", isBestFriend)
	fmt.Println("friend2Name", friend2Name)
	fmt.Println("friend2Age", friend2Age)
	fmt.Println("friend2Height", friend2Height)
	fmt.Println("isRealFriend", isRealFriend)
	fmt.Println("isUndefined", isUndefined)
}
