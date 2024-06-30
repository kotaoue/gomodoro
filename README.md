# gomodoro
Pomodoro Timer

## Links
- [効果音ラボ](https://soundeffect-lab.info/)

## Note
- Q. `fatal error: 'KHR/khrplatform.h' file not found` って怒られたらどうすれば良い？
  - A. M1 Mac 使ってたら、`go get fyne.io/fyne/v2` で v2 を使うと行ける気がする 

- Q. Progress Bar付けないの？
 - A. Progress Barはゴルーチンの中で動かす必要あるけどタイマーで使ってるからだめみたい。いつかチャレンジする
