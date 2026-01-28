#include <iostream>
#include <string>
#include <sstream>
#include <vector>

struct Node {
    long long key;
    std::string value;
    Node* left = nullptr;
    Node* right = nullptr;
    Node* father = nullptr;
};

struct Tree {
    Node* root = nullptr;

    void rightRotate(Node* y) {
        Node* x = y->left;
        if (x == nullptr) {
            return;
        }
        y->left = x->right;
        if (x->right) {
            x->right->father = y;
        }
        x->father = y->father;
        if (y->father == nullptr) {
            root = x;
        } else if (y->father->left == y) {
            y->father->left = x;
        } else {
            y->father->right = x;
        }
        x->right = y;
        y->father = x;
    }

    void leftRotate(Node* x) {
        Node* y = x->right;
        if (y == nullptr) {
            return;
        }
        x->right = y->left;
        if (y->left) {
            y->left->father = x;
        }
        y->father = x->father;
        if (x->father == nullptr) {
            root = y;
        } else if (x->father->left == x) {
            x->father->left = y;
        } else {
            x->father->right = y;
        }
        y->left = x;
        x->father = y;
    }

    void splay(Node* x) {
        while (x->father) {
            Node* p = x->father;
            Node* g = p->father;
            if (g == nullptr) { // zig
                if (x == p->left) {
                    rightRotate(p);
                } else  {
                    leftRotate(p);
                }
            } else if (x == p->left && p == g->left) { // zig-zig левый
                rightRotate(g);
                rightRotate(p);
            } else if (x == p->right && p == g->right) { // zig-zig правый
                leftRotate(g);
                leftRotate(p);
            } else if (x == p->right && p == g->left) { // zig-zag 2 штуки
                leftRotate(p);
                rightRotate(g);
            } else {
                rightRotate(p);
                leftRotate(g);
            }
        }
    }

    Node* search(long long key) {
        if (root == nullptr) {
            return nullptr;
        }
        Node* cur = root;
        Node* lastNotNull = nullptr;
        while (cur) {
            lastNotNull = cur;
            if (key == cur->key) {
                splay(cur);
                return cur;
            } else if (key < cur->key) {
                cur = cur->left;
            } else {
                cur = cur->right;
            }
        }
        if (lastNotNull) {
            splay(lastNotNull);
        }
        return nullptr;
    }

    void add(long long key, const std::string& value) {
        if (root == nullptr) {
            root = new Node{key, value};
            return;
        }
        Node* cur = root;
        Node* lastNotNull = nullptr;
        while (cur) {
            lastNotNull = cur;
            if (key == cur->key) {
                std::cout << "error" << std::endl;
                splay(cur);
                return;
            }
            if (key < cur->key) {
                cur = cur->left;
            } else {
                cur = cur->right;
            }
        }
        Node* newNode = new Node{key, value, nullptr, nullptr, lastNotNull};
        if (key < lastNotNull->key)
            lastNotNull->left = newNode;
        else
            lastNotNull->right = newNode;
        splay(newNode);
    }

    void set(long long key, const std::string& value) {
        Node* node = search(key);
        if (node == nullptr) {
            std::cout << "error" << std::endl;
            return;
        }
        node->value = value;
    }

    void del(long long key) {
        Node* node = search(key);
        if (node == nullptr) {
            std::cout << "error" << std::endl;
            return;
        }

        Node* leftTree = node->left;
        Node* rightTree = node->right;

        if (leftTree) {
            leftTree->father = nullptr;
        }
        if (rightTree) {
            rightTree->father = nullptr;
        }

        delete node;

        if (leftTree == nullptr) {
            root = rightTree;
        } else {
            Node* maxLeft = leftTree;
            while (maxLeft->right) maxLeft = maxLeft->right;
            splay(maxLeft);
            maxLeft->right = rightTree;
            if (rightTree) rightTree->father = maxLeft;
            root = maxLeft;
        }
    }

    Node* min() {
        if (root == nullptr) {
            return nullptr;
        }
        Node* cur = root;
        while (cur->left) {
            cur = cur->left;
        }
        splay(cur);
        return cur;
    }

    Node* max() {
        if (root == nullptr) {
            return nullptr;
        }
        Node* cur = root;
        while (cur->right) {
            cur = cur->right;
        }
        splay(cur);
        return cur;
    }

    void print() {
        if (root == nullptr) {
            std::cout << "_" << std::endl;
            return;
        }

        std::vector<Node*> level = {root};
        std::cout << "[" << root->key << " " << root->value << "]" << std::endl;

        while (true) {
            std::vector<Node*> next;
            bool any = false;
            for (Node* n : level) {
                if (n) {
                    next.push_back(n->left);
                    next.push_back(n->right);
                    if (n->left || n->right) any = true;
                } else {
                    next.push_back(nullptr);
                    next.push_back(nullptr);
                }
            }
            if (!any) return;
            bool first = true;
            for (Node* n : next) {
                if (!first) std::cout << " ";
                first = false;
                if (n == nullptr) {
                    std::cout << "_";
                } else {
                    std::cout << "[" << n->key << " " << n->value << " " << n->father->key << "]";
                }
            }
            std::cout << std::endl;
            level.swap(next);
        }
    }
};

int main() {
    Tree tree;
    std::string line;

    while (std::getline(std::cin, line)) {
        if (line.empty()) continue;
        if (line.front() == ' ' || line.back() == ' ' || line.find("  ") != std::string::npos) {
            std::cout << "error" << std::endl;
            continue;
        }

        std::stringstream ss(line);
        std::string cmd;
        ss >> cmd;

        if (cmd == "add") {
            long long key;
            std::string val;
            if (!(ss >> key >> val)) {
                std::cout << "error" << std::endl;
                continue;
            }
            tree.add(key, val);
        }
        else if (cmd == "set") {
            long long key;
            std::string val;
            if (!(ss >> key >> val)) {
                std::cout << "error" << std::endl;
                continue;
            }
            tree.set(key, val);
        }
        else if (cmd == "delete") {
            long long key;
            if (!(ss >> key)) {
                std::cout << "error" << std::endl;
                continue;
            }
            tree.del(key);
        }
        else if (cmd == "search") {
            long long key;
            if (!(ss >> key)) {
                std::cout << "error" << std::endl;
                continue;
            }
            Node* res = tree.search(key);
            if (!res) std::cout << "0" << std::endl;
            else std::cout << "1 " << res->value << std::endl;
        }
        else if (cmd == "min") {
            Node* res = tree.min();
            if (!res) std::cout << "error" << std::endl;
            else std::cout << res->key << " " << res->value << std::endl;
        }
        else if (cmd == "max") {
            Node* res = tree.max();
            if (!res) std::cout << "error" << std::endl;
            else std::cout << res->key << " " << res->value << std::endl;
        }
        else if (cmd == "print") {
            tree.print();
        }
        else {
            std::cout << "error" << std::endl;
        }
    }

    return 0;
}
