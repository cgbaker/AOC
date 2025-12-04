#include <string>
#include <iostream>
#include <ranges>
#include <algorithm>

using std::string;
using std::cout;
using std::endl;

auto char_to_digit = [](char c) { return c - '0'; };

int64_t max_joltage(string bank, int num_batteries = 2) {
    int64_t joltage = 0;
    for (int battery_num = num_batteries; battery_num > 0; battery_num--) {
        auto max_val = std::ranges::max(
            bank.substr(0, bank.length() - battery_num + 1) | std::views::transform(char_to_digit));
        auto max_pos = bank.find_first_of(max_val + '0');
        joltage = joltage*10 + max_val;
        if (battery_num > 1) {
            bank = bank.substr(max_pos + 1);
        }
    }
    return joltage;
}

int main() {
    string bank;

    int64_t joltage2 = 0;
    int64_t joltage12 = 0;
    while (std::getline(std::cin, bank)) {
        joltage2 += max_joltage(bank);
        joltage12 += max_joltage(bank, 12);
    }

    cout << "Part 1: " << joltage2 << endl;
    cout << "Part 2: " << joltage12 << endl;

    return 0;
}