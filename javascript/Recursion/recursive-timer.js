// A countdown timer using recursion

const time = new Date().getTime();

const endTime = time + 10000;

function startTimer (end) {
  let timerID = setTimeout(() => {
    
    // These lines are simply for illustrative purposes
    let sound = (timerID % 2 === 0) ? 'tick' : 'tock';
    console.log(sound);
    // 
    
    if (new Date().getTime() >= endTime) {
      console.log('finished!');
    } else {
      return startTimer(end);
    }
  }, 1000);
}

startTimer(endTime);
