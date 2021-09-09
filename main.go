package main
 
import (
    "fmt"
    "strconv"
    "os"
  )
var menu int
var ID int = 0
var title, description, saver string
var rating float64
var data = [][4]string{{"ID", "Title", "Rating", "Description"}}

func add(){
  fmt.Println("Input game title : ")
  fmt.Scan(&title)
  ID++
  fmt.Println("Input your rating : ")
  fmt.Scanf("%f", &rating)
  if rating >=5.0 {
    rating = 5.0
    description = "Good"
  }else if rating > 4.0 {
    description = "Good" 
  } else if rating >= 2.0 && rating <= 4.0 {
    description = "Average"
  } else if rating < 2.0 {
    description = "Poor"
  } else if rating < 1.0 {
    rating = 1.0
    description = "Poor"
  }
  newdata := [4]string{(strconv.Itoa(ID)), title, (strconv.FormatFloat(rating, 'f', 1,64)), description}
 data = append(data, newdata)
  newdata = [4]string{}
}
func showAll(){
  fmt.Println("List of game available : \n", len(data))
  for i:=0;i<len(data);i++{
    for j:=0;j<4;j++{
      fmt.Print(data[i][j], "  \t")
    }
    fmt.Println("")
    }
}

func searchID(ID int) int{
  var index int
  search := strconv.Itoa(ID)
  for i:=0;i<len(data);i++{
    if data[i][0] == search{
      index = i
      break
    }  
  }
  return index
}

func deleteID(ID int){
    var before, after [][4]string
    index := searchID(ID)
    before = data[:index]
    after = data[index+1:]
    data = [][4]string{}
    for i:=0;i<len(before);i++{
      data = append(data, before[i])
    }
    for i:=0;i<len(after);i++{
      data = append(data, after[i])
    }
}

func searchTitle(title string){
  fmt.Println("Search result(s) : ")
  for i:=0;i<len(data);i++{
    if data[i][1] == title{
      fmt.Println("ID : ",data[i][0])
      fmt.Println("Title : ",data[i][1])
      fmt.Println("Rating : ",data[i][2])
      fmt.Println("Description : ",data[i][3])
      }  
}
}
func stringtoFloat(input string) float64{
  var save float64
  if s, err := strconv.ParseFloat(input, 64); err == nil {
    save = s
}
  return save
}
func topthree(){
    if len(data) < 4 {
      fmt.Println("Sorry the data is too short, please add more")
    } else {
    var third,first,second [4]string
    for i:=1;i<len(data);i++{
      if stringtoFloat(data[i][2]) > stringtoFloat(first[2]) {
        third = second
        second = first
        first = data[i]
      } else if stringtoFloat(data[i][2]) > stringtoFloat(second[2]) {
        third = second
        second = data[i]
      } else if stringtoFloat(data[i][2]) > stringtoFloat(third[2]) {
        third = data[i]
      }
    }
    fmt.Print("Top 3 rated games : \n")
    fmt.Print(data[0][0], "\t", data[0][1], "\t", data[0][2], "\t", data[0][3], "\n")
    fmt.Print(first[0], "\t", first[1], "\t", first[2], "\t", first[3], "\n")
    fmt.Print(second[0], "\t", second[1], "\t", second[2], "\t", second[3], "\n")
    fmt.Print(third[0], "\t", third[1], "\t", third[2], "\t", third[3], "\n")
    }
}

func onlyfour(){
  fmt.Println("List of games with >4.0 rating")
  for i:=1;i<len(data);i++{
    if stringtoFloat(data[i][2]) > 4.0 {
      fmt.Print(data[i][0], "\t", data[i][1], "\t", data[i][2], "\t", data[i][3], "\n")
    }
  }
}

func main(){
  fmt.Println("Welcome to Game Review! Please select menu :\n1. Add new game data\n2. Delete a game data\n3. Show all game data\n4. Search game data (full title search)\n5. Top 3 rated games\n6. Game data with >4.0 rating\n7. Exit")
  fmt.Scan(&menu)
  if menu == 1 {
    add()
    main()
    } else if menu == 2 {
    fmt.Println("What ID you want to delete? : ")
    fmt.Scanf("%s", &saver)
    var num, err = strconv.Atoi(saver)
    if err == nil && num > 0 {
    deleteID(num)
    }
    showAll()
    main()
  } else if menu == 3 {
    showAll()
    main()
  } else if menu == 4 {
    fmt.Println("What title do you wish to search? : ")
    fmt.Scanf("%s", &saver)
    searchTitle(saver)
    main()
  } else if menu == 5 {
    topthree()
    main()
  } else if menu == 6 {
    onlyfour()
    main()
  } else if menu == 7 {
     fmt.Println("Thanks")
     os.Exit(1)
  }   
  
}