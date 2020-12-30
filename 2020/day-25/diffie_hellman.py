import math
from itertools import count

# For an explanation of the variable names used here, see
# https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange#Cryptographic_explanation

p = 20201227
g = 7
A, B = 11404017, 13768789  # my input


def discrete_log(β, α, n):
    """
    β = α^x mod n

    Given β, α and n, this function finds x. This is the Baby-step giant-step
    algorithm detailed here:

    https://en.wikipedia.org/wiki/Baby-step_giant-step#The_algorithm
    """
    m = math.ceil(math.sqrt(n))
    α_m_inverse = pow(α, -m, n)
    table = {}

    for j in range(m):
        table[pow(α, j, n)] = j

    γ = β

    for i in range(m):
        if γ in table:
            return i * m + table[γ]
        γ = (γ * α_m_inverse) % n
