use std::collections::HashMap;

struct Node {
    parent: Option<String>,
    children: Vec<String>,
}

impl Node {
    fn new(parent: Option<&str>) -> Self {
        let parent_value = match parent {
            Some(x) => Some(x.to_string()),
            None => None,
        };

        return Self {
            parent: parent_value,
            children: Vec::new(),
        };
    }

    fn set_parent(&mut self, name: &str) {
        if let Some(parent) = &self.parent {
            panic!("Found existing parent {} when setting {}", parent, name);
        }

        self.parent = Some(name.to_string());
    }
}

pub struct NamedTree {
    nodes: HashMap<String, Node>,
}

impl NamedTree {
    pub fn new() -> Self {
        return Self {
            nodes: HashMap::new(),
        };
    }

    pub fn add_child(&mut self, parent: &str, child: &str) {
        if let Some(child_node) = self.nodes.get_mut(child) {
            child_node.set_parent(parent);
        } else {
            self.nodes
                .insert(child.to_string(), Node::new(Some(parent)));
        }

        if !self.nodes.contains_key(parent) {
            self.nodes.insert(parent.to_string(), Node::new(None));
        }

        if let Some(parent_node) = self.nodes.get_mut(parent) {
            parent_node.children.push(child.to_string());
        } else {
            panic!("Parent could not be found!");
        }
    }

    pub fn get_parent(&self, name: &str) -> Option<&str> {
        if let Some(node) = self.nodes.get(name) {
            return match &node.parent {
                Some(x) => Some(&x),
                None => None,
            };
        } else {
            panic!("Invalid parent name!");
        }
    }

    pub fn get_children(&self, name: &str) -> &Vec<String> {
        if let Some(node) = self.nodes.get(name) {
            return &node.children;
        } else {
            panic!("Invalid parent name!");
        }
    }
}
