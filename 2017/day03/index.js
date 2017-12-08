function findCoordinates (val) {
  const length = Math.ceil(Math.sqrt(val))
  const halfLength = Math.floor(length / 2)

  if (val === length * length) {
    return { x: halfLength, y: -halfLength }
  }

  const location = val - Math.pow(length - 2, 2)
  const side = Math.floor(location / (length - 1))
  const pos = location % (length - 1)

  switch (side) {
    case 0:
      return { x: halfLength, y: -halfLength + pos }

    case 1:
      return { x: halfLength - pos, y: halfLength }

    case 2:
      return { x: -halfLength, y: halfLength - pos }

    case 3:
      return { x: -halfLength + pos, y: -halfLength }
  }
}

let { x, y } = findCoordinates(312051)
console.log(`Part 1: ${Math.abs(x) + Math.abs(y)}`)
