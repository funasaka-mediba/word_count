# word-count
ファイルの単語をカウントする。
io.Readerの勉強用。

# Example
```
# Buildせず使用
$ go run main.go mody.txt
"mody.txt": 215822 words

# Buildして使用
$ go build main.go
$ ./main mody.txt
"mody.txt": 215822 words
```

# 備考
```
# 速度検証した
# bufioを使用しなかったコードに比べると格段に速くなった。

# before
$ go build main.go
$ time ./main moby.txt
"moby.txt": 215822 words
./main moby.txt  0.61s user 0.76s system 97% cpu 1.404 total

# after
$ time ./main -w moby.txt
  215822 moby.txt
./main mody.txt  0.02s user 0.01s system 82% cpu 0.033 total
```
