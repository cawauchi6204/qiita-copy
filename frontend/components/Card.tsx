import Link from "next/link"
import { User } from "react-feather"


import useTag from "../hooks/useTag"
import { Post } from "../types/Post"
import Tags from "./Tags"

type Props = {
  post: Post
}

const Card: React.FC<Props> = ({ post }) => {
  return (
    <div key={post.id} className="w-full bg-white flex flex-col p-4">
      <div className="flex">
        <User size={40} />
        <div className="flex flex-col">
          <Link href={`/${post.postedBy}`}>
            <a className="hover:underline inline-block">
              {post.postedBy}
            </a>
          </Link>
          <time className="text-xs text-gray">
            {post.createdAt}
          </time>
        </div>
      </div>
      <div className="flex flex-col ml-10">
        <Link href={`/${post.postedBy}/items/${post.id}`}>
          <a className="text-2xl mt-2 font-bold hover:underline">
            {post.title}
          </a>
        </Link>
      </div>
      <Tags postId={post.id} />
    </div>
  )
}

export default Card
