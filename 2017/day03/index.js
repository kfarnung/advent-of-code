function distance (search) {
  let side = 3
  let count = 1
  let x = 0
  let y = 0

  if (search === 1) {
    return {x, y}
  }

  for (;; side += 2) {
    x++
    if (++count === search) {
      return {x, y}
    }

    for (let i = 2; i < side; i++) {
      y++
      if (++count === search) {
        return {x, y}
      }
    }

    for (let i = 1; i < side; i++) {
      x--
      if (++count === search) {
        return {x, y}
      }
    }

    for (let i = 1; i < side; i++) {
      y--
      if (++count === search) {
        return {x, y}
      }
    }

    for (let i = 1; i < side; i++) {
      x++
      if (++count === search) {
        return {x, y}
      }
    }
  }
}

let {x, y} = distance(312051)
console.log(Math.abs(x) + Math.abs(y))
