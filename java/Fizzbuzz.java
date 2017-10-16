import java.util.*;
public class Fizzbuzz{
	public static void main(String[] args){
		System.out.println("Enter the number:");
		int n=100;
		for(int i=1;i<=n;i++){
			if(i%5 == 0){
				System.out.println("Buzz");
			}
			else if(i%3 ==0){
				System.out.println("Fizz");
			}
			else if((i%3== 0) && (i%5==0)){
                System.out.println("FizzBuzz");
            }
            else{
            	System.out.println(i);
            }

		}

	}
}