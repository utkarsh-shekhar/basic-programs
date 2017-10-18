fun main(args: Array<String>) {
    if (args.size == 0 || args.size != 2) {
        println("Please provide first and second number")
        return
    }

    val first = args.get(0).toIntOrNull()
    val second = args.get(1).toIntOrNull()

    if (first !is Int || second !is Int) {
        println("Invalid arguments $first $second")
        return
    }

        val gcd = gcd(first, second)

        println("GCD of $first and $second is $gcd.")
}

fun gcd(first: Int, second: Int): Int {
    if (second != 0)
        return gcd(second, first % second)
    else
        return first
}