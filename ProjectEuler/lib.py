# PRIME NUMBERS
def prime(n):
    if n == 1:
        return False
    if n % 2 == 0 or n % 3 == 0 or n % 5 == 0:
        return False
    for i in range(6, n):
        if n % i == 0:
            return False
    return True
def tentobin(n):
    if n==0: return ''
    else:
        return tentobin(n/2) + str(n%2)
def set_primes(up):
    piv = []
    for i in range(2,up+1):
        piv.append(i)
    to_rem = []
    start = 0
    while piv[start]**2 < up+1:
        print(len(piv))
        for i in range(start + 1, len(piv)):
            if piv[i] % piv[start] == 0:
                to_rem.append(piv[i])
        for num in to_rem:
            piv.remove(num)
        to_rem = []
        start+=1
    f = open("data", "w")
    for num in piv:
        f.write(str(num) + '\n')
    f.close()
def get_primes():
    f = open("data", "r")
    res = dict()
    for line in f:
        a = list(map(int,line.split()))
        res[a[0]] = a[1]
        print(a[0], a[1])
    f.close()
    return res
def GCD(m,n):
    mult = 1
    while True:
        if m == 0 or n == 0 or m == n:
            return mult*max(n,m)
        if m == 1 or n == 1:
            return mult
        if m % 2 == 0 and n % 2 == 0:
            mult *= 2
            m = m//2
            n = n//2
        if m % 2 == 0 and n % 2 != 0:
            m = m//2
        if m % 2 != 0 and n % 2 == 0:
            n = n//2
        if m % 2 != 0 and n % 2 != 0:
            if n > m:
                piv = (n-m)//2
                n = m
                m = piv
            elif n < m:
                m = (m-n)//2
def eratosthenes(n):
    numbers = list(range(2, n + 1))
    for number in numbers:
        if number != 0:
            for candidate in range(2 * number, n+1, number):
                numbers[candidate-2] = 0
    return list(filter(lambda x: x != 0, numbers))
class Pandig:
    # class for making pan dig numbers in recursive way
    # generate and return list of all dig
    # from 0 to number can be changed to 1 etc
    def __init__(self, number, ban=[]):
        self.number = []
        for i in range(0,number + 1):
            if i not in ban:
                self.number.append(i)
        self.child = []
        for i in self.number:
            self.child.append(Pandig(number, ban + [i]))
    def get_n(self):
        res = []
        if len(self.number) == 0:
            return ['']
        for i in range(len(self.number)):
            low = self.child[i].get_n()
            for j in low:
                res.append(str(self.number[i]) + j)
        return res
    def get_number(self):
        r = list(map(int, self.get_n()))
        return r
        
def pandig(n):
    # make all pandig numbers from n
    # return list of all of them
    import itertools
    res = []
    for num in list(itertools.permutations(str(n))):
        piv = ''
        for d in num:
            piv += d
        res.append(int(piv))
    return res
    
    
