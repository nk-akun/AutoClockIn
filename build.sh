if [ -d "output" ]; then
    rm -r output
fi

if [ -f "auto_clock_in" ]; then
    rm auto_clock_in
fi

go build -o auto_clock_in
mkdir output
if [ -f "auto_clock_in" ]; then
    mv auto_clock_in output
    echo "编译成功!"
else
    echo "编译失败!"
fi