class MethodOverloading
{
  // function that is called on program start
  public static void main(String[] args)
  {
    Adder adder = new Adder();
    int array[] = {1, 2, 3, 4};

    // call and print the results of our functions
    System.out.print("Calling 'adder' with 2 integers: ");
    System.out.println(adder.add(1, 2));

    System.out.print("Calling 'adder' with 3 integers: ");
    System.out.println(adder.add(1, 2, 3));

    System.out.print("Calling 'adder' with an array: ");
    System.out.println(adder.add(array));
  }
}

class Adder
{
  public int add(int a, int b)
  {
    return a + b;
  }

  public int add(int a, int b, int c)
  {
    return a + b + c;
  }

  public int add(int array[])
  {
    int sum = 0;

    // loop through our array and add each item to our sum
    for (int i : array)
    sum += i;

    return sum;
  }
}
