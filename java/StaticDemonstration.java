import java.util.Date;

// Static imports are used to import static variables and methods as below
import static java.lang.Math.PI; // Static variable
import static java.lang.Math.abs; // Static method

/**
 * This class intends to demonstrate the usage of static keyword in Java
 */
public class StaticDemonstration {

    private final String name;

    private static final String time = new Date().toString();

    /* Static block gets executed when the class is loaded by the JVM.
        This happens at the beginning of the execution.
     */
    static {
        System.out.println("Class being loaded");
    }

    public StaticDemonstration(String name) {
        this.name = name;
    }


    /* Main method is a static method meaning that it can be
        called without requiring an instance of the containing class.
     */
    public static void main(String[] args) {
        System.out.println("Value of PI: " + PI);
        System.out.println("Value of abs(-1): " + abs(-1));
        System.out.println("Time is: " + time); //static main method can use static variables in the class without an instance
        greet(); // static main method can call other static methods in the class without an instance


        StaticDemonstration staticDemonstration = new StaticDemonstration("Java Developer");
        staticDemonstration.greetUser(); // To call a non-static method, an instance is needed
        System.out.println("Name is: " + staticDemonstration.name); // To access non-static members, an instance is needed

        // Static nested class can be initialized without requiring an instance of the enclosing class
        StaticNestedClass staticNestedClass = new StaticNestedClass();
        // Non static inner class requires an instance of enclosing class to be instantiated
        NonStaticInnerClass nonStaticInnerClass = staticDemonstration.new NonStaticInnerClass();
    }

    private static void greet() {
        System.out.println("Hello user. Have a good day.");
    }

    private void greetUser() {
        System.out.println("Hello " + name + ". Have a good day.");
    }

    private static class StaticNestedClass {
        private StaticNestedClass() {
            System.out.println("StaticNestedClass constructor");
        }
    }

    private class NonStaticInnerClass {

        private NonStaticInnerClass() {
            System.out.println("NonStaticInnerClass constructor");
        }
    }

}