/*
 An example of method overloading in Swift v 3 & 4
 sayHi method is overloaded here
*/
import Foundation

class Greeter {
	
	func sayHi(name : String) -> String {
		return "Hi! \(name)"
	}
	
	func sayHi(name: String, country: String) -> String {
		var greetings = "";
		switch country {
		case "India":
			greetings = "Namaste! \(name)"
			
		case "Portugal", "Brazil":
			greetings = "Ola! \(name)"
			
		case "Spain":
			greetings = "Hola! \(name)"
			
		case "France":
			greetings = "Bonjour! \(name)"
			
		default:
			greetings = "Hi! \(name)"
		}
		//trim trailing and leading white space characters
		greetings = greetings.trimmingCharacters(in: .whitespaces)
		return greetings
	}
	
}

let greeter = Greeter()
    
print(greeter.sayHi(name:"John"));
print(greeter.sayHi(name:"Alex", country:"France"))
print(greeter.sayHi(name:"Adarsh", country:"India"))
print(greeter.sayHi(name:"David",country:"Brazil"))

