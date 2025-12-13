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

struct Point {
    int x, y;
    int64_t area(const Point &b) const {
        int64_t w = int64_t(b.x) - int64_t(x);
        if (w < 0) {
            w *= -1;
        }
        w++;
        int64_t h = int64_t(b.y) - int64_t(y);
        if (h < 0) {
            h *= -1;
        }
        h++;
        return w*h;
    }
};

auto part1(const vector<Point> &points)
{
    int64_t ans = 0;
    for (auto p1 = points.begin(); p1 != points.end(); p1++) {
        auto p2 = p1;
        ++p2;
        for (; p2 != points.end(); p2++) {
            int64_t candidate = p1->area(*p2);
            if (candidate > ans) {
                ans = candidate;
            }
        }
    }
    return ans;
}

auto part2(const vector<Point> &points) {
    int64_t ans = 0;
    return ans;
}

int main(int argc, char *argv[])
{
    std::ifstream fin(argv[1]);
    string ln;
    vector<Point> points;
    while (std::getline(fin,ln)) {
        int x, y;
        char comma;
        std::stringstream ss(ln);
        ss >> x >> comma >> y;
        points.push_back({x,y});
    }
    std::println("Read {} points",points.size());
    std::println("Part 1: {}",part1(points));
    std::println("Part 2: {}",part2(points));
    return 0;
}