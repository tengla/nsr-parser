
bucket=$(aws s3 ls | grep nsr-stops-import | cut -d' ' -f3)
file=$(aws s3 ls "$bucket/valid/" | cut -d' ' -f5)
aws s3 cp "s3://$bucket/valid/$file" $file
