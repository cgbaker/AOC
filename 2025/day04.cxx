#include <string>
#include <iostream>
#include <ranges>
#include <algorithm>

using std::string;
using std::cout;
using std::endl;

auto readPadded() {
    string ln;
    if (std::getline(std::cin, ln))
    {
        return "." + ln + ".";
    }
    return string("");
}

bool canAccess(const string &top, const string &middle, const string &bottom) {
    if (middle[1] != '@') return false;
    int total = std::ranges::count(top.substr(0, 3), '@') + std::ranges::count(middle.substr(0, 3), '@') + std::ranges::count(bottom.substr(0, 3), '@');
    return total < 5;
}

int main() {
    string l1, l2, l3;
    bool eof = false;

    int results = 0;
    l2 = readPadded();
    l1.assign(l2.length(),'.');
    do {
        l3 = readPadded();
        if (l3.empty()) {
            l3.assign(l2.length(), '.');
            eof = true;
        } 
        string top = l1,
                middle = l2,
                bottom = l3;
        while (middle.length() >= 3) {
            if (canAccess(top,middle,bottom)) {
                results++;
            }
            top = top.substr(1);
            middle = middle.substr(1);
            bottom = bottom.substr(1);
        }
        l1 = l2;
        l2 = l3;
    } while (!eof);

    cout << "Part 1: " << results << endl;
    // cout << "Part 2: " << joltage12 << endl;

    return 0;
}