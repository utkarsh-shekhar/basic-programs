import java.util.*

fun main(arg: Array<String>) {
    val scanner = Scanner(System.`in`)
    val number = scanner.nextInt()
    var factorial = 1
    (2..number).forEach { factorial *= it }
    System.out.println("Factorial of $number is: $factorial")
}