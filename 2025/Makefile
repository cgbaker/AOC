# Find all dayNN.cxx files, extract NN, sort numerically, take the largest
LATEST := $(shell ls day*.cxx 2>/dev/null | sed -E 's/day([0-9]+)\.cxx/\1/' | sort -n | tail -1)

# Default target: run the latest day
default: run$(LATEST)

# Pattern rule: build dayNN from dayNN.cxx
day%: day%.cxx 
	@echo Compiling $@...
	@g++ -std=c++23 $< -o $@

# Pattern rule: run sample + real inputs
run%: day% input%.txt sample%.txt
	@echo "# Sample"
	@./$< < sample$*.txt
	@echo "# Real"
	@./$< < input$*.txt
