'''
Assuming numbers between 1 and 100,000
'''
# from random import randint

def countingSort(arr):
    countarr= [0]*100001
    for num in arr:
        countarr[num]+= 1
    for i in range(1,len(countarr)):
        countarr[i]= countarr[i-1] + countarr[i]
    sortedarr= [0]*len(arr)
    for i in range(len(arr)):
        sortedarr[countarr[arr[i]]-1]= arr[i]
        countarr[arr[i]]-= 1
    return sortedarr


# arr= [randint(1,100000) for i in range(10000)]
# print (countingSort(arr))