import java.util.Scanner;
public class Insertarray
{
	public static void main(String[] args)
	{
	int n,pos,x;
	Scanner s=new Scanner(System.in);
	System.out.println("Enter no of elements you want in array:");
	n=s.nextInt();
	int a[]=new int[n+1];
	System.out.println("Enter alll the elements:");
	for(int i=0;i<n;i++)
		{
		a[i]=s.nextInt();
		}
	System.out.println("Enter the position where you want to insert   element:");
	pos=s.nextInt();
	System.out.println("Enter the element you want to insert:");
	x=s.nextInt();

	for(int i=(n-1);i>=(pos-1);i--)
	{
	a[i+1]=a[i];
	}
	a[pos-1]=x;
	System.out.print("After inserting:");
	for(int i=0;i<n;i++)
	{
	System.out.print(a[i]+",");
	}
	System.out.println(a[n]);
}
}









