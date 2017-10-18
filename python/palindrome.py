def is_palindrome(word):
    elements = list(word)    
    is_palindrome = True
    while len(elements) > 0 and is_palindrome:       
        if elements[0] != elements[-1]:
            is_palindrome = False
        else:
            elements.pop(0)
            if len(elements) > 0:
                elements.pop(-1)
    return is_palindrome
while True:
    target = input("Enter a word which needs to check whether a palindrome or not:")
    print ("  The word '"+target+"' is a palindrome :"+ str(is_palindrome(target)))

