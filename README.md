# cmd_cache

> Cache the results of your command within the specified time
> 
> Currently only supports Linux
> 
> Don't use it in serious projects


# Example

> $ time ./main -c 'seq 1000000 |count' -t 10
>
> 500000500000
> 
> ./main -c 'seq 1000000 |count' -t 10  0.80s user 0.09s system 102% cpu 0.872 total
> 
> $ time ./main -c 'seq 1000000 |count' -t 10
>
> 500000500000
> 
> ./main -c 'seq 1000000 |count' -t 10  0.00s user 0.00s system 0% cpu 0.003 total
> 


