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

const pos = {}

pos[0] = {}
pos[0][0] = 1

let x = 0
let y = 0

let tx = 0
let ty = 0

let score = 1

for (var i=0; i< lines.length; i++) {
  const line = lines[i]

  const [d, amount] = line.split(' ')

  for (let i=0; i < ~~amount; i++) {
    let x2 = x
    let y2 = y
    if (d == 'L') {
      x--
      if (Math.abs(x - tx) > 1 || Math.abs(y - ty) > 1) {
        ty = y
        tx = x2
      }
    }
    if (d == 'R') {
      x++
      if (Math.abs(x - tx) > 1 || Math.abs(y - ty) > 1) {
        ty = y
        tx = x2
      }
    }
    if (d == 'D') {
      y++
      if (Math.abs(x - tx) > 1 || Math.abs(y - ty) > 1) {
        tx = x
        ty = y2
      }
    }
    if (d == 'U') {
      y--
      if (Math.abs(x - tx) > 1 || Math.abs(y - ty) > 1) {
        tx = x
        ty = y2
      }
    }

    if (!pos[tx]) {
      pos[tx]={}
    }
    if (!pos[tx][ty]) {
      pos[tx][ty] = 1
      score++
      // console.log(d, i, x, y)
    }
  }
} 

console.log(score)