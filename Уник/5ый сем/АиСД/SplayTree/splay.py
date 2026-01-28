import sys

sys.setrecursionlimit(2048)

class Node:
    def __init__(self, father, value, key):
        self.father = father
        self.value = value
        self.key = key
        self.right = None
        self.left = None

    def hasleft(self):
        return self.left is not None

    def hasright(self):
        return self.right is not None

    def isRoot(self):
        return self.father is None

class splayTree:
    def __init__(self):
        self.root = None

    # обозначения, как на картинках в википедии
    def __zig(self, x):

        p = x.father
        if p.isRoot():
            # просто подтянем и всё
            self.root = x
            x.father = None
        else:
            # если есть дед, то просто не забыть у него параметры поменять
            g = p.father
            if p == g.left:
                g.left = x
            else:
                g.right = x
            x.father = g

        if x == p.left:
            p.left = x.right
            if x.right:
                x.right.father = p
            x.right = p
        else:
            p.right = x.left
            if x.left:
                x.left.father = p
            x.left = p

        p.father = x

    def __zig_zig(self, x):
        self.__zig(x.father)
        self.__zig(x)

    def __zig_zag(self, x):
        self.__zig(x)
        self.__zig(x)

    def __splay(self, node):
        while not node.isRoot():
            p = node.father
            if p.isRoot():
                self.__zig(node)
            elif (node == p.left) == (p == p.father.left):
                self.__zig_zig(node)
            else:
                self.__zig_zag(node)

    # как и в прошлом номере все функции писал попарно: для ноды и для дерева
    def __find(self, cur_node, key):
        last_visited = None
        while cur_node:
            last_visited = cur_node
            if cur_node.key == key:
                return cur_node, True
            elif key < cur_node.key:
                cur_node = cur_node.left
            else:
                cur_node = cur_node.right
        return last_visited, False
    
    def search(self, key):
        if not self.root:
            return False
        node, found = self.__find(self.root, key)
        self.__splay(node)
        if not found:
            return False
        return node.value

    def __insert(self, cur_node, key, value):
        if cur_node.key == key:
            self.__splay(cur_node)
            return cur_node, False
        elif key < cur_node.key:
            if cur_node.hasleft():
                return self.__insert(cur_node.left, key, value)
            else:
                new_node = Node(cur_node, value, key)
                cur_node.left = new_node
                return new_node, True
        else:
            if cur_node.hasright():
                return self.__insert(cur_node.right, key, value)
            else:
                new_node = Node(cur_node, value, key)
                cur_node.right = new_node
                return new_node, True

    def add(self, key, value):
        if not self.root:
            self.root = Node(None, value, key)
        else:
            newNode, created = self.__insert(self.root, key, value)
            if not created:
                print('error')
            else:
                self.__splay(newNode)

    def set(self, key, value):
        if not self.root:
            raise KeyError()
        node, found = self.__find(self.root, key)
        self.__splay(node)
        if not found:
            raise KeyError()
        node.value = value
    
    # решил вынести
    def __merge(self, left, right):
        if not left:
            return right
        if not right:
            return left
        max_node = self.__find_max(left) # левосторонняя
        self.__splay(max_node)
        max_node.right = right
        right.father = max_node
        return max_node
    
    def delete(self, key):
        if not self.root:
            raise KeyError()
        node, found = self.__find(self.root, key)
        self.__splay(node)
        if not found:
            raise KeyError()
        self.root = self.__merge(node.left, node.right)
        if self.root:
            self.root.father = None

    # Основная проблема этой задачи заключается в выводе, ибо хранить тысячи nil'ов очень затратно
    # поэтому было решено использовать Run-Length Encoding
    # т.е. вместо тысячи nil nil nil nil .... nil В очереди
    # Мы будем хранить 1000 (Число подряд идущих nil'ов)
    def print(self, stream):
        if self.root is None:
            stream.write('_\n')
            return
        stream.write(f'[{self.root.key} {self.root.value}]\n')
        queue = []
        if self.root.hasleft():
            queue.append(self.root.left)
        else:
            queue.append(1)
        if self.root.hasright():
            queue.append(self.root.right)
        else:
            if (queue[-1] == 1):
                return
            else:
                queue.append(1)

        anyOne = True
        while anyOne:
            out = ""
            anyOne = False
            next = []
            for i in range(len(queue) - 1):
                if type(queue[i]) == int:
                    out += '_ ' * queue[i]
                    if len(next) == 0 or type(next[-1]) != int:
                        next.append(2 * queue[i])
                    else:
                        next[-1] += 2 * queue[i]
                else:
                    out += f'[{queue[i].key} {queue[i].value} {queue[i].father.key}] '
                    if queue[i].hasleft():
                        next.append(queue[i].left)
                        anyOne = True
                    else:
                        if len(next) == 0:
                            next.append(1)
                        elif type(next[-1]) == int:
                            next[-1] += 1
                        else:
                            next.append(1)
                    if queue[i].hasright():
                        next.append(queue[i].right)
                        anyOne = True
                    else:
                        if len(next) == 0:
                            next.append(1)
                        elif type(next[-1]) == int:
                            next[-1] += 1
                        else:
                            next.append(1)
            # теперь вывод последних
            temp = queue[-1]
            if type(temp) == int:
                out += '_ ' * (temp - 1)
                out += '_\n'
                if len(next) == 0 or type(next[-1]) != int:
                    next.append(2 * temp)
                else:
                    next[-1] += 2 * temp
            else:
                out += f'[{temp.key} {temp.value} {temp.father.key}]\n'
                if temp.hasleft():
                    anyOne = True
                    next.append(temp.left)
                else:
                    if len(next) == 0:
                        next.append(1)
                    elif type(next[-1]) == int:
                        next[-1] += 1
                    else:
                        next.append(1)
                if temp.hasright():
                    anyOne = True
                    next.append(temp.right)
                else:
                    if len(next) == 0:
                        next.append(1)
                    elif type(next[-1]) == int:
                        next[-1] += 1
                    else:
                        next.append(1)
            queue = next
            stream.write(out)

    # То же самое, что и в прошлом задании,
    # только не рекурсивно
    def __find_min(self, cur_node):
        while cur_node.hasleft():
            cur_node = cur_node.left
        return cur_node

    def __find_max(self, cur_node):
        while cur_node.hasright():
            cur_node = cur_node.right
        return cur_node

    def min(self):
        if not self.root:
            raise KeyError()
        min_node = self.__find_min(self.root)
        self.__splay(min_node)
        return min_node.key, min_node.value

    def max(self):
        if not self.root:
            raise KeyError()
        max_node = self.__find_max(self.root)
        self.__splay(max_node)
        return max_node.key, max_node.value

tree = splayTree()
for line in sys.stdin:
    line = line.rstrip('\n')
    if not line:
        continue
    fields = line.split()
    cmd = fields[0]
    try:
        if cmd == 'add':
            if len(fields) == 3:
                tree.add(int(fields[1]), fields[2])
            elif len(fields) == 2:
                tree.add(int(fields[1]), '')
            else:
                print('error')
        elif cmd == 'set':
            if len(fields) == 3:
                tree.set(int(fields[1]), fields[2])
            elif len(fields) == 2:
                tree.set(int(fields[1]), '')
            else:
                print('error')
        elif cmd == 'delete':
            if len(fields) != 2:
                print('error')
            else:
                tree.delete(int(fields[1]))
        elif cmd == 'search':
            if len(fields) != 2:
                print('error')
            else:
                answ = tree.search(int(fields[1]))
                if answ != False:
                    print(f'1 {answ}')
                else:
                    print(0)
        elif cmd == 'min':
            if len(fields) != 1:
                print('error')
            else:
                node = tree.min()
                print(f'{node[0]} {node[1]}')
        elif cmd == 'max':
            if len(fields) != 1:
                print('error')
            else:
                node = tree.max()
                print(f'{node[0]} {node[1]}')
        elif cmd == 'print':
            if len(fields) != 1:
                print('error')
            else:
                tree.print(sys.stdout)
        else:
            print('error')
    except KeyError:
        print('error')
