const fs = require('fs')

class Condition {
  constructor (registerName, comparison, value) {
    this._registerName = registerName
    this._comparison = comparison
    this._value = value
  }

  get registerName () {
    return this._registerName
  }

  getComparisonFunction () {
    switch (this._comparison) {
      case '<':
        return this._lessThan.bind(this)

      case '>':
        return this._greaterThan.bind(this)

      case '<=':
        return this._lessThanEqual.bind(this)

      case '>=':
        return this._greaterThanEqual.bind(this)

      case '==':
        return this._strictEqual.bind(this)

      case '!=':
        return this._notStrictEqual.bind(this)
    }

    return null
  }

  _lessThan (val) {
    return val < this._value
  }

  _greaterThan (val) {
    return val > this._value
  }

  _lessThanEqual (val) {
    return !this._greaterThan(val)
  }

  _greaterThanEqual (val) {
    return !this._lessThan(val)
  }

  _strictEqual (val) {
    return val === this._value
  }

  _notStrictEqual (val) {
    return !this._strictEqual(val)
  }
}

class Instruction {
  constructor (registerName, modifier, value, condition) {
    this._registerName = registerName
    this._modifier = modifier
    this._value = value
    this._condition = condition
  }

  get registerName () {
    return this._registerName
  }

  get condition () {
    return this._condition
  }

  getModificationFunction () {
    switch (this._modifier) {
      case 'inc':
        return this._increment.bind(this)

      case 'dec':
        return this._decrement.bind(this)
    }

    return null
  }

  _increment (val) {
    return val + this._value
  }

  _decrement (val) {
    return val - this._value
  }
}

class Registers {
  constructor () {
    this._map = new Map()
    this._maxValue = Number.MIN_SAFE_INTEGER
  }

  getLargestValue () {
    if (this._map.size === 0) {
      return 0
    }

    let max = Number.MIN_SAFE_INTEGER

    for (const val of this._map.values()) {
      max = Math.max(max, val)
    }

    return max
  }

  getMaxValue () {
    if (this._map.size === 0) {
      return 0
    }

    return this._maxValue
  }

  getRegisterValue (name) {
    const value = this._map.get(name)
    if (value !== undefined) {
      return value
    }

    return 0
  }

  setRegisterValue (name, value) {
    this._maxValue = Math.max(this._maxValue, value)
    this._map.set(name, value)
  }
}

class Processor {
  constructor () {
    this._registers = new Registers()
  }

  getLargestRegisterValue () {
    return this._registers.getLargestValue()
  }

  getMaxRegisterValue () {
    return this._registers.getMaxValue()
  }

  processInstruction (instruction) {
    if (this.evaluateCondition(instruction.condition)) {
      const regName = instruction.registerName
      const func = instruction.getModificationFunction()

      const value = this._registers.getRegisterValue(regName)
      this._registers.setRegisterValue(regName, func(value))
    }
  }

  evaluateCondition (condition) {
    const regName = condition.registerName
    const func = condition.getComparisonFunction()

    return func(this._registers.getRegisterValue(regName))
  }
}

class Day08 {
  static * parseRows (str) {
    let current = ''
    let row = []

    for (const ch of str) {
      if (ch === '\n' || ch === ' ') {
        row.push(current)
        current = ''

        if (ch === '\n') {
          yield row
          row = []
        }
      } else {
        current += ch
      }
    }
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')
    const processor = new Processor()

    for (const row of this.parseRows(fileContent)) {
      const instruction = new Instruction(
      row[0], row[1], Number.parseInt(row[2]), new Condition(row[4], row[5],
      Number.parseInt(row[6])))

      processor.processInstruction(instruction)
    }

    return [
      processor.getLargestRegisterValue(),
      processor.getMaxRegisterValue()
    ]
  }
}

module.exports = Day08
