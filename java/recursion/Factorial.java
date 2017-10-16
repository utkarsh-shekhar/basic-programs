

import java.util.Scanner;

/**
 * Created by buddhimah on 10/16/2017.
 */
public class Factorial {

    public static  int fac(int num){

        int r;

        if(num==1){
            return 1;

        }
        r= fac(num-1)*num;

        return r;




    }

    public static void main(String args[]){

        Scanner sc = new Scanner(System.in);
        System.out.println("Enter a number: ");

        int num =sc.nextInt();

        System.out.println(fac(num));




    }
}
