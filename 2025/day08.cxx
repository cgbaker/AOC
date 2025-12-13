#include <sstream>
#include <string>
#include <iostream>
#include <algorithm>
#include <list>
#include <numeric>
#include <ranges>
#include <functional>
#include <vector>
#include <format>
#include <iostream>
#include <print>
#include <set>
#include <string>
#include <string_view>
#include <fstream>
#include <map>
#include <set>
#include <queue>

using std::cout;
using std::endl;
using std::list;
using std::pair;
using std::string;
using std::vector;
using std::map;
using std::set;
using std::tuple;

struct Box {
    int64_t x, y, z;
};

using T = std::tuple<Box, Box, int64_t>;

auto distance = [](Box const &b1, Box const &b2)
{
    int64_t dx = (b1.x - b2.x),
            dy = (b1.y - b2.y),
            dz = (b1.z - b2.z);
    return dx*dx + dy*dy + dz*dz;
};

bool operator<(Box const &b1, Box const &b2)
{
    return b1.x < b2.x ||
           (b1.x == b2.x && b1.y < b2.y) ||
           (b1.x == b2.x && b1.y == b2.y && b1.z < b2.z);
};

struct ShortestDistance {
    bool operator()(const T& a, const T& b) const {
        return std::get<2>(a) < std::get<2>(b);
    }
};

struct LargestDistance {
    bool operator()(const T& a, const T& b) const {
        return std::get<2>(a) > std::get<2>(b);
    }
};

int64_t part1(int K, const vector<Box> &boxes)
{
    std::priority_queue<T, std::vector<T>, ShortestDistance> heap;
    for (auto p1 = boxes.begin(); p1 != boxes.end(); p1++) {
        auto p2 = p1;
        ++p2;
        for (; p2 != boxes.end(); p2++) {
            auto new_pair = std::make_tuple(*p1, *p2, distance(*p1, *p2));
            if (heap.size() < K) {
                heap.push(new_pair);
            } else if (std::get<2>(new_pair) < std::get<2>(heap.top())) {
                heap.pop();
                heap.push(new_pair);
            }
        }
    }
    vector<set<Box>> circuits;
    for (; !heap.empty(); heap.pop()) {
        auto t = heap.top();
        auto[b1,b2,d] = t;
        auto setb1 = std::find_if(circuits.begin(), circuits.end(), [b1](auto circuit){
            return circuit.contains(b1);
        });
        auto setb2 = std::find_if(circuits.begin(), circuits.end(), [b2](auto circuit){
            return circuit.contains(b2);
        });
        if (setb1 == circuits.end() && setb2 == circuits.end()) {
            circuits.push_back({b1,b2});
        } else if (setb1 == circuits.end() && setb2 != circuits.end()) {
            setb2->insert(b1);
        } else if (setb1 != circuits.end() && setb2 == circuits.end()) {
            setb1->insert(b2);
        } else if (setb1 != setb2) {
            setb1->insert(setb2->begin(), setb2->end());
            circuits.erase(setb2);
        }
    }
    std::sort(circuits.begin(), circuits.end(), [](const set<Box> &s1, const set<Box> &s2){
        return s1.size() > s2.size();
    });
    return circuits[0].size() * circuits[1].size() * circuits[2].size();
}

int64_t part2(const vector<Box> &boxes) {
    std::priority_queue<T, std::vector<T>, LargestDistance> heap;
    for (auto p1 = boxes.begin(); p1 != boxes.end(); p1++) {
        auto p2 = p1;
        ++p2;
        for (; p2 != boxes.end(); p2++) {
            auto new_pair = std::make_tuple(*p1, *p2, distance(*p1, *p2));
            heap.push(new_pair);
        }
    }
    vector<set<Box>> circuits;
    for (; !heap.empty(); heap.pop()) {
        auto t = heap.top();
        auto[b1,b2,d] = t;
        auto setb1 = std::find_if(circuits.begin(), circuits.end(), [b1](auto circuit){
            return circuit.contains(b1);
        });
        auto setb2 = std::find_if(circuits.begin(), circuits.end(), [b2](auto circuit){
            return circuit.contains(b2);
        });
        if (setb1 == circuits.end() && setb2 == circuits.end()) {
            circuits.push_back({b1,b2});
        } else if (setb1 == circuits.end() && setb2 != circuits.end()) {
            setb2->insert(b1);
        } else if (setb1 != circuits.end() && setb2 == circuits.end()) {
            setb1->insert(b2);
        } else if (setb1 != setb2) {
            setb1->insert(setb2->begin(), setb2->end());
            circuits.erase(setb2);
        } 
        if (circuits.size() == 1 && circuits[0].size() == boxes.size()) {
            std::println("full connect with ({},{},{}) and ({},{},{})", b1.x, b1.y, b1.z, b2.x, b2.y, b2.z);
            return b1.x * b2.x;
        }
    }
    return -1;
}

int main(int argc, char *argv[])
{
    std::ifstream fin(argv[1]);
    string ln;
    vector<Box> boxes;
    int K;
    {
        std::getline(fin, ln);
        std::stringstream ss(ln);
        ss >> K;
    }
    while (std::getline(fin,ln)) {
        int x, y, z;
        char comma;
        std::stringstream ss(ln);
        ss >> x >> comma >> y >> comma >> z;
        boxes.push_back({x,y,z});
    }
    std::sort(boxes.begin(), boxes.end());

    std::println("Read {} boxes",boxes.size());
    std::println("Part 1: {}",part1(K, boxes));
    std::println("Part 2: {}",part2(boxes));
    return 0;
}