<?php

function fizz_buzz_test()
{
    for($i=1; $i <= 100; $i++){
        if($i % 3 == 0 && $i % 5 == 0){
            echo 'FizzBuzz';
            echo '</br>';
        }
        else if($i % 3 == 0){
            echo 'Fizz';
            echo '</br>';
        }else if($i % 5 == 0){
            echo 'Buzz';
            echo '</br>';
        }else{
            echo $i;
            echo '</br>';            
        }
    }
}

fizz_buzz_test();

?>