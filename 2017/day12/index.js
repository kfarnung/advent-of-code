const fs = require('fs');

class ProgramNode {
  constructor (id) {
    this._id = id;
    this._neighbors = [];
  }

  get id () {
    return this._id;
  }

  get neighbors () {
    return this._neighbors;
  }

  addNeighbor (node) {
    this._neighbors.push(node);
  }
}

class ProgramGraph {
  constructor () {
    this._map = [];
  }

  addRawData (rawData) {
    const node = this.findOrCreateNode(rawData[0]);

    for (const childId of rawData.slice(1)) {
      const neighbor = this.findOrCreateNode(childId);
      node.addNeighbor(neighbor);
      neighbor.addNeighbor(node);
    }
  }

  findOrCreateNode (id) {
    let node = this._map[id];

    if (node === undefined) {
      node = new ProgramNode(id);
      this._map[id] = node;
    }

    return node;
  }

  countNodesInGroup (id, visited = []) {
    const processingQueue = [];

    const startingNode = this.findOrCreateNode(id);
    visited.push(startingNode.id);
    processingQueue.push(startingNode);

    while (processingQueue.length > 0) {
      const current = processingQueue.shift();

      for (const neighbor of current.neighbors) {
        if (visited.indexOf(neighbor.id) < 0) {
          visited.push(neighbor.id);
          processingQueue.push(neighbor);
        }
      }
    }

    return visited.length;
  }

  countGroups () {
    const visited = [];
    let groupCount = 0;

    for (const startingNode of this._map) {
      if (visited.indexOf(startingNode.id) < 0) {
        this.countNodesInGroup(startingNode.id, visited);
        groupCount++;
      }
    }

    return groupCount;
  }
}

class Day12 {
  static * parseRows (str) {
    let current = '';
    let row = [];

    for (const ch of str) {
      if (ch === '\n' || ch === ' ' || ch === ',') {
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

    const graph = new ProgramGraph();
    for (const row of this.parseRows(fileContent)) {
      graph.addRawData(row);
    }

    return [
      graph.countNodesInGroup(0),
      graph.countGroups()
    ];
  }
}

module.exports = Day12;
