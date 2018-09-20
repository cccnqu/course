# C 語言

## C 語言的陷阱

資工系老師若使用 C/C++ 語言作為第一門語言時，會有幾個效應：

1. 很多學生會掛在陷阱上撐不過去
2. 但是撐過去的人，將會是《能夠學任何一門語言》的人

這就好像第一次當兵，就直接到真正的戰場去，沒有掛掉的人，就會變成實戰經驗豐富的士兵，不會再那麼容易掛點了 ....

問題是、有多少學生撐得過去，就是個未知數了 ....

* Facebook 討論 -- https://www.facebook.com/ccckmit/posts/10156016824741893

### 避免緩衝區溢位

檔案： [scanf_no_overflow.c](scanf_no_overflow.c)

```c
#include <stdio.h>

int main() {
  char input[5];
  scanf("%4s", input);
  printf("Your input : %s\n", input);
}
```

解析：如果將 %4s 改為 %s，就會有緩衝區溢位的問題 (輸入大於 4 個字時，會超過 input[5] 的大小，導致發生作業系統錯誤，或者覆蓋掉後續資料或程式碼。 (這種覆蓋使用得當，將有可能竊取你電腦的控制權，像是取得 super user 的權力)

執行結果

```
$ gcc scanf_no_overflow.c -o scanf_no_overflow
$ ./scanf_no_overflow
1234567890
Your input : 1234
```

### 連續 scanf

檔案： [scanf_loop.c](scanf_loop.c)

```c
#include <stdio.h>

#define SIZE 3

int main() {
  int a[SIZE];
  for (int i=0; i<SIZE; i++) {
    scanf("%d", &a[i]);    
    printf("%d\n", a[i]);    
  }
}
```

執行結果

```
$ gcc scanf_loop.c -std=c99 -o scanf_loop
$ ./scanf_loop
1 2 3
1
2
3
```

問題：為何輸入 1 + 空白之後，不會馬上呼叫 printf() 輸出 1 呢？

解答： 這個問題有點複雜，因為 scanf 只會在你打入換行字元之後，才會真正開始執行。

在還沒打入換行字元前，整個行程 process 會卡在《等待輸入》狀態。

因此才會在你打完換行後，才繼續執行。

我們可以在執行另一次 scanf_loop 觀察看看：

```
$ ./scanf_loop
1 2
1 
2
3
3
```

上述結果應該可以證明，只有換行打上去之後，才會真正驅動 scanf 執行。

