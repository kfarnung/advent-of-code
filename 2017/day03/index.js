class Day03 {
  static getSideLength (count) {
    const length = Math.ceil(Math.sqrt(count));
    if (length % 2 === 0) {
      return length + 1;
    }

    return length;
  }

  static findCoordinates (count) {
    const length = this.getSideLength(count);
    const halfLength = Math.floor(length / 2);

    if (count === length * length) {
      return { x: halfLength, y: -halfLength };
    }

    const location = count - Math.pow(length - 2, 2);
    const side = Math.floor(location / (length - 1));
    const pos = location % (length - 1);

    switch (side) {
      case 0:
        return { x: halfLength, y: -halfLength + pos };

      case 1:
        return { x: halfLength - pos, y: halfLength };

      case 2:
        return { x: -halfLength, y: halfLength - pos };

      case 3:
        return { x: -halfLength + pos, y: -halfLength };
    }
  }

  static manhattanDistance (count) {
    const { x, y } = this.findCoordinates(count);
    return Math.abs(x) + Math.abs(y);
  }

  static findIndex (x, y) {
    let length = 0;
    let side = 0;
    let pos = 0;

    const absX = Math.abs(x);
    const absY = Math.abs(y);

    if (absX > absY) {
      length = absX * 2 + 1;
      if (x > 0) {
        side = 0;
        pos = y + absX;
      } else {
        side = 2;
        pos = absX - y;
      }
    } else {
      length = absY * 2 + 1;
      if (y > 0) {
        side = 1;
        pos = absY - x;
      } else {
        side = 3;
        pos = x + absY;
      }
    }

    return side * (length - 1) + pos + Math.pow(length - 2, 2) - 1;
  }

  static getValidIndex (arr, index) {
    if (index < arr.length) {
      return arr[index];
    }

    return 0;
  }

  static calculateAdjacentSum (arr, x, y) {
    let sum = 0;

    sum += this.getValidIndex(arr, this.findIndex(x + 1, y));
    sum += this.getValidIndex(arr, this.findIndex(x + 1, y + 1));
    sum += this.getValidIndex(arr, this.findIndex(x, y + 1));
    sum += this.getValidIndex(arr, this.findIndex(x - 1, y + 1));
    sum += this.getValidIndex(arr, this.findIndex(x - 1, y));
    sum += this.getValidIndex(arr, this.findIndex(x - 1, y - 1));
    sum += this.getValidIndex(arr, this.findIndex(x, y - 1));
    sum += this.getValidIndex(arr, this.findIndex(x + 1, y - 1));

    arr.push(sum);
    return sum;
  }

  static * generateSums () {
    const arr = [];
    let side = 3;
    let x = 0;
    let y = 0;

    arr.push(1);

    for (;; side += 2) {
      x++;
      yield this.calculateAdjacentSum(arr, x, y);

      for (let i = 2; i < side; i++) {
        y++;
        yield this.calculateAdjacentSum(arr, x, y);
      }

      for (let i = 1; i < side; i++) {
        x--;
        yield this.calculateAdjacentSum(arr, x, y);
      }

      for (let i = 1; i < side; i++) {
        y--;
        yield this.calculateAdjacentSum(arr, x, y);
      }

      for (let i = 1; i < side; i++) {
        x++;
        yield this.calculateAdjacentSum(arr, x, y);
      }
    }
  }

  static searchSums (search) {
    for (const sum of this.generateSums()) {
      if (sum > search) {
        return sum;
      }
    }
  }

  static run (input) {
    const num = Number.parseInt(input);

    return [
      this.manhattanDistance(num),
      this.searchSums(num)
    ];
  }
}

module.exports = Day03;
