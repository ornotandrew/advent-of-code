from expression import Node


with open("input.txt", "r") as f:
    lines = f.read().splitlines()


print(sum(Node(line).evaluate(op_precedence=["+", "*"]) for line in lines))
