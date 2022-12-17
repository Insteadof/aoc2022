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

function makeArray(n, value = 0) {
  const result = []
  for (let i=0;i<n;i++) {
    result.push(value)
  }
  return result
}

let score = 0

const map = []
for (let i=0;i<1000;i++) {
  map.push(makeArray(1000, 0))
}

function drawline(a, b) {
  if (a[0] == b[0]) {
    for (let i=Math.min(a[1], b[1]); i <= Math.max(a[1], b[1]); i++) {
      map[a[0]][i] = 2
    }
  } else {
    for (let i=Math.min(a[0], b[0]); i <= Math.max(a[0], b[0]); i++) {
      map[i][a[1]] = 2
    }
  }
}

for (var i=0; i< lines.length; i++) {
  const line = lines[i]

  const pos = line.split(' -> ').map(n => n.split(',').map(n => ~~ n))
  for (k=0; k < pos.length-1; k++) {
    drawline(pos[k], pos[k+1])
  }
} 

let running = true


while (running) {
  score++
  let x=500
  let y=0

  while (y < 950) {
    if (map[x][y+1] == 0) {
      y++
    } else {
      if (map[x-1][y+1] == 0) {
        x--
        y++
      } else if (map[x+1][y+1] == 0) {
        x++
        y++
      } else {
        break
      }
    }
  }
  map[x][y] = 1

  if (y > 940) {
    running = false
    score--
  }
}

console.log(score)