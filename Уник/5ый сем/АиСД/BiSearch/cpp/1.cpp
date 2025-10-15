#include <iostream>
#include <vector>
#include <sstream>
#include <string>

int biSearchRec(const std::vector<int>& nums, int target, int low, int high) {
    if (low > high) {
        return -1;
    }
    int mid = low + (high - low) / 2;

    if (nums[mid] == target) {
        int left = biSearchRec(nums, target, low, mid - 1);
        if (left == -1)
            return mid;
        return left;
    } else if (nums[mid] < target) {
        return biSearchRec(nums, target, mid + 1, high);
    } else {
        return biSearchRec(nums, target, low, high - 1);
    }
}

int main() {
    std::string line;
    getline(std::cin, line);
    std::stringstream ss(line);
    std::vector<int> nums;
    int temp;
    while (ss >> temp) {
        nums.push_back(temp);
    }

    while (getline(std::cin, line)) {
        if (line.empty()) continue;

        std::string search_word;
        int key;
        std::stringstream ss_2(line);
        ss_2 >> search_word >> key;
        std::cout << biSearchRec(nums, key, 0, (int)(nums.size() - 1)) << '\n';
    }
}