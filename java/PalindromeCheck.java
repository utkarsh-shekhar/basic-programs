import java.util.Scanner;

/**
 * To check if a given string is a palindrome
 */
public class PalindromeCheck {

    public static void main(String[] args) {
        // Accept input from the console
        System.out.print("Please enter a string: ");
        String s = null;
        try {
            Scanner sc = new Scanner(System.in);
            s = sc.next();
            sc.close();
        } catch (Exception e) {
            outputAndEnd("Unexpected error, please try to re-run the program."); // Error message for any unknown issues while reading the input
        }

        char[] chars = s.toCharArray();
        int len = chars.length;

        /*
        Loop through all the characters and evaluate if given string is a palindrome
        For even length strings - Say SaaS (or Java)
        S(J) and a(a) from beginning are compared against S(a) and a(v) from the end

        For odd length strings - Say madam (or fruit)
        m(f) and a(r) from beginning are compared against m(t) and a(i) from the end, the middle character d(u) is ignored
         */
        for (int i = 0; i < len / 2; i++) {
            if (chars[i] != chars[len - 1 - i]) {
                outputAndEnd(s + " is not a palindrome"); // Report failure when the first mismatch is found.
            }
        }
        // If no mismatch is found and the loop is successfully complete, report the success.
        outputAndEnd(s + " is a palindrome");
    }

    private static void outputAndEnd(String message) {
        System.out.println(message); // output the given message
        System.exit(0); // exit the program
    }
}
