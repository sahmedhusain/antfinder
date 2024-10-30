#!/bin/zsh
clear
echo -e "v--------------example00.txt--------------v\n"
time go run main.go examples/example00.txt
echo -e "\nv--------------example01.txt--------------v\n"
time go run main.go examples/example01.txt
echo -e "\nv--------------example02.txt--------------v\n"
time go run main.go examples/example02.txt
echo -e "\nv--------------example03.txt--------------v\n"
time go run main.go examples/example03.txt
echo -e "\nv--------------example04.txt--------------v\n"
time go run main.go examples/example04.txt
echo -e "\nv-------------example05.txt-------------v\n"
time go run main.go examples/example05.txt
echo -e "\nv-------------example06.txt-------------v\n"
time go run main.go examples/example06.txt
echo -e "\nv-------------example07.txt-------------v\n"
time go run main.go examples/example07.txt
echo -e "\nv-----------badexample00.txt------------v\n"
time go run main.go examples/badexample00.txt
echo -e "\nv-----------badexample01.txt------------v\n"
time go run main.go examples/badexample01.txt