use std::{cmp::Ordering, str::FromStr};

#[derive(Debug, Clone)]
enum PacketContents {
    Literal(usize),
    Operator(Vec<Packet>),
}

#[derive(Debug, Clone)]
pub struct Packet {
    raw: String,
    version: usize,
    type_id: usize,
    contents: PacketContents,
}

impl Packet {
    pub fn get_value(&self) -> usize {
        if self.type_id == 4 {
            return match self.contents {
                PacketContents::Literal(value) => value,
                _ => unreachable!(),
            };
        }

        let children_values = match &self.contents {
            PacketContents::Operator(children) => children,
            _ => unreachable!(),
        }
        .iter()
        .map(|p| p.get_value());

        match self.type_id {
            0 => children_values.sum(),
            1 => children_values.product(),
            2 => children_values.min().unwrap(),
            3 => children_values.max().unwrap(),
            _ => {
                let first = children_values.clone().nth(0).unwrap();
                let second = children_values.clone().nth(1).unwrap();

                match self.type_id {
                    5 => match first.cmp(&second) {
                        Ordering::Greater => 1,
                        _ => 0,
                    },
                    6 => match first.cmp(&second) {
                        Ordering::Less => 1,
                        _ => 0,
                    },
                    7 => match first.cmp(&second) {
                        Ordering::Equal => 1,
                        _ => 0,
                    },
                    _ => unreachable!(),
                }
            }
        }
    }
}

impl Packet {
    fn parse(raw: &str) -> (usize, Packet) {
        let version = from_bin(&raw[..3]);
        let type_id = from_bin(&raw[3..6]);
        let mut content_length = 0;

        let contents = match type_id {
            4 => {
                let mut bin_value = String::new();
                let mut offset = 6;
                loop {
                    let group = &raw[offset..offset + 5];
                    bin_value += &group[1..];
                    if &group[0..1] == "0" {
                        break;
                    }
                    offset += 5;
                }
                content_length = offset + 5;
                PacketContents::Literal(from_bin(&bin_value))
            }
            _ => {
                let length_type_id = from_bin(&raw[6..7]);
                match length_type_id {
                    0 => {
                        let children_start_idx = 7 + 15;
                        let total_children_length = from_bin(&raw[7..children_start_idx]);
                        let children_raw =
                            &raw[children_start_idx..children_start_idx + total_children_length];
                        content_length += children_start_idx + total_children_length;
                        let mut children: Vec<Packet> = vec![];

                        let mut offset = 0;
                        while offset < children_raw.len() {
                            let (child_length, child) = Packet::parse(&children_raw[offset..]);
                            offset += child_length;
                            children.push(child);
                        }

                        PacketContents::Operator(children)
                    }
                    1 => {
                        let children_start_idx = 7 + 11;
                        let num_children = from_bin(&raw[7..children_start_idx]);
                        let children_raw = &raw[children_start_idx..];
                        let mut children: Vec<Packet> = vec![];

                        let mut offset = 0;
                        while children.len() < num_children {
                            let (child_length, child) = Packet::parse(&children_raw[offset..]);
                            offset += child_length;
                            children.push(child);
                        }
                        content_length += children_start_idx + offset;

                        PacketContents::Operator(children)
                    }
                    _ => {
                        unreachable!();
                    }
                }
            }
        };

        (
            content_length,
            Packet {
                version,
                type_id,
                raw: raw.to_string(),
                contents,
            },
        )
    }

    pub fn sum_version_numbers(&self) -> usize {
        self.version
            + match &self.contents {
                PacketContents::Literal(_) => 0,
                PacketContents::Operator(children) => {
                    children.iter().map(|p| p.sum_version_numbers()).sum()
                }
            }
    }
}

pub fn from_bin(s: &str) -> usize {
    usize::from_str_radix(s, 2).unwrap()
}

#[derive(Debug)]
pub struct ParseError;

impl FromStr for Packet {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut raw: String = String::new();
        // these numbers are way too large for a u64
        for c in s.trim().chars() {
            raw += &format!("{:0>4b}", c.to_digit(16).unwrap());
        }
        Ok(Packet::parse(&raw).1)
    }
}
