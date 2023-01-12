# `GolangDojo`

## `Why Go was created? & What differentiates it from other languages?`

- มีการใช้ multicore processors กับ cloud infrastructure เป็น 100,1000 server ในการ deploy app
- ทำให้ infrastructure สามารถ scalable,distributed ได้มากขึ้น
- มีความ dynamic,capacity ที่เยอะขึ้น
- programming language ส่วนใหญ่ไม่สามารถใช้ประโยชน์จาก infrastructure นี้ได้อย่างเต็มที่
- execute 1 task infrastructure improvements จะทำการรัน `parallel task` ขึ้นมาที่สามารถทำอะไรหลายอย่างได้ในครั้งเดียวพร้อมกัน ทำให้ app เร็วขึ้น แล้วทำให้ user ใช้งานได้อย่างลื่นไหล เช่น การที่เราใช้ google drive เราสามารถ download file, upload file พร้อมกัน หรือในขณะที่เรากำลังทำ upload,download file อยู่เราก็สามารถ navigate ไปหน้าอื่นได้ โดยที่ทุกอย่างยังทำงานพร้อมกัน`(parallel)`ไปไม่ติดขัดหรือเกิดปัญหา `(concept multi-threading)`
  ![parallel](/img/parallel.PNG)
- การทำ `multi-threading` ก็จะมีปัญหาที่ต้อง handle คือ `concurrency(การทำงานหลายๆอย่างพร้อมกันในครั้งเดียว) `ที่เกิดขึ้นของการรัน `parallel task` ที่มีการ update data เดียวกัน`(access shared data)` เช่น app จองตั๋ว มี user 2 คนมาจองตั๋วใบสุดท้ายที่เหลืออยู่พร้อมกัน โดยที่ต้องไม่เกิดการจองเบิ้ล ซึ่งในส่วนนี้ dev ต้องเป็นคนเขียน code prevent conflict พวกนี้เอง
- programming language ส่วนใหญ่ทำได้ แต่ code จะมีความ complex สูง ยากต่อการจัดการ และราคาแพงแล้วก็ช้า
- Go ถูกออกแบบมาเพื่อรันบน multiple cores และ support concurrency
- Go ถูกและจัดการง่ายกว่า

## `Characteristic of Go`

- simple and readable
- fast build time, start up ,run
- require fewer resources
- efficiency and safety
- Go use for server-side or back-end side

## `Init Go Mudule`

```
go mod init <module path>
```

## `Run Go`

```
go run <name go file>
go run . => run ทุกไฟล์ที่อยู่ใน directory นี้
```

## `File Structure Go`

1. package
   - ทุก Go file จะต้องมี package
   - first statement ของ Go file ก็คือ package
2. main function
   - เป็น entrypoint ของ Go app
   - เพื่อบอกให้ Go รู้ว่าจะต้อง start execute ที่ไหน
3. import built-in package
   - Go program จะมี built-in package ที่จำเป็นมาให้เรียบร้อย

`Note :` https://pkg.go.dev/std
