#include <iostream>
#include <vector>
#include <string>
#include <sstream>

using namespace std;

struct Deque {
    vector<string> strings;
    int first = 0;
    int last = 0;
    int size = 0;
    int cap = 0;

    void print() const {
        if (size == 0) {
            cout << "empty\n";
            return;
        }
        for (int i = 0; i < size; i++) {
            int idx = (first + i) % cap;
            cout << strings[idx] << ' ';
        }
        cout << '\n';
    }

    string popb() {
        if (size == 0)
            return "underflow";
        last = (last + cap - 1) % cap;
        string temp = strings[last];
        size--;
        return temp;
    }
    string popf() {
        if (size == 0)
            return "underflow";
        string temp = strings[first];
        first = (first + 1) % cap;
        size--;
        return temp;
    }
    void pushb(const string& str) {
        if (size == cap) {
            cout << "overflow\n";
            return;
        }
        size++;
        strings[last] = str;
        last = (last + 1) % cap;
    }
    void pushf(const string& str) {
        if (size == cap) {
            cout << "overflow\n";
            return;
        }
        size++;
        first = (first + cap - 1) % cap;
        strings[first] = str;
    }
};

int main() {
    string line;
    Deque deque;
    bool flag = false;

    while (true) {
        if (!getline(cin, line))
            break;
        if (line.empty())
            continue;

        int count = 0;
        for (const char c : line)
            if (c == ' ')
                count++;

        stringstream ss(line);
        vector<string> command;
        string temp;
        while (ss >> temp)
            command.push_back(temp);

        if (!flag) {
            if (command[0] != "set_size" || command.size() != 2) {
                cout << "error\n";
                continue;
            }
            int size;
            try {
                size = stoi(command[1]);
            } catch (...) {
                cout << "error\n";
                continue;
            }
            if (size < 0) {
                cout << "error\n";
                continue;
            }
            deque = Deque{vector<string>(size), 0, 0, 0, size};
            flag = true;
            continue;
        }

        const string& cmd = command[0];

        if (cmd == "print") {
            if (command.size() != 1 || count > 0) {
                cout << "error\n";
                continue;
            }
            deque.print();
        }
        else if (cmd == "pushf") {
            if (command.size() != 2) {
                cout << "error\n";
                continue;
            }
            deque.pushf(command[1]);
        }
        else if (cmd == "pushb") {
            if (command.size() != 2) {
                cout << "error\n";
                continue;
            }
            deque.pushb(command[1]);
        }
        else if (cmd == "popf") {
            if (command.size() != 1 || count > 0) {
                cout << "error\n";
                continue;
            }
            cout << deque.popf() << '\n';
        }
        else if (cmd == "popb") {
            if (command.size() != 1 || count > 0) {
                cout << "error\n";
                continue;
            }
            cout << deque.popb() << '\n';
        }
        else {
            cout << "error\n";
        }
    }
}
