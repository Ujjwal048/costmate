package constants

import "time"

// UseDummyData determines whether to use dummy data or fetch real costs from AWS
const UseDummyData = false
const DefaultCurrency = "USD"

var CurrentMonth = time.Now()

// ServiceCosts contains dummy data for testing
var ServiceCosts = []struct {
	ServiceName string
	Cost        float64
	Unit        string
	Percent     float64
}{
	{"EC2", 156.78, "USD", 0},
	{"S3", 45.32, "USD", 0},
	{"RDS", 89.45, "USD", 0},
	{"Lambda", 12.67, "USD", 0},
	{"CloudFront", 34.21, "USD", 0},
	{"DynamoDB", 23.45, "USD", 0},
	{"Elastic Beanstalk", 28.90, "USD", 0},
	{"ECS", 42.15, "USD", 0},
	{"EKS", 73.20, "USD", 0},
	{"ElastiCache", 67.89, "USD", 0},
	{"API Gateway", 15.43, "USD", 0},
	{"Route 53", 0.50, "USD", 0},
	{"CloudWatch", 8.75, "USD", 0},
	{"SNS", 3.25, "USD", 0},
	{"SQS", 4.80, "USD", 0},
	{"KMS", 2.50, "USD", 0},
	{"IAM", 0.00, "USD", 0},
	{"VPC", 0.00, "USD", 0},
	{"CloudTrail", 5.20, "USD", 0},
	{"WAF", 18.30, "USD", 0},
	{"Athena", 100.00, "USD", 0},
	{"ECR", 100.00, "USD", 0},
}

const TotalCost = 1000.00
