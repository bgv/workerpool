## workerpool
Simple [Go](http://golang.org) routine pool inspired by: http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url]

## Why another one?
Because I needed it and was good exercise for me, oh ... and was fun as well :)

## Get it

```bash
go get github.com/bgv/workerpool
```

## Use it

```go
import (
    "fmt"
    
    "github.com/bgv/workerpool"
)

...

// Create 10 workers and a queue with size 50
pool := workerpool.New(10, 50)

// Create and submit 100 jobs
for i := 0; i < 100; i++ {
    count := i

    job := func() {
    	fmt.Printf("I am job number %d!\n", count)
    	pool.JobDone()
    }

    pool.AddJob(job)
}

// Wait for all jobs and then stop the worker pool
pool.Stop()

```

## License

MIT License

Copyright (c) 2016 Boris Varbanov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/badge/license-MIT-blue.svg
[Build-Status-Url]: https://travis-ci.org/bgv/workerpool
[Build-Status-Image]: https://api.travis-ci.org/bgv/workerpool.svg?branch=master
[ReportCard-Url]: https://goreportcard.com/report/github.com/bgv/workerpool
[ReportCard-Image]: https://goreportcard.com/badge/github.com/bgv/workerpool
[GoDoc-Url]: https://godoc.org/github.com/bgv/workerpool
[GoDoc-Image]: https://godoc.org/github.com/bgv/workerpool?status.svg
