const fs = require('fs');

class Board {
  constructor(input) {
    this._rows = Board._parseBoard(input);
  }

  getValue(x, y) {
    if (y < this._rows.length && x < this._rows[y].length) {
      return this._rows[y][x];
    }

    return ' ';
  }

  findStart() {
    const row = this._rows[0];
    for (let x = 0; x < row.length; x++) {
      if (row[x] === '|') {
        return [x, 0];
      }
    }

    return null;
  }

  findNextDelta(x, y, deltaX, deltaY) {
    if (deltaX !== 0) {
      if (this.getValue(x, y - 1) !== ' ') {
        return [0, -1];
      } else if (this.getValue(x, y + 1) !== ' ') {
        return [0, 1];
      } else {
        throw new Error('Unexpected');
      }
    } else if (deltaY !== 0) {
      if (this.getValue(x - 1, y) !== ' ') {
        return [-1, 0];
      } else if (this.getValue(x + 1, y) !== ' ') {
        return [1, 0];
      } else {
        throw new Error('Unexpected');
      }
    } else {
      throw new Error('Unexpected direction');
    }
  }

  static _parseBoard(str) {
    const rows = [];
    let row = [];
    for (const ch of str) {
      if (
        ch === '-' ||
        ch === '|' ||
        ch === '+' ||
        ch === ' ' ||
        (ch >= 'A' && ch <= 'Z')
      ) {
        row.push(ch);
      } else if (ch === '\n') {
        rows.push(row);
        row = [];
      } else if (ch !== '\r') {
        throw new Error('Unexpected character!');
      }
    }

    return rows;
  }
}

class Day19 {
  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const board = new Board(fileContent);
    const letters = [];
    let numSteps = 0;

    let [x, y] = board.findStart();
    let deltaX = 0;
    let deltaY = 1;

    for (;;) {
      const current = board.getValue(x, y);
      if (current >= 'A' && current <= 'Z') {
        letters.push(current);
      } else if (current === '+') {
        [deltaX, deltaY] = board.findNextDelta(x, y, deltaX, deltaY);
      } else if (current === ' ') {
        break;
      } else if (current !== '|' && current !== '-') {
        throw new Error('Unexpected');
      }

      x += deltaX;
      y += deltaY;
      numSteps++;
    }

    return [letters.join(''), numSteps];
  }
}

module.exports = Day19;
