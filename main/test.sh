#!/bin/bash

#* test file for multiple diff cases

go run . --output test00.txt banana standard 

echo

go run . --output=test00.txt "First\nTest" shadow

cat test00.txt

go run . --output=test01.txt "hello" standard

cat test01.txt

go run . --output=test02.txt "123 -> #$%" standard

cat test02.txt


go run . --output=test03.txt "432 -> #$%&@" shadow

cat test03.txt

go run . --output=test04.txt "There" shadow

cat test04.txt

go run . --output=test05.txt "123 -> \"#$%@" thinkertoy

cat test05.txt

go run . --output=test06.txt "2 you" thinkertoy

cat test06.txt

go run . --output=test07.txt 'Testing long output!' standard

cat test07.txt


go run . "{|}~"

echo 

go run . "1a\"#FdwHywR&/()="

echo


go run . ":;<=>?@"

echo

go run . "[\]^_ 'a"

echo

go run . '\!" #$%&'"'"'()*+,-./'

echo


go run . "RGB"

echo

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

echo

go run . "abcdefghijklmnopqrstuvwxyz"