export type Post = {
	id: string
	title: string
	body: string
	postedBy: string
	organizationId: string
	isDeleted: number
	isDraft: number
	createdAt: string
	updatedAt: string
}

export type PostResponse = {
	id: string
	title: string
	body: string
	posted_by: string
	organization_id: string
	is_deleted: number
	is_draft: number
	created_at: string
	updated_at: string
}
