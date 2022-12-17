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

let beacons = lines.map(l => {
  const [s,b] = l.replace(' closest beacon is at x=', '')
  .replace('Sensor at x=', '')
  .replace(/ y=/gi, '')
  .split(':').map(i => i.split(',').map(k => ~~k))
  return {
    s,b,dist: dist(s[0], s[1], b[0], b[1])
  }
})

let minX = 0 // Math.min.apply(null, beacons.map(i => i.s[0] - i.dist-1))
let maxX = 4000000 // Math.max.apply(null, beacons.map(i => i.s[0] + i.dist+1))

// let minY = Math.min.apply(null, beacons.map(i => Math.min(i.b[1], i.s[1])))
// let maxY = Math.max.apply(null, beacons.map(i => Math.max(i.b[1], i.s[1])))

const map = {}
for (let x=minX; x<=maxX; x++) {
  map[x] = 0
}

function dist(x1,y1,x2,y2) {
  return Math.abs(x1-x2) + Math.abs(y1-y2)
}

let y = 2000000

function inSensor(b, x, y) {
  return dist(b.s[0], b.s[1], x, y) <= b.dist
}

function check(minX, minY, maxX, maxY) {
  if (minX == maxX || minY == maxY) return

  if (beacons.some(b => 
    inSensor(b, minX, minY) && 
    inSensor(b, maxX-1, minY) && 
    inSensor(b, minX, maxY-1) && 
    inSensor(b, maxX-1, maxY-1))) {
      return
  }
  if (minX>=maxX-1 && minY>=maxY-1) {
    console.log('spot: ', minX, minY, minX*4000000+minY)
    return
  }
  let cX = Math.round((minX + maxX) / 2)
  let cY = Math.round((minY + maxY) / 2)

  check(minX, minY, cX, cY)
  check(cX, minY, maxX, cY)
  check(minX, cY, cX, maxY)
  check(cX, cY, maxX, maxY)
}

check(0,0,4000000,4000000)
// check(14,11,15,12)
