def is_prime(num: int) -> bool:
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            return False
    return True

def solution(n: int) -> int:
    prime_counter = 1
    number = 3
    while prime_counter <= n:
        if is_prime(number):
            prime_counter += 1
        number += 1
    return number - 1


print(solution(10001))
