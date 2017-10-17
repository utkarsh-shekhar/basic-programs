/* This program demonstrate the method overloading
 */

class OverloadDemo
{
    void test()
    {
        System.out.println("No parameters");
    }

    /* Overload the test for one integer parameter */
    void test(int a)
    {
        System.out.println("a : " + a);
    }

    /* Overload the test for two integer parameters */
    void test(int a, int b)
    {
        System.out.println("a is " + a + "\nb is " + b);
    }

    /* Overload the test for a double parameter */
    double test(double a)
    {
        System.out.println("double a : " + a);
        return a*a;
    }
}

class Overload
{
    public static void main(String args[])
    {

        OverloadDemo ob = new OverloadDemo();
        double result;

        /* all the version of method test() */
        ob.test();
        ob.test(100);
        ob.test(100, 200);

        result = ob.test(123.25);

        System.out.println("Result of ob.test(123.25) is " + result);

    }
}
