const fs = require('fs')

class ProgramDance {
  constructor () {
    this._programs = ProgramDance._generatePrograms()
  }

  processDanceMoves (danceMoves) {
    for (const danceMove of danceMoves) {
      this.processDanceMove(danceMove)
    }
  }

  processDanceMove (danceMove) {
    switch (danceMove[0]) {
      case 's':
        this._spin(Number.parseInt(danceMove[1]))
        break

      case 'x':
        this._exchange(Number.parseInt(danceMove[1]), Number.parseInt(danceMove[2]))
        break

      case 'p':
        this._partner(danceMove[1], danceMove[2])
        break
    }
  }

  toString () {
    return this._programs.join('')
  }

  static _getCharacter (index) {
    return String.fromCharCode('a'.charCodeAt(0) + index)
  }

  static _generatePrograms () {
    const programs = []
    for (let i = 0; i < 16; i++) {
      programs.push(this._getCharacter(i))
    }

    return programs
  }

  _spin (num) {
    for (let i = 0; i < num; i++) {
      this._programs.unshift(this._programs.pop())
    }
  }

  _exchange (index1, index2) {
    const temp = this._programs[index1]
    this._programs[index1] = this._programs[index2]
    this._programs[index2] = temp
  }

  _partner (name1, name2) {
    this._exchange(this._programs.indexOf(name1), this._programs.indexOf(name2))
  }
}

class Day16 {
  static * parseDanceMoves (str) {
    const regexp = /([sxp])([a-z0-9]+)(?:\/([a-z0-9]+))?/g
    let result = null

    do {
      result = regexp.exec(str)
      if (result !== null) {
        yield result.slice(1)
      }
    } while (result != null)
  }

  static getIteratorIndex (iterator, index) {
    let i = 0
    for (const val of iterator) {
      if (i++ === index) {
        return val
      }
    }

    return null
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')
    const danceMoves = Array.from(this.parseDanceMoves(fileContent))
    const cycleSet = new Set()

    const dance = new ProgramDance()
    while (true) {
      dance.processDanceMoves(danceMoves)

      const currentState = dance.toString()

      if (cycleSet.has(currentState)) {
        break
      }

      cycleSet.add(currentState)
    }

    return [
      this.getIteratorIndex(cycleSet.keys(), 0),
      this.getIteratorIndex(cycleSet.keys(), (1000000000 % cycleSet.size) - 1)
    ]
  }
}

module.exports = Day16
