import {
	S3Client,
	PutObjectCommand,
	PutObjectCommandInput,
	ListObjectsCommand,
	ListObjectsCommandInput,
	GetObjectCommand,
	GetObjectCommandInput,
	DeleteObjectsCommand,
	DeleteObjectsCommandInput,
} from '@aws-sdk/client-s3'

const REGION = 'ap-northeast-1'
const s3Client = new S3Client({
	region: REGION,
	credentials: {
		accessKeyId: process.env.NEXT_PUBLIC_AWS_ACCESS_KEY!,
		secretAccessKey: process.env.NEXT_PUBLIC_AWS_SECRET_KEY!,
	},
})

const BUCKET_NAME = process.env.NEXT_PUBLIC_S3_BUCKET!

export const upload = async (fileName: string, body: any) => {
	const uploadParams: PutObjectCommandInput = {
		Bucket: BUCKET_NAME,
		Key: `images/${fileName}`,
		Body: body,
	}
	return await s3Client.send(new PutObjectCommand(uploadParams))
}

export const get = async (fileName: string) => {
	const getParams: GetObjectCommandInput = {
		Bucket: BUCKET_NAME,
		Key: `images/${fileName}`,
	}
	return await s3Client.send(new GetObjectCommand(getParams))
}

export const list = async () => {
	const listParams: ListObjectsCommandInput = {
		Bucket: BUCKET_NAME,
		Prefix: 'images',
	}
	return await s3Client.send(new ListObjectsCommand(listParams))
}

export const _delete = async () => {
	const deleteParams: DeleteObjectsCommandInput = {
		Bucket: BUCKET_NAME,
		Delete: { Objects: [{ Key: 'your/upload/location/file.txt' }] },
	}
	return await s3Client.send(new DeleteObjectsCommand(deleteParams))
}

export default s3Client
