#!/bin/bash

# 初始化选项
verbose=0
show_help() {
    echo "Usage: $0 [-v]"
    echo "Run Go tests matching a keyword pattern"
    echo "Options:"
    echo "  -v    Show verbose test output"
    exit 0
}

# 解析命令行选项
while getopts "vh" opt; do
    case $opt in
        v) verbose=1 ;;
        h) show_help ;;
        *) exit 1 ;;
    esac
done

# 读取用户输入
read -p "Enter test rpc: " keyword

# 目标目录
target_dir="result/code/Grok-3/"

# 查找匹配的测试文件
test_files=($(find "$target_dir" -type f -name "*_${keyword}_*test.go"))

# 检查文件是否存在
if [ ${#test_files[@]} -eq 0 ]; then
    echo "Error: No test files found containing '$keyword'"
    exit 1
fi

# 初始化计数器
total=0
success=0
fail=0

# 运行测试
for file in "${test_files[@]}"; do
    ((total++))

    if [ $verbose -eq 1 ]; then
        echo "--------------------------------------------------"
        echo "Running test file: $file"
        go test -v "$file"
    else
        echo -n "Testing ${file}... "
        go test "$file" &> /dev/null
    fi

    exit_status=$?

    # 显示简洁结果
    if [ $exit_status -eq 0 ]; then
        [ $verbose -eq 0 ] && echo "OK"
        ((success++))
    else
        [ $verbose -eq 0 ] && echo "FAIL"
        ((fail++))
    fi
done

# 计算结果
if [ $total -gt 0 ]; then
    success_rate=$(echo "scale=2; $success * 100 / $total" | bc)
    fail_rate=$(echo "scale=2; $fail * 100 / $total" | bc)
else
    success_rate=0
    fail_rate=0
fi

# 输出统计
echo "=================================================="
echo "Summary:"
echo "Total files: $total"
echo "Success:     $success (${success_rate}%)"
echo "Failed:      $fail (${fail_rate}%)"