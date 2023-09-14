//Go Master Learning List//

// Golang code for printing current working directory
  
package main
  
import (
    "fmt"
    "os"
)
  
func main() {
  
    // using the function
    mydir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(mydir)
}


// use of time
package main
import (
  "fmt"
  "time"
)

func main () {
  fmt.Println("  __      _")
  fmt.Println("o'')}____//")
  fmt.Println("o'')}____//")
  fmt.Println(" (_(_/-(_/ ")
  p := fmt.Println
  now := time.Now()
  p(now)
}


// multi assignment
package main

import "fmt"

func main () {
  var magicNum, powerLevel int32
  magicNum=2048
  powerLevel=9001
  fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
  amount,unit := 10,"doll hairs"
  fmt.Println(amount, unit, ", that's expensive...")
}

//User input
package main

import "fmt"

func main() {
  fmt.Println("What would you like for lunch?")
  
  // Add your code below:
  var food string
  fmt.Scan(&food)
  fmt.Printf("Sure, we can have %v for lunch.", food)
}

//Sprint
package main

import "fmt"

func main() {
  template := "I wish I had a %v."
  pet := "puppy"
  
  var wish string
  
  // Add your code below:
  wish = fmt.Sprintf(
    template, pet)
  
  fmt.Println(wish)
}

//verbs
package main

import "fmt"

func main() {
  floatExample := 1.75
  // Edit the following Printf for the FIRST step
  fmt.Printf("Working with a %T", floatExample) 
  
  fmt.Println("\n***") // Added for spacing
  
  yearsOfExp := 3
  reqYearsExp := 15
  // Edit the following Printf for the SECOND step
  fmt.Printf("I have %d years of Go experience and this job is asking for %d years", yearsOfExp, reqYearsExp) 
  
  fmt.Println("\n***") // Added for spacing
  
  stockPrice := 3.50
  // Edit the following Printf for the THIRD step
  fmt.Printf("Each share of Gopher feed is $ %.2f", stockPrice) 
}

//If,Else
package main
import "fmt"
func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"
if lockCombo==robAttempt{
  fmt.Println("The vault is now opened.")
} else {
  fmt.Println("Try again!")
}
  
}

// &&, || (and, or)
package main
import "fmt"
func main() {
	rightTime := true
	rightPlace := true
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}
	enoughRobbers := false
	enoughBags := true
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
}

// ! (not)
package main
import "fmt"
func main() {
	readyToGo := true
	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}

//else if
package main
import "fmt"
func main() {
	amountStolen := 64650
	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
  } else if amountStolen >= 5000 {
    fmt.Println("Think of all the candy we can buy!")
  } else {
		fmt.Println("Why did we even do this?")
	}
}

//Switch, case, default
package main
import "fmt"
func main() {
	name := "H. J. Simp"
	// Add your switch statement below:
switch name {
  case "Butch":
  fmt.Println("Head to Robbers Roost.")
  case "Bonnie":
  fmt.Println("Stay put in Joplin")
  case "H. J. Simp":
  fmt.Println("That's the man, GET HIM!")
  default:
  fmt.Println("Just hide!")
}
}
// Short Dec
package main
import "fmt"
func main() {
  if success := true; success {		
    fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}
	amountStolen := 50000
	switch numOfThieves := 5; numOfThieves{
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
	  fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
	  fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}
// Random Num Gen
package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
	// Add your code below:
  
  rand.Seed(time.Now().UnixNano())
  amountLeft := rand.Intn(10000)
  
  fmt.Println("amountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}


//Heist story idea
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  isHesitOn:=true
  rand.Seed(time.Now().UnixNano())
//Sneak
  eludedGuards:=rand.Intn(100)
  if eludedGuards>=50{
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember this is the first step.")
  } else { 
    isHesitOn=false 
    fmt.Println("Plan a better disguise next time?")
  }
//Crack
  openedVault:=rand.Intn(100)
  if isHesitOn==true&&openedVault>=70{
    fmt.Println("Grab and GO!")
  }else if isHesitOn==true&&openedVault<70{
    isHesitOn=false
    fmt.Println("The vault can't be opened")
    }
//Leaving
  LeftSafely:=rand.Intn(5)
  if isHesitOn==true{
    switch LeftSafely{
      case 0:
      isHesitOn=false
      fmt.Println("You set off the alarm!")
      case 1:
      isHesitOn=false
      fmt.Println("You we're found by the guards!")
      case 2:
      isHesitOn=false
      fmt.Println("You left the door locked!")
      case 3:
      isHesitOn=false
      fmt.Println("You set off the alarm!")
      default:
      fmt.Println("Start the getaway car!")
    }
  }
  //How much $$$
  if isHesitOn==true{
    amtStolen:=10000+rand.Intn(10000000)
    fmt.Println("We got $", amtStolen, "Drinks are on me!")
  }
  fmt.Println("isHesitOn is currently:", isHesitOn)
}
//Go Function examples
func multiplier(x int32, y int32) int32 {
  return x * y
}
func multiplier(x, y int32) int32 {
  return x * y
}
func main() {
  var product int32
  product = multiplier(25, 4)
  fmt.Println(product) // Prints: 100
}
func main() {
  var mainX, mainY, newProduct int32
  mainX = 6
  mainY = 7
  newProduct = multiplier(mainX, mainY)
  fmt.Println(newProduct) // Prints: 42
}
package main
import "fmt"

func computeMarsYears(earthYears int) int {
  
  earthDays := earthYears * 365
  marsYears := earthDays / 687
  return marsYears
}

func main() {
  myAge := 25
  earthYears := 45
  myMartianAge := computeMarsYears(myAge)
  fmt.Println(myMartianAge)
}
// Multi return function
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
  if averageGrade > 90 {
    gradeLetter = "A"
  } else if averageGrade > 80 {
    gradeLetter = "B"
  } else if averageGrade > 70 {
    gradeLetter = "C"
  } else if averageGrade > 60 {
    gradeLetter = "D"
  } else {
    gradeLetter = "F"
  }
 
  return gradeLetter, averageGrade 
}
//Another example multi function
package main
import "fmt"

func getLikesAndShares(postId int) (int, int) {
  var likesForPost, sharesForPost int
  switch postId {
  case 1:
    likesForPost = 5
    sharesForPost = 7
  case 2:
    likesForPost = 3
    sharesForPost = 11
  case 3:
    likesForPost = 22
    sharesForPost = 1
  case 4:
    likesForPost = 7
    sharesForPost = 9
  }
  fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
  return likesForPost, sharesForPost 
}

func main() {
  var likes, shares int
  likes, shares=getLikesAndShares(4)
  if likes > 5 {
    fmt.Println("Woohoo! We got some likes.")
  }
  if shares > 10 {
    fmt.Println("We went viral!")
  }
}
//connect to db
package main
import "fmt"
func queryDatabase(query string) string {
  var result string
  connectDatabase()
  defer disconnectDatabase()  
  if query == "SELECT * FROM coolTable;" {
    result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
  }  
  fmt.Println(result)
  return result
}
func connectDatabase() {
  fmt.Println("Connecting to the database.")
}
func disconnectDatabase() {
  fmt.Println("Disconnecting from the database.")
}
func main() {
  queryDatabase("SELECT * FROM coolTable;")
}
//scan input
package main
import (
  "fmt"
  "time"
)

func main() {
  var w float64
  var h float64
  fmt.Println("Type two floats ->: ")
  fmt.Scanln(&w, &h)
  waitForIt(joinTwoStrings("Hi", " there"))
  squareAP(w, h)
}
func joinTwoStrings(prefix string, suffix string) string {
  return prefix + suffix
}
func printOutDate() {
  t := time.Now()
  fmt.Println(t)
}
func waitForIt(message string) {
  defer fmt.Println("Done!")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println("Waiting")
  fmt.Println(message)
}
func squareAP(w float64, h float64) {
  sq:=w*h
  pr:=w+w+w+w
  fmt.Println("Area is",sq, "m2", "", "permeter is", pr, "m")
}
//value address find
func main() {
	treasure := "The friends we make along the way."
  fmt.Println(&treasure)
}
//Address pointer 
var pointerForInt *int
 
minutes := 525600
 
pointerForInt = &minutes
 
fmt.Println(pointerForInt) // Prints 0xc000018038

//Derefernce address and asign new
lyrics := "Moments so dear" 
pointerForStr := &lyrics//implict assignement
 
*pointerForStr = "Journeys to plan" //*used for address refernce
 
fmt.Println(lyrics) // Prints: Journeys to plan

//another Derefernce example 
//Since addHundred() expects a pointer (and pointers are variables that store an address) 
//the final touch was to provide addHundred() the address of x. With that, x is now 101
func addHundred (numPtr *int) {
  *numPtr += 100
}
 
func main() {
  x := 1
  addHundred(&x)
  fmt.Println(x) // Prints 101
}
//Another example
package main
import "fmt"
// Change brainwash to have a pointer parameter
func brainwash(saying *string) {
	// Dereference saying below: 
	*saying = "Beep Boop."
}
func main() {
	greeting := "Hello there!"
	// Call your brainwash() below:
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)
}
//simple loop
package main
import (
    "fmt"
)
func main() {
  for count := 1; count <= 12; count++ {
    fmt.Println(count)
  }
}

//
number := 0 // Initialize a variable to be used inside the loop
for number < 5 {
  fmt.Println(number)
  number++ // Update the variable being used
}

//sphinx.go while loop 'for'
package main
import (
  "fmt"
)

func ask() (int) {
  var input int
  fmt.Print("I am thinking of a number between 1-100: ")
  fmt.Scan(&input)
  return input
}

func main() {
	var guess int
	for guess != 56 {
	  guess = ask() 
	   if guess == 56 {
		fmt.Println("You are correct! You may go through to the treasure!")
	} else if guess != 56 {
	  fmt.Println("nah nigga that ain't it!")}
	}
}
// infinte loop
package main
import (
  "fmt"
)
var number int
func count() {
  var input int
  fmt.Print("Number of guests: ")
  fmt.Scan(&input)
  number = number + input
  fmt.Println("Total guests:", number)
}

func main() {
for {
  count()
  }
}

//
animals := []string{"Cat", "Dog", "Fish", "Turtle"}
for index := 0; index < len(animals); index++ {
  if animals[index] == "Dog" {
    fmt.Println("Found the perfect animal!")
    break // Stop searching the array
  }
}

jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "red"}
for index := 0; index < len(jellybeans); index++ {
  if jellybeans[index] == "green" {
    continue
  }
  fmt.Println("You ate the", jellybeans[index], "jellybean!")
}

//loop break conitune
package main
import ("fmt")
func main() {
  for count := 0; count < 20; count++ {
    if count == 8 {
      continue
    }
    if count == 15 {
      break
    }
    fmt.Println(count)
  }
}
// array (list)
letters := []string{"A", "B", "C", "D"}
for index, value := range letters {
  fmt.Println("Index:", index, "Value:", value)
}
// map (dict)
addressBook := map[string]string{
  "John": "12 Main St",
  "Janet": "56 Pleasant St",
  "Jordan": "88 Liberty Ln",
}
for key, value := range addressBook {
  fmt.Println("Name:", key, "Address:", value)
}

// looping though maps and arrays
package main
import ("fmt")
func main() {
  menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}
  fmt.Println("The menu:")
  for index, value := range menu {
    fmt.Println("Number:", index, "Food item:", value)
  }
  orders := map[string]string{
    "John": "Cheeseburgers",
    "Janet": "Hamburgers",
    "Jordan": "Fries",
  }
  orders["James"] = "Chicken Sandwich"
  
    fmt.Println("\nThe friend's orders:")
  for key, value := range orders {
    fmt.Println(key, "will have a", value)
  }
	
//implict array creation
var triangleNumber [3]int
	triangleNumber[0]=1
	triangleNumber[1]=2
	triangleNumber[2]=3
triangleSides := [3]int{15, 26, 30}
triangleAngles := [...]int{30, 60, 90}

}  

//Slices
// Each of the following creates an empty slice
var numberSlice []int
stringSlice := []string{}
 
// The following creates a slice with elements
names := []string{"Kathryn", "Martin", "Sasha", "Steven"}

array := [5]int{2, 5, 7, 1, 3}
// This is a slice of the whole array
sliceVersion := array[:]
fmt.Println(sliceVersion)
// [2 5 7 1 3]
// This is a slice containing the elements at indices 2, 3, and 4
partialSlice := array[2:5]
fmt.Println(partialSlice)
// [7 1 3]


//append,cap,len
books := []string{"Tom Sawyer", "Of Mice and Men"}
books = append(books, "Frankenstein")
books = append(books, "Dracula")
fmt.Println(books)
fmt.Println(cap(books)
fmt.Pritnln(len(books)
// [Tom Sawyer Of Mice and Men Frankenstein Dracula]

	    
	    func printFirstLastArray(array [4]int) {
    fmt.Println("First", array[0])
    fmt.Println("Last", array[3])
}
 
func printFirstLastSlice(slice []int) {
    length := len(slice)
    if (length > 0) {
        fmt.Println("First", slice[0])
        fmt.Println("Last", slice[length-1])
    }
}
	    
package main
import "fmt"

func main() {
    myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
    changeLastElement(myTutors, "Bobby")
    fmt.Println(myTutors)
}

func changeLastElement(myTutors []string, newName string) {
  myTutors[len(myTutors)-1]=newName
  fmt.Println(myTutors)
}
//Finish for homework, updated fuel vairable
package main
import ("fmt"
    "os"
    "bufio"
    "strconv"
    "strings")
var fuelRemaining int
var fuel int


func fuelGauge(fuel int) {
  fmt.Println("fuel level:",fuel)
}


func calculateFuel(planet string) int {
  const fuel1 int
  switch planet {
    case "Venus":
    fuel1=300000
    case "Mercury":
    fuel1=500000
    case "Mars":
    fuel1=700000
    case "Planet-X":
    fuel1=800000
    default:
    fuel1=0
}
  return(fuel1)
}


func greetPlanet(planet string) {
  fmt.Println("You have arrived at",planet)
}


func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}


func flyToPlanet(planet string, fuel int) int {
  var fuelCost int
  fuelRemaining=fuel
  fuelCost=calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining-=fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}


func fuelCounter(fuel int, fuelRemaining int) int {
    filename := "f.txt"
    var fh *os.File
   // counter := fuel
    _, err := os.Stat(filename)
    if !os.IsNotExist(err) {
        fh, _ = os.Open(filename)
        reader := bufio.NewReader(fh)
        var line string
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("error")
            fmt.Println(err)
            os.Exit(1)
        }
        line = strings.TrimSuffix(line, "\n")
        fuel, err = strconv.Atoi(line)
        if err != nil {
            fmt.Println("error")
            fmt.Println(err)
            os.Exit(1)
        }
    }
 
    fuelUpdate:= fuelRemaining-fuel
    fmt.Printf("%d\n", fuelUpdate)
 
    fh, err = os.Create(filename)
    if err == nil {
        fh.WriteString(fmt.Sprintf("%d\n", fuel))
        fh.Close()
    }
    return(fuelUpdate)
}



func deBug(fuel int) {
  fmt.Println(fuel)
  fmt.Println(fuelRemaining)
}



func main() {  
  fuel:=10000000
  planetChoice:="Mercury"
  fuel=flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  fuelCounter(fuel, fuelRemaining)
  deBug(fuel)

}
