use std::{
    collections::{HashMap, HashSet},
    str::FromStr,
};

type Node = String;
type Path = Vec<Node>;

fn head(path: &Path) -> Node {
    path.iter().last().unwrap().to_string()
}

fn get_duplicate_small_node(path: &Path) -> Option<Node> {
    let mut seen: HashSet<&Node> = HashSet::new();
    for node in path {
        if is_big_cave(node) {
            continue;
        }
        if seen.contains(node) {
            return Some(node.clone());
        }
        seen.insert(node);
    }
    None
}

fn is_big_cave(node: &Node) -> bool {
    node.chars().nth(0).unwrap().is_uppercase()
}

#[derive(Debug)]
pub struct Cave {
    edges: HashMap<Node, Vec<Node>>,
}

impl Cave {
    pub fn all_possible_paths(&self, allow_dupicate_small_cave: bool) -> Vec<Path> {
        let mut paths: Vec<Path> = vec![];
        let mut queue: Vec<Path> = vec![vec!["start".to_string()]];

        while queue.len() > 0 {
            let current_path = queue.pop().unwrap();
            let current_node = head(&current_path);

            // if we're at an end node, stop
            if current_node == "end" {
                paths.push(current_path);
                continue;
            }

            let next_nodes: Vec<Node> = self.edges[&current_node]
                .iter()
                .cloned()
                .filter(|node| {
                    if node == "start" {
                        return false;
                    }
                    if is_big_cave(node) {
                        return true;
                    }
                    if !allow_dupicate_small_cave {
                        return !current_path.contains(node);
                    }

                    match get_duplicate_small_node(&current_path) {
                        Some(dup) => dup != *node && !current_path.contains(node),
                        None => true,
                    }
                })
                .collect();

            for node in next_nodes {
                let mut new_path = current_path.clone();
                new_path.push(node.clone());
                queue.push(new_path);
            }
        }

        paths
    }
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Cave {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut edges: HashMap<Node, Vec<Node>> = HashMap::new();

        for line in s.trim().split("\n") {
            let parts: Vec<&str> = line.split("-").collect();
            let left = parts[0].to_string();
            let right = parts[1].to_string();

            edges
                .entry(left.clone())
                .or_insert(vec![])
                .push(right.clone());
            edges
                .entry(right.clone())
                .or_insert(vec![])
                .push(left.clone());
        }

        Ok(Cave { edges })
    }
}
