using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApplication3
{
   
    class Program
    {
        private static int gcd(int n1, int n2)
        {
            int rem = 5;
            while (n2 > 0)
            {
                rem = n1 % n2;
                if (rem == 0)
                    return n2;
                n1 = n2;
                n2 = rem;

            }
            return n1;
        }

            static void Main(string[] args)
        {

            int tc = int.Parse(Console.ReadLine());
            while (tc > 0)
            {
                List<int> inp = Console.ReadLine().Split(' ').Select(s => int.Parse(s)).ToList<int>();
                int s = inp[0];
                int d = inp[1];
                int n = inp[1];

                List<int> intarr = Console.ReadLine().Split(' ').Select(s => int.Parse(s)).ToList<int>();
                int max = intarr.Max();
                if (intarr.Count(x => x == max) > 1)
                {
                    Console.WriteLine("Many Roads");
                }
                else {
                    Console.WriteLine(intarr.IndexOf(max));
                }
                tc--;

            }
        }
    }
}
