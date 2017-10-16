function factorial(n) {
  if (n < 0 ) throw new Error('Argument must be a positive integer.');

  if (n === 0 || n === 1) return 1;

  return n * factorial(n - 1);
}

const inputs = [4, 6, 8, 2, -1];

inputs.forEach(n => {
  console.log(`Factorial of ${n}: ${factorial(n)}`);
});