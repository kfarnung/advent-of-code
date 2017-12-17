class Day17 {
  static part1 (stepSize) {
    const spinlock = [ 0 ]
    let currentPosition = 0
    for (let i = 1; i < 2018; i++) {
      currentPosition = ((currentPosition + stepSize) % i) + 1
      spinlock.splice(currentPosition, 0, i)
    }

    return spinlock[currentPosition + 1]
  }

  static part2 (stepSize) {
    let currentPosition = 0
    let afterZero = 0
    for (let i = 1; i < 50000000; i++) {
      currentPosition = ((currentPosition + stepSize) % i) + 1
      if (currentPosition === 1) {
        afterZero = i
      }
    }

    return afterZero
  }

  static run (input) {
    const stepSize = Number.parseInt(input)

    return [
      this.part1(stepSize),
      this.part2(stepSize)
    ]
  }
}

module.exports = Day17
