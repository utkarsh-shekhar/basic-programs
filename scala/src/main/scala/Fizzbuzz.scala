object Fizzbuzz extends App {

  def fizzbuzz(n: Int): String = n match {
    case _ if n % 15 == 0 => "FizzBuzz"
    case _ if n % 3 == 0 => "Fizz"
    case _ if n % 5 == 0=> "Buzz"
    case _ => n.toString
  }

  println((1 to 100).map(fizzbuzz(_)).mkString(" "))
}

