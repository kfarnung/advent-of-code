const fs = require('fs');

class Instruction {
  constructor(operation, registerName, value) {
    this._operation = operation;
    this._registerName = registerName;
    this._value = value;
  }

  get operation() {
    return this._operation;
  }

  get registerName() {
    return this._registerName;
  }

  get value() {
    return this._value;
  }
}

class Registers {
  constructor() {
    this._map = new Map();
  }

  getValue(name) {
    const value = this._map.get(name);
    if (value !== undefined) {
      return value;
    }

    return 0;
  }

  setValue(name, value) {
    this._map.set(name, value);
  }
}

class Process {
  constructor(instructions) {
    this._registers = new Registers();
    this._instructions = Array.from(instructions);
    this._instructionIndex = 0;
    this._countMul = 0;
  }

  execute() {
    while (this._instructionIndex < this._instructions.length) {
      this._processInstruction(this._instructions[this._instructionIndex]);
    }

    return this._countMul;
  }

  _processInstruction(instruction) {
    switch (instruction.operation) {
      case 'set':
        this._set(instruction.registerName, instruction.value);
        break;

      case 'sub':
        this._sub(instruction.registerName, instruction.value);
        break;

      case 'mul':
        this._mul(instruction.registerName, instruction.value);
        break;

      case 'jnz':
        this._jnz(instruction.registerName, instruction.value);
        break;

      default:
        throw new Error('Invalid operation');
    }
  }

  _set(registerName, value) {
    this._registers.setValue(registerName, this._toNumValue(value));
    this._instructionIndex++;
  }

  _sub(registerName, value) {
    const current = this._registers.getValue(registerName);
    this._registers.setValue(registerName, current - this._toNumValue(value));
    this._instructionIndex++;
  }

  _mul(registerName, value) {
    this._countMul++;
    const current = this._registers.getValue(registerName);
    this._registers.setValue(registerName, current * this._toNumValue(value));
    this._instructionIndex++;
  }

  _jnz(registerName, value) {
    let current = Number.parseInt(registerName);
    if (Number.isNaN(current)) {
      current = this._registers.getValue(registerName);
    }

    if (current !== 0) {
      this._instructionIndex += this._toNumValue(value);
    } else {
      this._instructionIndex++;
    }
  }

  _toNumValue(value) {
    let numValue = Number.parseInt(value);
    if (Number.isNaN(numValue)) {
      numValue = this._registers.getValue(value);
    }

    return numValue;
  }
}

class Day23 {
  static *parseInstructions(str) {
    const regexp = /^([a-z]+) ([0-9a-z]+)(?: (-?[0-9]+|[a-z]+))?$/gm;
    let result = null;

    do {
      result = regexp.exec(str);
      if (result !== null) {
        yield new Instruction(result[1], result[2], result[3]);
      }
    } while (result != null);
  }

  static runPart1(instructions) {
    const process = new Process(instructions);
    return process.execute();
  }

  static runPart2() {
    let b = 67 * 100 + 100000;
    const c = b + 17000;
    let g = 0;
    let h = 0;

    do {
      let f = 1;
      let d = 2;

      do {
        // let e = 2;

        // do {
        //   g = d;
        //   g *= e;
        //   g -= b;

        //   if (g === 0) {
        //     f = 0;
        //   }

        //   e -= -1;
        //   g = e;
        //   g -= b;
        // } while (g !== 0);

        if (b % d === 0) {
          f = 0;
          break;
        }

        d -= -1;
        g = d;
        g -= b;
      } while (g !== 0);

      if (f === 0) {
        h -= -1;
      }

      g = b;
      g -= c;

      b -= -17;
    } while (g !== 0);

    return h;
  }

  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const instructions = Array.from(this.parseInstructions(fileContent));

    return [this.runPart1(instructions), this.runPart2()];
  }
}

module.exports = Day23;
