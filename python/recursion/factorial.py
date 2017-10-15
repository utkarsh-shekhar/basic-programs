def factorial(n: int) -> int:  # Python >= 3.5
    if n == 0:
        return 1
    if n < 0:
        return n * factorial(n + 1)
    return n * factorial(n - 1)

if __name__ == "__main__":  # executes only upon import
    print("Factorial of -3 is: "+str(factorial(-3)))
    print("Factorial of 6 is: "+str(factorial(6)))
    print("Factorial of -9 is: "+str(factorial(-9)))
