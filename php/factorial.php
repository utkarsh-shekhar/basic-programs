<?php

/**
* Calculate a factorial given a number.
* For different results, change the function call.
*/
function factorial($number) { 

    if ($number < 2) { 
        return 1; 
    } else { 
        return ($number * factorial($number-1)); 
    } 
}

$res = factorial(4);
echo $res;
