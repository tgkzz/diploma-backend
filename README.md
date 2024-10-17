
# Finkomek backend application

Finkomek is a platform that provides financial education for citizens of Kazakhstan. This repository is a monolith, to this project. The production version is available at [link](https://finkomek-web.vercel.app/) (if it hasn't been cut off yet).


## Realisation

Since the project was a thesis and the deadline was tight, the project was done exactly as it is now. However, in my understanding there is a reference realisation 

![schema](https://iili.io/d91WlTu.md.png)

Services in the reference implementation should communicate via gRPC/HTTP or message brokers (gRPC, Kafka, RabbitMQ), but in my understanding it would be necessary to deploy several servers, which is not favourable in terms of diploma project development. 
## Future developement

At the moment, I have no plans to further develop this project, as the team is more than likely not going to pursue it. However, PRs are always welcome
