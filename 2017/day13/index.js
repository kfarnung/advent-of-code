const fs = require('fs');

class FirewallLayer {
  constructor (depth, range) {
    this._depth = depth;
    this._range = range;
  }

  get severity () {
    return this._depth * this._range;
  }

  isAtTop (time) {
    return (time % ((this._range - 1) * 2)) === 0;
  }
}

class Firewall {
  constructor () {
    this._layers = [];
    this._maxDepth = 0;
  }

  addRawLayer (rawData) {
    const depth = rawData[0];
    this._maxDepth = Math.max(this._maxDepth, depth);
    this._layers[depth] = new FirewallLayer(depth, rawData[1]);
  }

  getSeverity () {
    let severity = 0;

    for (let i = 0; i <= this._maxDepth; i++) {
      const layer = this._layers[i];
      if (layer !== undefined) {
        if (layer.isAtTop(i)) {
          severity += layer.severity;
        }
      }
    }

    return severity;
  }

  isCaught (delay = 0) {
    for (let i = 0; i <= this._maxDepth; i++) {
      const layer = this._layers[i];
      if (layer !== undefined) {
        if (layer.isAtTop(i + delay)) {
          return true;
        }
      }
    }

    return false;
  }

  findSafeDelay () {
    let delay = -1;

    while (this.isCaught(++delay)) {}

    return delay;
  }
}

class Day13 {
  static * parseRows (str) {
    let current = '';
    let row = [];

    for (const ch of str) {
      if (ch === '\n' || ch === ' ' || ch === ':') {
        if (current.length > 0) {
          row.push(Number.parseInt(current));
          current = '';
        }

        if (ch === '\n') {
          yield row;
          row = [];
        }
      } else if (ch >= '0' && ch <= '9') {
        current += ch;
      }
    }
  }

  static run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const firewall = new Firewall();

    for (const row of this.parseRows(fileContent)) {
      firewall.addRawLayer(row);
    }

    return [
      firewall.getSeverity(0),
      firewall.findSafeDelay()
    ];
  }
}

module.exports = Day13;
