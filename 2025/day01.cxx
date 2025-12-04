#include <fstream>
#include <string>
#include <functional>
#include <iostream>

int parse_rotation(const std::string &line) {
    int num = std::atoi(line.substr(1).data());
    if (line[0] == 'L') {
        return -num;
    } else {
        return num;
    }
}

int main() {
    std::string line;

    auto reducer = [](const auto& acc, const std::string& input) {
        auto [pos, zeros, passes] = acc;
        auto rot = parse_rotation(input);

        auto oldPos = pos;
        pos += rot;
        if (pos <= 0) {
            // we were positive before, so we've passed zero at least once
            if (oldPos > 0) {
                passes++;
            }
            passes += -(pos/100);
            pos = ((pos % 100) + 100) % 100;
        } else {
            passes += pos/100;
            pos = pos % 100;
        }

        if (0 == pos) zeros++;

        return std::make_tuple(pos, zeros, passes);
    };

    auto results = std::make_tuple<int,int,int>(50,0,0);
    while (std::getline(std::cin,line)) {
        results = reducer(results, line);
    }

    std::cout << "Part 1:  " << std::get<1>(results) << std::endl;
    std::cout << "Part 2:  " << std::get<2>(results) << std::endl;

    return 0;
}