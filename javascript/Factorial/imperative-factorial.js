function factorial(n) {
  if (n < 0 ) throw new Error('Argument must be a positive integer.');

  let result = 1;

  for (let i = 2; i <= n; i++) {
    result = result * i;
  }

  return result;
}

const inputs = [4, 6, 8, 2, -1];

inputs.forEach(n => {
  console.log(`Factorial of ${n}: ${factorial(n)}`);
});