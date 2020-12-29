const fs = require('fs');

class InfiniteGrid {
  constructor(defaultValue = '.') {
    this._defaultValue = defaultValue;
    this._ne = [];
    this._nw = [];
    this._se = [];
    this._sw = [];
  }

  getValue(x, y) {
    const { grid, adjX, adjY } = this._translateCoodinates(x, y);

    const row = grid[adjX];
    if (row !== undefined) {
      const cell = row[adjY];
      if (cell !== undefined) {
        return cell;
      }
    }

    return this._defaultValue;
  }

  setValue(x, y, value) {
    const { grid, adjX, adjY } = this._translateCoodinates(x, y);

    let row = grid[adjX];
    if (row === undefined) {
      row = [];
      grid[adjX] = row;
    }

    row[adjY] = value;
  }

  _translateCoodinates(x, y) {
    let grid = null;
    if (x >= 0) {
      if (y >= 0) {
        grid = this._ne;
      } else {
        y = Math.abs(y) - 1;
        grid = this._se;
      }
    } else {
      x = Math.abs(x) - 1;
      if (y >= 0) {
        grid = this._nw;
      } else {
        y = Math.abs(y) - 1;
        grid = this._sw;
      }
    }

    return { grid, adjX: x, adjY: y };
  }
}

class Vector2d {
  constructor(x, y) {
    this._x = x;
    this._y = y;
  }

  get x() {
    return this._x;
  }

  get y() {
    return this._y;
  }

  add(other) {
    this._x += other._x;
    this._y += other._y;
  }

  rotateLeft() {
    const newX = this._y * -1;
    this._y = this._x;
    this._x = newX;
  }

  rotateRight() {
    const newX = this._y;
    this._y = this._x * -1;
    this._x = newX;
  }

  reverse() {
    this._x *= -1;
    this._y *= -1;
  }
}

class InfectionAgent {
  constructor(grid, evolution = false) {
    this._grid = grid;
    this._evolution = evolution;
    this._current = new Vector2d(0, 0);
    this._direction = new Vector2d(0, 1);
    this._infectedCount = 0;
  }

  get infectedCount() {
    return this._infectedCount;
  }

  burst() {
    const currentValue = this._grid.getValue(this._current.x, this._current.y);
    if (currentValue === '.') {
      this._direction.rotateLeft();

      if (this._evolution) {
        this._grid.setValue(this._current.x, this._current.y, 'W');
      } else {
        this._grid.setValue(this._current.x, this._current.y, '#');
        this._infectedCount++;
      }
    } else if (currentValue === 'W') {
      this._grid.setValue(this._current.x, this._current.y, '#');
      this._infectedCount++;
    } else if (currentValue === '#') {
      this._direction.rotateRight();

      if (this._evolution) {
        this._grid.setValue(this._current.x, this._current.y, 'F');
      } else {
        this._grid.setValue(this._current.x, this._current.y, '.');
      }
    } else if (currentValue === 'F') {
      this._direction.reverse();
      this._grid.setValue(this._current.x, this._current.y, '.');
    } else {
      throw new Error('Unexpected value');
    }

    this._current.add(this._direction);
  }
}

class Day22 {
  static *parseGrid(str) {
    const regexp = /^[.#]+$/gm;
    let result = null;

    do {
      result = regexp.exec(str);
      if (result !== null) {
        yield result[0].split('');
      }
    } while (result != null);
  }

  static fillGrid(grid, data) {
    const deltaX = Math.floor(data[0].length / 2);
    const deltaY = Math.floor(data.length / 2);

    for (let y = deltaY; y >= deltaY * -1; y--) {
      for (let x = deltaX * -1; x <= 12; x++) {
        grid.setValue(x, y, data[data.length - (y + deltaY) - 1][x + deltaX]);
      }
    }
  }

  static runInfection(startingGrid, numBursts = 10000, evolution = false) {
    const grid = new InfiniteGrid();
    this.fillGrid(grid, startingGrid);
    const agent = new InfectionAgent(grid, evolution);

    for (let i = 0; i < numBursts; i++) {
      agent.burst();
    }

    return agent.infectedCount;
  }

  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const startingGrid = Array.from(this.parseGrid(fileContent));

    return [
      this.runInfection(startingGrid),
      this.runInfection(startingGrid, 10000000, true),
    ];
  }
}

module.exports = Day22;
