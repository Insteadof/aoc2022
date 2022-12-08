fs = require('fs')

const file = fs.readFileSync("./input.txt", "utf8");
const lines = file.split('\n')

const grid = []
lines.forEach( i => grid.push(Array(lines[0].length)))

const ll = lines[0].length

function spot(x, y) {
  if (y < 0 || y >= lines.length) {
    return 1000
  }
  return ~~(lines[y][x])
}

function senique(x, y) {
  const h = spot(x,y)
  let left = 0
  let top = 0
  let right = 0
  let bottom = 0
  let l = x
  let r = x
  let t = y
  let b = y

  do  {
    l--
    left++
  } while (l > 0 && spot(l, y) < h)
  do  {
    r++
    right++
  } while (r < lines[0].length - 1 && spot(r, y) < h)
  do  {
    t--
    top++
  } while (t > 0 && spot(x, t) < h)
  do  {
    b++
    bottom++
  } while (b < lines.length -1 && spot(x, b) < h)

  return left * right * bottom * top
}

let max = 0
for (var i=1; i< lines.length-1; i++) {
  const line = lines[i]
  for (k=1; k < line.length-1; k++) {
    const s = senique(k, i)
    if (s > max) {
      console.log(k, i)
      max = s
    }
  }
} 

console.log(max)
console.log(lines.length)
console.log(lines[0].length)