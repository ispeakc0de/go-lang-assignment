package main
import "fmt"

var m map[string]string

//init method
func init(){
m=make(map[string]string)
}

//function add
func add(key,value string) {
if lan, ok := m[key]; ok {
		fmt.Println("duplicate data with value : ",lan)
		}else {
		m[key]=value
		}
 }


//update map
func update(key,updatedValue string) {
if _, ok := m[key]; ok {
		m[key]=updatedValue
	} else {
		fmt.Println("\nKey doesn't exists")
	}
}


//get value

func get(key string) {
if lan, ok := m[key]; ok {
		fmt.Println(lan)
	} else {
		fmt.Println("\nKey doesn't exists")
	}
}

//get all value

func getall() {
for k, v := range m{
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
}

//delete value

func deleteValue(key string) {
delete(m, key)
}

//main function
func main() {
add("name","shubham")
getall()
add("name","shubham")
update("name","shubh")
getall()
add("roll","1605598")
get("roll")
deleteValue("name")
getall()

}
