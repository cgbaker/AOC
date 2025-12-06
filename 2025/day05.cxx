#include <sstream>
#include <string>
#include <iostream>
#include <algorithm>
#include <list>
#include <numeric>
#include <ranges>
#include <functional>

using std::cout;
using std::endl;
using std::list;
using std::pair;
using std::string;

typedef int64_t ID;
typedef std::pair<int64_t, int64_t> Range;
auto contains = [](const Range &range, ID id)
{
    return range.first <= id && id <= range.second;
};
auto overlaps = [](const Range &r1, const Range &r2)
{
    return r1.first <= r2.second && r2.first <= r1.second;
};

int main()
{
    list<Range> ranges;
    string line;
    while (std::getline(std::cin, line))
    {
        if (line.empty())
            break;
        std::stringstream ss(line);
        ID begin, end;
        char dash;
        ss >> begin >> dash >> end;
        ranges.push_back(Range(begin, end));
    }

    ranges.sort([](const Range &r1, const Range &r2) { 
        return r1.first < r2.first || (r1.first == r2.first && r1.second < r2.second); 
    });

    auto iter = ranges.begin();
    auto next(iter);
    ++next;
    while (next != ranges.end())
    {
        if (overlaps(*iter, *next))
        {
            Range combined( std::min(iter->first, next->first), std::max(iter->second, next->second) );
            (*iter) = combined;
            next = ranges.erase(next);
        }
        else
        {
            ++iter;
            next = iter;
            ++next;
        }
    }

    auto finder = [&ranges](ID id)
    {
        auto it = std::find_if(ranges.begin(), ranges.end(), [id](const Range &r)
                               { return contains(r, id); });
        return it != ranges.end();
    };

    ID id;
    ID part1 = 0;
    while (std::cin >> id)
    {
        if (finder(id))
            part1++;
    };

    ID part2 = std::accumulate(ranges.begin(), ranges.end(), ID(0), [](ID acc, const Range &r)
                               { return acc + r.second - r.first + 1; });

    cout << "Part 1: " << part1 << endl;
    cout << "Part 2: " << part2 << endl;
    return 0;
}