fs = require('fs')

const file = fs.readFileSync("./data.txt", "utf8");
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

function hash(a) {
  const obj = {}
  for (var i in a) {
    obj[ord(a[i])] = 1
  }
  return obj
}

var score = 0;

function point(a) {
  const t = a - ord('a')
  if (t>=0 && t<=27) {
    return t + 1
  }
  return a - ord('A') + 27
}

function cmp(a, obj) {
  const h = hash(a)
  for (var i in a) {
    if (obj[ord(a[i])] && h[ord(a[i])]) {
      obj[ord(a[i])] += 1
      h[ord(a[i])] = 0
    }
  }
}

for (var i = 0; i < lines.length; i+=3) {
  const a = lines[i]
  const b = lines[i+1]
  const c = lines[i+2]

  const o = hash(a)
  cmp(b, o)
  cmp(c, o)
  score += sum(Object.keys(o).filter(k => o[k] > 2).map(point))
  console.log(o)
}

console.log(score)
// console.log(point('z'))
// console.log(point('Z'))