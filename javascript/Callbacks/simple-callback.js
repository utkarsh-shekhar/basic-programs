// callbacks as simple functions
function sayHello (name) {
  console.log(`Hello ${name}`);
}

function askQuestion (name) {
  console.log(`${name}, how are you today?`);
}

/**
 * @param  {string}    Whatever name you like
 * @param  {Function}  Just another function to "extend" the functionality of this function
 */
function main (name, callback) {
  callback(name);
}

main('John Doe', sayHello);
main('Jane Roe', askQuestion);