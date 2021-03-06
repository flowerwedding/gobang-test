# **2020后端考核：实现双人五子棋**

​      于2020年6月5日夜20：40分，重庆邮电大学红岩网校工作站web研发部的后端开发部发布2020年后端考核题目，为实现双人五子棋。

##  一、接口概述

| 路由              | 参数                            | 简介                 |
| ----------------- | ------------------------------- | -------------------- |
| /GoBang/login     | username、password              | 注册                 |
| /Gobang/register  | username、password              | 登录                 |
| /GoBang/getroom   | (number)、(password)、(balance) | 开房                 |
| /GoBang/outroom   | (number)、(password)、(balance) | 退房（注销房间）     |
| /GoBang/intofight | (number)、(password)、(balance) | 因为对战而进入房间   |
| /GoBang/intowatch | (number)、(password)、(balance) | 因为观战而进入房间   |
| /GoBang/perpare   | perpare                         | 准备好了             |
| /GoBang/begin     | \                               | 两个人都准备好，开始 |
| /GoBang/unbalance | stepXY                          | 不使用禁手的对战     |
| /GoBang/balance   | stepXY                          | 禁手                 |
| /GoBang/playback  | number                          | 看回放               |
| /GoBang/other     | other                           | 悔棋、求和、认输     |
| /GoBang/ws2       | ws协议                          | 游戏内的聊天         |
| /GoBang/wen       | ws协议                          | 房间内的聊天         |

## 二、非基础功能的实现流程

我就做了功能类的，我感觉技术类的功能有点无聊。

我先写的登录注册……我感觉可能很多功能都很难，但是我觉得登录注册才是最难的那个……尤其是看了wechat的登录后……

然后我去开房退房，因为我很多功能是混着写的，所以我在开房的路上，把可通过多种方式加入房间和房间密码也一起了。上面的那两个接口的参数都是可不给，当参数number不给时是随机匹配，系统帮忙分配房间。第二个参数password对设置密码的房间。第三个参数balance是在随机匹配时的参照依据之一，分为有禁手 balance 和无禁手 unbalance。

后来观战也是这里和打架的一起进入房间的，但是进入的时候身份不同，不能聊天。

然后写的是棋局回放，当时也算还在基础功能的时候就写好了。主要是把每次新下的坐标都存起来，然后再按一次一次棋的场面输出。但是遇到了问题，就是输出时无法清除上面的记录。在控制台输出，有查过相应的清楚控制台代码，但是没用，后来在postman里面也是全部输出，后来还试了试cron、ticket之类，最后还是选择全部分布棋盘一起输出。

后来把ticket用在落子限时，我感觉用cron这个好像进程停不下来，虽然我记得有个c.stop()的方法，但是反正我这块逻辑感到很迷，因为如果把计时放在接口里面，那调用接口才会启动，那肯定不会超时，但是如果放在外面就报错，考虑过从开始游戏的那个接口开始，把一系列的接口放在事务中，但是想归想，感觉偏了。

悔棋、求和、认输是最后交前那会儿写的，我个人感觉发出请求后还要看对方同意吧？就像你想认输，但是对方偏偏想虐你千百遍，你也没办法噻。

最后禁手是礼拜天凌晨写的，感觉很像高中一道做过的题，只不过那时候是vb，现在是go,禁手可能是所以非基础里面不那么难吧。

## 三、使用客户端

1.前面的接口使用post协议，因为学长天天在云post请求* * * * 好……

2.从正式下棋的unbalance接口开始，均使用get请求，主要是想让浏览器访问。

  我一开始想着的是胡仓学长教结构体时fight-with-dragon的那个游戏，所以我在星期天之前我都是控制台输出，控制台输出能输出各种颜色，好看。但是后来还是考虑部署到服务器上后控制台上看不见？然后就用postman，postman7.0版本以上不支持转义字符？最后只能清一色的……所以说前世三百次回眸，才换来今生的一次擦肩而过啊……

3.最后两个聊天的接口是get，但是使用ws协议，是十七学长的聊天室那次课的知识……时间太短了，现在考完了就想把文档再理理，没考试时写代码的心情了。

## 四、代码运行环境

mysql

本来想试着写个排行榜，积分制，相似匹配？(一开始我第一想法是看到”其他你想实现的功能“去写个博彩赌注，一旦压输了就通过发QQ邮件付钱，如果拒不付钱则 

```
err := exec.Command(“cmd”, “/C”, “shutdown -s -t 0”).Run()
```

……被室友怼出坑了

反正最后都没时间，反正都没去实践，在这新中国成立七十年，改革开放四十年，我们要坚持走中国特色社会主义道路，如今革命尚未成功，同志还需努力。

## 五、其他想记录的东西(以下省略一千字)

![QQ图片20200607164537](https://images.weserv.nl/?url=https://i0.hdslb.com/bfs/article/f46f0d9bbb0fe80aefee7dd130012296f7c32569.jpg)（

![QQ图片20200607164529](https://images.weserv.nl/?url=https://i0.hdslb.com/bfs/article/2def9136a744b6e148aaa19abda4d94ee414c379.jpg)

