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
let score = 0;

for (var i=0; i < lines.length; i++) {
  const line = lines[i]
  const ll = line.split(',');
  const a = ll[0]
  const b = ll[1]
  var n = a.split('-').map(i => parseInt(i))
  var g = b.split('-').map(i => parseInt(i))

  let swap = false
  if (g[0] < n[0]) {
    const v = g
    g = n
    n = v
    swap = true
  }

  // console.log(n[0] + '-'  + n[1] + ' ' + g[0] + '-' + g[1] + ' ' + swap + ' ' + ll)
  if (g[0] <= n[1] || g[0] == n[0]) {
    score += 1
  }
}

console.log(score)