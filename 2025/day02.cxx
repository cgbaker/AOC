#include <string>
#include <iostream>

using std::string;

bool splitAndCompare(const string& s, int num_pieces) {
    if (s.length() % num_pieces != 0) {
        return false;
    }
    const auto chunk_size = s.length() / num_pieces;
    string last = "";
    for (size_t i = 0; i < s.length(); i += chunk_size)
    {
        auto piece = s.substr(i, chunk_size);
        if (last != "" && piece != last) {
            return false;
        }
        last = piece;
    }
    return true;
}


int64_t sum_invalid(int64_t begin, int64_t end, bool limitSplits = true) {
    int64_t sum = 0;
    for (int64_t i = begin; i <= end; i++) {
        const auto s = std::to_string(i);
        if (limitSplits) {
            if (splitAndCompare(s, 2)) {
                sum += i;
            }
        } else {
            for (int numPieces=2; numPieces <= s.length(); numPieces++) {
                if (splitAndCompare(s, numPieces)) {
                    sum += i;
                    break;
                }
            }
        }
    }
    return sum;
}

int main() {
    string line;

    int64_t part1 = 0, part2 = 0;

    int64_t begin, end;
    char dash, comma;
    while (std::cin >> begin >> dash >> end) {
        part1 += sum_invalid(begin,end);
        part2 += sum_invalid(begin,end, false);
        if (!(std::cin >> comma)) {
            break;
        }
    }

    std::cout << "Part 1: " << part1 << std::endl;
    std::cout << "Part 2: " << part2 << std::endl;

    return 0;
}