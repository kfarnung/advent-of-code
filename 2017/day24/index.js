const fs = require('fs');

class Component {
  constructor(pinsA, pinsB) {
    this._pinsA = pinsA;
    this._pinsB = pinsB;

    this._connectionsA = [];
    this._connectionsB = [];
  }

  get pinsA() {
    return this._pinsA;
  }

  get pinsB() {
    return this._pinsB;
  }

  addConnection(component) {
    if (component._pinsA === this._pinsA || component._pinsB === this._pinsA) {
      this._connectionsA.push(component);
    }

    if (component._pinsA === this._pinsB || component._pinsB === this._pinsB) {
      this._connectionsB.push(component);
    }
  }

  getMaxStrength(incomingPin, visited) {
    let max = Number.MIN_SAFE_INTEGER;
    const connections = this._getConnections(incomingPin);
    const outgoingPin = this._getOutgoingPin(incomingPin);

    if (visited.indexOf(this) >= 0) {
      return 0;
    }

    visited.push(this);

    for (const connection of connections) {
      max = Math.max(max, connection.getMaxStrength(outgoingPin, visited));
    }

    visited.pop();

    return this._pinsA + this._pinsB + max;
  }

  getLongestStrength(incomingPin, visited) {
    let maxStrength = Number.MIN_SAFE_INTEGER;
    let maxLength = Number.MIN_SAFE_INTEGER;
    const connections = this._getConnections(incomingPin);
    const outgoingPin = this._getOutgoingPin(incomingPin);

    if (visited.indexOf(this) >= 0) {
      return { length: visited.length, strength: 0 };
    }

    visited.push(this);

    for (const connection of connections) {
      const result = connection.getLongestStrength(outgoingPin, visited);
      if (result.length > maxLength) {
        maxLength = result.length;
        maxStrength = result.strength;
      } else if (result.length === maxLength) {
        if (result.strength > maxStrength) {
          maxStrength = result.strength;
        }
      }
    }

    visited.pop();

    return {
      length: maxLength,
      strength: this._pinsA + this._pinsB + maxStrength,
    };
  }

  _getConnections(incomingPin) {
    if (incomingPin === this._pinsA) {
      return this._connectionsB;
    } else if (incomingPin === this._pinsB) {
      return this._connectionsA;
    } else {
      throw new Error('Invalid pin');
    }
  }

  _getOutgoingPin(incomingPin) {
    if (incomingPin === this._pinsA) {
      return this._pinsB;
    } else if (incomingPin === this._pinsB) {
      return this._pinsA;
    } else {
      throw new Error('Invalid pin');
    }
  }
}

class ComponentGraph {
  constructor() {
    this._map = new Map();
  }

  addNode(component) {
    this._connectNodes(component.pinsA, component);
    this._connectNodes(component.pinsB, component);
  }

  getMaxStrength() {
    let max = Number.MIN_SAFE_INTEGER;
    const roots = this._getFromMap(0);
    const visited = [];

    for (const root of roots) {
      max = Math.max(max, root.getMaxStrength(0, visited));
    }

    return max;
  }

  getLongestStrength() {
    let maxLength = Number.MIN_SAFE_INTEGER;
    let maxStrength = Number.MIN_SAFE_INTEGER;
    const roots = this._getFromMap(0);
    const visited = [];

    for (const root of roots) {
      const result = root.getLongestStrength(0, visited);
      if (result.length > maxLength) {
        maxLength = result.length;
        maxStrength = result.strength;
      } else if (result.length === maxLength) {
        if (result.strength > maxStrength) {
          maxStrength = result.strength;
        }
      }
    }

    return maxStrength;
  }

  _connectNodes(numPins, component) {
    const components = this._getFromMap(numPins);
    for (const c of components) {
      c.addConnection(component);
      component.addConnection(c);
    }

    this._addToMap(numPins, component);
  }

  _getFromMap(numPins) {
    const components = this._map.get(numPins);
    if (components !== undefined) {
      return components;
    }

    return [];
  }

  _addToMap(numPins, component) {
    let components = this._map.get(numPins);
    if (components === undefined) {
      components = [];
      this._map.set(numPins, components);
    }

    components.push(component);
  }
}

class Day24 {
  static *parseComponents(str) {
    const regexp = /^([0-9]+)\/([0-9]+)$/gm;
    let result = null;

    do {
      result = regexp.exec(str);
      if (result !== null) {
        yield new Component(
          Number.parseInt(result[1]),
          Number.parseInt(result[2])
        );
      }
    } while (result != null);
  }

  static *run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const graph = new ComponentGraph();

    for (const component of this.parseComponents(fileContent)) {
      graph.addNode(component);
    }

    yield graph.getMaxStrength();
    yield graph.getLongestStrength();
  }
}

module.exports = Day24;
