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
beacons.forEach(beacon => {
  for (let x=minX; x<=maxX; x++) {
    if (!beacons.some(h => h.b[0] === x && h.b[1] === y) && dist(beacon.s[0], beacon.s[1], x, y) <= 
        beacon.dist) {
      map[x] = 1
    }
  }
});

console.log(sum(Object.values(map)))

console.log(minX)
console.log(maxX)

// console.log(beacons)