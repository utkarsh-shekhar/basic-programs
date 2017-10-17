/**
 * An example of method overloading with JavaScript, also known as 'monkey patching'.
 * JavaScript as a language does not inherently support method overloading, but we can
 * implement our own solution as we have access to the array of arguments supplied to the function.
 * This gives us substantially more control over the method overloading interface.
 */
class Adder {
  /**
   * Entry function to handle addition with different arguments
   *
   * @params {int} a First integer
   * @params {int} b Second integer
   * Or;
   * @params {int} a First integer
   * @params {int} b Second integer
   * @params {int} c Third integer
   * Or;
   * @params {Array} a Array of integers
   *
   * @returns {int} Sum of arguments
   * @throws {error} If arguments do not match one of the above params
   */
  add(a, b) {
    const argArray = Array.prototype.slice.call(arguments);
    // If the length is 2, we have `a` and `b` as arguments
    if (argArray.length === 2) {
      return a + b;
    }

    // If the length is 3, we have `a`, `b`, and `c` as arguments
    if (argArray.length === 3) {
      return this.addThreeArgs.apply(this, arguments);
    }
    
    // If `a` is an Array, we can sum using `addArray`
    if (Array.isArray(a)) {
      return this.addArray.apply(this, arguments);
    }

    // If none of the above match, we were given invalid input
    throw new Error('Please use one of the following formats: add(1, 2), add(1, 2, 3) or add([1, 2, 3]).');
  }

  /**
   * Sum three numbers
   *
   * @params {int} a First integer
   * @params {int} b Second integer
   * @params {int} c Third integer
   * @returns {int} Sum of a, b, and c
   */
  addThreeArgs(a, b, c) {
    return a + b + c;
  }

  /**
   * Reduce array values to a single sum value of array items
   *
   * @params {Array} a Array of integers
   * @returns {int} Sum of items in array
   */
  addArray(arr) {
    return arr.reduce((acc, cur) => acc + cur);
  }
}

const adder = new Adder();

const addWithTwo = adder.add(1, 1); // 2
console.log(addWithTwo);

const addWithThree = adder.add(1, 2, 3); // 6
console.log(addWithThree);

const addWithArray = adder.add([1, 2, 3, 4]); // 10
console.log(addWithArray);
