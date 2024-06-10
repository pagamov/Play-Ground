from lib import eratosthenes, prime, Pandig, pandig
from time import time

def chain(n):
    res = [n]
    while True:
        piv = sum(list(map(lambda x: int(x)**2,[x for x in str(res[-1])])))
        if piv == 89:
            return True
        elif piv == 1:
            return False
        res.append(piv)

r = []
for i in range(2, 10000000+1):
    if chain(i):
        r.append(i)
print(len(r))
