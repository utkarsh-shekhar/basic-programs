class Armstrong{  
  public static void main(String[] args)  {  
    int c=0,a,temp;  
    int n=153;//It is the number to check armstrong  
    temp=n;  
    d0 {  
    a=n%10;  
    c=c+(a*a*a); 
    n=n/10;  
    } while(n>0);
    if(temp==c)  
        System.out.println("armstrong number");   
    else  
        System.out.println("Not armstrong number");   
   }  
}  
