# goperf
A go-based cross-platform CPU multicore performance evaluation tool
## Installation
``git clone https://github.com/lincvic/goperf``

``cd goperf``

``go install``

``go build``

## How to use
#### Parameters:
* -msg : The size (number) of message send to goroutines, **default is 100**.
> In goperf, each message handled by one goroutines.
* -workers : This parameter indicated how many time will a goroutine calculate FACT when receiving a message, **default is 10**.
* -fact : The number that will be calculated with factorial, **default is 1234**.
* -buffer : The channel buffer size of message, **default is 1**
* -time : This parameter decide how many times the test will run, **default is 1**
* -o : Output csv file path, **default is "result.csv"**

## Example
On Azure ARM VM Standard D4ps v5 (4 vcpus, 16 GiB memory)

``go run . -workers 100 -buffer 1 -msg 100 -fact 1234567 -time 5``

Output :

> goperf MultiCore CPU Test starting ...
>
>CPU Cores 8, Message Size 100, 100 workers inside each goroutine, Channel Buffer Size 1, Test will run 5 times
>
>Testing Round 1/5 :
> 
>Total time is 21.513752 s
> 
>8 CPU cores process 4.648190 msg/s with 100 goroutines
>
>Testing Round 2/5 :
> 
>Total time is 20.947001 s
> 
>8 CPU cores process 4.773953 msg/s with 100 goroutines
> 
>Testing Round 3/5 :
> 
>Total time is 20.935914 s
> 
>8 CPU cores process 4.776481 msg/s with 100 goroutines
> 
>Testing Round 4/5 :
> 
>Total time is 21.549876 s
> 
>8 CPU cores process 4.640398 msg/s with 100 goroutines
> 
>Testing Round 5/5 :
> 
>Total time is 21.912863 s
> 
>8 CPU cores process 4.563530 msg/s with 100 goroutines
> 
>Testing complete, Saving result into result.csv
## Things to Know
* goperf **Always implemented every CPU core**
