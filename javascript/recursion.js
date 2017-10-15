// Basic function using recursion to compute the sum of an array of integers

function SumOfArray(myArray) {
    
    // Exit condition
    if(myArray.length === 1) {
        return myArray[0];
    } 
    else {
        return myArray.pop() + SumOfArray(myArray);
    }
};