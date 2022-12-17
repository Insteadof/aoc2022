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

const pos = Array(25)
for (let i=0; i< pos.length; i++) {
  pos[i] = Array(25).map(k => 0)
}

pos[12][12] = 1

let rope = []
for (let k=0; k < 10; k++) {
  rope.push({
    x:0,
    y:0
  })
}

let score = 1


function calc(m, depth, initial) {
  let {x,y} = rope[depth]
  let x2 = x
  let y2 = y

  x += m[0]
  y += m[1]

  rope[depth].x = x
  rope[depth].y = y

  if (depth >= 9) {
    return
  }

  let t = rope[depth+1]
  // if (Math.abs(x - t.x) > 1 && Math.abs(y - t.y) > 1) {
  //   let m2
  //   if (m[1] != 0) {
  //     let m2 = [x2 - t.x, y2 - t.y]
  //   }
    
  //   calc(m, depth+1)
  //   return
  // }

  if (Math.abs(x - t.x) > 1 || Math.abs(y - t.y) > 1) {
    let m2 = [x - t.x, y - t.y]
    if (Math.abs(m2[0]) >= 2) {
      m2[0] = parseInt(m2[0] / 2)
    }
    if (Math.abs(m2[1]) >= 2) {
      m2[1] = parseInt(m2[1] / 2)
    }
    calc(m2, depth+1)
  }
}

for (var i=0; i< lines.length; i++) {
  const line = lines[i]

  const [d, amount] = line.split(' ')

  let m = [0, 0] 
  if (d == 'L') {
    m[0] = -1
  }
  if (d == 'R') {
    m[0] = 1
  }
  if (d == 'D') {
    m[1] = 1
  }
  if (d == 'U') {
    m[1] = -1
  }

  for (let i=0; i < ~~amount; i++) {
    calc(m, 0, true)

    tx = rope[9].x + 12
    ty = rope[9].y + 12

    if (!pos[ty]) {
      pos[ty]=[]
    }
    if (!pos[ty][tx]) {
      pos[ty][tx] = 1
      score++
      // console.log(d, i, x, y)
    }
  }
} 

console.log(score)