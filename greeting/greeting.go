package greeting

// import "fmt"

type Salutation struct {
	Name string
	Greeting string
} // some comment

type Printer func(string) ()


func GetPrefix(name string) (Prefix string){
	switch name {
		case "Bob": Prefix = "Mr "
		default: Prefix = "Dude"
	}
	return
}

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	// println(alternate, message)
	if prefix:= "Mr"; isFormal {
		do(prefix+message)
	}
	do(alternate)
}
func CreateMessage(name string, greeting string) (message string, alternate string) {
 println(len(greeting))
 alternate = "alternate..." + message
 return
}
func Print(s string) {
	println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		println(s + custom)
	}
}