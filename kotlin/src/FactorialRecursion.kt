import java.util.*

fun main(arg: Array<String>) {
    val scanner = Scanner(System.`in`)
    val number = scanner.nextInt()
    System.out.println("Factorial of $number is: ${factorial(number.toLong())}")
}

fun factorial(number: Long): Long {
    if (number < 0) {
        throw IllegalArgumentException("Number can not be negative")
    }
    return if (number <= 2) {
        number
    } else {
        number * factorial(number - 1)
    }
}