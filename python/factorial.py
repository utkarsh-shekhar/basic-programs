def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

# change the value for a different result
print(factorial(5))

##### an alternative model for the negative number to put them equal to one. because does not exist Factorial for negative numbers. ####

n=int(input("Enter number:"))
k=1
while(n>0):
    k=k*n
    n=n-1
print("Factorial of the number is: ")
print(k)
