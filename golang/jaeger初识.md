## Jager是干啥的

​		分布式链路追踪系统，底层用golang实现，兼容opentracing标准。

## Jager怎么做的

​		一套完整的Jager追踪系统包括Jaeger-client、Jaeger-agent、Jaeger-collector、Database和Jaeger-query UI等基本组件，如下图架构图所示，Jaeger客户端支持多种语言，jaeger-agent与客户端进行数据交互，并把数据push到Jaeger-collector组件，Jaeger-collector将数据持久化到数据库，Jaeger-query是一个web服务，用于展示跟踪链路。以下为Jaeger容器化部署的基本流程: 分为测试环境和正式环境两种方式。

![img](https://img2018.cnblogs.com/common/826988/202001/826988-20200115112659501-1441836958.png)

## Jager咋用的

### 1、安装必要的包

```
"github.com/opentracing/opentracing-go"
"github.com/uber/jaeger-client-go"
"github.com/uber/jaeger-client-go/config"
```

### 2、**安装部署jaeger整套**

为了快速上手, 官方提供了”All in One”的docker镜像, 启动Jaeger服务只需要jaegertracing/all-in-one一个镜像源

查看docker版本

```
docker --version
```

搜索jager版本

```
docker search jaegertracing/all-in-one
```

拉取镜像(不加版本号默认是latest)

```
docker pull jaegertracing/all-in-one
```

起一个jager容器

```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one
```

查看容器是否启动

```
docker ps
```

![image-20210220171133904](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220171133904.png)

访问 http://localhost:16686/search查看是否访问成功

![image-20210220171303805](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220171303805.png)

### 3、jager demo上手试试吧伙计们

**编写demo**

初始化jaeger tracer的initJaeger方法，reporter中配置jaeger Agent的ip与端口，以便将tracer的信息发布到agent中。

配置LocalAgentHostPort参数为127.0.0.1:6381，6381接口是接受压缩格式的thrift协议数据。

采样率暂且设置为1。

```
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort:"127.0.0.1:6831",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
```

在main函数中创建调用InitJaeger，并创建一个root span，调用两个函数，分别表示调用两个分布式服务。

我们用ContextWithSpan来创建一个新的ctx，将span的信息与context关联，传到foo3中时，需要创建一个子span，父span是ctx中的span。

我们在foo3中调用StartSpanFromContext时，忽略了第二个参数，这是利用子span创建的新的context，当我们在foo3中再调用别的比如foo5时，我们应该使用新的context，而不是传入的ctx。

注意StartSpanFromContext会用到opentracing.SetGlobalTracer()来启动新的span，所以在main函数中需要调用。

```
func foo3(req string, ctx context.Context) (reply string){
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo3")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()
 
	println(req)
	//2.模拟处理耗时
	time.Sleep(time.Second/2)
	//3.返回reply
	reply = "foo3Reply"
	return
}
//跟foo3一样逻辑
func foo4(req string, ctx context.Context) (reply string){
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo4")
	defer func() {
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()
 
	println(req)
	time.Sleep(time.Second/2)
	reply = "foo4Reply"
	return
}
 
func main() {
	//初始化jager
	tracer, closer := initJaeger("jaeger-demo")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)//StartspanFromContext创建新span时会用到
 
	span := tracer.StartSpan("span_root")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	r1 := foo3("Hello foo3", ctx)
	r2 := foo4("Hello foo4", ctx)
	fmt.Println(r1, r2)
	span.Finish()
}
```

运行结果：

![image-20210221171331342](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210221171331342.png)



​		查看详情，可以看到这个响应的reply、request、ip、service-name、开始时间、结束时间、开销时间等信息。主要是这个demo写的比较简单，看不到太多信息。在实际工作时你会发现，你所处的是一个工作流上的一环，你所依赖的上下游也是整个环中的一部分罢了，尤其是你本身这一环中，各个微服务之间相互依赖调用的，比如前端访问你后端接口，先要去做um的权限校验，查看用户是否注册合法，若校验通过的话会携带token信息去到你的网关服务，网关服务再根据你的请求体找到你需要访问的服务，转发给实际的业务服务接口，真正的业务服务接口可能又有用户管理权限的鉴权，然后再去做数据库的crud，db记录下来之后再去调用他的下游服务。这还是一个比较简单的模型，若是在中间加入一些redis、es之类的、服务依赖再等一些的，是不是会更复杂。你的调用链越多，你接入链路追踪的必要性和收益就越大。

​		再次感慨技术是一个不断循环发展的螺旋状式现状体结构，向一个不断向上盘旋的树，在向上延伸的过程中会涌现出很多新鲜的事物，但新事物会带来技术创新、革命的同时也会导致一些问题，然后聪明的技术儿们又想出更多方式来弥补这种漏洞，不断的完善和完成。在一直打补丁的过程中，这棵树也越来越壮大和蓬勃。但是树再高、再壮也离不了它的根。其实技术层出不穷，但是核心还是那些基础，数据结构、计算机网络、操作系统、编译原理，很多东西你在仔细回头去看，会发现很久之前已经被提出来了，只不过因为当时的某些原因被搁置了，现在又突然再次被发展了起来。就像机器学习，也不是什么新鲜概念了。所以说啊，学无止境。

​		当然，回道正题，恭喜你又解锁了一个新技能--jaeger的链路追踪。