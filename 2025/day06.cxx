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

using std::cout;
using std::endl;
using std::list;
using std::pair;
using std::string;
using std::vector;

int64_t part1() {
    vector<vector<int64_t>> operands;

    string ln;
    bool first_line = true;
    while (std::getline(std::cin,ln)) {
        if (ln.find_first_of("+*") != std::string::npos) break;
        std::stringstream ss(ln);
        int64_t val;
        if (first_line) {
            while (ss >> val) {
                vector<int64_t> ops = {val};
                operands.push_back(ops);
            }
        } else {
            auto it = operands.begin();
            while (ss >> val) {
                it->push_back(val);
                it++;
            }
        }
        first_line = false;
    }

    int64_t part1 = 0; 
    std::stringstream ss(ln);
    string op;
    auto it = operands.begin();
    while (ss >> op) {
        if (op == "+") {
            part1 += std::accumulate(it->begin(), it->end(), int64_t(0), std::plus<int64_t>{});
        } else if (op == "*") {
            part1 += std::accumulate(it->begin(), it->end(), int64_t(1), std::multiplies<int64_t>{});
        }
        it++;
    }
    return part1;
}

int64_t part2()
{
    int64_t part2 = 0;

    vector<string> lines;
    string ln;
    while (std::getline(std::cin, ln))
    {
        lines.push_back(ln);
    }
    auto ops = lines.back();
    lines.pop_back();
    auto op_ptr = ops.end();
    --op_ptr;

    // get point to last entry in string
    vector<string::iterator> pointers;
    for (auto it = lines.begin(); it != lines.end(); it++)
    {
        auto end = it->end();
        --end;
        pointers.push_back(end);
    }

    // scan from the back to the front
    do {
        // build operands
        vector<int64_t> operands;
        while (*op_ptr == ' ') {
            --op_ptr;
        }
        // apply op to operands
        cout << *op_ptr;
        switch (*op_ptr)
        {
        case '*':
            break;
        case '+':
            break;
        }
        if (op_ptr != ops.begin()) --op_ptr;
    } while (op_ptr != ops.begin());
    cout << endl;

    return part2;
}

int main()
{
    // cout << "Part 1: " << part1() << endl;
    cout << "Part 2: " << part2() << endl;
    return 0;
}