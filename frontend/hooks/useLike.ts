import axios from 'axios'
import { useEffect, useState } from 'react'
import { Like } from '../types/Like'

const useLike = (postId: string, userId?: string) => {
	const [likes, setLikes] = useState<Like[]>([])
	const [count, setCount] = useState(0)
	const [isLike, setIsLike] = useState(false)
	const getLikes = async () => {
		try {
			const likes = await axios.get<Like[]>(`/api/post/${postId}/likes`)
			setLikes(likes.data)
			setCount(likes.data.length)
		} catch (e) {
			console.error(e)
		}
	}
	useEffect(() => {
		getLikes()
		if (userId) {
			updateIsLike(userId)
		}
	}, [postId, userId, count])

	const updateIsLike = (userId: string) => {
		likes.map((like) => {
			like.likeUserId === userId && setIsLike(true)
		})
	}

	const updateLike = async (userId: string) => {
		try {
			await axios.put('/api/like', {
				postId: postId,
				userId: userId,
			})
		} catch (e) {
			console.error(e)
		}
	}
	const addCount = () => {
		setCount((c) => c + 1)
	}
	const subtractCount = () => {
		setCount((c) => c - 1)
	}
	return { likes, count, isLike, updateLike, addCount, subtractCount }
}

export default useLike
