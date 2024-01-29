package helper

var version = "1.0.0"       //private
var Applications = "golang" //public

func sayGoodBye(name string) string { //private
	return "Good Bye " + name
}
func SayHello(name string) string { //public
	return "Hello, " + name
}
