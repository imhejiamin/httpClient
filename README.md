# httpClient-异步http原理

HTTP Reactive Client 是一个典型的消息（事件）驱动的案例。

参考：

https://jersey.github.io/documentation/latest/rx-client.html#d0e5556

http://blog.csdn.net/pmlpml/article/details/78850661


练习要求：

1、依据文档图6-1，用中文描述 Reactive 动机。

2、使用 go HTTPClient 实现图 6-2 的 Naive Approach。

3、为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3。

4、对比两种实现，用数据说明 go 异步 REST 服务协作的优势。

5、思考： 是否存在一般性的解决方案？

result：

![result](https://github.com/imhejiamin/httpClient/blob/master/pic/result.png)

