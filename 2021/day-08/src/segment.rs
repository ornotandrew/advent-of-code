//   0:6     1:2     2:5     3:5     4:4
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....

//   5:5     6:6     7:3     8:7     9:6
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

#[derive(Debug, Clone, PartialEq, Eq, Hash, Copy)]
pub enum Segment {
    A,
    B,
    C,
    D,
    E,
    F,
    G,
}

impl Segment {
    pub fn from_str(s: char) -> Segment {
        match s {
            'a' => Segment::A,
            'b' => Segment::B,
            'c' => Segment::C,
            'd' => Segment::D,
            'e' => Segment::E,
            'f' => Segment::F,
            'g' => Segment::G,
            _ => panic!("Segment should be a-g"),
        }
    }
}
