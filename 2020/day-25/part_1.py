from diffie_hellman import A, B, g, p, discrete_log

a = discrete_log(A, g, p)
print(pow(B, a, p))
