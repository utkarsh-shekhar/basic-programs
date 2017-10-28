# A Python implementation of the old coding interview fizz buzz challenge
# featuring user input
# Coded by John D. Close III
# fizz_buzz.py

def fizzbuzz(fizz, buzz, count):
    """ This is my implementation of the FizzBuzz algorithm. The twist here is that the user enters
        the value of fizz, buzz, and the number of iterations. Fizz and Buzz are meant to be integers
        and trditionally expected to fit the following constraints
        1 < fizz < buzz
        fizz != buzz
        fizz * buzz <= the number of iterations.
        Though you can legitimately enter the higher number as fizz with the program still working,
        any variation from the constraints may result in unexpected results """
        
    fb = fizz*buzz
    for i in range(1, count + 1): #range to start at one, go to 101
        # because python is zero based
        if i%fb == 0:
            print('FizzBuzz')
            continue
        elif i%buzz == 0:
            print('Buzz')
            continue
        elif i%fizz == 0:
            print('Fizz')
            continue
        else:
            print(i)


print('Welcome to the FizzBuzz program (Python 3x version).')
fz = int(input("Enter the Fizz number: "))
bz = int(input("Enter the Buzz number: "))
size = int(input("Enter the number of iterations: "))
fizzbuzz(fz, bz, size)
