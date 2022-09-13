import Link from "next/link"
import React from "react"
import { Heart } from "react-feather"
import useLike from "../../hooks/useLike"
import { Post } from "../../types/Post"

type Props = {
  post: Post
  userId?: string
}

const LikeButton: React.FC<Props> = ({ post, userId }) => {
  const { count, addCount, subtractCount, updateLike, isLike } = useLike(post.id, userId)
  return (
    <div>
      {userId ? (
        <>
          {isLike ? (
            <>
              <Heart onClick={() => {
                updateLike(userId)
                addCount()
              }} />
              <Link href={`/${post.postedBy}/items/${post.id}/likes`}>
                <a>
                  <span className="ml-2">{count}</span>
                </a>
              </Link>
            </>
          )
            :
            (
              <>
                <Heart strokeWidth={3} onClick={() => {
                  updateLike(userId)
                  subtractCount()
                }} />
                <Link href={`/${post.postedBy}/items/${post.id}/likes`}>
                  <a>
                    <span className="ml-2">{count}</span>
                  </a>
                </Link>
              </>
            )
          }
        </>
      ) :

        (
          <>
            <Heart strokeWidth={3} />
            <Link href={`/${post.postedBy}/items/${post.id}/likes`}>
              <a>
                <span className="ml-2">{count}</span>
              </a>
            </Link>
          </>
        )
      }
    </div>
  )
}

export default LikeButton
