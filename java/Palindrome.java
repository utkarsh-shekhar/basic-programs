import java.util.Scanner;
public class Palindrome {
    static boolean isPalindrome(String word) {
        char[] letterArray = word.toCharArray();
        int startingIndex = 0;
        int lastIndex = letterArray.length - 1;
        while (lastIndex > startingIndex) {
            if (letterArray[startingIndex] != letterArray[lastIndex]) {
                return false;
            }
            ++startingIndex;
            --lastIndex;
        }
        return true;
    }
    public static void main(String args[]) {
        Palindrome palindrome = new Palindrome();
        while (true) {
            System.out.printf("Enter a word which needs to check whether a palindrome or not: ");
            Scanner scan = new Scanner(System.in);
            String input = scan.next();
            System.out.printf("The word '%s' is a palindrome : %s %n", input, palindrome.isPalindrome(input));
        }
    }
}
