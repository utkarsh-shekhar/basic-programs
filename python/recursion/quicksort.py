from random import randint

def _quicksort(arr, lo, hi):
    if lo < hi:
        p = partition(arr, lo, hi)
        _quicksort(arr, lo, p-1)
        _quicksort(arr, p+1, hi)

def partition(arr, lo, hi):
    #pivot is chosen as first element of list
    #shuffle pivot ie first element with any other element randomly
    randidx = randint(lo,hi)
    arr[randidx], arr[lo] = arr[lo], arr[randidx]
    pivot = arr[lo]
    left, right = lo, hi
    
    while True:
        
        while arr[left] <= pivot:
            left += 1
            if left == hi:
                break
                
        while arr[right] > pivot:
            right -= 1
            if right == lo:
                break
                
        if (left >= right):
            break
            
        arr[left], arr[right] = arr[right], arr[left]
        

    
    arr[lo], arr[right] = arr[right], arr[lo]
    return right
    


def quicksort(arr):
    if type(arr) is list:
        _quicksort(arr, 0, len(arr)-1)
            
        
            