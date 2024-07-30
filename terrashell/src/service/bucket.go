package service

import (
	"terrashell/src/model"
	"terrashell/src/proxy"
	"terrashell/src/report"
)

const BucketModule = "aws_s3_bucket"
const BucketResource = "aws_s3_bucket"
const BucketInstance = "aws_s3_bucket"

type BucketService struct {
	TerraformProxy *proxy.TerraformProxy
	TrivyProxy     *proxy.TrivyProxy
	BucketReport   *report.BucketReport
}

func (bs *BucketService) Create(Bucket *model.Bucket) *model.Bucket {
	bs.TrivyProxy.Config(
		bs.TerraformProxy.InitAndPlanAndShow(
			BucketModule, BucketResource, BucketInstance, map[string]interface{}{
				"bucket": Bucket.Id,
			},
		),
	)
	options := bs.TerraformProxy.Apply(BucketModule, map[string]interface{}{
		"bucket": Bucket.Id,
	})
	result := &model.Bucket{
		Id:     bs.TerraformProxy.Output(options, "s3_bucket_id"),
		Arn:    bs.TerraformProxy.Output(options, "s3_bucket_arn"),
		Region: bs.TerraformProxy.Output(options, "s3_bucket_region"),
	}
	bs.BucketReport.Generate(result)
	return result
}

func (bs *BucketService) Read(Bucket *model.Bucket) *model.Bucket {
	bs.TerraformProxy.Init()
	options := bs.TerraformProxy.Import(BucketModule, BucketResource, BucketInstance, Bucket.Id)
	result := &model.Bucket{
		Id:     bs.TerraformProxy.Output(options, "s3_bucket_id"),
		Arn:    bs.TerraformProxy.Output(options, "s3_bucket_arn"),
		Region: bs.TerraformProxy.Output(options, "s3_bucket_region"),
	}
	bs.BucketReport.Generate(result)
	return result
}

func (bs *BucketService) Update(Bucket *model.Bucket) *model.Bucket {
	bs.TrivyProxy.Config(
		bs.TerraformProxy.InitAndPlanAndShow(
			BucketModule, BucketResource, BucketInstance, map[string]interface{}{
				"bucket": Bucket.Id,
			},
		),
	)
	options := bs.TerraformProxy.ApplyWithReplace(BucketModule, BucketResource, BucketInstance, map[string]interface{}{
		"bucket": Bucket.Id,
	})
	result := &model.Bucket{
		Id:     bs.TerraformProxy.Output(options, "s3_bucket_id"),
		Arn:    bs.TerraformProxy.Output(options, "s3_bucket_arn"),
		Region: bs.TerraformProxy.Output(options, "s3_bucket_region"),
	}
	bs.BucketReport.Generate(result)
	return result
}

func (bs *BucketService) Delete(Bucket *model.Bucket) *model.Bucket {
	bs.TerraformProxy.Init()
	options := bs.TerraformProxy.Destroy(BucketModule, BucketResource, BucketInstance, map[string]interface{}{
		"bucket": Bucket.Id,
	})
	result := &model.Bucket{
		Id:     bs.TerraformProxy.Output(options, "s3_bucket_id"),
		Arn:    bs.TerraformProxy.Output(options, "s3_bucket_arn"),
		Region: bs.TerraformProxy.Output(options, "s3_bucket_region"),
	}
	bs.BucketReport.Generate(result)
	return result
}
