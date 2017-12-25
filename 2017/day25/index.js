const fs = require('fs');

class BlueprintStateCondition {
  constructor (currentValue, newValue, direction, nextState) {
    this._currentValue = currentValue;
    this._newValue = newValue;
    this._direction = direction;
    this._nextState = nextState;
  }

  get currentValue () {
    return this._currentValue;
  }

  get newValue () {
    return this._newValue;
  }

  get nextState () {
    return this._nextState;
  }

  getDelta () {
    if (this._direction === 'left') {
      return -1;
    } else if (this._direction === 'right') {
      return 1;
    } else {
      throw new Error('Invalid direction');
    }
  }
}

class BlueprintState {
  constructor (name) {
    this._name = name;
    this._conditions = [];
  }

  get name () {
    return this._name;
  }

  addCondition (condition) {
    this._conditions[condition.currentValue] = condition;
  }

  getCondition (value) {
    return this._conditions[value];
  }
}

class Blueprint {
  constructor (str) {
    this._lines = str.split('\n');
    this._lineIndex = 0;
    this._beginState = null;
    this._checksumInterval = -1;
    this._stateMap = new Map();
  }

  get beginState () {
    return this._beginState;
  }

  get checksumInterval () {
    return this._checksumInterval;
  }

  static parse (str) {
    const blueprint = new Blueprint(str);
    blueprint._parseHeader();
    while (blueprint._parseState()) {}

    return blueprint;
  }

  getState (name) {
    return this._stateMap.get(name);
  }

  _getNextLine () {
    return this._lines[this._lineIndex++];
  }

  _parseState () {
    let result = null;

    const stateRegex = /In state ([A-Z]+):/;
    result = stateRegex.exec(this._getNextLine());
    if (result === null) {
      return false;
    }

    const state = new BlueprintState(result[1]);
    this._parseCondition(state);
    this._parseCondition(state);

    if (this._getNextLine() !== '') {
      throw new Error('Expected a blank line');
    }

    this._stateMap.set(state.name, state);

    return true;
  }

  _parseCondition (state) {
    let result = null;

    const conditionRegex = / {2}If the current value is ([0-1]):/;
    result = conditionRegex.exec(this._getNextLine());
    const currentValue = Number.parseInt(result[1]);

    const valueRegex = / {4}- Write the value ([0-1])./;
    result = valueRegex.exec(this._getNextLine());
    const newValue = Number.parseInt(result[1]);

    const moveRegex = / {4}- Move one slot to the (left|right)./;
    result = moveRegex.exec(this._getNextLine());
    const direction = result[1];

    const nextStateRegex = / {4}- Continue with state ([A-Z])./;
    result = nextStateRegex.exec(this._getNextLine());
    const nextState = result[1];

    state.addCondition(new BlueprintStateCondition(currentValue, newValue, direction, nextState));
  }

  _parseHeader () {
    let result = null;

    const beginRegex = /Begin in state ([A-Z])./;
    result = beginRegex.exec(this._getNextLine());
    this._beginState = result[1];

    const checksumRegex = /Perform a diagnostic checksum after ([0-9]+) steps./;
    result = checksumRegex.exec(this._getNextLine());
    this._checksumInterval = Number.parseInt(result[1]);

    if (this._getNextLine() !== '') {
      throw new Error('Expected a blank line');
    }
  }
}

class InfiniteTape {
  constructor () {
    this._map = new Map();
  }

  getValue (index) {
    const value = this._map.get(index);
    if (value !== undefined) {
      return value;
    }

    return 0;
  }

  setValue (index, value) {
    if (value !== 0) {
      this._map.set(index, value);
    } else {
      this._map.delete(index);
    }
  }

  countNonZero () {
    return this._map.size;
  }
}

class TuringMachine {
  constructor (blueprint) {
    this._blueprint = blueprint;
    this._tape = new InfiniteTape();
    this._position = 0;
    this._currentState = blueprint.beginState;
    this._tick = 0;
  }

  run (numTicks = 1) {
    for (let i = 0; i < numTicks; i++) {
      const state = this._blueprint.getState(this._currentState);
      const value = this._tape.getValue(this._position);

      const condition = state.getCondition(value);
      this._tape.setValue(this._position, condition.newValue);
      this._position += condition.getDelta();
      this._currentState = condition.nextState;
    }
  }

  checksum () {
    return this._tape.countNonZero();
  }
}

class Day25 {
  static * run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const parsed = Blueprint.parse(fileContent);

    const machine = new TuringMachine(parsed);
    machine.run(parsed.checksumInterval);
    yield machine.checksum();
  }
}

module.exports = Day25;
