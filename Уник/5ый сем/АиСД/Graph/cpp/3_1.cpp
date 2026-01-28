#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <string>
#include <vector>

void UndirAdd(std::map<std::string, std::vector<std::string>>& graph,
                std::vector<std::string>& edges) {
    if (std::find(graph[edges[0]].begin(), graph[edges[0]].end(), edges[1]) == graph[edges[0]].end()) {
        graph[edges[0]].push_back(edges[1]);
    }
    if (std::find(graph[edges[1]].begin(), graph[edges[1]].end(), edges[0]) == graph[edges[1]].end()) {
        graph[edges[1]].push_back(edges[0]);
    }
}

void DirAdd(std::map<std::string, std::vector<std::string>>& graph,
            std::vector<std::string>& edges) {
    if (std::find(graph[edges[0]].begin(), graph[edges[0]].end(), edges[1]) == graph[edges[0]].end()) {
        graph[edges[0]].push_back(edges[1]);
    }
}

void Glubinu(std::map<std::string, std::vector<std::string>>& graph,
        const std::string& startVertex,
        std::map<std::string, bool>& visited) {

    if (visited[startVertex]) return;
    visited[startVertex] = true;

    std::cout << startVertex << std::endl;

    std::vector<std::string> neighbors = graph[startVertex];
    std::sort(neighbors.begin(), neighbors.end());
    for (const auto& next : neighbors) {
        Glubinu(graph, next, visited);
    }
}

void Shirinu(std::map<std::string, std::vector<std::string>>& graph,
        const std::string& startVertex) {

    std::map<std::string, bool> visited;
    std::queue<std::string> q;

    q.push(startVertex);

    while (q.size() != 0) {
        std::string node = q.front();
        q.pop();

        if (visited[node]) {
            continue;
        }
        visited[node] = true;

        std::cout << node << std::endl;

        std::vector<std::string> neighbors = graph[node];
        std::sort(neighbors.begin(), neighbors.end());

        for (const auto& next : neighbors) {
            if (visited[next] == false) {
                q.push(next);
            }
        }
    }
}

int main() {
    std::map<std::string, std::vector<std::string>> graph;

    std::string graphType, startNode, mode;
    if (!(std::cin >> graphType >> startNode >> mode))
        return 0;

    std::string left, right;
    while (std::cin >> left >> right) {
        std::vector<std::string> edge{left, right};

        if (graphType == "d") {
            DirAdd(graph, edge);
        } else if (graphType == "u") {
            UndirAdd(graph, edge);
        } else {
            std::cout << "error" << std::endl;
        }
    }

    if (mode == "d") {
        std::map<std::string, bool> visited;
        Glubinu(graph, startNode, visited);
    } else if (mode == "b") {
        Shirinu(graph, startNode);
    } else {
        std::cout << "error" << std::endl;
    }
}
