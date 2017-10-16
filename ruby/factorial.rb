puts 'Hello there, Can you tell me the number you want factotial of?'
number = gets.to_i
fact = 1;

(1..number).each do |i|
    fact = fact*i;
end

print "Factorial of ", number ," is: ", fact
