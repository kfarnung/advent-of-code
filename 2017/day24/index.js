const fs = require('fs');

class Component {
  constructor (pinsA, pinsB) {
    this._pinsA = pinsA;
    this._pinsB = pinsB;
    this._inUse = false;
  }

  get pinsA () {
    return this._pinsA;
  }

  get pinsB () {
    return this._pinsB;
  }

  get inUse () {
    return this._inUse;
  }

  set inUse (value) {
    this._inUse = value;
  }

  getOutgoingPin (incomingPin) {
    if (incomingPin === this._pinsA) {
      return this._pinsB;
    } else if (incomingPin === this._pinsB) {
      return this._pinsA;
    } else {
      throw new Error('Invalid pin');
    }
  }
}

class ComponentManager {
  constructor () {
    this._map = [];
  }

  addNode (component) {
    this._addToMap(component.pinsA, component);

    if (component.pinsA !== component.pinsB) {
      this._addToMap(component.pinsB, component);
    }
  }

  getMaxStrength (current = null, incomingPin = 0) {
    let max = 0;

    for (const next of this._getFromMap(incomingPin)) {
      if (!next.inUse) {
        next.inUse = true;
        max = Math.max(max, this.getMaxStrength(next, next.getOutgoingPin(incomingPin)));
        next.inUse = false;
      }
    }

    if (current !== null) {
      max += current.pinsA;
      max += current.pinsB;
    }

    return max;
  }

  getLongestStrength () {
    return this._getLongestStrength().strength;
  }

  _getLongestStrength (current = null, incomingPin = 0, currentLength = 0) {
    let maxStrength = 0;
    let maxLength = currentLength;

    for (const next of this._getFromMap(incomingPin)) {
      if (!next.inUse) {
        next.inUse = true;

        const result = this._getLongestStrength(next, next.getOutgoingPin(incomingPin), currentLength + 1);
        if (result.length > maxLength) {
          maxLength = result.length;
          maxStrength = result.strength;
        } else if (result.length === maxLength) {
          if (result.strength > maxStrength) {
            maxStrength = result.strength;
          }
        }

        next.inUse = false;
      }
    }

    if (current !== null) {
      maxStrength += current.pinsA;
      maxStrength += current.pinsB;
    }

    return { length: maxLength, strength: maxStrength };
  }

  _getFromMap (numPins) {
    const components = this._map[numPins];
    if (components !== undefined) {
      return components;
    }

    return [];
  }

  _addToMap (numPins, component) {
    let components = this._map[numPins];
    if (components === undefined) {
      components = [];
      this._map[numPins] = components;
    }

    components.push(component);
  }
}

class Day24 {
  static * parseComponents (str) {
    const regexp = /^([0-9]+)\/([0-9]+)$/gm;
    let result = null;

    do {
      result = regexp.exec(str);
      if (result !== null) {
        yield new Component(Number.parseInt(result[1]), Number.parseInt(result[2]));
      }
    } while (result != null);
  }

  static * run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const manager = new ComponentManager();

    for (const component of this.parseComponents(fileContent)) {
      manager.addNode(component);
    }

    yield manager.getMaxStrength();
    yield manager.getLongestStrength();
  }
}

module.exports = Day24;
