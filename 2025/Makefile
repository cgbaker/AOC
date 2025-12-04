
default: run01

run01: day01
	@echo "# Sample"
	@./day01 < sample01.txt
	@echo "# Real"
	@./day01 < input01.txt

day01: day01.cxx input01.txt sample01.txt
	g++ -std=c++23 day01.cxx -o day01
