fs = require('fs')

const file = fs.readFileSync("./input.txt", "utf8");
const lines = file.split('\n\n')

function sort(arr, fc) {
  return arr.sort((a, b) => fc(b)-fc(a))
}
function ord(a) {
  return a.charCodeAt(0);
}
function sum(n) {
  return n.reduce((partialSum, a) => partialSum + a, 0);
}

function isArray(b) {
  return b && b.hasOwnProperty('length')
}

function compare(a, b) {
  console.log('---')
  console.log(a)
  console.log(b)
  if (a.length === 0) {
    return b.length === 0 ? undefined : true
  }

  for (let i=0; i<a.length;i++) {
    if (i >= b.length) {
      return false
    }
    let inner = undefined
    if (isArray(a[i]) && isArray(b[i])) {
      inner = compare(a[i],b[i])
    } else if (isArray(a[i]) || isArray(b[i])) {
      if (!isArray(a[i])) {
        inner = compare([a[i]], b[i])
      } else 
        inner = compare(a[i], [b[i]])
    } else {
      if (a[i] != b[i]) {
        return a[i] < b[i]
      }
    }
    if (inner !== undefined) {
      return inner
    }
  }
  return b.length>a.length ? true : undefined
}

console.log('score', sum(lines.map((v, index) => {
  console.log('line')
  const g = compare.apply(null, v.split('\n').map(i => eval(i)))
  console.log('g',index+1, g)
  return g ? (index+1) : 0
})))
