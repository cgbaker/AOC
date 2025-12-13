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
#include <charconv>

using std::cout;
using std::endl;
using std::list;
using std::pair;
using std::string;
using std::vector;
using std::map;
using std::set;
using std::tuple;

struct Machine {
    int lights;
    vector<vector<int>> buttons;
    vector<int> joltages;
};

struct Node {
    int value;
    int depth;
};

int push_button(int value, const vector<int> &button) {
    int flips = 0;
    for (auto l : button) {
        flips += (1 << l);
    }
    return value ^ flips;
}

int fewest_presses(const Machine &m) {
    set<int> visited;
    list<Node> tovisit = {Node{0,0}};
    while (!tovisit.empty()) {
        Node n = tovisit.front();
        // std::println("looking at {} after {} presses",n.value,n.depth);
        tovisit.pop_front();
        visited.insert(n.value);
        for (auto button : m.buttons) {
            int new_value = push_button(n.value, button);
            if (m.lights == new_value) {
                return n.depth+1;
            }
            if (!visited.contains(new_value)) {
                tovisit.push_back({new_value,n.depth+1});
            }
        }
   }

    return 0;
}

auto part1(const vector<Machine> &machines)
{
    int64_t ans = 0;
    for (auto m : machines) {
        ans += fewest_presses(m);
    }
    return ans;
}

auto part2(const vector<Machine> &machines) {
    int64_t ans = 0;
    return ans;
}

auto parse_lights(const string &s) {
    int pos = 0;
    int value = 0;
    for (char c : s.substr(1, s.length()-2)) {
        if (c == '#') {
            value += (1 << pos);
        }
        ++pos;
    }
    return value;
}


std::vector<int> split_to_ints(std::string_view s)
{
    std::vector<int> result;
    for (auto part : s.substr(1,s.length()-2) | std::views::split(',')) {
        int value{};
        auto* begin = &*part.begin();
        auto* end   = begin + std::ranges::distance(part);
        if (std::from_chars(begin, end, value).ec == std::errc()) {
            result.push_back(value);
        }
    }
    return result;
}

int main(int argc, char *argv[])
{
    std::ifstream fin(argv[1]);
    string ln;
    vector<Machine> machines;
    while (std::getline(fin,ln)) {
        std::stringstream ss(ln);
        string s;
        ss >> s;
        auto lights = parse_lights(s);
        vector<vector<int>> buttons;
        while (ss >> s) {
            if (s[0] == '(') {
                buttons.push_back(split_to_ints(s));
            } else {
                machines.push_back({lights,buttons,split_to_ints(s)});
                break;
            }
        }
    }
    std::println("Read {} machines",machines.size());
    std::println("Part 1: {}",part1(machines));
    std::println("Part 2: {}",part2(machines));
    return 0;
}