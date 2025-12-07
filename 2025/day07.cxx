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

using std::cout;
using std::endl;
using std::list;
using std::pair;
using std::string;
using std::vector;
using std::map;

const char START = 'S';
const char SPLITTER = '^';

int64_t part1(std::ifstream &fin) {
    int part1 = 0;
    string ln;
    std::getline(fin, ln);
    const int LEN = ln.length();
    std::vector<bool> slots(LEN);
    slots[ln.find_first_of(START)] = true;
    while (std::getline(fin, ln)) {
        for (int i=0; i<LEN; i++) {
            if (ln[i] == SPLITTER && slots[i]) {
                part1++;
                slots[i] = false;
                if (i < LEN-1) slots[i+1] = true;
                if (i > 0) slots[i-1] = true;
            }
        }
    }
    return part1;
}

map<pair<int,int>, int64_t> cache;

int64_t DFS(int pos, int depth, vector<string> lines) {
    if (pos < 0 || pos >= lines[0].length()) return 0;
    if (depth >= lines.size()) return 1;

    auto it = cache.find({pos, depth});
    if (it != cache.end())
    {
        return it->second;
    }

    string ln = lines[depth];
    int64_t ans;
    if (ln[pos] == SPLITTER) {
        ans = DFS(pos-1,depth+1,lines) + DFS(pos+1,depth+1,lines);
    } else {
        ans = DFS(pos, depth+1, lines);
    }
    cache[{pos,depth}] = ans;
    return ans;
}

int64_t part2(std::ifstream &fin)
{
    int64_t part2 = 0;
    string ln;
    vector<string> lines;
    std::getline(fin,ln);
    int pos = ln.find_first_of(START);
    while (std::getline(fin, ln)) {
        lines.push_back(ln);
    } 
    return DFS(pos, 0, lines);
}

int main(int argc, char *argv[])
{
    std::ifstream fin(argv[1]);
    // std::ifstream fin("input07.txt");
    // cout << "Part 1: " << part1(fin) << endl;
    cout << "Part 2: " << part2(fin) << endl;
    return 0;
}