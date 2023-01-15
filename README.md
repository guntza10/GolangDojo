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

## `Goroutine - Concurrency`

- ปกติเวลาเรา execute app มันจะรันเป็น single thread ตั้งแต่ต้นจนจบตามลำดับ
- ถ้ามี process ใดที่ใช้เวลาในการ execute นาน จะทำให้ thread ติด block
- ทำให้ code process ในบรรทัดต่อไปต้องรอ process ก่อนหน้าให้ execute เสร็จก่อน
- เราจะใช้ Goroutine(keyword go) เข้ามาช่วย
- โดยที่ Goroutine จะแยก thread ของ process ที่ใช้เวลานานออกมา execute ต่างหาก (run in background)

  ![goroutine_1](/img/goroutine_1.PNG)

  ## `Synchronizing the Goroutines`

- ปัญหาคือ main thread execute เสร็จก่อนที่ตัว Goroutine จะเริ่มและ execute code
- main thread ไม่ได้รอ Goroutine execute
- พอ main thread execute เสร็จมัน app ก็จบการทำงานพร้อมกับ terminate ทุก thread ทิ้ง
- เราใช้ `WaitGroup` เข้ามาช่วย เพื่อบอกให้ตัว main thread รอ Goroutine execute

  ![goroutine_2](/img/goroutine_2.PNG)

### WaitGroup => มี 3 function

- Add(no) => จะเป็นการบอก main thread ว่ามีจำนวน thread เท่าไรที่มันต้องรอและต้อง execute ก่อนที่จะ สร้าง new thread
- Wait() => จำเป็นต้อง execute ไว้ตอนสุดท้ายของ main thread มันจะรอจนกว่าทุก thread จะ done (WaitGroup list เป็น 0 ไม่เหลือ thread)
  จะทำงานก่อนที่ app จะ exit terminate ทุกอย่าง
- Done() => จำเป็นต้อง execute ไว้ตอนสุดท้ายของ separated thread(Goroutine thread) เมื่อถูก call มันจะทำการลดจำนวน WaitGroup ลง 1
  (remove thread ที่อยู่ใน WaitGroup list ออกไป 1)

  ![goroutine_3](/img/goroutine_3.PNG)
  ![goroutine_4](/img/goroutine_4.PNG)

## `Comparison to other languages`

### other languages

- ภาษาอื่นก็สามารถทำ concurrency ได้ แต่ code complex มาก

  ![compare_1](/img/compare_1.PNG)

- ใช้ค่าใช้จ่ายเยอะกว่า
- การสร้าง thread เยอะๆแพง และใช้เวลาเยอะในตอนเริ่ม
- ต้องใช้ memory,space ในการจัดการเยอะมาก

### Golang

- มี Goroutine(`Operation System Thread`) ในการจัดการ thread
- ถูกและ lightweight (ใช้ memory,space น้อย)

![compare_2](/img/compare_2.PNG)

![compare_3](/img/compare_3.PNG)
