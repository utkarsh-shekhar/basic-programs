patch-1
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
num = int(input("Enter a number: "))

factorial = 1

# check if the number is negative, positive or zero
if num < 0:
   print("Sorry, factorial does not exist for negative numbers")
elif num == 0:
   print("The factorial of 0 is 1")
else:
   for i in range(1,num + 1):
       factorial = factorial*i
   print("The factorial of",num,"is",factorial)
