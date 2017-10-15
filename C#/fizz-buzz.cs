static public void Main()
    {
        for (int i =1; i < 100; i++)
        {
            Console.WriteLine((i%3==0 && i%5==0)?"Fizz Buzz":i%3==0 ?"Fizz": i%5==0 ?"Buzz":i.ToString());
        }

    }
