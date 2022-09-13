import Link from "next/link"
import { User, Pocket, Tag } from "react-feather"
import { Post } from "../types/Post"
import LikeButton from "./LikeButton.tsx"

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
      <div className="ml-10">
        <Tag className="inline" size={16} />
        <Link href="">
          <a>
            Git
          </a>
        </Link>
        <div className="flex justify-between mt-2">
          <LikeButton post={post} />
          <Pocket size={40} />
        </div>
      </div>
    </div >
  )
}

export default Card
