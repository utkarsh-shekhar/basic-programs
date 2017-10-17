import java.util.Scanner;

/*
 * Find the Greatest Common Dividor of 2 numbers
 */
class GCD {

  public static void main(String[] args) {

    int num1 = 0, num2 = 0;

    // Accept input from the console
    System.out.println("Enter 2 integers:");

    try {
      Scanner scanner = new Scanner(System.in);

      // Gets the two inputs as integers
      num1 = scanner.nextInt();
      num2 = scanner.nextInt();

      scanner.close();

      // Calls the gcd function
      System.out.println("The GCD is: " + gcd(num1, num2));

    } catch (Exception e) {
      // Display error message here
      System.out.println("Error whilst reading user input");
    }

  }

  // The gcd recursive function
  public static int gcd(int a, int b) {
    if (b == 0) {
      return a;
    } else {
      return gcd(b, a % b);
    }
  }
}
