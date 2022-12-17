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

let monkeys = []

let allDivs = 1

for (var i=0; i< lines.length; i += 7) {
  const line = lines[i+1]
  let items = line.split(': ')[1]
  items = items.split(', ').map(i => ~~i)
  let operations = lines[i+2].split('= ')[1].split(' ')
  let div = ~~lines[i+3].split('divisible by ')[1]
  allDivs *= div

  let t = ~~lines[i+4].split('throw to monkey ')[1]
  let f = ~~lines[i+5].split('throw to monkey ')[1]

  const m = {
    monkeyNumber: i / 7,
    items,
    operation: function(old) {
      a = old
      b = isNaN(parseInt(operations[2])) ? old : parseInt(operations[2])
      if (operations[1] == '+') {
        return a + b
      }
      if (operations[1] == '*') {
        return a * b
      }
    },
    div,
    t,
    f,
    inspect: 0
  }
  monkeys.push(m)
}

// console.log(monkeys[0].operation(4))
for (let round=0; round < 10000; round++) {
  monkeys.forEach(m => {
    let item
    while ((item = m.items.shift()) != undefined) {
      m.inspect += 1,
      item = m.operation(item)
      // item = parseInt(item / 3)
      item %= allDivs

      if (item % m.div == 0) {
        // item /= m.div
        monkeys[m.t].items.push(item)
      } else {
        monkeys[m.f].items.push(item)
      }
    }
  })
  if (round == 1000-1) {
   console.log('round ' + round)
    monkeys.forEach(m => {
      console.log(m.monkeyNumber + ' ' + m.inspect)
    })
  }
  // monkeys.forEach(m => {
  //   console.log(m.items)
  // })
}

const k = sort(monkeys, i => i.inspect).slice(0,2).map(i => i.inspect)
console.log(k[0]*k[1])