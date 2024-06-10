#include <iostream>
#include <ctime>

double diffclock(clock_t clock1,clock_t clock2) {
    double diffticks=clock2-clock1;
    double diffms=(diffticks)/(CLOCKS_PER_SEC);
    return diffms;
}

bool prime(long long i) {
    for (long long n = 2; n < i; n++) {
        if (i % n == 0) {
            return false;
        }
    }
    return true;
}

int main() {
    clock_t start = clock();

    long long odd = 9;

    while (true) {
        if (odd % 2 != 0 && !prime(odd)) {
            for (long long i = 2; i < odd; i++) {
                if (prime(i)) {
                    for (long long j = 1; j < odd; j++) {
                        if (odd == i + 2 * j * j) {
                            goto next;
                        }
                    }
                }
            }

            std::cout << "res = " << odd << std::endl;
            goto ext;
        }
        next:
        odd += 2;
    }
    ext:
    clock_t end = clock();
    std::cout << diffclock(start,end)<<std::endl;
    return 0;
}
