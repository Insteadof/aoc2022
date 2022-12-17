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
let x=1
let i=0

let inst = ''
let V
let row = ''

for (var cycle=1; cycle < 241; cycle++) {
  let b
  if (!inst) {
    const a = lines[i++]
    b = a.split(' ')
    inst = b[0]
  }

  row += (Math.abs(row.length - x) < 2) ? '#' : '.'

  if (inst == 'noop') {
    inst = ''
  } else if (inst == 'addx') {
    V = ~~b[1]
    inst = 's'
  } else if (inst == 's') {
    x += V
    inst = ''
    V = 0
  }
  if ((cycle) % 40 == 0) {
    console.log(row)
    row = ''
  }
}


console.log(score)