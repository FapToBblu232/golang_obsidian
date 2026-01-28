
def generate_matrix_bits(n, m):
    matrix = []
    for i in range(n):
        row = []
        block_size = 1 << (i + 1)  # 2^(i+1) через битовый сдвиг
        half_block = block_size >> 1  # block_size // 2
        
        for j in range(m):
            # Позиция внутри паттерна
            pos_in_block = (j + 1) % block_size
            # Если позиция во второй половине блока - ставим 1, иначе 0
            row.append(1 if pos_in_block >= half_block else 0)
        
        matrix.append(row)
    return matrix

m = int(input('Введите разрядность (aka наше m)\n'))
n = 2**m - 1
H = generate_matrix_bits(m, n)
for row in H:
    print(' '.join(f'{x:2}' for x in row))

mode = int(input(f"Если уже дана строка длины {n}, то введите 1\nЕсли нужно закодить, а потом поменять индекс, то 2\n"))

if mode == 1:
    V = str(input("Введи свою строку\n"))
    answ = []
    for i in range(m):
        bit = 0
        for j in range(n):
            bit ^= H[i][j] * int(V[j])
        answ.append(bit)

    print(answ)
    ind = 0
    for i in range(len(answ)):
        if answ[i] == 1:
            ind += 2 ** i
    print(f"Бит под номером {ind} изменён")
    
else:
