### Advent of code 

#### Create contest setup 

```
# run from <>/contests/aoc
YEAR=2024
mkdir $YEAR
cd $YEAR
for i in $(seq -f "%02g" 1 25)
do
  echo "creating day$i"
  mkdir day$i
  cd day$i
  cp ../../samples/code.go day$i.go
  cp ../../samples/tst.go day$i\_test.go
  sed -i "" s/samples/day$i/g day$i.go
  sed -i "" s/samples/day$i/g day$i\_test.go   
  sed -i "" s/20XX/$YEAR/g day$i.go
  sed -i "" s/20XX/$YEAR/g day$i\_test.go 
  sed -i "" s/dayXX/day$i/g day$i.go 
  sed -i "" s/dayXX/day$i/g day$i\_test.go 
  touch input_small.txt 
  touch input_final.txt
  cd ..
done
```