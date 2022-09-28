# goperf
A go-based cross-platform CPU multicore performance evaluation tool
## Installation
``git clone https://github.com/lincvic/goperf``

``go build``

## How to use
#### Parameters:
* -msg : The size (number) of message send to goroutines, **default is 100**.
> In goperf, each message handled by one goroutines.
* -workers : This parameter indicated how many time will a goroutine calculate FACT when receiving a message, **default is 10**.
* -fact : The number that will be calculated with factorial, **default is 1234**.
* -buffer : The channel buffer size of message, **default is 1**

## Example
On Azure ARM VM Standard D4ps v5 (4 vcpus, 16 GiB memory)

``go run . -workers 100 -buffer 1 -msg 100 -fact 1234567``

Output is :

> goperf MultiCore CPU Test starting ...
> 
> CPU Cores 4, Message Size 100, 100 workers inside each goroutine, Channel Buffer Size 1
> 
> Total time is 26740 ms
> 
> 4 CPU cores process 0.003740 message/ms with 100 goroutines

## Things to Know
* goperf **Always implemented every CPU core**