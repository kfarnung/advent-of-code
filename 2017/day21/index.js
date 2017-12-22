const fs = require('fs');

class Image {
  constructor (pixels = []) {
    this._pixels = pixels;
  }

  get size () {
    return this._pixels.length;
  }

  static fromString (str) {
    var arr = [];

    for (const strRow of str.split('/')) {
      const row = [];
      for (const ch of strRow) {
        row.push(ch);
      }

      arr.push(row);
    }

    return new Image(arr);
  }

  getPixel (row, column) {
    if (row > this._pixels.length - 1) {
      throw new Error('Invalid row value');
    }

    const pixelRow = this._pixels[row];

    if (column > pixelRow.length - 1) {
      throw new Error('Invalid row value');
    }

    return pixelRow[column];
  }

  setPixel (row, column, value) {
    let pixelRow = this._pixels[row];
    if (pixelRow === undefined) {
      pixelRow = [];
      this._pixels[row] = pixelRow;
    }

    pixelRow[column] = value;
  }

  getSubImage (row, column, height, width) {
    const subImage = [];
    for (let i = row; i < row + height; i++) {
      subImage.push(this._pixels[i].slice(column, column + width));
    }

    return new Image(subImage);
  }

  * getPermutations () {
    for (let i = 0; i < 4; i++) {
      this.flipVertical();
      yield this.toString();

      this.flipHorizontal();
      yield this.toString();

      this.flipVertical();
      yield this.toString();

      this.flipHorizontal();
      yield this.toString();

      this.rotateRight();
    }
  }

  flipHorizontal () {
    const size = this._pixels.length;
    for (let i = 0; i < size; i++) {
      for (let j = 0; j < Math.floor(size / 2); j++) {
        this._swap(i, j, i, size - j - 1);
      }
    }
  }

  flipVertical () {
    const size = this._pixels.length;
    for (let i = 0; i < Math.floor(size / 2); i++) {
      for (let j = 0; j < size; j++) {
        this._swap(i, j, size - i - 1, j);
      }
    }
  }

  rotateRight () {
    const size = this._pixels.length;
    for (let i = 0; i < Math.floor(size / 2); i++) {
      for (let j = i; j < size - i - 1; j++) {
        this._swap(i, j, size - j - 1, i);
        this._swap(size - j - 1, i, size - i - 1, size - j - 1);
        this._swap(size - i - 1, size - j - 1, j, size - i - 1);
      }
    }
  }

  countPixels (ch = '#') {
    let count = 0;
    for (const row of this._pixels) {
      for (const pixel of row) {
        if (pixel === ch) {
          count++;
        }
      }
    }

    return count;
  }

  toString () {
    return this._pixels.map((value) => value.join('')).join('/');
  }

  _swap (aRow, aColumn, bRow, bColumn) {
    const temp = this._pixels[aRow][aColumn];
    this._pixels[aRow][aColumn] = this._pixels[bRow][bColumn];
    this._pixels[bRow][bColumn] = temp;
  }
}

class EnhancementRule {
  constructor (search, replacement) {
    this._search = Image.fromString(search);
    this._replacement = Image.fromString(replacement);
  }

  get replacement () {
    return this._replacement;
  }

  getPermutations () {
    return Array.from(this._search.getPermutations());
  }
}

class EnhancementRunner {
  constructor (rules) {
    this._rules = new Map();

    for (const rule of rules) {
      const permutations = rule.getPermutations();
      for (const permutation of permutations) {
        if (this._rules.has(permutation) && this._rules.get(permutation) !== rule) {
          throw new Error('Duplicate pattern');
        }

        this._rules.set(permutation, rule);
      }
    }
  }

  enhance (image) {
    const imageSize = image.size;
    const squareSize = EnhancementRunner._getSquareSize(imageSize);
    const numSquares = imageSize / squareSize;
    const newImage = new Image();

    for (let i = 0; i < numSquares; i++) {
      for (let j = 0; j < numSquares; j++) {
        this._enhanceSegment(image, i, j, squareSize, newImage);
      }
    }

    return newImage;
  }

  static _getSquareSize (imageSize) {
    if (imageSize % 2 === 0) {
      return 2;
    } else if (imageSize % 3 === 0) {
      return 3;
    } else {
      throw new Error('Unexpected image size');
    }
  }

  _enhanceSegment (image, row, column, squareSize, newImage) {
    const subImage = image.getSubImage(row * squareSize, column * squareSize, squareSize, squareSize).toString();
    const replacement = this._rules.get(subImage).replacement;
    const newSquareSize = replacement.size;

    for (let i = 0; i < newSquareSize; i++) {
      for (let j = 0; j < newSquareSize; j++) {
        newImage.setPixel(row * newSquareSize + i, column * newSquareSize + j, replacement.getPixel(i, j));
      }
    }
  }
}

class Day21 {
  static * parseRules (str) {
    const regexp = /^([.#/]+) => ([.#/]+)$/gm;
    let result = null;

    do {
      result = regexp.exec(str);
      if (result !== null) {
        yield new EnhancementRule(result[1], result[2]);
      }
    } while (result != null);
  }

  static * run (input) {
    const fileContent = fs.readFileSync(input, 'utf8');
    const rules = this.parseRules(fileContent);
    const runner = new EnhancementRunner(rules);

    const startPosition = '.#./..#/###';
    let image = Image.fromString(startPosition);

    for (let i = 0; i < 18; i++) {
      image = runner.enhance(image);

      if (i === 4) {
        yield image.countPixels();
      }
    }

    yield image.countPixels();
  }
}

module.exports = Day21;
