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

var root = {

}
var inLs = false
var dir = root

for (var i=0; i< lines.length; i++) {
  const line =lines[i]

  if (line[0] == '$') {
    inLs = false
    var cmd = line.substr(2).split(' ')
    var param = cmd[1]
    cmd = cmd[0]
    console.log(cmd)
    if (cmd == 'cd') {
      if (param == '/') {
        dir = root
      } else if (param == '..') {
        dir = dir._back
      } else {
        dir = dir[param]
      }
    }
    if (cmd == 'ls') {
      inLs = true
    }
  } else if (inLs) {
    let size
    const h = line.split(' ')
    if (h[0] === 'dir') {
      console.log('dir ' + h[1])
    } else {
      size = parseInt(h[0])
    }
    dir[h[1]] = {
      _size: size,
      // _dir: h[1],
      // _ls: [],
      _back: dir
    }
  }
} 

const h= []

function sizer(dir, alter  = false) {
  var s = 0
  for (var key in dir) {
    if (key[0] != '_') {
        if (!dir[key]._size) {
          const ls = sizer(dir[key], alter)
          if (ls >= 1) {
            if (alter) {
              h.push({dir: dir[key], size: ls})
            }
          }
          s += ls
        } else {
          s += dir[key]._size
        }
    }
  }
  // console.log(s)
  return s
}

const total = sizer(root, false)

sizer(root, true)

const free = 70000000 - total
const needed = 30000000 - free

console.log('needed ' + needed)
r = sort(h.filter(i => i.size > needed), i => i.size)

r.forEach(i => {
  console.log(i.size, )
})
// console.log(sum(h.map(u => sizer(u, false))))