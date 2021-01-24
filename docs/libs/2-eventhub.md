# Simple Event Bus

EventBus in pub/sub pattern
![img](https://miro.medium.com/max/1022/1*6jeHWE0f2Mgd2CWJLTAZjg.png)

The traditional approach to implementing an event bus involves using callbacks. Subscribers usually implement an interface and then the event bus propagates data via the interface.

With go’s concurrency model we know that channels can be used in most places to replace callbacks. In this article, we focus on how to use channels for implementing an event bus.
topic based events

