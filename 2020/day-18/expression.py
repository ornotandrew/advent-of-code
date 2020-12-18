class Node:
    operations = {"+": lambda a, b: a + b, "*": lambda a, b: a * b}

    def __init__(self, line):
        if isinstance(line, str):
            line = list([int(c) if c.isdigit() else c for c in line.replace(" ", "")])

        self.terms = []

        i = 0
        while i < len(line):
            next_open = line.index("(", i) if "(" in line[i:] else None
            if next_open is not None:
                # find the corresponding closing bracket
                nesting = 1
                for matching_close in range(next_open + 1, len(line)):
                    if line[matching_close] == "(":
                        nesting += 1
                    elif line[matching_close] == ")":
                        nesting -= 1
                    if nesting == 0:
                        break

                child = Node(line[next_open + 1 : matching_close])
                self.terms += line[i:next_open] + [child]
                i = matching_close + 1
            else:  # there are no more sub-expressions
                self.terms += line[i:]
                break

    def evaluate(self, op_precedence=None):
        flat = [t.evaluate(op_precedence) if isinstance(t, Node) else t for t in self.terms]

        def apply_op_at_pos(flat, pos):
            op = flat[pos]
            result = self.operations[op](flat[pos - 1], flat[pos + 1])
            return flat[: pos - 1] + [result] + flat[pos + 2 :]

        if op_precedence:
            for op in op_precedence:
                while op in flat:
                    pos = flat.index(op)
                    flat = apply_op_at_pos(flat, pos)
        else:  # just go left-to-right
            while len(flat) > 1:
                pos = next(i for i, op in enumerate(flat) if op in self.operations.keys())
                flat = apply_op_at_pos(flat, pos)

        return flat[0]

    def __repr__(self):
        result = ""
        for term in self.terms:
            result += f"{term}"

        return f"Node<{result}>"
