import axios from 'axios'
import { useEffect, useState } from 'react'
import { BlogTag } from '../types/BlogTag'

const useTag = (postId: string) => {
	const [tags, setTags] = useState<BlogTag[]>([])
	const getTagsByPostId = async () => {
		const tags = await axios.get<BlogTag[]>(
			`http://localhost/post/${postId}/tags`
		)
		setTags(tags.data)
	}
	useEffect(() => {
		getTagsByPostId()
	}, [postId])
	return { tags }
}

export default useTag
