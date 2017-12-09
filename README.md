# Cost Explorer
This is my first attempt at using Go.  The goal is to explore the new AWS Cost Explorer API to generate daily spend messages to slack.

Outputs a table like the one blow:

```
+--------------------------------+-------+-------+-------+-------+-------+-------+-------+
|          AWS SERVICE           | 12-01 | 12-02 | 12-03 | 12-04 | 12-05 | 12-06 | 12-07 |
+--------------------------------+-------+-------+-------+-------+-------+-------+-------+
| Amazon Elastic Compute Cloud - |  9.48 |  8.62 |  8.74 |  9.25 |  9.30 |  9.25 |  9.43 |
| Compute                        |       |       |       |       |       |       |       |
| Amazon Elastic Block Store     |  7.35 |  6.81 |  6.86 |  6.79 |  6.71 |  6.86 |  6.10 |
| Amazon Relational Database     |  3.79 |  3.78 |  3.78 |  3.78 |  3.78 |  3.78 |  3.75 |
| Service                        |       |       |       |       |       |       |       |
| Amazon ElastiCache             |  1.04 |  1.04 |  1.04 |  1.10 |  1.06 |  1.07 |  0.93 |
| Amazon Simple Storage Service  |  0.82 |  0.90 |  0.89 |  1.01 |  0.93 |  0.93 |  0.86 |
| Amazon Elastic Load Balancing  |  0.32 |  0.31 |  0.43 |  0.41 |  0.36 |  0.34 |  0.34 |
| AmazonCloudWatch               |  0.14 |  0.14 |  0.14 |  0.14 |  0.15 |  0.15 |  0.15 |
| Amazon Virtual Private Cloud   |  0.13 |  0.06 |  0.13 |  0.11 |  0.07 |  0.06 |  0.06 |
| Amazon Route 53                |  0.10 |  0.04 |  0.06 |  0.07 |  0.06 |  0.05 |  0.05 |
| Amazon CloudFront              |  0.05 |  0.04 |  0.04 |  0.04 |  0.04 |  0.04 |  0.04 |
| Amazon Kinesis Analytics       |  0.03 |  0.03 |  0.03 |  0.03 |  0.03 |  0.03 |  0.03 |
| Amazon Kinesis                 |  0.02 |  0.02 |  0.02 |  0.02 |  0.02 |  0.02 |  0.02 |
| AWS Directory Service          |  0.01 |  0.01 |  0.01 |  0.01 |  0.01 |  0.01 |  0.01 |
+--------------------------------+-------+-------+-------+-------+-------+-------+-------+
```
