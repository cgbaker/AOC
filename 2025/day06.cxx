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

int64_t getOperand(vector<string::iterator> &pointers)
{
    int64_t operand = 0;
    for (auto p = pointers.begin(); p != pointers.end(); ++p)
    {
        if (**p != ' ')
        {
            operand = 10 * operand + (**p - '0');
        }
        --(*p);
    }
    return operand;
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
    bool done = false;
    do {
        // build operands
        vector<int64_t> operands;
        char op;
        do {
            operands.push_back(getOperand(pointers));
            op = *op_ptr;
            if (op_ptr == ops.begin()) {
                done = true;
            } else {
                --op_ptr;
            }
        } while (op == ' ');
        if (!done) {
            // skip the blank column
            --op_ptr;
            auto op = getOperand(pointers);
            if (op != 0) {
                EXIT_FAILURE;
            }
        }        
        // apply op to operands
        switch (op)
        {
        case '+':
            part2 += std::accumulate(operands.begin(), operands.end(), int64_t(0), std::plus<int64_t>{});
            break;
        case '*':
            part2 += std::accumulate(operands.begin(), operands.end(), int64_t(1), std::multiplies<int64_t>{});
            break;
        }
    } while (!done);

    return part2;
}

int main()
{
    // cout << "Part 1: " << part1() << endl;
    cout << "Part 2: " << part2() << endl;
    return 0;
}