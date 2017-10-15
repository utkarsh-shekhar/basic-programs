def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

# change the value for a different result
print(factorial(5))
