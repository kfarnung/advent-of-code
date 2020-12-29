const fs = require('fs');

class ProgramNode {
  constructor(name) {
    this._name = name;

    this._weight = -1;
    this._parent = null;
    this._childNodes = [];
  }

  get name() {
    return this._name;
  }

  get weight() {
    return this._weight;
  }

  set weight(val) {
    this._weight = val;
  }

  get parent() {
    return this._parent;
  }

  set parent(node) {
    if (this._parent !== null) {
      throw new Error("parent can't be set twice");
    }

    this._parent = node;
  }

  get childNodes() {
    return this._childNodes;
  }

  addChildNode(node) {
    this._childNodes.push(node);
  }

  getTotalWeight() {
    let sum = 0;

    for (const child of this._childNodes) {
      sum += child.getTotalWeight();
    }

    return sum + this._weight;
  }

  getUnbalancedChild() {
    const map = new Map();
    for (const childNode of this._childNodes) {
      const totalWeight = childNode.getTotalWeight();

      let weightList = map.get(totalWeight);
      if (weightList === undefined) {
        weightList = [];
        map.set(totalWeight, weightList);
      }

      weightList.push(childNode);
    }

    if (map.size > 1) {
      for (const list of map.values()) {
        if (list.length === 1) {
          const unbalancedChild = list[0].getUnbalancedChild();

          if (unbalancedChild !== null) {
            return unbalancedChild;
          }

          return list[0];
        }
      }
    }

    return null;
  }
}

class ProgramGraph {
  constructor() {
    this._map = new Map();
  }

  addRawData(rawData) {
    const node = this.findOrCreateNode(rawData[0]);
    node.weight = rawData[1];

    for (const childName of rawData.slice(2)) {
      const childNode = this.findOrCreateNode(childName);
      childNode.parent = node;
      node.addChildNode(childNode);
    }
  }

  findOrCreateNode(name) {
    let node = this._map.get(name);

    if (node === undefined) {
      node = new ProgramNode(name);
      this._map.set(name, node);
    }

    return node;
  }

  findRootNode() {
    const first = this._map.values().next();
    if (!first.done) {
      let rootNode = first.value;
      while (rootNode.parent !== null) {
        rootNode = rootNode.parent;
      }

      return rootNode;
    }
  }

  getCorrectedWeight() {
    const unbalancedChild = this.findRootNode().getUnbalancedChild();

    for (const childNode of unbalancedChild.parent.childNodes) {
      if (childNode !== unbalancedChild) {
        const diff =
          childNode.getTotalWeight() - unbalancedChild.getTotalWeight();
        return unbalancedChild.weight + diff;
      }
    }
  }
}

class Day07 {
  static *parseRows(str) {
    let current = '';
    let row = [];

    for (const ch of str) {
      if (ch === '\n' || ch === ' ' || ch === ')') {
        if (current.length > 0) {
          if (ch === ')') {
            current = Number.parseInt(current);
          }

          row.push(current);
          current = '';
        }

        if (ch === '\n') {
          yield row;
          row = [];
        }
      } else if ((ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')) {
        current += ch;
      }
    }
  }

  static run(input) {
    const fileContent = fs.readFileSync(input, 'utf8');

    const graph = new ProgramGraph();
    for (const row of this.parseRows(fileContent)) {
      graph.addRawData(row);
    }

    return [graph.findRootNode().name, graph.getCorrectedWeight()];
  }
}

module.exports = Day07;
