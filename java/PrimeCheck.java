import java.util.InputMismatchException;
import java.util.Scanner;

/**
 * To check whether a given number is prime
 */
public class PrimeCheck {

    public static void main(String[] args) {

        // Accept input from the console
        System.out.print("Please enter a natural number: ");
        int num = 0;
        try {
            Scanner sc = new Scanner(System.in);
            num = sc.nextInt();
            sc.close();
        } catch (InputMismatchException e) {
            outputAndEnd("Input not valid, expected a natural number."); // Error message for all non numerical inputs
        }

        if (num < 0) {
            outputAndEnd("Input not valid, natural number expected"); // Error message for all negative inputs
        } else if (num < 2) {
            outputAndEnd("0 and 1 are neither prime nor composite"); // Info message for 0 and 1
        }

        // Check if the given number is divisible from any number between 2 and num/2, if so, its composite
        for (int i = 2; i <= num / 2; i++) {
            if (num % i == 0) {
                outputAndEnd(num + " is not a prime number."); // Output result and exit as we found one divisor other than 1 and itself
            }
        }
        // The number has exit the loop with no divisors, which means it is a prime number, print success message.
        outputAndEnd(num + " is a prime number.");
    }

	// A method to print the given message and exit the program
    private static void outputAndEnd(String message) {
        System.out.println(message); // output the given message
        System.exit(0); // exit the program
    }
}
