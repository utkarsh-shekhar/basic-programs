
class Fibonacci
{
    //recursive function
	public static int fibonacci(int n) { 
		if (n <= 0) { 
			return 0; }
		if (n == 1) { 
			return 1; }
		if (n == 2) { 
			return 1; }
		return fibonacci(n - 1) + fibonacci(n - 2);
	}

	public static void main (String[] args)
	{
		int number=10;	//number represents the number of terms in the series
		for (int i = 0; i < number; i++) {
			System.out.printf("%s ", fibonacci(i));
			}
	}
}
