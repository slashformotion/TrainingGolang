mkdir $1
cd $1
echo "//$1.o"> $1.go
echo "package main" >> $1.go
echo "import(" >> $1.go
echo "	\"fmt\"" >> $1.go
echo ")" >> $1.go
echo "" >> $1.go
echo "" >> $1.go
echo "func main() {" >> $1.go
echo "	fmt.Println(\"--- Training.go => $1 ---\")" >> $1.go
echo "}"  >> $1.go

cd ..

echo "Section \"$1\" added successfully"