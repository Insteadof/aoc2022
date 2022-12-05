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

let stacks = []

for (var k=0; k < 9; k++) {
  const stack = []
  stacks.push(stack)
  for (var i=0; i < 8; i++) {
     var a = lines[i][k*4+1]
      if (a != ' ') {
        stack.unshift(a)
      }
  }
}

lines.splice(0, 10)

lines.forEach(line => {
  line = line.replace('move ', '')
  var m = line.split(' ').map(i => parseInt(i))
  console.log(m)
  const u = []
  for (var k=0; k < m[0]; k++) {
    const t = stacks[m[2]-1].pop()
    u.push(t)
  }
  for (var k=0; k < m[0]; k++) {
    stacks[m[4]-1].push(u.pop())
  }
})


var r = ''
for (var k=0; k < 9; k++) {
  var h = stacks[k].pop()
  r += h
}

console.log(r)