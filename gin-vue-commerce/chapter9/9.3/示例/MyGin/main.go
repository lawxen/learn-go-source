package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	// 1.初始化limiter，每秒10个令牌，令牌桶容量为20。
	// 参数r以Limit表示，Limit是float64，代表每秒向桶中产生令牌数。
	// 参数b以int表示，代表桶的容量大小，也是最大并发数
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 20)

	// 2.获取指定时间内指定数量的令牌，获取成功则返回true
	limiter.AllowN(time.Now(), 2)
	// 2.获取1个令牌，获取到返回true，否则false
	// Allow()内部调用是AllowN()
	bo := limiter.Allow()
	if bo {
		fmt.Println("获取令牌成功")
	}

	// 3.阻塞直到获取足够的令牌或者上下文取消
	// 创建一个带有10秒超时的上下文context
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// 等待获取一个令牌
	// 如果当前获取令牌数超过最大限制或者通过context超时，返回错误
	limiter.Wait(ctx)
	err := limiter.WaitN(ctx, 20)
	if err != nil {
		fmt.Println("error", err)
	}

	// 4.预定令牌数量
	limiter.ReserveN(time.Now(), 1)
	// 4.预定令牌
	// 当调用Reserve后不管是否存在有效令牌都会返回Reservation指针对象
	// 通过返回的Reservation进行指定操作
	// Reserve()内部调用是ReserveN()
	reservation := limiter.Reserve()
	if 0 == reservation.Delay() {
		fmt.Println("获取令牌成功")
	}

	// 5.修改令牌生成速率
	limiter.SetLimit(rate.Every(time.Millisecond * 100))
	limiter.SetLimitAt(time.Now(), rate.Every(time.Millisecond*100))

	// 6.修改令牌桶大小，即生成令牌的最大数量限制
	limiter.SetBurst(50)
	limiter.SetBurstAt(time.Now(), 50)

	// 7.获取限流的速率，即rate.NewLimiter()的参数r的值
	// 每秒允许处理事件数，即每秒处理事件频率
	l := limiter.Limit()
	fmt.Printf("每秒允许处理事件数，即每秒处理事件频率为: %v\n", l)

	// 8.获取令牌桶的容量大小，即rate.NewLimiter()的参数n的值
	limiter.Burst()
}
