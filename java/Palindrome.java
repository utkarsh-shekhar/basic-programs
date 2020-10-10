import java.util.Scanner;
public class Palindrome {
    static boolean isPalindrome(String word) {
        char[] letterArray = word.toCharArray();
        int i = 0; //i is the starting index value which is initially 0 in all cases
        int l,la; //la is for storing last index value 
        l=letterArray.length; //for storing length of word
        int la = l - 1;
        while (la > i) {
            if (letterArray[i] != letterArray[la]) {
                return false;
            }
            ++i;
            --la;
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
