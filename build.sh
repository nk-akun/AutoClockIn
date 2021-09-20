if [-d "output"]; then
    rm -r output
fi

if [-d "auto_clock_in"]; then
    rm auto_clock_in
fi

go build -o auto_clock_in
mkdir output
if [-d "auto_clock_in"]; then
    mv auto_clock_in output
fi

if [-d "auto_clock_in"]; then
    rm auto_clock_in
else
    echo "编译失败!"
fi