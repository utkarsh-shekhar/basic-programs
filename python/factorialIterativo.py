
"""
 Program performs the factorial of n, in iterative form. (n!)
"""

def factorial(n):
    result = 1
 
    for i in range(n):
        result *= (i + 1)
 
    return result
  


value = input("Enter with value: ")
fact = factorial(value)

print ("\nThe fatorial the ", value, " is: ", fact)