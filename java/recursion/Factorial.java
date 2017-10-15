            import java.util.*;
     
            class Factorial{  
             static int factorial(int n){    
              if (n == 0)    
                return 1;    
              else    
                return(n * factorial(n-1));    
             }    
             public static void main(String args[]){  
              Scanner sc = new Scanner(System.in);
              int i,fact=1;  
              int number = sc.nextInt();//It is the number to calculate factorial    
              fact = factorial(number);   
              System.out.println("Factorial of "+number+" is: "+fact);    
             }  
            }  
