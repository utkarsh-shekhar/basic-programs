def factorial(n: int) -> int:  # Python >= 3.5
    if n == 0:
        return 1
    if n < 0:
        raise ValueError('Input needs to be a non-negative integer: ' + str(n))
    return n * factorial(n - 1)

if __name__ == "__main__":  # executes only upon import
    try:
        print("Factorial of 3 is: "+str(factorial(3)))
        print("Factorial of 6 is: "+str(factorial(6)))
        print("Factorial of -9 is: "+str(factorial(-9)))
    except Exception as ex:
        print(ex)
