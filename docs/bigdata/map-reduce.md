# Map Reduce

1. MapReduce library会把输入文件划分成多个16到64MB大小的分片（大小可以通过参数调节），
   然后在一组机器上启动程序。
2. 其中比较特殊的程序是master，剩下的由master分配任务的程序叫worker。
   总共有M个map任务和R个reduce任务需要分配，master会选取空闲的worker，
   然后分配一个map任务或者reduce任务。
    
3. 处理map任务的worker会从输入分片读入数据，解析出输入数据的K/V对，
   然后传递给Map函数，生成的K/V中间结果会缓存在内存中。 
   
4. map任务的中间结果会被周期性地写入到磁盘中，以partition函数来分成R个部分。
   R个部分的磁盘地址会推送到master，然后由它转发给响应的reduce worker。 
   
5. 当reduce worker接收到master发送的地址信息时，它会通过RPC来向map worker读取对应的数据。
   当reduce worker读取到了所有的数据，它先按照key来排序，方便聚合操作。
   
6. reduce worker遍历排序好的中间结果，对于相同的key，把其所有数据传入到Reduce函数进行处理，
   生成最终的结果会被追加到结果文件中。 当所有的map和reduce任务都完成时，master会唤醒用户程序，
   然后返回到用户程序空间执行用户代码。
   

## Hadoop

## Hbase

## Flink

## Spark