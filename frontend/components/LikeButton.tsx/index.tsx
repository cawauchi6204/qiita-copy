import React from "react"
import { Heart } from "react-feather"
import useLike from "../../hooks/useLike"

type Props = {
  postId: string
  userId?: string
}

const LikeButton: React.FC<Props> = ({ postId, userId }) => {
  const { count, addCount, subtractCount, updateLike, isLike } = useLike(postId, userId)
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
              <span className="ml-2">{count}</span>
            </>
          )
            :
            (
              <>
                <Heart strokeWidth={3} onClick={() => {
                  updateLike(userId)
                  subtractCount()
                }} />
                <span className="ml-2">{count}</span>
              </>
            )
          }
        </>
      ) :

        (
          <>
            <Heart strokeWidth={3} />
            <span className="ml-2">{count}</span>
          </>
        )
      }
    </div>
  )
}

export default LikeButton
