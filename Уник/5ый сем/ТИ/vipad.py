def norma(string):
    count = 0
    for i in string:
        if i == '1':
            count += 1
    return count

def index(string):
    answ = 0
    for i in range(0, len(string)):
        if string[i] == '1':
            answ += (i + 1)
    return answ

def vipad():
    string = str(input("Введи строку после выпадения\n"))
    n = len(string) + 1
    k = n + 1
    W_b = index(string)
    if W_b == 0:
        print("Выпал последний \"0\"")
        string += '0'
        print(string)
        return
    
    norm = norma(string)
    delta = (k - W_b % k) % k
    print(f"Количество еденичеГ - {norm};\nW(b) - {W_b};\nDelta - {delta};\n")
    if norm >= delta:
        print("Выпал \"0\", справа", delta, "\"1\"")
    else:
        print("Выпал \"1\", справа", n - delta, "\"0\"")



def vstavka():
    string = str(input("Введи строку после вставки\n"))
    k = len(string)
    n = k - 1
    W_b = index(string)
    delta = W_b % k
    norm = norma(string)
    print(f"Количество еденичеГ - {norm};\nW(b) - {W_b};\nDelta - {delta};\n")
    if delta == norm:
        print("Первый символ лишний")
    elif delta == 0:
        print("Последний символ лишний")
    elif norm > delta:
        print("Вставка - \"0\", справа", delta, "\"1\"")
    else:
        print("Вставка - \"1\", справа", n + 1 - delta, "\"0\"")


mode = int(input("1 - выпадение, 2 - вставка\n"))
if mode == 1:
    vipad()
else:
    vstavka()