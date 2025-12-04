
default: run02

run01: day01
	@echo "# Sample"
	@./day01 < sample01.txt
	@echo "# Real"
	@./day01 < input01.txt

day01: day01.cxx input01.txt sample01.txt
	g++ -std=c++23 day01.cxx -o day01

run02: day02
	@echo "# Sample"
	@./day02 < sample02.txt
	@echo "# Real"
	@./day02 < input02.txt

day02: day02.cxx input02.txt sample02.txt
	g++ -std=c++23 day02.cxx -o day02
