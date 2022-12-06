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

const line =lines[0]

for (var i=14; i< line.length; i++) {
  v = []
  let o = true
  for (var t=0; t<14;t++) {
    const u = line[i-t]
    if (v.indexOf(u) >= 0) {
      o = false
    }
    v.push(u)
  }
  if (o) {
    console.log(i+1)
    break
  }
} 

