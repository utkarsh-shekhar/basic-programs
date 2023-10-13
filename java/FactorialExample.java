import java.util.*;

public class FactorialExample {

	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		
		System.out.print("Enter the integer you want the factorial of: ");
		int factorial = scanner.nextInt();
		
		while(factorial <= 0) {
			System.out.println("Please enter a positive integer greater than 0.");
			factorial = scanner.nextInt();
		}
		System.out.println("The factorial of " + factorial + " is " + factorialCalculator(factorial));
		scanner.close();
	}
	
	private static int factorialCalculator(int x) {
		if(x <= 1) {
			return 1;
		}
		else {
			return x *= factorialCalculator(x - 1);
		}
	}

}
