#include <sstream>
#include <string>
#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
#include <ranges>
#include <functional>

using std::cout;
using std::endl;
using std::pair;
using std::string;
using std::vector;

typedef int64_t ID;
typedef std::pair<int64_t, int64_t> Range;
auto cmp = [](const Range &r1, const Range &r2)
{
    return r1.second < r2.first || r1.first < r2.first;
};
auto contains = [](const Range &range, ID id)
{
    return range.first <= id && id <= range.second;
};
auto overlaps = [](const Range &r1, const Range &r2)
{
    return contains(r1, r2.first) || contains(r1, r1.second);
};

int main()
{
    // nothing clever for now, just keep a vector of the ranges
    vector<Range> ranges;
    ID part1 = 0;
    string line;
    while (std::getline(std::cin, line))
    {
        if (line.empty())
            break;

        std::stringstream ss(line);
        ID begin, end;
        char dash;
        ss >> begin >> dash >> end;
        Range r(begin, end);
        auto it = std::lower_bound(ranges.begin(), ranges.end(), r, cmp);
        if (it != ranges.end() && overlaps(*it, r))
        {
            cout << std::format("Combining [{},{}] and [{},{}]",r.first,r.second,it->first,it->second) << endl;
            (*it).first = std::min(it->first, r.first);
            (*it).second = std::max(it->second, r.second);
        }
        else
        {
            ranges.insert(it, r);
        }
    }

    auto reducer = [&ranges](ID id)
    {
        auto it = std::find_if(ranges.begin(), ranges.end(), [id](const Range &r)
                               { return contains(r, id); });
        return it != ranges.end();
    };

    // we got the first one above
    ID id;
    while (std::cin >> id)
    {
        if (reducer(id))
            part1++;
    };

    std::for_each(ranges.begin(), ranges.end(), [](const Range &r){
        cout << std::format("[{},{}]", r.first, r.second) << endl;
    });
    ID part2 = std::accumulate(ranges.begin(), ranges.end(), ID(0), [](ID acc, const Range &r) {
        return acc + r.second - r.first + 1;
    });

    cout << "Part 1: " << part1 << endl;
    cout << "Part 2: " << part2 << endl;
    return 0;
}
