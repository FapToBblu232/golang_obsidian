#include <iostream>
#include <map>
#include <vector>
#include <string>
#include <algorithm>
#include <queue>

using namespace std;

bool contains(const vector<string>& vertexes, const string& target) {
    return find(vertexes.begin(), vertexes.end(), target) != vertexes.end();
}

void addNotOrientir(map<string, vector<string>>& graph, const vector<string>& verts) {
    const string& v1 = verts[0];
    const string& v2 = verts[1];
    if (!contains(graph[v1], v2))
        graph[v1].push_back(v2);
    if (!contains(graph[v2], v1))
        graph[v2].push_back(v1);
}

void addOrientir(map<string, vector<string>>& graph, const vector<string>& verts) {
    const string& v1 = verts[0];
    const string& v2 = verts[1];
    if (!contains(graph[v1], v2))
        graph[v1].push_back(v2);
}

void dfs(map<string, vector<string>>& graph, const string& start, map<string, bool>& visited) {
    if (visited[start]) return;
    visited[start] = true;
    cout << start << "\n";

    vector<string> sosedi = graph[start];
    sort(sosedi.begin(), sosedi.end());
    for (const auto& vert : sosedi) {
        dfs(graph, vert, visited);
    }
}

void bfs(map<string, vector<string>>& graph, const string& start) {
    map<string, bool> visited;
    queue<string> q;
    q.push(start);

    while (!q.empty()) {
        string node = q.front();
        q.pop();

        if (visited[node]) continue;
        visited[node] = true;
        cout << node << "\n";

        vector<string> sosedi = graph[node];
        sort(sosedi.begin(), sosedi.end());
        for (const auto& vert : sosedi) {
            if (!visited[vert]) {
                q.push(vert);
            }
        }
    }
}

int main() {
    map<string, vector<string>> graph;
    string type, start, traversalType;
    if (!(cin >> type >> start >> traversalType)) return 0;

    string v1, v2;
    while (cin >> v1 >> v2) {
        vector<string> pair = {v1, v2};
        if (type == "d") {
            addOrientir(graph, pair);
        } else if (type == "u") {
            addNotOrientir(graph, pair);
        }
    }

    if (traversalType == "d") {
        map<string, bool> visited;
        dfs(graph, start, visited);
    }
    else if (traversalType == "b") {
        bfs(graph, start);
    }
}
