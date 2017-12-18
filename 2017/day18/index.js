const fs = require('fs')

class Instruction {
  constructor (operation, registerName, value) {
    this._operation = operation
    this._registerName = registerName
    this._value = value
  }

  get operation () {
    return this._operation
  }

  get registerName () {
    return this._registerName
  }

  get value () {
    return this._value
  }
}

class Registers {
  constructor () {
    this._map = new Map()
  }

  getValue (name) {
    const value = this._map.get(name)
    if (value !== undefined) {
      return value
    }

    return 0
  }

  setValue (name, value) {
    this._map.set(name, value)
  }
}

class ProcessBase {
  constructor (instructions) {
    this._registers = new Registers()
    this._instructions = Array.from(instructions)
    this._instructionIndex = 0
  }

  _processInstruction (instruction) {
    switch (instruction.operation) {
      case 'snd':
        this._snd(instruction.registerName)
        break

      case 'set':
        this._set(instruction.registerName, instruction.value)
        break

      case 'add':
        this._add(instruction.registerName, instruction.value)
        break

      case 'mul':
        this._mul(instruction.registerName, instruction.value)
        break

      case 'mod':
        this._mod(instruction.registerName, instruction.value)
        break

      case 'rcv':
        this._rcv(instruction.registerName)
        break

      case 'jgz':
        this._jgz(instruction.registerName, instruction.value)
        break
    }
  }

  _set (registerName, value) {
    this._registers.setValue(registerName, this._toNumValue(value))
    this._instructionIndex++
  }

  _add (registerName, value) {
    const current = this._registers.getValue(registerName)
    this._registers.setValue(registerName, current + this._toNumValue(value))
    this._instructionIndex++
  }

  _mul (registerName, value) {
    const current = this._registers.getValue(registerName)
    this._registers.setValue(registerName, current * this._toNumValue(value))
    this._instructionIndex++
  }

  _mod (registerName, value) {
    const current = this._registers.getValue(registerName)
    this._registers.setValue(registerName, current % this._toNumValue(value))
    this._instructionIndex++
  }

  _jgz (registerName, value) {
    let current = Number.parseInt(registerName)
    if (Number.isNaN(current)) {
      current = this._registers.getValue(registerName)
    }

    if (current > 0) {
      this._instructionIndex += this._toNumValue(value)
    } else {
      this._instructionIndex++
    }
  }

  _toNumValue (value) {
    let numValue = Number.parseInt(value)
    if (Number.isNaN(numValue)) {
      numValue = this._registers.getValue(value)
    }

    return numValue
  }
}

class Part1Process extends ProcessBase {
  constructor (instructions) {
    super(instructions)
    this._lastFrequency = -1
    this._returnValue = -1
  }

  execute () {
    while (this._instructionIndex < this._instructions.length) {
      this._processInstruction(this._instructions[this._instructionIndex])
    }

    return this._returnValue
  }

  _snd (registerName) {
    const value = this._registers.getValue(registerName)
    this._lastFrequency = value
    this._instructionIndex++
  }

  _rcv (registerName) {
    const current = this._registers.getValue(registerName)
    if (current !== 0) {
      this._exit(this._lastFrequency)
    }

    this._instructionIndex++
  }

  _exit (result) {
    this._returnValue = result
    this._instructionIndex = this._instructions.length
  }
}

class Part2Process extends ProcessBase {
  constructor (instructions, id) {
    super(instructions)
    this._messageQueue = []
    this._isWaiting = false
    this._sendHandler = null

    this._registers.setValue('p', id)
  }

  get isBlocked () {
    return this._isWaiting && this._messageQueue.length === 0
  }

  runNext () {
    do {
      this._processInstruction(this._instructions[this._instructionIndex])
    } while (!this._isWaiting)
  }

  setSendHandler (func) {
    this._sendHandler = func
  }

  onMessageReceived (msg) {
    this._messageQueue.push(msg)
  }

  _snd (registerName) {
    const value = this._registers.getValue(registerName)
    this._sendHandler(value)
    this._instructionIndex++
  }

  _rcv (registerName) {
    if (this._messageQueue.length > 0) {
      this._isWaiting = false
      this._registers.setValue(registerName, this._messageQueue.shift())
      this._instructionIndex++
    } else {
      this._isWaiting = true
    }
  }
}

class Day18 {
  static * parseInstructions (str) {
    const regexp = /^([a-z]+) ([0-9a-z]+)(?: (-?[0-9]+|[a-z]+))?$/gm
    let result = null

    do {
      result = regexp.exec(str)
      if (result !== null) {
        yield new Instruction(result[1], result[2], result[3])
      }
    } while (result != null)
  }

  static runPart1 (instructions) {
    const incorrectProcess = new Part1Process(instructions)
    return incorrectProcess.execute()
  }

  static runPart2 (instructions) {
    let sendCount = 0

    const correctProcess0 = new Part2Process(instructions, 0)
    correctProcess0.setSendHandler((msg) => {
      correctProcess1.onMessageReceived(msg)
    })

    const correctProcess1 = new Part2Process(instructions, 1)
    correctProcess1.setSendHandler((msg) => {
      sendCount++
      correctProcess0.onMessageReceived(msg)
    })

    while (!correctProcess0.isBlocked || !correctProcess1.isBlocked) {
      correctProcess0.runNext()
      correctProcess1.runNext()
    }

    return sendCount
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8')
    const instructions = Array.from(this.parseInstructions(fileContent))

    return [
      this.runPart1(instructions),
      this.runPart2(instructions)
    ]
  }
}

module.exports = Day18
