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

let width = lines[0].length
let height = lines.length

function makeArray(n) {
  const result = []
  for (let i=0;i<n;i++) {
    result.push(0)
  }
  return result
}

let map = []
for (let y=0; y< height;y++) {
  map.push(makeArray(width))
}

map[20][77] = 1

let step = 1
let running = true

function checkStep(x, y, ox, oy, level, step) {
  if (y+oy < 0 || y+oy >= height) {
    return
  }
  if (x+ox < 0 || x+ox >= width) {
    return
  }

  if (map[y+oy][x+ox] === 0 && (ord(lines[y+oy][x+ox].replace('E', 'z'))) >= (level-1)) {
    map[y+oy][x+ox] = step
    if (lines[y+oy][x+ox] === 'a') {
      console.log('stappen: ' + (step - 1))
      running = false
    }
  }
}

while (running) {
  for (let y=0; y< height;y++) {
    for (let x=0; x < width;x++) {
      if (map[y][x] === step) {
        const level = ord(lines[y][x].replace('E', 'z'))
        checkStep(x, y, 1, 0, level, step+1)
        checkStep(x, y, -1, 0, level, step+1)
        checkStep(x, y, 0, 1, level, step+1)
        checkStep(x, y, 0, -1, level, step+1)
      }
    }
  }
  // console.log(step)
  // if (step === 5) {
  print()
  
  // var waitTill = new Date(new Date().getTime() + 600);
  // while(waitTill > new Date()){}

  // }
  step++
}

function print() {
  for (let y=0; y< height;y++) {
    h = ''
    for (let x=0; x < width;x++) {
      h += String.fromCharCode(map[y][x] + ord('a'))
    }
    //console.log(h)
  }
}
