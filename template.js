fs = require('fs')

const file = fs.readFileSync("./input.txt", "utf8");
const lines = file.split('\n')

function sort(arr, fc) {
  return arr.sort((a, b) => fc(b)-fc(a))
}
function ord(a) {
  return a.charCodeAt(0);
}
function sum(n) {
  return n.reduce((partialSum, a) => partialSum + a, 0);
}

let score = 0

for (var i=0; i< lines.length; i++) {
  const line = lines[i]

  
} 

console.log(score)