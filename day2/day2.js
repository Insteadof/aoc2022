fs = require('fs')

const file = fs.readFileSync("./data.txt", "utf8");
const lines = file.split('\n')

var score = 0



lines.forEach(line => {
  const [a , b] = line.split(' ')

  //rock
  if (a === 'A' && b === 'X') {
    score += 3
  }
  if (a === 'B' && b === 'X') {
    score += 1
  }
  if (a === 'C' && b === 'X') {
    score += 2
  }

  //paper
  if (a === 'A' && b === 'Y') {
    score += 4
  }
  if (a === 'B' && b === 'Y') {
    score += 5
  }
  if (a === 'C' && b === 'Y') {
    score += 6
  }

//schaar
  if (a === 'A' && b === 'Z') {
    score += 8
  }
  if (a === 'B' && b === 'Z') {
    score += 9
  }
  if (a === 'C' && b === 'Z') {
    score += 7
  }
}) 

console.log(score)