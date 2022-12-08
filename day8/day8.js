fs = require('fs')

const file = fs.readFileSync("./input.txt", "utf8");
const lines = file.split('\n')

function sum(n) {
  return n.reduce((partialSum, a) => partialSum + a, 0);
}

const grid = []
lines.forEach( i => grid.push(Array(lines[0].length)))

const ll = lines[0].length

for (var i=0; i< lines.length; i++) {
  const line = lines[i]

  let left = 0
  let right = 0
  for (k=0; k < line.length; k++) {
    const a = ~~line[k]
    const b = ~~line[line.length - 1 - k]
    if (a > left || k == 0) {
      left = a
      grid[i][k] = 1
    }
    if (b > right || k == 0) {
      right = b
      grid[i][line.length - 1 - k] = 1
    }
  }

} 

for (k=0; k < ll; k++) {
  let left = 0
  let right = 0
  for (var i=0; i< lines.length; i++) {
    const line = lines[i]
    const line2 = lines[lines.length - 1 - i]

    const a = ~~line[k]
    const b = ~~line2[k]
    if (a > left || i == 0) {
      left = a
      grid[i][k] = 1
    }
    if (b > right || i == 0) {
      right = b
      grid[lines.length - 1 - i][k] = 1
    }
  }
}

let h = sum(grid.map(i => sum(i)))
console.log(h)