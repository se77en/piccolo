Piccolo [![wercker status](https://app.wercker.com/status/c72691aa9e5d90dd78fad2c63f787905/s "wercker status")](https://app.wercker.com/project/bykey/c72691aa9e5d90dd78fad2c63f787905)  [![Build status](https://ci.appveyor.com/api/projects/status/85peroqvohwt71yd)](https://ci.appveyor.com/project/se77en/piccolo)
=======

Piccolo is a tool that run a command periodically. It can be used both as a command line utility and a library

![Piccolo](http://fc08.deviantart.net/fs70/i/2012/042/8/2/majin_piccolo_by_brolyeuphyfusion9500-d4pcxl5.png)

## Installation

~~~
$ go get github.com/se77en/piccolo
~~~

### For Command Line

~~~
$ go install piccolo.go
~~~

### For Application

~~~
import github.com/se77en/piccolo/piccolo
~~~

## Usage

### For Command Line

~~~
Usage: piccolo <interval> <command>
~~~

 Examples

~~~
piccolo 700ms echo hello
~~~

The time can be such as "ns", "us", "ms", "s", "m", "h".


### For Application

~~~
piccolo.AddTimingFunc("test", -1, func(){fmt.Println("do something")})
piccolo.StartTiming(7 * time.Second)
~~~


# License

MIT
