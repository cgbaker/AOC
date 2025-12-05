#include <string>
#include <iostream>
#include <ranges>
#include <algorithm>
#include <list>
#include <span>

using std::cout;
using std::endl;
using std::list;
using std::span;
using std::string;

const char EMPTY = '.';
const char FULL = '@';
const char CANMOVE = 'x';

auto readPadded()
{
    string ln;
    if (std::getline(std::cin, ln))
    {
        return EMPTY + ln + EMPTY;
    }
    return string("");
}

void printMap(list<string> map) {
    for (auto it = map.begin(); it != map.end(); it++)
    {
        cout << *it << endl;
    }
}

bool canAccess(span<char> &prev, span<char> &cur, span<char> &next)
{
    if (cur[1] == EMPTY) return false;

    int total = std::ranges::count(prev.subspan(0, 3), EMPTY) + std::ranges::count(cur.subspan(0, 3), EMPTY) + std::ranges::count(next.subspan(0, 3), EMPTY);
    return total >= 5;
}

int main()
{
    // read the input, padding the left and right border
    list<string> lines;
    do
    {
        string ln = readPadded();
        if (ln.length() == 0)
        {
            break;
        }
        lines.push_back(ln);
    } while (true);
    const auto width = lines.begin()->length();
    // pad the top and bottom border
    lines.push_front(string(width, EMPTY));
    lines.push_back(string(width, EMPTY));

    int round = 0;
    int part1 = 0, part2 = 0;
    bool movedSomething;
    do
    {
        round++;
        // reset/init
        movedSomething = false;
        
        // clear movable spots
        for (auto it=lines.begin(); it != lines.end(); it++) {
            std::replace(it->begin(), it->end(), CANMOVE, EMPTY);
        }

        auto it = lines.begin();
        auto l1 = it++;
        auto l2 = it++;
        auto l3 = it++;

        // check all rows
        while (l3 != lines.end())
        {
            auto prev = span<char>(l1->data(), l1->size());
            auto cur = span<char>(l2->data(), l2->size());
            auto next = span<char>(l3->data(), l3->size());
            // check all columns
            while (cur.size() >= 3)
            {
                if (canAccess(prev, cur, next))
                {
                    cur[1] = CANMOVE;
                    if (round == 1) part1++;
                    part2++;
                    movedSomething = true;
                }
                prev = prev.subspan(1);
                cur = cur.subspan(1);
                next = next.subspan(1);
            }
            // advance to next row
            l1 = l2;
            l2 = l3;
            l3++;
        }
    } while (movedSomething);

    cout << "Part 1: " << part1 << endl;
    cout << "Part 2: " << part2 << endl;

    return 0;
}